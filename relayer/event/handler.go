package event

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/MaxeASN/maxe-core/relayer/client"
	"github.com/MaxeASN/maxe-core/service/txmgr"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

const (
	// l1 tx gnerated
	Generated uint8 = iota
	// l1 tx send to the node
	Send
	// l1 tx received by the node
	Pending
	Successful
	Failed
)

type Handler interface {
	Handle(ctx context.Context, abi *abi.ABI)
	Submit(ctx context.Context, event *TxEvent, receiptCh chan *TxReceipt, signer *Signer) (Submitter, error)
	UpdateState(ctx context.Context, signer *Signer, receipt *TxReceipt) (Updater, error)
	Stop(ctx context.Context)
}

const eventsfile = "events.rlp"

type EventHandler struct {
	// layer 2 chainId
	chainId uint64
	// websocket client
	Client client.Client
	// pending events, it will writes the events.rlp before exit.
	Events   map[string]TxEvent
	EventsCh chan *TxEvent
	EventsMu sync.RWMutex

	// tx manager
	txMgr *txmgr.DefaultTxManager

	// reconn
	reConn chan struct{}
	// host
	host string
	// timeout
	timeout time.Duration
	// lock
	lock sync.RWMutex
}

func NewEventHandler(host string, chainId uint64, timeout time.Duration, txmgrCfg *txmgr.Config) Handler {
	// parse timeout
	endpoint, err := client.NewWsClient(host, "maxe-core event handler", timeout)
	//

	if err != nil {
		log.Error("failed to connect to layer2 node", "err", err)
	}
	log.Info("Successfully connected to layer2 node", "chainId", chainId)

	return &EventHandler{
		chainId:  chainId,
		Client:   endpoint,
		timeout:  timeout,
		Events:   make(map[string]TxEvent),
		EventsCh: make(chan *TxEvent),
		reConn:   make(chan struct{}),
		host:     host,
		txMgr:    txmgr.NewDefaultTxManager(txmgrCfg),
	}
}

func (eh *EventHandler) Handle(ctx context.Context, abi *abi.ABI) {
	connCheckTimer := time.NewTimer(eh.timeout)
	go eh.heartBeat(ctx, connCheckTimer)

	// load events.rlp
	eh.loadEvents()

	defer close(eh.reConn)

	for {
		select {
		case <-ctx.Done():
			// todo
			eh.writeEvents()
			return
		case e := <-eh.EventsCh:
			// todo
			// store event
			eh.EventsMu.Lock()
			eh.Events[e.TxHash] = *e
			eh.EventsMu.Unlock()

		case <-eh.reConn:
			endpoint, err := client.NewWsClient(eh.host, "maxe-core event handler", eh.timeout)
			if err != nil {
				log.Error("failed to connect to layer2 node", "err", err)
			}

			eh.lock.Lock()
			eh.Client = endpoint
			eh.lock.Unlock()
		}
	}
}

func (eh *EventHandler) Submit(ctx context.Context, event *TxEvent, receiptCh chan *TxReceipt, signer *Signer) (Submitter, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case eh.EventsCh <- event:
		return newSubmitter(ctx, event, signer, eh.txMgr), nil
	}
}

func (eh *EventHandler) UpdateState(ctx context.Context, signer *Signer, receipt *TxReceipt) (Updater, error) {
	txStateAbi, _ := MetaData.GetAbi()
	eh.EventsMu.RLock()
	event, ok := eh.Events[receipt.L2TxHash]
	eh.EventsMu.RUnlock()
	if !ok {
		return nil, errors.New("l2 tx not found")
	}

	tsp := &TxStatePara{
		L2Account: event.TypedData.Account,
		From:      event.TypedData.From,
		SeqNum:    event.TypedData.TxInfo.SeqNum,
		ChainId:   eh.chainId,
		TxReceipt: *receipt,
	}
	return newUpdater(ctx, eh.chainId, tsp, txStateAbi, signer, eh.txMgr), nil
}

func (eh *EventHandler) Stop(ctx context.Context) {
	err := eh.Client.Close()
	if err != nil {
		log.Error("failed to close websocket client", "err", err)
	}
}

