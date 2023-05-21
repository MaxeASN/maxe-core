package client

import (
	"context"
	"encoding/json"
	"errors"
)

type subscription struct {
	subId          string
	notificationCh chan *Notification
	client         Requester
	stopCh         chan struct{}
}

func (s *subscription) ID() string {
	return s.subId
}

// fixme: 修复
func (s *subscription) NotificationCh() <-chan *Notification {
	return s.notificationCh
}

func (s *subscription) Unsubscribe() error {
	req := &Request{
		ID:      1,
		Jsonrpc: "2.0",
		Method:  "eth_unsubscribe",
		Params:  makeParams(s.subId),
	}

	resp, err := s.client.Request(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return errors.New(string(resp.Error.Message))
	}
	return nil
}

func newSubscription(subId string, client Requester) *subscription {
	s := &subscription{
		client:         client,
		subId:          subId,
		notificationCh: make(chan *Notification, 1),
		stopCh:         make(chan struct{}),
	}

	go func() {
		// close the notification channel
		defer close(s.notificationCh)
		defer close(s.stopCh)

		for {
			select {
			case <-s.stopCh:
				s.Unsubscribe()
				return
			}
		}
	}()
	return s
}

func (s *subscription) Stop() {
	s.stopCh <- struct{}{}
}

type subRequest struct {
	req   *Request
	subCh chan *subscription
	errCh chan error
}

type subscriptionResult struct {
	Subscription string          `json:"subscription"`
	Result       json.RawMessage `json:"result"`
}

type outRequest struct {
	req   *Request
	resCh chan *Response
	errCh chan error
}
