package terraform

import (
	"github.com/hashicorp/terraform/states"
	"github.com/hashicorp/terraform/states/statefile"
	"github.com/isan-rivkin/kubefigure/common"
)

// OutputStatus outputs the values of Terraform output
type OutputStatus struct {
	// Attribute name in module
	// +optional
	Key string `json:"key"`
	// Value
	// +optional
	Value string `json:"value"`
}

func GetOutputsValuesFromStateFile(file *statefile.File) map[string]*states.OutputValue {
	return file.State.Modules[""].OutputValues
}

func GetOutputsFromStateFile(file *statefile.File) ([]*OutputStatus, error) {

	outputValues := file.State.Modules[""].OutputValues
	outputs := []*OutputStatus{}
	for key, value := range outputValues {
		if !value.Sensitive {
			statusValue := common.ConvertValueToString(value.Value)
			if statusValue != "" {
				outputs = append(outputs, &OutputStatus{Key: key, Value: statusValue})
			}
		}
	}
	return outputs, nil
}