func (eh *EventHandler) loadEvents() {
	f, err := os.OpenFile(eventsfile, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Error("failed to open events file", "err", err)
		return
	}
	defer f.Close()
	// decode events.rlp

	log.Info("events.rlp", "file", "loaded")

}

func (eh *EventHandler) writeEvents() {
	f, err := os.OpenFile(eventsfile, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Error("failed to open events file", "err", err)
		return
	}
	defer f.Close()
	// locker
	eh.lock.RLock()
	defer eh.lock.RUnlock()
	// encode events.rlp
	// write events.rlp

}

// heartBeat check the connection is lost or timeout
func (eh *EventHandler) heartBeat(ctx context.Context, connectionCheckTimer *time.Timer) {
	for range connectionCheckTimer.C {
		// check if the node is connected
		if err := eh.Client.Check(ctx); err != nil {
			log.Error("========= event handler: ws-tunnel ==========", "status", "check failed", "error", err)
			eh.reConn <- struct{}{}
		} else {
			log.Trace("========= event handler: ws-tunnel ==========", "status", "ok")
		}

		// reset timer
		connectionCheckTimer.Reset(eh.timeout)
	}
}

// Loop send signed tx to the workerpool, listen for tx receipt,
// decode and submit to the receipt channel
func Loop(work *Work, receiptCh chan *TxReceipt, receiptMgr *txmgr.DefaultTxManager) error {
	err := receiptMgr.Bankend[work.ChainId].SendRawTransaction(context.Background(), work.SignedTx.L1RawTx)
	if err != nil {
		return err
	}
	for {
		// wait for tx receipt
		receipt, err := receiptMgr.Bankend[work.ChainId].TransactionReceipt(context.Background(), work.SignedTx.L1txHash)
		if errors.Is(err, ethereum.NotFound) {
			log.Trace("Transaction not yet mined", "hash", work.SignedTx.L1txHash)
			time.Sleep(5 * time.Second)
			continue
		} else if err != nil {
			log.Info("Receipt retrieval failed", "hash", work.SignedTx.L1txHash, "err", err)
			time.Sleep(5 * time.Second)
			continue
		} else if receipt == nil {
			log.Warn("Receipt and error are both nil", "hash", work.SignedTx.L1txHash)
			time.Sleep(5 * time.Second)
			continue
		}
		r := receipt.(*types.Receipt)
		log.Info("Got receipt", "L1_tx_hash", work.SignedTx.L1txHash, "blk_num", r.BlockNumber)
		// check tx status
		if r.Status == 1 {
			receiptCh <- &TxReceipt{
				L2TxHash: work.L2txHash,
				TxHash:   work.SignedTx.L1txHash,
				Status:   Successful,
			}
			return nil
		}
		if r.Status == 0 {
			receiptCh <- &TxReceipt{
				L2TxHash: work.L2txHash,
				TxHash:   work.SignedTx.L1txHash,
				Status:   Failed,
			}
			return nil
		}

	}
}

type Work struct {
	L2txHash string `json:"l2txhash"`
	SignedTx *SignedTx
	ChainId  uint64 `json:"chainid"`
}

type SignedTx struct {
	L1txHash string `json:"l1txhash"`
	L1RawTx  string `json:"tx"`
}

type Submitter interface {
	Result() <-chan *Work
	Err() <-chan error
}

type SimpleSubmitter struct {
	resultCh chan *Work
	err      chan error
}

func newSubmitter(ctx context.Context, e *TxEvent, signer *Signer, txmgr *txmgr.DefaultTxManager) Submitter {
	s := &SimpleSubmitter{
		resultCh: make(chan *Work),
		err:      make(chan error),
	}

	go s.callForSign(ctx, e, signer, txmgr)
	return s
}

