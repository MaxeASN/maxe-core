package event

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/MaxeASN/maxe-core/relayer/client"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

type Signer struct {
	Client *rpc.Client
	Online bool
}

func NewSigner(host string, tlsConfig *client.TLSConfig) *Signer {
	//
	var client *rpc.Client
	var err error
	// check if tls config is nil or not
	if tlsConfig == nil {
		client, err = rpc.Dial(host)
	} else {
		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{},
			},
		}
		client, err = rpc.DialOptions(context.Background(), host, rpc.WithHTTPClient(httpClient))
	}
	if err != nil {
		log.Error("failed to connect to layer2 signer", "err", err)
		return nil
	}
	log.Info("Init sign service success", "host", host)
	return &Signer{
		Client: client,
		Online: true,
	}
}

func (s *Signer) Ping() {
	// todo: update signer status

}

type ClientWithSigner struct {
	ChainId int
	Online  bool
	Address string
	Client  *rpc.Client
}

func NewClientWithSigner(
	host string,
	chainId int,
	address string,
	tlsConfig *client.TLSConfig) *ClientWithSigner {
	signer := NewSigner(host, tlsConfig)
	// return
	return &ClientWithSigner{
		ChainId: chainId,
		Address: address,
		Client:  signer.Client,
		Online:  signer.Online,
	}
}

type RepSignData struct {
	Signature string `json:"signature"`
}
