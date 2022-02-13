package sources

import (
	"github.com/hashicorp/consul/api"
	"github.com/isan-rivkin/kubefigure/common"
	"github.com/isan-rivkin/kubefigure/sources/consuls"
	"github.com/zclconf/go-cty/cty"
)

type ConsulPayload struct {
	Payload          *consuls.ConsulPayload
	dataSourceConfig DataSourceConfig
}

func NewConsulPayload(p *consuls.ConsulPayload, input DataSourceConfig) SourcePayload {
	return &ConsulPayload{
		Payload:          p,
		dataSourceConfig: input,
	}
}

func (p *ConsulPayload) Value() (cty.Value, error) {

	jPath := p.dataSourceConfig.ConsulConf.KV.ValueJsonPath

	if jPath != "" {
		val, err := common.FindValInJson(jPath, p.Payload.Pair.Value)
		return val, err
	}

	return p.Payload.AsVal, nil
}

type ConsulDataSource struct {
	Client           consuls.Client
	Path             string
	Opts             *api.QueryOptions
	dataSourceConfig DataSourceConfig
}

func (c *ConsulDataSource) Get() (SourcePayload, error) {
	payload, err := c.Client.ReadKVAsVal(c.Path, c.Opts)

	if err != nil {
		return nil, err
	}

	return NewConsulPayload(payload, c.dataSourceConfig), nil
}

func NewConsulDataSource(c consuls.Client, input DataSourceConfig) DataSource {
	var opts *api.QueryOptions

	inputOpts := input.ConsulConf.KV.Opts

	if inputOpts != nil {
		opts = &api.QueryOptions{}

		if inputOpts.Namespace != "" {
			opts.Namespace = inputOpts.Namespace
		}

		if inputOpts.Datacenter != "" {
			opts.Datacenter = inputOpts.Datacenter
		}

		if inputOpts.Partition != "" {
			opts.Datacenter = inputOpts.Partition
		}
	}

	return &ConsulDataSource{
		Path:             input.ConsulConf.KV.Path,
		Client:           c,
		dataSourceConfig: input,
		Opts:             opts,
	}
}
