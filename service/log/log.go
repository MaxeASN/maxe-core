package log

import (
	"io"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
)

const logFile = "console.log"

func multiWriter(ctx *cli.Context) io.Writer {
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	var writers []io.Writer
	// check if we need to add os.Stdout
	if enabled(ctx) {
		writers = append(writers, os.Stdout)
	}
	// out to file
	writers = append(writers, f)
	return io.MultiWriter(writers...)
}

func SetGlobalDefaults(ctx *cli.Context) {
	cfg := NewCliConfig(ctx)
	mw := multiWriter(ctx)

	log.Root().SetHandler(
		log.LvlFilterHandler(
			level(cfg.Level),
			log.StreamHandler(mw, format(cfg.Format)),
		),
	)
}
