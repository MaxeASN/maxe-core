package client

import (
	"context"
)

type Requester interface {
	Request(ctx context.Context, req *Request) (*Response, error)
}
type Subscriber interface {
	Subscribe(ctx context.Context, req *Request) (Subscription, error)
}

type Heartbeater interface {
	Check(ctx context.Context) error
}

type Closer interface {
	Close() error
}

type Client interface {
	//
	Requester
	//
	Subscriber
	//
	Heartbeater
	//
	Closer
	// ChainId returns the id of the chain
	ChainId(ctx context.Context) (uint64, error)
	// EstimateGas returns the execution gas of the transaction
	EstimateGas(ctx context.Context, tx any) (uint64, error)
	// MaxPriorityFeePerGas
	MaxPriorityFeePerGas(ctx context.Context, tx any) (uint64, error)
	// GasPrice returns the suggested tip for block
	GasPrice(ctx context.Context) (uint64, error)
	// TransactionCount returns nonce of the address
	TransactionCount(ctx context.Context, address string, numberOrTag NumberOrTag) (uint64, error)
	// SendRawTransaction send the raw signed tx to the client
	SendRawTransaction(ctx context.Context, rawTx string) (string, error)
	//
	TransactionByHash(ctx context.Context, hash string) (any, error)
	//
	TransactionReceipt(ctx context.Context, hash string) (any, error)
	// SubscribeNewTopics
	SubscribeNewTopics(ctx context.Context, address []string, topics []string) (Subscription, error)
	//
	Logs(ctx context.Context, filter any) (any, error)
}

type Subscription interface {
	ID() string
	NotificationCh() chan *Notification
	Unsubscribe() error
	Stop()
}
