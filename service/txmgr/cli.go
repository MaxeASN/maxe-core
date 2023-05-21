package txmgr

import (
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli"
)

const (
	HostsFlagName                = "txmgr.hosts"
	TxSendTimeoutFlagName        = "txmgr.send.timeout"
	TxCheckIntervalFlagName      = "txmgr.check.interval"
	ReceiptCheckIntervalFlagName = "txmgr.receipt.interval"
)

type Config struct {
	// hosts
	Hosts map[string]string
	// timeout for sending a transaction
	TxSendTimeout time.Duration
	// timeout for checking transaction status
	TxCheckInterval time.Duration
	// timeout for retrieving transaction receipt
	ReceiptCheckInterval time.Duration
	//
}

// CliFlags returns the cli flags for the txmgr module.
func CliFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringSliceFlag{
			Name:     HostsFlagName,
			Usage:    "txmgr hosts flag",
			Required: true,
		},
		cli.DurationFlag{
			Name:  TxSendTimeoutFlagName,
			Value: 10 * time.Second,
			Usage: "tx send timeout, default 10s",
		},
		cli.DurationFlag{
			Name:  TxCheckIntervalFlagName,
			Value: 10 * time.Second,
			Usage: "Interval for checking tx status, default 10s",
		},
		cli.DurationFlag{
			Name:  ReceiptCheckIntervalFlagName,
			Value: 6 * time.Second,
			Usage: "Interval for checking tx receipt, default 6s",
		},
	}
}

// ReadCliConfig reads the txmgr config from the command line context.
func ReadCliConfig(ctx *cli.Context) *Config {
	hosts := ctx.GlobalStringSlice(HostsFlagName)
	cfg := &Config{
		Hosts:                make(map[string]string),
		TxSendTimeout:        ctx.GlobalDuration(TxSendTimeoutFlagName),
		TxCheckInterval:      ctx.GlobalDuration(TxCheckIntervalFlagName),
		ReceiptCheckInterval: ctx.GlobalDuration(ReceiptCheckIntervalFlagName),
	}

	for _, h := range hosts {
		res := strings.Split(h, ",")
		if len(res) != 2 {
			panic(fmt.Errorf("invalid txmgr host: %s", h))
		}
		cfg.Hosts[res[0]] = res[1]
	}
	return cfg
}
