package client

import (
	"crypto/tls"

	"github.com/ethereum/go-ethereum/rpc"
)

// newClient returns a new rpc client
func newClient(host string) (*rpc.Client, error) {
	endpoint, err := rpc.Dial(host)
	if err != nil {
		return nil, err
	}
	return endpoint, nil
}

type ClientWithTls struct {
	Endpoint  *rpc.Client
	ChainId   int
	TLSConfig *tls.Config
}

// NewTlsClient returns a new rpc client with tls
func NewTlsClient(host string, chainId int, tlsConfig *tls.Config) *ClientWithTls {
	endpoint, err := newClient(host)
	if err != nil {
		panic(err)
	}
	// todo: need to use tls config
	_ = endpoint

	return &ClientWithTls{}
}

type ClientWithSigner struct {
	Endpoint      *rpc.Client
	ChainId       int
	SignerAddress string
}

func NewL2Signer(host string, chainId int, signerAddress string) *ClientWithSigner {
	endpoint, err := newClient(host)
	if err != nil {
		panic(err)
	}
	return &ClientWithSigner{
		Endpoint:      endpoint,
		ChainId:       chainId,
		SignerAddress: signerAddress,
	}
}

func (c *ClientWithSigner) Sign(data any) ([]byte, error) {
	var result []byte
	err := c.Endpoint.Call(&result, "sign_tx", data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
