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

func (s *TFState) OutputsStatus() ([]*OutputStatus, error) {
	outputs, err := GetOutputsFromStateFile(s.File)

	if err != nil {
		return nil, fmt.Errorf("failed parsing outputs from state file, Error %v", err)
	}

	return outputs, err
}
