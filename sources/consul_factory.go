package sources

import (
	"fmt"

	"github.com/isan-rivkin/kubefigure/sources/consuls"
)

type ConsulFactory struct {
}

func (cf *ConsulFactory) Create(input DataSourceConfig) (DataSource, error) {
	var c consuls.Client
	authType := input.ConsulConf.ConsulAuthType

	switch authType {
	case consuls.HTTPAuthType:
		addr := input.ConsulConf.HttpAuth.Address
		port := input.ConsulConf.HttpAuth.Port
		c = consuls.NewClient(addr, port)
	default:
		return nil, fmt.Errorf("no such auth method consul %s", authType)
	}

	return NewConsulDataSource(c, input), nil
}
