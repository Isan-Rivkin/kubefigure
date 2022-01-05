package terraform

import (
	"fmt"

	"github.com/hashicorp/terraform/states/statefile"
)

type RemoteStateConnector interface {
	Download() ([]byte, error)
	DownloadAsStatefile() (*statefile.File, error)
}

type TFState struct {
	File *statefile.File
}

func NewStateFromFile(File *statefile.File) *TFState {
	return &TFState{
		File: File,
	}
}

// returns all addresses of modules including full addresses of internal nested modules of other modules
// list include empty string as well which is the root module. i.e resources inside the state with no actual module
// module.A
// module.B
// module.A.module.C
// module.B.module.D
func (s *TFState) ListModulesAddrs() []string {
	var addrs []string
	for addr := range s.File.State.Modules {
		addrs = append(addrs, addr)
	}
	return addrs
}

// example
// data.terraform_remote_state.my_important_state\
// aws_security_group_rule.my_sg
// module.mod-name.module.internal-mod-name.datadog_monitor.my_dd_alert
func (s *TFState) ListAllResources() []string {
	var resources []string
	for _, module := range s.File.State.Modules {
		for _, r := range module.Resources {
			resources = append(resources, r.Addr.String())
		}
	}
	return resources
}

//func (s *TFState) GetAllResourceInstances()

func (s *TFState) OutputsStatus() ([]*OutputStatus, error) {
	outputs, err := GetOutputsFromStateFile(s.File)

	if err != nil {
		return nil, fmt.Errorf("failed parsing outputs from state file, Error %v", err)
	}

	return outputs, err
}
