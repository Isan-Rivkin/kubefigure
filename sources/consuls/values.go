package consuls

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

type ConsulPayload struct {
	Pair  *api.KVPair
	AsVal cty.Value
}

// consul pair value is always slice of bytes and therefore the implied type from cty perspective if always
// a list of numbers. due to that it always creates it instead of json or string
// tldr; value in consul assumed json or string
func (c *BaseClient) consulBytesToVal(value []byte) (cty.Value, error) {

	val := cty.Value{}
	var err error

	ty, _ := ctyjson.ImpliedType(value)

	if ty.IsObjectType() {
		val, err = ctyjson.Unmarshal(value, ty)
	} else {
		val, err = gocty.ToCtyValue(string(value), cty.String)
	}

	return val, err
}
func (c *BaseClient) ReadKVAsVal(path string, opts *api.QueryOptions) (*ConsulPayload, error) {
	pair, err := c.ReadKV(path, opts)

	if err != nil {
		return nil, err
	}

	if pair == nil {
		return nil, fmt.Errorf("key not found %s", path)
	}

	payload := &ConsulPayload{
		Pair: pair,
	}

	val, err := c.consulBytesToVal(pair.Value)

	if err != nil {
		return nil, fmt.Errorf("failed convering consul bytes result %s - %v", path, err)
	}

	payload.AsVal = val

	return payload, nil
}
