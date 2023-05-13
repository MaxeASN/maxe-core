package event

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/MaxeASN/maxe-core/relayer/client"
	"github.com/MaxeASN/maxe-core/service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
			eh.Events[e.TxHash] = *e

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
		return newSubmitter(event, signer), nil
	}
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
func Loop(signedTx *SignedTx, receiptCh chan *TxReceipt) error {
	receiptCh <- &TxReceipt{TxHash: signedTx.Hash}
	return nil
}

type SignedTx struct {
	Hash  string `json:"hash"`
	RawTx string `json:"tx"`
}

type Submitter interface {
	Result() <-chan *SignedTx
	Err() <-chan error
}

type SimpleSubmitter struct {
	resultCh chan *SignedTx
	err      chan error
}

func newSubmitter(e *TxEvent, signer *Signer) Submitter {
	s := &SimpleSubmitter{
		resultCh: make(chan *SignedTx),
		err:      make(chan error),
	}

	go s.callForSign(e, signer)
	return s
}

func (s *SimpleSubmitter) callForSign(e *TxEvent, signer *Signer) {
	var data = "abcd181823182f39786656f159cfb99fa5181823182f3978656f159cfb99fa55"
	var res RepSignData
	err := signer.Client.Call(&res, "maxe_sign",
		e.TypedData.TxInfo.ChainId,
		e.TypedData.From,
		data)
	// response error
	if err != nil {
		s.err <- err
		return
	}
	s.resultCh <- &SignedTx{
		Hash:  e.TxHash,
		RawTx: res.Signature,
	}
}

func (s *SimpleSubmitter) Result() <-chan *SignedTx {
	return s.resultCh
}

func (s *SimpleSubmitter) Err() <-chan error {
	return s.err
}
