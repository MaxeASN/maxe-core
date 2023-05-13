package relayer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/MaxeASN/maxe-core/relayer/config"
	"github.com/MaxeASN/maxe-core/relayer/event"
	maxelog "github.com/MaxeASN/maxe-core/service/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
)

func Main(appCtx *cli.Context) error {
	// welcome
	message := "Welcome to Maxe Relayer console: "
	message += `
	    __  ______   _  __ ______
	   /  |/  /   | | |/ // ____/
	  / /|_/ / /| | |   // __/
	 / /  / / ___ |/   |/ /___
	/_/  /_/_/  |_/_/|_/_____/
		`
	fmt.Println(message)

	// set global log env
	maxelog.SetGlobalDefaults(appCtx)

	//
	ctx, cancel := context.WithCancel(context.Background())
	// _ = ctx

	// init relayer
	relayer := NewRelayer(appCtx)
	go relayer.Loop(ctx)
	// go relayer.MockData()

	// event listener loop
	go relayer.EventListener.Listen(
		ctx,
		relayer.TxsEventCh,
	)
	// event handler loop
	go relayer.EventHandler.Handle(ctx, relayer.txStateABI)
	// workerpool
	go relayer.WorkerPool.Start()

	log.Info("relayer", "status", "started")

	// defer
	defer func() {
		relayer.EventListener.Stop(ctx)
		relayer.EventHandler.Stop(ctx)
		relayer.WorkerPool.Stop()
	}()
	// wait for the system signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	}...)
	<-interrupt
	cancel()
	return nil
}

type Relayer struct {
	Config *config.Config

	WorkerPool *WorkerPool

	TxsEventCh    chan *event.TxEvent
	EventListener event.Listener
	EventHandler  event.Handler
	Signer        *event.Signer

	// txstate
	txReceiptCh chan *event.TxReceipt
	txStateABI  *abi.ABI

	// Metrics *Metrics

	Lock sync.RWMutex
}

func NewRelayer(ctx *cli.Context) *Relayer {
	cfg := config.NewConfig(ctx)

	txStateAbi, err := event.MetaData.GetAbi()
	if err != nil {
		log.Error("Failed to get abi", "err", err)
	}

	// event listener
	el := event.NewEventListener(
		cfg.L2Config.Host,
		cfg.L2Config.L2TxStateOracleAddr.String(),
		cfg.L2Config.L2TxStateOracleTopics,
		txStateAbi,
		cfg.L2Config.NetworkTimeout)

	// event handler
	eh := event.NewEventHandler(config.DefaultL2WsHost, cfg.L2Config.ChainId, cfg.L2Config.NetworkTimeout, cfg.TxMgr)

	// signer client, currently not used tls
	signer := event.NewSigner(cfg.Signer.Host, nil)

	return &Relayer{
		Config:        cfg,
		TxsEventCh:    make(chan *event.TxEvent, config.TxEventChannelLength),
		WorkerPool:    NewWorkerPool(event.Loop, 50, cfg.L1Config.Timeout),
		EventListener: el,
		EventHandler:  eh,
		Signer:        signer,
		txReceiptCh:   make(chan *event.TxReceipt, config.TxReceiptChannelLength),
		txStateABI:    txStateAbi,
	}
}

func (r *Relayer) Loop(ctx context.Context) {

	for {
		select {
		case e := <-r.TxsEventCh:
			go func() {
				// add to handler
				submitter, err := r.EventHandler.Submit(ctx, e, r.txReceiptCh, r.Signer)
				if err != nil {
					log.Error("Failed to submit event", "err", err)
					return
				}
				select {
				case err := <-submitter.Err():
					log.Info("Error submitting event", "err", err)
					return
				case signedTx := <-submitter.Result():
					log.Info("Event submitted", "result", signedTx)
					// got signed tx
					// serve signed tx, waiting for tx receipt
					r.WorkerPool.Serve(signedTx, r.txReceiptCh)
				}

				log.Info("submitted txEvent to worker pool", "Layer2_tx_hash", e.TxHash)
			}()
		case receipt := <-r.txReceiptCh:
			_ = receipt
			log.Info(">>>>>>>>>>>>>>> received tx receipt", "Layer2_tx_hash", receipt.TxHash)
		case <-ctx.Done():
			return
		}
	}
}

func (r *Relayer) MockData() {
	for {
		time.Sleep(10 * time.Second)
		r.TxsEventCh <- &event.TxEvent{}
	}
}
