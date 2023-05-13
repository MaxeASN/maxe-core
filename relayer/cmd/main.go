package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MaxeASN/maxe-core/relayer"
	"github.com/MaxeASN/maxe-core/relayer/flags"
	"github.com/urfave/cli"
)

var (
	Version   = "0.1.0"
	GitCommit = ""
	GitDate   = ""
)

func main() {
	// init app
	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s - %s", Version, GitCommit)
	app.Name = "maxe-core-relayer"
	app.Usage = "relay layer1 transactions"
	app.Description = "hello world."
	app.Action = entry()
	app.Commands = []cli.Command{}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Maxe relayer failed: %v", err)
	}
}

func entry() func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return relayer.Main(ctx)
	}
}
