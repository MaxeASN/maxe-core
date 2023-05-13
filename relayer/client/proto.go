package client

import (
	"encoding/json"
)

type Request struct {
	Jsonrpc string            `json:"jsonrpc"`
	ID      uint64            `json:"id"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
}

type Response struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      uint64          `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *Error          `json:"error,omitempty"`
}

type YouKnowWho struct {
	Jsonrpc *string          `json:"jsonrpc"`
	ID      *json.RawMessage `json:"id"`
	Method  *string          `json:"method"`
	Params  *json.RawMessage `json:"params"`
	Result  *json.RawMessage `json:"result,omitempty"`
	Error   *json.RawMessage `json:"error,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	// Earliest block is the genesis block.
	TagEarliest = "earliest"
	// Latest block is the latest block.
	TagLatest = "latest"
	// Pending block is a block that is pending.
	TagPending = "pending"
)

// NumberOrTag returns the block number or tag.
// tag is optional, can see the block tag above.
type NumberOrTag struct {
	Number int64
	Tag    string
}

type SubscriptionTopic struct {
	Address []string `json:"address"`
	Topics  []string `json:"topics"`
}

type SubscriptionResult struct {
	Subscription string  `json:"subscription"`
	Result       *result `json:"result"`
	// Result *types.Log `json:"result,omitempty"`
}

// todo: can replace with types.log
type result struct {
	Address     string   `json:"address"`
	Topics      []string `json:"topics"`
	Data        string   `json:"data"`
	BlockNumber string   `json:"blockNumber"`
	BlockHash   string   `json:"blockHash"`
	TxHash      string   `json:"transactionHash"`
	TxIndex     string   `json:"transactionIndex"`
	LogIndex    string   `json:"logIndex"`
	Removed     bool     `json:"removed"`
}

// type topic string
type Notification struct {
	Jsonrpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}
