package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/websocket"
)

type WsClient struct {
	Host     string
	Name     string
	Count    uint64
	Endpoint *websocket.Conn
	Timeout  time.Duration

	// request channel
	requestCh chan Request

	// request for data, channel
	outRequestCh chan *outRequest

	// subscribe channel
	subRequestCh chan *subRequest

	// subscribe sets
	subRequests   map[uint64]*subRequest
	requestMu     sync.Mutex
	outRequests   map[uint64]*outRequest
	subMu         sync.RWMutex
	subscriptions map[string]*subscription
}

func NewWsClient(
	host string,
	name string,
	timeout time.Duration) (Client, error) {
	// dial webscoket client
	dialer := websocket.Dialer{
		HandshakeTimeout: timeout,
	}
	endpint, _, err := dialer.Dial(host, nil)
	if err != nil {
		log.Error("failed to dial websocket")
		panic(err)
	}

	// return new websocket client
	ws := &WsClient{
		Host:          host,
		Name:          name,
		Endpoint:      endpint,
		Timeout:       timeout,
		requestCh:     make(chan Request),
		outRequestCh:  make(chan *outRequest),
		subRequestCh:  make(chan *subRequest),
		subRequests:   make(map[uint64]*subRequest),
		outRequests:   make(map[uint64]*outRequest),
		subscriptions: make(map[string]*subscription),
	}

	// loop
	go ws.looping()

	return ws, nil
}

func (ws *WsClient) looping() {
	var wg sync.WaitGroup
	wg.Add(3)

	defer close(ws.requestCh)
	defer close(ws.outRequestCh)
	defer close(ws.subRequestCh)

	// for processing subscription requests
	go func() {
		for {
			select {
			case sub := <-ws.subRequestCh:
				// get id
				id := sub.req.ID
				// mutex
				ws.requestMu.Lock()
				ws.subRequests[id] = sub
				ws.requestMu.Unlock()
				// sub
				ws.requestCh <- *sub.req
			case out := <-ws.outRequestCh:
				ws.requestMu.Lock()
				ws.outRequests[out.req.ID] = out
				ws.requestMu.Unlock()

				// copy
				reqCopy := *out.req
				// request
				ws.requestCh <- reqCopy

			}
		}
	}()

	// reader loop
	go func() {
		for {
			t, r, err := ws.Endpoint.NextReader()
			if err != nil {
				// wg.Done()
				log.Error("error reading from websocket client", "sleep", "20s", "err", err)
				time.Sleep(20 * time.Second)
				continue
			}
			// check message type
			if t != websocket.TextMessage {
				continue
			}
			// read message
			message, err := io.ReadAll(r)
			if err != nil {
				// wg.Done()
				log.Error("error reading from websocket reader", "err", err)
				return
			}
			// unmarshal message
			msg, err := unmarshal(message)
			if err != nil {
				log.Error("unable to unmarshal message from websocket client", "err", err)
			}

			switch msg := msg.(type) {
			case *Response:
				ws.requestMu.Lock()
				if subReq, ok := ws.subRequests[msg.ID]; ok {
					delete(ws.subRequests, msg.ID)
					ws.requestMu.Unlock()

					var id string
					err := json.Unmarshal(msg.Result, &id)
					if err != nil {
						log.Error("unable to unmarshal id from websocket data", "err", err)
						return
					}
					// new subscription
					sub := newSubscription(id, ws)
					// store
					ws.subMu.Lock()
					ws.subscriptions[id] = sub
					ws.subMu.Unlock()

					go func() {
						subReq.subCh <- sub
					}()
					continue
				}
				if outReq, ok := ws.outRequests[msg.ID]; ok {
					delete(ws.outRequests, msg.ID)
					ws.requestMu.Unlock()
					//
					msgCopy := *msg
					go func() {
						outReq.resCh <- &msgCopy
					}()
				}
			case *Request:
				continue
			case *Notification:
				if msg.Method != "eth_subscription" {
					continue
				}
				//
				subRes := subscriptionResult{}
				err := json.Unmarshal(msg.Params, &subRes)
				if err != nil {
					wg.Done()
					log.Error("unable to unmarshal subscription result from websocket data", "err", err)
					continue
				}
				msgCopy := *msg
				ws.subMu.RLock()
				if sub, ok := ws.subscriptions[subRes.Subscription]; ok {
					ws.subMu.RUnlock()
					go func() {
						sub.notificationCh <- &msgCopy
					}()
				}
			default:
				log.Error("WsClient Looping", "Read.msg.unknown", msg)
			}

		}
	}()

	// writer loop
	go func() {
		for request := range ws.requestCh {
			data, err := json.Marshal(&request)
			if err != nil {
				wg.Done()
				log.Error("error marshaling request to websocket data", "err", err)
				return
			}
			// write message
			err = ws.Endpoint.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				wg.Done()
				log.Error("error writing to websocket client", "err", err)
				return
			}
		}
	}()
	log.Info("Webscoket client loop started", "name", ws.Name)
	wg.Wait()

}

