package log

import (
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
)

const (
	EnableFlagName = "log"
	LevelFlagName  = "log.level"
	FormatFlagName = "log.format"
)

const prefix = "maxe-log"

func CliFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  EnableFlagName,
			Usage: "Enable out to terminal, use: '--log' to enable, if not set, log will be printed to the file.",
		},
		cli.StringFlag{
			Name:  LevelFlagName,
			Value: "info",
			Usage: "Set log level, use: '--log.level=info', support: 'trace', 'debug', 'info', 'warn', 'error'",
		},
		cli.StringFlag{
			Name:  FormatFlagName,
			Value: "terminal",
			Usage: "Format log output, use: '--log.format=terminal', support: 'json', 'terminal', 'fmt'",
		},
	}
}

type CliType struct {
	Level  string
	Format string
}

func NewCliConfig(ctx *cli.Context) *CliType {
	cfg := defaults()
	cfg.Level = ctx.GlobalString(LevelFlagName)
	cfg.Format = ctx.GlobalString(FormatFlagName)
	return cfg
}

func enabled(ctx *cli.Context) bool {
	log.Output("this is a debug log", 3, 0, ctx)
	return ctx.Bool(EnableFlagName)
}

func defaults() *CliType {
	return &CliType{
		Level:  "info",
		Format: "terminal",
	}
}

func level(l string) log.Lvl {
	l = strings.ToLower(l)
	lvl, err := log.LvlFromString(l)
	if err != nil {
		panic("Cannot parse log level: " + err.Error())
	}
	return lvl
}

func format(f string) log.Format {
	switch f {
	case "terminal":
		return log.TerminalFormat(true)
	case "json":
		return log.JSONFormat()
	case "fmt":
		return log.LogfmtFormat()
	default:
		return log.LogfmtFormat()
	}
}
