package vault

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

// OutputStatus outputs the values of Vault output
type VaultPayload struct {
	// Attribute name in module
	// +optional
	Key string `json:"key"`
	// Value
	// +optional
	_         cty.Value `json:"value"`
	JsonVal   *ctyjson.SimpleJSONValue
	JsonBytes []byte
}

func (v *Vaultclient) secretToVal(secret interface{}) (*ctyjson.SimpleJSONValue, []byte, error) {
	bytes, err := json.Marshal(secret)

	if err != nil {
		return nil, nil, fmt.Errorf("failed marshling json vault %s", err)
	}

	// t.IsObjectType() is true
	t, err := ctyjson.ImpliedType(bytes)

	if err != nil {
		return nil, nil, fmt.Errorf("failed getting implied type vault %s", err)
	}

	val, err := ctyjson.Unmarshal(bytes, t)

	if err != nil {
		return nil, nil, err
	}

	jsonVal := &ctyjson.SimpleJSONValue{
		Value: val,
	}
	// m := jsonVal.AsValueMap()
	// for k, v := range m {
	// 	fmt.Println("kkk ", k)
	// 	fmt.Println("vvv ", v.AsString())
	// }

	return jsonVal, bytes, nil
}

func (v *Vaultclient) ReadAsVal(secretPath, optionalSecretVersion string) (*VaultPayload, error) {

	raw, err := v.Read(secretPath, optionalSecretVersion)

	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, errors.New("ErrNilSecret")
	}

	secretPayload, found := raw["data"]

	if !found {
		return nil, errors.New("ErrNoSecretData")
	}

	jsonVal, asBytes, err := v.secretToVal(secretPayload)

	if err != nil {
		return nil, fmt.Errorf("ErrConvertingVaultCtyJsonVal - %s", err)
	}

	return &VaultPayload{
		JsonVal:   jsonVal,
		JsonBytes: asBytes,
	}, nil
}
