package sources

import (
	"github.com/isan-rivkin/kubefigure/sources/consuls"
	"github.com/isan-rivkin/kubefigure/sources/terraform"
	"github.com/isan-rivkin/kubefigure/sources/vault"
	"github.com/zclconf/go-cty/cty"
)

const (
	TerraformSource SourceType = "terraform"
	VaultSource     SourceType = "vault"
	ConsulSource    SourceType = "consulkv"
)

type DataSourceConfig struct {
	TerraformConf *terraform.SourceInput
	VaultConf     *vault.SourceInput
	ConsulConf    *consuls.SourceInput
}

type SourcePayload interface {
	Value() (cty.Value, error)
}

type DataSourceFactory interface {
	Create(input DataSourceConfig) (DataSource, error)
}

type DataSource interface {
	Get() (SourcePayload, error)
}
