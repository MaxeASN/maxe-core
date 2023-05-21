package txmgr

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/MaxeASN/maxe-core/relayer/client"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type Txpack struct {
	ChainId uint64
	// raw tx input data, can retrieve from tx input
	Input []byte
	// to address is the address of the recipient, or nil for contract creation
	To common.Address
	//
	Value *big.Int
	//
	From common.Address
}

type DefaultTxManager struct {
	// need to regist different backend
	Bankend map[uint64]Backend
	config  *Config
	//
	// Signer map[uint64]any

	// lock
	lock sync.RWMutex
}

func NewDefaultTxManager(cfg *Config) *DefaultTxManager {
	txMgr := &DefaultTxManager{
		Bankend: make(map[uint64]Backend),
		config:  cfg,
		// Signer:  make(map[uint64]any),
	}
	ctx := context.Background()
	// todo: add const params for chainIds
	txMgr.register(10950, NewEthBackend(ctx, 10950, cfg.Hosts["eth"]))
	txMgr.register(23328, NewEthBackend(ctx, 23328, cfg.Hosts["eth"]))
	// txMgr.register(10, NewOpBackend(ctx, 10, cfg.Hosts["op"]))
	return txMgr
}

func (m *DefaultTxManager) register(chainId uint64, backend Backend) {
	m.Bankend[chainId] = backend
}

func (m *DefaultTxManager) Send(ctx context.Context, chainId uint64, tx *Txpack) (*types.Receipt, error) {
	return nil, nil
}

func (m *DefaultTxManager) Craft(ctx context.Context, tx *Txpack) (*types.Transaction, error) {
	log.Info(">>>>>>>", "tx craft", tx.ChainId)
	m.lock.RLock()
	backend, ok := m.Bankend[tx.ChainId]
	m.lock.RUnlock()
	if !ok {
		return nil, errors.New("not registered chainId")
	}

	tip, basefee, err := m.suggestGasPrice(ctx, tx.ChainId)
	if err != nil {
		return nil, err
	}
	gasFeeCap := new(big.Int).Add(tip, new(big.Int).Mul(basefee, big.NewInt(2)))
	nonce, _ := backend.NonceAt(ctx, tx.From, client.NumberOrTag{Tag: client.TagPending})
	log.Info("?>?>?>?>?", "nonce", nonce)
	log.Info("?>?>?>?>?", "maxfeepergas", gasFeeCap.Uint64(), "tip", tip.Uint64())
	rawTx := &types.DynamicFeeTx{
		ChainID:    big.NewInt(int64(tx.ChainId)),
		Nonce:      nonce,
		To:         &tx.To,
		Value:      tx.Value,
		GasFeeCap:  gasFeeCap,
		GasTipCap:  tip,
		Data:       tx.Input,
		AccessList: nil,
	}

	gas, err := m.Bankend[tx.ChainId].EstimateGas(ctx, eth.CallMsg{
		From:      tx.From,
		To:        &tx.To,
		GasFeeCap: gasFeeCap,
		GasTipCap: tip,
		Data:      tx.Input,
	})
	if err != nil {
		return nil, err
	}

	rawTx.Gas = gas
	log.Info("<<<", "gas", gas)

	newTx := types.NewTx(rawTx)

	return newTx, nil
}

func (m *DefaultTxManager) WithSignature(ctx context.Context, tx *types.Transaction, signature string) (*types.Transaction, error) {
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(types.LatestSignerForChainID(tx.ChainId()), sig)
}

func (m *DefaultTxManager) suggestGasPrice(ctx context.Context, chainId uint64) (*big.Int, *big.Int, error) {
	m.lock.RLock()
	backend, ok := m.Bankend[chainId]
	m.lock.RUnlock()
	if !ok {
		return nil, nil, errors.New("not registered chainId")
	}
	tip, baseFee, err := backend.SuggestGasPrice(ctx)
	if err != nil {
		return nil, nil, err
	}
	log.Info(">>>>>> suggestGasPriceCaps", "suggestGasPriceCaps", tip, "baseFee", baseFee)

	return tip, baseFee, nil
}

func (m *DefaultTxManager) Raw(ctx context.Context, tx any) (string, error) {
	signedTx := tx.(*types.Transaction)
	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		log.Error("Failed to serialize transaction: %v", err)
		return "", err
	}
	rawTxHex := hexutil.Encode(rawTxBytes)
	fmt.Printf("Raw transaction: %s\n", rawTxHex)
	return rawTxHex, nil
}

// packTx returns the constructed tx
func (m *DefaultTxManager) packTx(ctx context.Context, tx *Txpack) (any, error) {
	panic("not implemented")
}

// send send a tx to the client
func (m *DefaultTxManager) send(ctx context.Context, tx any) (any, error) {
	panic("not implemented")
}
