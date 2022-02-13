/*
Copyright © 2022 NAME HERE isanrivkin@gmail.com

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
	"fmt"
	"io/ioutil"

	"github.com/isan-rivkin/kubefigure/common"
	"github.com/isan-rivkin/kubefigure/sources"
	"github.com/isan-rivkin/kubefigure/sources/consuls"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type ConsulInput struct {
	Addr       string `yaml:"address"`
	Port       int    `yaml:"port"`
	Key        string `yaml:"key"`
	ValuePath  string `yaml:"valuePath"`
	Namespace  string `yaml:"namespace"`
	Partition  string `yaml:"partition"`
	Datacenter string `yaml:"datacenter"`
}

var (
	consulInput *ConsulInput = &ConsulInput{}
)

// consulCmd represents the consul command
var consulCmd = &cobra.Command{
	Use:   "consul",
	Short: "get kv from consul",
	Long:  `go run main.go input consul --input=consul_input.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consul called")

		if err := renderConsulInputFile(); err != nil {
			panic(fmt.Errorf("failed rending consul input - %v", err))
		}

		sourcesController := sources.NewDataSourceController()
		ds, err := sourcesController.GetDataSource(sources.ConsulSource, sources.DataSourceConfig{
			ConsulConf: &consuls.SourceInput{
				ConsulAuthType: consuls.HTTPAuthType,
				HttpAuth: &consuls.HTTPAuth{
					Address: consulInput.Addr,
					Port:    consulInput.Port,
				},
				KV: consuls.KVInfo{
					Path:          consulInput.Key,
					ValueJsonPath: consulInput.ValuePath,
					Opts: &consuls.QueryOpts{
						Datacenter: consulInput.Datacenter,
					},
				},
			},
		})

		if err != nil {
			panic(err)
		}

		payload, err := ds.Get()

		if err != nil {
			panic(err)
		}

		val, err := payload.Value()
		if err != nil {
			panic(err)
		}

		log.Info(common.ConvertValueToString(val))

	},
}

func renderConsulInputFile() error {
	yamlFile, err := ioutil.ReadFile(fileInput)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, consulInput)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	return nil
}

func init() {
	inputCmd.AddCommand(consulCmd)
	consulCmd.PersistentFlags().StringVar(&fileInput, "input", "", "read input from file")
}
