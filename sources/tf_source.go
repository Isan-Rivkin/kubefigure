package sources

import (
	"fmt"

	"github.com/isan-rivkin/kubefigure/common"
	"github.com/isan-rivkin/kubefigure/sources/terraform"
	"github.com/zclconf/go-cty/cty"
)

type TerraformPayload struct {
	state            *terraform.TFState
	dataSourceConfig DataSourceConfig
}

type TerraformDataSource struct {
	connector        terraform.RemoteStateConnector
	dataSourceConfig DataSourceConfig
}

func NewTFStatePayload(s *terraform.TFState, input DataSourceConfig) SourcePayload {
	return &TerraformPayload{
		state:            s,
		dataSourceConfig: input,
	}
}

func (tfp *TerraformPayload) Value() (cty.Value, error) {
	p := tfp.dataSourceConfig.TerraformConf.Value.StateValuePath
	val, err := common.FindValInJson(p, tfp.state.Raw)

	if err != nil {
		return val, fmt.Errorf("failed finding val jsonpath %s terraform - %s", p, err)
	}

	return val, err
}

func NewTerraformDataSource(connector terraform.RemoteStateConnector, input DataSourceConfig) *TerraformDataSource {
	return &TerraformDataSource{
		connector:        connector,
		dataSourceConfig: input,
	}
}

func (tfs *TerraformDataSource) Get() (SourcePayload, error) {
	c := terraform.NewTerraformClient()
	state, err := c.GetState(tfs.connector)

	if err != nil {
		return nil, fmt.Errorf("failed getting tf state source %s", err)
	}

	return NewTFStatePayload(state, tfs.dataSourceConfig), nil
}
