package sources

import (
	"fmt"

	"github.com/isan-rivkin/kubefigure/common"
	"github.com/isan-rivkin/kubefigure/sources/vault"
	"github.com/zclconf/go-cty/cty"
)

type VaultPayload struct {
	dataSourceConfig DataSourceConfig
	payload          *vault.VaultPayload
}

func (vp *VaultPayload) Value() (cty.Value, error) {
	jp := vp.dataSourceConfig.VaultConf.Secret.SecretValuePath
	val, err := common.FindValInJson(jp, vp.payload.JsonBytes)

	if err != nil {
		return val, fmt.Errorf("failed getting jsonpath %s for vault value - %s", jp, err)
	}

	return val, err
}

type VaultDataSource struct {
	client           vault.Client
	dataSourceConfig DataSourceConfig
}

func (vds *VaultDataSource) Get() (SourcePayload, error) {

	secretPath := vds.dataSourceConfig.VaultConf.Secret.Path
	version := vds.dataSourceConfig.VaultConf.Secret.OptionalSecretVersion
	secretPayload, err := vds.client.ReadAsVal(secretPath, version)

	if err != nil {
		return nil, fmt.Errorf("failed read path %s from vault %s", secretPath, err)
	}

	return &VaultPayload{
		dataSourceConfig: vds.dataSourceConfig,
		payload:          secretPayload,
	}, nil
}

func NewVaultDataSource(c vault.Client, input DataSourceConfig) DataSource {
	return &VaultDataSource{
		client:           c,
		dataSourceConfig: input,
	}
}
