package config

const (
	DefaultSignerHTTPHost = "http://localhost:23655"
	DefaultSignerWSHost   = "ws://localhost:23656"
)

const (
	DefaultL2HTTPHost = "http://localhost:8545"
	DefaultL2WsHost   = "ws://localhost:8546"
)

const (
	// DefaultL2TxTimeout is the default timeout for transactions.
	// We set this to 12 seconds as the default for now, which is 6 blocks in OP
	DefaultL2TxTimeout = "12s"

	// DefaultNetworkTimeout is the default timeout for L2 rpc requests.
	// We set this to 10 seconds as the default for now
	DefaultNetworkTimeout = "10s"
)

const (
	//
	TxEventChannelLength = 1000
	//
	TxReceiptChannelLength = 1000
)
