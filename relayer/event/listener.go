package event

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/MaxeASN/maxe-core/relayer/client"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	// txState "github.com/MaxeASN/maxe-core/relayer/contracts/txstateoracle"
)

type Listener interface {
	Listen(ctx context.Context, eventCh chan *TxEvent)
	Stop(ctx context.Context)
}

// todo: event can find in ../contracts/.../state.go
type TxEvent struct {
	// raw tx hash in l2
	TxHash string
	// typed tx
	TypedData *TxEventParams
}

type EventListener struct {
	// webscoket client
	Client client.Client
	// transaction state oracle address on chain
	STOAddress string
	// event channel
	// EventCh chan *TxEvent
	// topic ids, subscribe at once
	topics []string
	// subid for unsub
	sub client.Subscription
	// timeout
	timeout time.Duration
	// error
	reConn chan struct{}
	// host
	host string
	// lock
	lock sync.RWMutex
	//
	txStateABI *abi.ABI
}

// NewEventListener creates a new EventListener, which can be used to listen for l2 tx events.
func NewEventListener(host string, stoAddress string, topics []string, abi *abi.ABI, timeout time.Duration) Listener {
	// parse timeout
	endpoint, err := client.NewWsClient(host, "maxe-core event listener", timeout)
	if err != nil {
		log.Error("Failed to connect to Layer2 node", "err", err)
	}
	log.Info("Successfully connected to Layer2 node", "host", host, "listening contract address", stoAddress)

	return &EventListener{
		Client:     endpoint,
		STOAddress: stoAddress,
		topics:     topics,
		txStateABI: abi,
		timeout:    timeout,
		reConn:     make(chan struct{}),
		host:       host,
	}
}

// Listen listen events, and transfer to the event channel
func (el *EventListener) Listen(ctx context.Context, eventCh chan *TxEvent) {
	// init new timer
	connectionCheckTimer := time.NewTimer(el.timeout)

	// check conn loop
	go el.heartBeat(ctx, connectionCheckTimer)

	// subscribe
	// fixme: logic need to be changed
	el.subscribe(ctx, el.topics)

	// for loop
	for {
		select {
		case notification := <-el.sub.NotificationCh():
			// parse notification
			sr := client.SubscriptionResult{}
			if err := json.Unmarshal(notification.Params, &sr); err != nil {
				log.Error("failed to parse notification", "err", err)
				continue
			}

			// parse log
			event, err := el.parseLog(ctx, sr)
			if err != nil {
				log.Error("failed to parse log data", "err", err)
				continue
			}

			// return to the relayer
			eventCh <- event

		case <-el.reConn:
			endpoint, err := client.NewWsClient(el.host, "maxe-core event listener", el.timeout)
			if err != nil {
				log.Error("failed to connect to Layer2 node", "err", err)
			}

			el.lock.Lock()
			el.Client = endpoint
			el.lock.Unlock()

			// try to unsubscribe
			el.sub.Unsubscribe()
			// re-sub
			el.subscribe(ctx, el.topics)
		case <-ctx.Done():
			return
		}
	}
}

func (el *EventListener) parseLog(ctx context.Context, sr client.SubscriptionResult) (*TxEvent, error) {
	var event TxEvent
	// parse log
	data := strings.Replace(sr.Result.Data, "0x", "", -1)
	dataByte, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	// parse TxEventParams
	var res TxEventParams
	err = el.txStateABI.UnpackIntoInterface(&res, "L1transferEvent", dataByte)
	if err != nil {
		return nil, err
	}
	// transfer hex string to uint64
	sn, err := strconv.ParseUint(sr.Result.Topics[3], 0, 64)
	if err != nil {
		return nil, err
	}
	// pack TxEvent
	event = TxEvent{
		TxHash: sr.Result.TxHash,
		TypedData: &TxEventParams{
			Account: common.HexToAddress(sr.Result.Topics[1]),
			From:    common.HexToAddress(sr.Result.Topics[2]),
			SeqNum:  sn,
			TxInfo:  res.TxInfo,
		},
	}
	return &event, nil
}

func (el *EventListener) Stop(ctx context.Context) {
	select {
	case <-ctx.Done():
		el.sub.Stop()
		log.Info("event listener unsubscribed")
	default:
		err := el.Client.Close()
		if err != nil {
			log.Error("failed to close websocket client", "err", err)
		}
	}
}

func (el *EventListener) subscribe(ctx context.Context, topics []string) {
	sub, err := el.Client.SubscribeNewTopics(ctx, []string{el.STOAddress}, topics)
	if err != nil {
		panic(fmt.Sprintf("subscribe txEvent failed, Error = %s", err))
	}
	if sub.ID() == "" {
		panic("failed to subscribe txEvent, subscription id is empty")
	}
	el.lock.Lock()
	el.sub = sub
	el.lock.Unlock()
	log.Info("subscribe txEvent successed", "subId", sub.ID())
}

func (el *EventListener) heartBeat(ctx context.Context, connectionCheckTimer *time.Timer) {
	for range connectionCheckTimer.C {
		// check if the node is connected
		if err := el.Client.Check(ctx); err != nil {
			log.Error("========= event listener: ws-tunnel ==========", "status", "check failed", "error", err)
			el.reConn <- struct{}{}
		} else {
			// conn ok
			log.Trace("========= event listener: ws-tunnel ==========", "status", "ok")
		}

		// reset timer
		connectionCheckTimer.Reset(el.timeout)
	}
}
