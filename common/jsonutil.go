package common

import (
	"encoding/json"
	"fmt"

	"github.com/yalp/jsonpath"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

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

	ty, err := GetImpliedType(out)

	if err != nil {
		return v, fmt.Errorf("failed json path imply type %s", err)
	}

	val, err := gocty.ToCtyValue(out, ty)

	if err != nil {
		return val, fmt.Errorf("to cty value failure - %s", err)
	}
	return val, err
}
