package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/isan-rivkin/kubefigure/cmd"
	"github.com/yalp/jsonpath"
)

func foo() {
	cmd.Execute()
}
func main() {

	jsonPath := "/Users/isanrivkin/workspace/terraform/states/tapi-ingress-endpoints/connectweb24-com_cross-region_production/input.json"
	raw, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		panic(err)
	}

	helloFilter, err := jsonpath.Prepare("$.source.module..source")
	if err != nil {
		panic(err)
	}

	var data interface{}
	if err = json.Unmarshal(raw, &data); err != nil {
		panic(err)
	}

	out, err := helloFilter(data)

	if err != nil {
		panic(err)
	}

	fmt.Print(out)
}
