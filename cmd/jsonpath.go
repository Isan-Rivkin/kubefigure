/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/yalp/jsonpath"
)

var (
	jsonPath string
)

// jsonpathCmd represents the jsonpath command
var jsonpathCmd = &cobra.Command{
	Use:   "jsonpath",
	Short: "test jsonpath",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(jsonpathCmd)
	jsonpathCmd.PersistentFlags().StringVar(&jsonPath, "file", "", "path to a json file")
}
