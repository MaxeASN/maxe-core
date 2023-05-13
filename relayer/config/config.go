package config

import (
	"time"

	"github.com/MaxeASN/maxe-core/relayer/client"
	"github.com/MaxeASN/maxe-core/service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

type Config struct {
	L1Config *L1Config
	L2Config *L2Config
	Signer   *SignerConfig
	TxMgr    *txmgr.Config
}

type L2Config struct {
	L2TxStateOracleAddr   common.Address
	L2TxStateOracleTopics []string
	L2TxTimeout           time.Duration
	NetworkTimeout        time.Duration
	ChainId               uint64
	Host                  string
	// disabled
	// PrivateKey          string
}

// todo: add diff signers
type SignerConfig struct {
	Host      string
	TlsConfig client.TLSConfig
}

type L1Config struct {
	Endpoints map[uint64]*client.Client
	Timeout   time.Duration
}

// NewConfig returns a new Config with ctx. If no specified, use the default config.
func NewConfig(ctx *cli.Context) *Config {
	// signer config
	var signerHost string
	if ctx.IsSet(SignerHostFlagName) {
		signerHost = ctx.String(SignerHostFlagName)
	} else {
		signerHost = DefaultSignerHTTPHost
	}

	// l2 config
	var td time.Duration
	var networkTd time.Duration
	if ctx.IsSet(L2TxTimeoutFlagName) {
		td, _ = time.ParseDuration(ctx.GlobalString(L2TxTimeoutFlagName))
	} else {
		td, _ = time.ParseDuration(DefaultL2TxTimeout)
	}

	if ctx.IsSet(NetworkTimeoutFlagName) {
		networkTd, _ = time.ParseDuration(ctx.GlobalString(NetworkTimeoutFlagName))
	} else {
		networkTd, _ = time.ParseDuration(DefaultNetworkTimeout)
	}

	l2Config := L2Config{
		L2TxStateOracleAddr:   common.HexToAddress(ctx.GlobalString(L2TxStateOracleAddrFlagName)),
		L2TxStateOracleTopics: ctx.GlobalStringSlice(L2TxStateOracleTopicFlagName),
		L2TxTimeout:           td,
		NetworkTimeout:        networkTd,
		ChainId:               ctx.GlobalUint64(L2NodeChainIdFlagName),
		Host:                  ctx.GlobalString(L2NodeHostFlagName),
	}

	l1Config := L1Config{
		Endpoints: make(map[uint64]*client.Client),
		Timeout:   networkTd,
	}

	// read tx manager config from args
	txmgrCfg := txmgr.ReadCliConfig(ctx)

	return &Config{
		L1Config: &l1Config,
		L2Config: &l2Config,
		Signer: &SignerConfig{
			Host: signerHost,
		},
		TxMgr: txmgrCfg,
	}
}

type TLSConfig struct {
}
