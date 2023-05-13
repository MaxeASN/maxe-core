package flags

import (
	"github.com/MaxeASN/maxe-core/relayer/config"
	maxelog "github.com/MaxeASN/maxe-core/service/log"
	"github.com/MaxeASN/maxe-core/service/txmgr"
	"github.com/urfave/cli"
)

var Flags []cli.Flag

var requiredFlags = []cli.Flag{
	SignerHostFlag,
	L2NodeChainId,
	L2NodeHostFlag,
	L2TxStateOracleAddrFlag,
	L2TxStateOracleTopicFlag,
	NetworkTimeoutFlag,
	TLSCaFlag,
}

var (
	//required flags
	SignerHostFlag = cli.StringFlag{
		Name:  config.SignerHostFlagName,
		Usage: "RPC Host for L1 signer",
		Value: config.DefaultSignerHTTPHost,
	}
	L2NodeHostFlag = cli.StringFlag{
		Name:  config.L2NodeHostFlagName,
		Usage: "RPC Host for L2 relayer",
		Value: config.DefaultL2WsHost,
	}
	L2NodeChainId = cli.Uint64Flag{
		Name:  config.L2NodeChainIdFlagName,
		Usage: "Chain ID for L2 relayer",
		Value: 10,
	}
	L2TxStateOracleAddrFlag = cli.StringFlag{
		Name:  config.L2TxStateOracleAddrFlagName,
		Usage: "RPC address for L2 tx state oracle",
		Value: "",
	}
	L2TxStateOracleTopicFlag = cli.StringSliceFlag{
		Name:     config.L2TxStateOracleTopicFlagName,
		Usage:    "Topic for L2 tx state oracle",
		Required: true,
	}
	NetworkTimeoutFlag = cli.StringFlag{
		Name:  config.NetworkTimeoutFlagName,
		Usage: "Network timeout, default: 10s",
		Value: config.DefaultNetworkTimeout,
	}
	TLSCaFlag = cli.StringFlag{
		Name:  config.TLSCaFlagName,
		Usage: "TLS CA file path",
		Value: "tls/tls.ca",
	}
)

func init() {
	Flags = append(Flags, requiredFlags...)
	Flags = append(Flags, maxelog.CliFlags()...)
	Flags = append(Flags, txmgr.CliFlags()...)
}
