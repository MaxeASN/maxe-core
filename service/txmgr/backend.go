package txmgr

import (
	"context"

	"github.com/MaxeASN/maxe-core/relayer/client"
)

type Backend interface {
	//
	ChainId(ctx context.Context) (uint64, error)
	//
	BlockNumber(ctx context.Context) (uint64, error)
	//
	NonceAt(ctx context.Context, account any, numberOrTag client.NumberOrTag) (uint64, error)
	// deprecated, use NonceAt instead
	PendingNonceAt(ctx context.Context, account any) (uint64, error)
	//
	EstimateGas(ctx context.Context, call any) (uint64, error)
	//
	SendTransaction(ctx context.Context, tx *Txpack) (string, error)
	//
	TransactionReceipt(ctx context.Context, txHash string) (any, error)
}
