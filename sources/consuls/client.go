package consuls

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func Test() {

}

type Client interface {
	ReadKV(path string, opts *api.QueryOptions) (*api.KVPair, error)
	ReadKVAsVal(path string, opts *api.QueryOptions) (*ConsulPayload, error)
}

type BaseClient struct {
	Conf *api.Config
}

func NewClient(addr string, port int) Client {
	c := api.DefaultConfig()

	// if empty than default config will be take from the standard env
	if addr != "" && port > 0 {
		c.Address = fmt.Sprintf("%s:%d", addr, port)
	}

	return &BaseClient{
		Conf: c,
	}
}

func (c *BaseClient) getClient() (*api.Client, error) {
	// Get a new client
	client, err := api.NewClient(c.Conf)
	return client, err
}

func (c *BaseClient) ReadKV(path string, opts *api.QueryOptions) (*api.KVPair, error) {

	client, err := c.getClient()

	if err != nil {
		return nil, fmt.Errorf("failed connecting consul client - %v", err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// Lookup the pair
	pair, _, err := kv.Get(path, opts)

	if err != nil {
		return nil, fmt.Errorf("failed getting kv from consul %s - %v", path, err)
	}

	return pair, nil
}
