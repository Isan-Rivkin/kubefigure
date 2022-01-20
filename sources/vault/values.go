package vault

import (
	"github.com/zclconf/go-cty/cty"
)

// OutputStatus outputs the values of Terraform output
type VaultPayload struct {
	// Attribute name in module
	// +optional
	Key string `json:"key"`
	// Value
	// +optional
	Value cty.Value `json:"value"`
}

func (v *Vaultclient) ReadAsVal(secretPath, optionalSecretVersion string) (*VaultPayload, error) {
	return nil, nil
	// raw, err := v.Read(secretPath, optionalSecretVersion)
	// if err != nil {
	// 	return nil, err
	// }

	// if raw == nil {
	// 	return nil, errors.New("ErrNilSecret")
	// }

	// jsonVal := ctyjson.SimpleJSONValue{cty.Value{v: raw
	// return &VaultPayload{}, nil
}