func unmarshal(data []byte) (any, error) {
	t := YouKnowWho{}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	if t.Method != nil {
		// that maybe request data
		if t.ID != nil {
			req := Request{}
			err := json.Unmarshal(data, &req)
			return &req, err
		}
		// that maybe Notification data
		notification := Notification{}
		err := json.Unmarshal(data, &notification)
		return &notification, err
	}
	// that maybe response data
	resp := Response{}
	err = json.Unmarshal(data, &resp)
	return &resp, err
}

func (ws *WsClient) Close() error {
	return ws.Endpoint.Close()
}

func (ws *WsClient) Check(ctx context.Context) error {
	_, err := ws.ChainId(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ws *WsClient) Request(ctx context.Context, req *Request) (*Response, error) {
	/*
		@dev: send and receive
	*/
	outReq := outRequest{
		req:   req,
		resCh: make(chan *Response),
		errCh: make(chan error),
	}
	defer close(outReq.resCh)
	defer close(outReq.errCh)

	ws.outRequestCh <- &outReq
	select {
	case res := <-outReq.resCh:
		return res, nil
	case err := <-outReq.errCh:
		return nil, err
	}
}

func (ws *WsClient) Subscribe(ctx context.Context, req *Request) (Subscription, error) {
	// make subscription request
	subRequest := subRequest{
		req:   req,
		subCh: make(chan *subscription),
		errCh: make(chan error),
	}

	defer close(subRequest.subCh)
	defer close(subRequest.errCh)

	ws.subRequestCh <- &subRequest

	log.Info("WsClient Subscribe", "id", req.ID)

	select {
	case sub := <-subRequest.subCh:
		return sub, nil
	case err := <-subRequest.errCh:
		return nil, err
	}
}

func (ws *WsClient) ChainId(ctx context.Context) (uint64, error) {
	request := Request{
		ID:      1,
		Jsonrpc: "2.0",
		Method:  "eth_chainId",
		Params:  []json.RawMessage{},
	}
	res, err := ws.Request(ctx, &request)
	if err != nil {
		return 0, err
	}
	if res.Error != nil {
		return 0, errors.New(string(res.Error.Message))
	}

	i, err := strconv.ParseInt(strings.Replace(string(res.Result), "\"", "", -1), 0, 16)
	if err != nil {
		return 0, err
	}
	return uint64(i), err
}

func (ws *WsClient) EstimateGas(ctx context.Context, tx any) (uint64, error) {
	panic("Not implemented")
}

func (ws *WsClient) MaxPriorityFeePerGas(ctx context.Context, tx any) (uint64, error) {
	panic("Not implemented")
}

func (ws *WsClient) GasPrice(ctx context.Context) (uint64, error) {
	panic("Not implemented")
}

func (ws *WsClient) TransactionCount(ctx context.Context, address string, numberOrTag NumberOrTag) (uint64, error) {
	panic("Not implemented")
}

func (ws *WsClient) PendingNonceAt(ctx context.Context, address string) (uint64, error) {
	panic("Not implemented")
}

func (ws *WsClient) SendRawTransaction(ctx context.Context, rawTx string) (string, error) {
	panic("Not implemented")
}

func (ws *WsClient) TransactionByHash(ctx context.Context, hash string) (any, error) {
	panic("Not implemented")
}

func (ws *WsClient) TransactionReceipt(ctx context.Context, txHash string) (any, error) {
	panic("Not implemented")
}

func (ws *WsClient) SubscribeNewTopics(ctx context.Context, address []string, topics []string) (Subscription, error) {
	// new subscribe
	subTopic := &SubscriptionTopic{
		Address: address,
		Topics:  topics,
	}
	req := &Request{
		ID:      ws.nextId(),
		Jsonrpc: "2.0",
		Method:  "eth_subscribe",
		Params:  makeParams("logs", subTopic),
	}

	return ws.Subscribe(ctx, req)
}

func (ws *WsClient) Logs(ctx context.Context, filter any) (any, error) {
	panic("Not implemented")
}

func makeParams(params ...any) []json.RawMessage {
	if len(params) == 0 {
		return nil
	}
	res := make([]json.RawMessage, len(params))
	for i, param := range params {
		b, err := json.Marshal(param)
		if err != nil {
			log.Error("failed to marshal params", "param", param, "err", err)
			return nil
		}
		res[i] = b
	}
	return res
}

func (ws *WsClient) nextId() uint64 {
	return atomic.AddUint64(&ws.Count, 1)
}
