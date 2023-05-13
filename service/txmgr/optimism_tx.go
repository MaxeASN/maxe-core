package txmgr

import (
	"context"

	"github.com/MaxeASN/maxe-core/relayer/client"
)

type optimism struct {
	ethWraper *ethereum
}

func NewOpBackend(ctx context.Context, chainId uint64, host string) *optimism {
	ethClient := NewEthBackend(ctx, chainId, host)
	return &optimism{
		ethWraper: ethClient,
	}
}

func (op *optimism) ChainId(ctx context.Context) (uint64, error) {
	return op.ethWraper.ChainId(ctx)
}

func (op *optimism) BlockNumber(ctx context.Context) (uint64, error) {
	return op.ethWraper.BlockNumber(ctx)
}

func (op *optimism) NonceAt(ctx context.Context, account any, numberOrTag client.NumberOrTag) (uint64, error) {
	return op.ethWraper.NonceAt(ctx, account, numberOrTag)
}

// deprecated, use NonceAt instead
func (op *optimism) PendingNonceAt(ctx context.Context, account any) (uint64, error) {
	panic("deprecated, use NonceAt instead")
}

func (op *optimism) EstimateGas(ctx context.Context, call any) (uint64, error) {
	return op.ethWraper.EstimateGas(ctx, call)
}

func (op *optimism) SendTransaction(ctx context.Context, tx *Txpack) (string, error) {
	return op.ethWraper.SendTransaction(ctx, tx)
}

func (op *optimism) TransactionReceipt(ctx context.Context, txHash string) (any, error) {
	return op.ethWraper.TransactionReceipt(ctx, txHash)
}
