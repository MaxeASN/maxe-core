package txmgr

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Txpack struct {
	ChainId uint64
	// raw tx input data, can retrieve from tx input
	Input []byte
	// to address is the address of the recipient, or nil for contract creation
	To *common.Address
	// gas limit to be used in the constructed tx
	GasLimit uint64
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
	txMgr.register(1, NewEthBackend(ctx, 1, cfg.Hosts["eth"]))
	// txMgr.register(10, NewOpBackend(ctx, 10, cfg.Hosts["op"]))
	return txMgr
}

func (m *DefaultTxManager) register(chainId uint64, backend Backend) {
	m.Bankend[chainId] = backend
}

func (m *DefaultTxManager) Send(ctx context.Context, chainId uint64, tx *Txpack) (*types.Receipt, error) {
	return nil, nil
}

// packTx returns the constructed tx
func (m *DefaultTxManager) packTx(ctx context.Context, tx *Txpack) (any, error) {
	panic("not implemented")
}

// send send a tx to the client
func (m *DefaultTxManager) send(ctx context.Context, tx any) (any, error) {
	panic("not implemented")
}