func (s *SimpleSubmitter) callForSign(ctx context.Context, e *TxEvent, signer *Signer, tm *txmgr.DefaultTxManager) {
	// craft tx
	pack := &txmgr.Txpack{
		ChainId: e.TypedData.TxInfo.ChainId,
		Input:   e.TypedData.TxInfo.Data,
		To:      e.TypedData.TxInfo.Receiver,
		Value:   e.TypedData.TxInfo.Amount,
		From:    e.TypedData.TxInfo.From,
	}

	tx, err := tm.Craft(ctx, pack)
	if err != nil {
		s.err <- err
		return
	}

	// get tx hash
	hash := strings.Replace(types.LatestSignerForChainID(tx.ChainId()).Hash(tx).Hex(), "0x", "", -1)
	// sign the tx hash
	var res RepSignData
	err = signer.Client.Call(&res, "maxe_sign",
		e.TypedData.TxInfo.ChainId,
		e.TypedData.From,
		hash)
	// response error
	if err != nil {
		s.err <- err
		return
	}

	signedTx, err := tm.WithSignature(ctx, tx, res.Signature)
	if err != nil {
		s.err <- err
		return
	}

	// rawtx
	raw, err := tm.Raw(ctx, signedTx)
	if err != nil {
		s.err <- err
		return
	}

	s.resultCh <- &Work{
		ChainId:  e.TypedData.TxInfo.ChainId,
		L2txHash: e.TxHash,
		SignedTx: &SignedTx{
			L1txHash: signedTx.Hash().String(),
			L1RawTx:  raw,
		},
	}

}

func (s *SimpleSubmitter) Result() <-chan *Work {
	return s.resultCh
}

func (s *SimpleSubmitter) Err() <-chan error {
	return s.err
}

type UpdateResult struct {
	L2TxHash string
	TxState  uint8
}
type Updater interface {
	Result() <-chan *UpdateResult
	Err() <-chan error
}

type SimpleUpdater struct {
	resultCh chan *UpdateResult
	err      chan error
}

func newUpdater(ctx context.Context, chainId uint64, txState *TxStatePara, abi *abi.ABI, signer *Signer, tm *txmgr.DefaultTxManager) Updater {
	sr := &SimpleUpdater{
		resultCh: make(chan *UpdateResult),
		err:      make(chan error),
	}
	txhash := strings.Replace(txState.TxReceipt.TxHash, "0x", "", -1)
	hashBytes, _ := hex.DecodeString(txhash)
	go func() {
		//pack
		d, err := abi.Pack("setL1TxState",
			hashBytes,
			txState.L2Account,
			txState.From,
			txState.SeqNum,
			big.NewInt(int64(txState.TxReceipt.Status)))
		if err != nil {
			sr.err <- err
			return
		}


		pack := &txmgr.Txpack{
			ChainId: chainId,
			Input:   d,
			To:      common.HexToAddress("0x933D5BfE7695AB56A400e0EC9783F7Eca6aC64ff"),
			Value:   big.NewInt(0),
			From:    common.HexToAddress("0x11Ab1F1E7CF1442b0D994313622f0Cc419E178E4"),
		}

		tx, err := tm.Craft(ctx, pack)
		if err != nil {
			sr.err <- err
			return
		}

		// get tx hash
		hash := strings.Replace(types.LatestSignerForChainID(tx.ChainId()).Hash(tx).Hex(), "0x", "", -1)
		// sign the tx hash
		var res RepSignData
		err = signer.Client.Call(&res, "maxe_sign",
			pack.ChainId,
			pack.From,
			hash)
		// response error
		if err != nil {
			sr.err <- err
			return
		}

		signedTx, _ := tm.WithSignature(ctx, tx, res.Signature)
		// rawtx
		raw, _ := tm.Raw(ctx, signedTx)

		err = tm.Bankend[pack.ChainId].SendRawTransaction(ctx, raw)
		if err != nil {
			sr.err <- err
			return
		}

		sr.resultCh <- &UpdateResult{
			L2TxHash: signedTx.Hash().String(),
			TxState:  Successful,
		}
	}()
	return sr
}

func (s *SimpleUpdater) Result() <-chan *UpdateResult {
	return s.resultCh
}

func (s *SimpleUpdater) Err() <-chan error {
	return s.err
}
