package event

import (
	"context"

	txState "github.com/MaxeASN/maxe-core/relayer/contracts/txstateoracle"
)

type Transaction interface {
	Create(ctx context.Context, args ...any) (any, error)
}

type TxReceipt struct {
	TxHash string `json:"tx_hash"`
}

type TxEventParams = txState.TxStateL1transferEvent
type TxInfoParams = txState.TxStateTransactionInfo

var MetaData = txState.TxStateMetaData
