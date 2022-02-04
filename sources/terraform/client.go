package terraform

import "fmt"

type TFClient interface {
	GetState(connector RemoteStateConnector) (*TFState, error)
}

type TerraformClient struct {
}

func NewTerraformClient() TFClient {
	return &TerraformClient{}
}

func (tc *TerraformClient) GetState(connector RemoteStateConnector) (*TFState, error) {
	file, raw, err := connector.DownloadAsStatefile()

	if err != nil {
		return nil, fmt.Errorf("failed downloading state file with connector Error %v", err)
	}

	return NewStateFromFile(file, raw), nil
}
