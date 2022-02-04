package common

import (
	"encoding/json"
	"fmt"

	"github.com/yalp/jsonpath"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

func getImpliedType(o interface{}) (cty.Type, error) {
	ty, err := gocty.ImpliedType(o)

	if err == nil {
		return ty, nil
	}

	bytes, err := json.Marshal(o)

	if err != nil {
		return cty.Type{}, fmt.Errorf("failed getting type not json %s", err)
	}

	ty, err = ctyjson.ImpliedType(bytes)

	if err != nil {
		return ty, fmt.Errorf("failed getting implied type %s", err)
	}

	return ty, nil
}

func FindValInJson(jPath string, raw []byte) (cty.Value, error) {
	v := cty.Value{}
	jsonPathFilter, err := jsonpath.Prepare(jPath)
	if err != nil {
		return v, fmt.Errorf("failed preparing json filter %s", err)
	}

	var data interface{}
	if err = json.Unmarshal(raw, &data); err != nil {
		return v, fmt.Errorf("failed parsing raw bytes to data %s", err)
	}

	out, err := jsonPathFilter(data)
	if err != nil {
		return v, fmt.Errorf("failed running json path query %s", err)
	}

	ty, err := getImpliedType(out)

	if err != nil {
		return v, fmt.Errorf("failed json path imply type %s", err)
	}

	val, err := gocty.ToCtyValue(out, ty)

	if err != nil {
		return val, fmt.Errorf("to cty value failure - %s", err)
	}
	return val, err
}
