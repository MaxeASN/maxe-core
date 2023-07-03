package txmgr

import (
	"context"
	"errors"
	"math/big"

	"github.com/MaxeASN/maxe-core/relayer/client"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

type ethereum struct {
	client  *ethclient.Client
	c       *rpc.Client
	chainId uint64
}

func NewEthBackend(ctx context.Context, chainId uint64, host string) *ethereum {
	client, err := ethclient.DialContext(ctx, host)
	r, err := rpc.DialContext(ctx, host)
	if err != nil {
		panic(err)
	}
	log.Info("Successfully connected to node", "host", host)
	return &ethereum{
		client:  client,
		chainId: chainId,
		c:       r,
	}
}

func (e *ethereum) ChainId(ctx context.Context) (uint64, error) {
	if e.chainId <= 0 {
		return 0, errors.New("chain id not set")
	}
	return e.chainId, nil
}

func (e *ethereum) BlockNumber(ctx context.Context) (uint64, error) {
	return e.client.BlockNumber(ctx)
}

func (e *ethereum) NonceAt(ctx context.Context, account any, numberOrTag client.NumberOrTag) (uint64, error) {
	// parse number or tag to number
	number := toBlockNum(numberOrTag)
	return e.client.NonceAt(ctx, account.(common.Address), number)
}

// deprecated, use NonceAt instead
func (e *ethereum) PendingNonceAt(ctx context.Context, account any) (uint64, error) {
	panic("deprecated, use NonceAt instead")
}

func (e *ethereum) EstimateGas(ctx context.Context, call any) (uint64, error) {
	callMsg := call.(eth.CallMsg)
	return e.client.EstimateGas(ctx, callMsg)
}

func (e *ethereum) SuggestGasPrice(ctx context.Context) (*big.Int, *big.Int, error) {
	gasTipCap, err := e.client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, err
	}
	header, err := e.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	// todo
	return gasTipCap, header.BaseFee, err
	// return gasTipCap, big.NewInt(50), nil
}

func (e *ethereum) SendTransaction(ctx context.Context, tx *Txpack) (string, error) {
	t := e.packTx(tx)
	err := e.client.SendTransaction(ctx, t)
	return t.Hash().String(), err
}

func (e *ethereum) SendRawTransaction(ctx context.Context, raw string) error {
	return e.c.CallContext(ctx, nil, "eth_sendRawTransaction", raw)
}

func (e *ethereum) TransactionReceipt(ctx context.Context, txHash string) (any, error) {
	th := common.HexToHash(txHash)
	return e.client.TransactionReceipt(ctx, th)
}

func (e *ethereum) packTx(tx *Txpack) *types.Transaction {
	// todo
	rawTx := &types.LegacyTx{}
	return types.NewTx(rawTx)
}

func toBlockNum(numberOrTag client.NumberOrTag) *big.Int {
	switch {
	case numberOrTag.Tag == "latest":
		return nil
	case numberOrTag.Tag == "pending":
		return big.NewInt(-1)
	case numberOrTag.Tag == "finalized":
		return big.NewInt(-3)
	case numberOrTag.Tag != "safe":
		return big.NewInt(-4)
	default:
		return big.NewInt(int64(numberOrTag.Number))
	}
}
