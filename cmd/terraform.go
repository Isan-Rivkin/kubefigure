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
	"os"
	"strings"

	"github.com/isan-rivkin/kubefigure/common"
	"github.com/isan-rivkin/kubefigure/sources"
	"github.com/isan-rivkin/kubefigure/sources/terraform"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	fileInput string
)

type TerraformInput struct {
	BackendType             string `yaml:"type"`
	BackendName             string `yaml:"name"`
	BackendKey              string `yaml:"key"`
	RemoteRegion            string `yaml:"region"`
	StateValuePath          string `yaml:"stateValuePath"`
	StatusOnly              bool   `yaml:"statusOnly"`
	ListAvailableValuePaths bool   `yaml:"listAvailableValuePaths"`
}

var tfInput *TerraformInput = &TerraformInput{}

// terraformCmd represents the terraform command
var terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "fetch terraform inputs",
	Long: `
$input terraform --backed-type=s3 --backend-name=bucket-name --backend-key=some/key/state --backend-region=us-east-1 --state-value-path=outputs.alb_sg --status-only
$input terraform --input=input.yaml
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("executing terraform inputs")
		if err := initTerraformInput(); err != nil {
			log.WithError(err).Error("failed rendering input")
			os.Exit(1)
		}

		if tfInput.BackendType != "s3" {
			panic("only s3 supported")
		}

		if !strings.HasPrefix(tfInput.StateValuePath, "$.outputs.") {
			panic("for now, only outputs path")
		}

		runWithController()
	},
}

func runWithController() {
	sourcesController := sources.NewDataSourceController()
	ds, err := sourcesController.GetDataSource(sources.TerraformSource, sources.DataSourceConfig{
		TerraformConf: &terraform.SourceInput{
			StorageType: terraform.S3Storage,
			S3Storage: &terraform.S3StateStorage{
				Bucket: tfInput.BackendName,
				Region: tfInput.RemoteRegion,
			},
			Value: terraform.StateValueInput{
				Key:            tfInput.BackendKey,
				StateValuePath: tfInput.StateValuePath,
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

	strVal := common.ConvertValueToString(val)

	fmt.Println(strVal)
}
func renderInputFile() error {
	yamlFile, err := ioutil.ReadFile(fileInput)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, tfInput)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	return nil
}
func initTerraformInput() error {
	if fileInput != "" {
		return renderInputFile()
	}
	return nil
}
func init() {
	inputCmd.AddCommand(terraformCmd)
	terraformCmd.PersistentFlags().StringVar(&tfInput.BackendType, "backend-type", "", "--backend-type s3")
	terraformCmd.PersistentFlags().StringVar(&tfInput.BackendName, "backend-name", "", "--backend-name my-s3-bucket")
	terraformCmd.PersistentFlags().StringVar(&tfInput.BackendKey, "backend-key", "", "--backend-key key/to/state/in/bucket")
	terraformCmd.PersistentFlags().StringVar(&tfInput.RemoteRegion, "remote-region", "us-east-1", "--remote-region us-east-1")
	terraformCmd.PersistentFlags().StringVar(&tfInput.StateValuePath, "state-value-path", "", "--state-value-path outputs.elb_sg")
	terraformCmd.PersistentFlags().StringVar(&fileInput, "input", "", "read input from file")
	terraformCmd.PersistentFlags().BoolVar(&tfInput.StatusOnly, "status-only", true, "--status-only will print the remote value info")
	terraformCmd.PersistentFlags().BoolVar(&tfInput.ListAvailableValuePaths, "list-value-paths", true, "--list-value-paths will list available paths in the state")
}
