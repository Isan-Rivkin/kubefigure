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

	"github.com/isan-rivkin/kubefigure/common"
	"github.com/isan-rivkin/kubefigure/inputs"
	"github.com/isan-rivkin/kubefigure/sources"
	"github.com/isan-rivkin/kubefigure/sources/consuls"
	"github.com/isan-rivkin/kubefigure/sources/terraform"
	"github.com/isan-rivkin/kubefigure/sources/vault"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var sourceToInput = map[sources.SourceType]func(s *inputs.Source) sources.DataSourceConfig{
	sources.VaultSource:     inputToVaultDataSourceInput,
	sources.TerraformSource: inputToTerraformDataSourceInput,
	sources.ConsulSource:    inputToConsulDataSourceInput,
}

// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "fetch input from data sources",
	Long: `use subcommand or directly
		go run main.go input --input=generic_input.yaml
	`,
	Run: func(cmd *cobra.Command, args []string) {
		sourcesController := sources.NewDataSourceController()
		fmt.Println("input called ", fileInput)
		inputRes, err := inputs.NewInputsFromYaml(fileInput)

		if err != nil {
			fmt.Errorf("failed parsing yaml input - %v", err)
			return
		}

		for _, s := range inputRes.Sources {
			fmt.Println(s.SourceType)
			if inputBuilder, found := sourceToInput[sources.SourceType(s.SourceType)]; found {
				dsInput := inputBuilder(s)
				ds, err := sourcesController.GetDataSource(sources.SourceType(s.SourceType), dsInput)
				if err != nil {
					log.WithError(err).Error("failed for datasouce")
					continue
				}
				payload, err := ds.Get()
				if err != nil {
					log.WithError(err).Error("failed fetching payload for ds")
					continue
				}
				val, err := payload.Value()

				if err != nil {
					log.WithError(err).Error("failed getting val from payload")
					continue
				}
				strVal := common.ConvertValueToString(val)
				log.WithFields(log.Fields{
					"source": s.SourceType,
				}).Info(strVal)
			}
		}
	},
}

func inputToVaultDataSourceInput(s *inputs.Source) sources.DataSourceConfig {
	return sources.DataSourceConfig{
		VaultConf: &vault.SourceInput{
			VaultAddr:     s.VaultConf.VaultAddr,
			VaultAuthType: s.VaultConf.VaultAuthType,
			Approle: &vault.ApproleAuth{
				RoleID: s.VaultConf.Approle.RoleID,
			},
			Secret: vault.SecretInfo{
				Path:            s.VaultConf.Secret.Path,
				SecretValuePath: s.VaultConf.Secret.SecretValuePath,
			},
		},
	}
}

func inputToTerraformDataSourceInput(s *inputs.Source) sources.DataSourceConfig {
	return sources.DataSourceConfig{
		TerraformConf: &terraform.SourceInput{
			StorageType: s.TerraformConf.StorageType,
			S3Storage: &terraform.S3StateStorage{
				Bucket: s.TerraformConf.S3Storage.Bucket,
				Region: s.TerraformConf.S3Storage.Region,
			},
			Value: terraform.StateValueInput{
				Key:            s.TerraformConf.Value.Key,
				StateValuePath: s.TerraformConf.Value.StateValuePath,
			},
		},
	}
}

func inputToConsulDataSourceInput(s *inputs.Source) sources.DataSourceConfig {
	return sources.DataSourceConfig{
		ConsulConf: &consuls.SourceInput{
			ConsulAuthType: s.ConsulConf.ConsulAuthType,
			HttpAuth: &consuls.HTTPAuth{
				Address: s.ConsulConf.HttpAuth.Address,
				Port:    s.ConsulConf.HttpAuth.Port,
			},
			KV: consuls.KVInfo{
				Path:          s.ConsulConf.KV.Path,
				ValueJsonPath: s.ConsulConf.KV.ValueJsonPath,
				Opts: &consuls.QueryOpts{
					Datacenter: s.ConsulConf.KV.Opts.Datacenter,
				},
			},
		},
	}
}

func init() {
	rootCmd.AddCommand(inputCmd)
	inputCmd.PersistentFlags().StringVar(&fileInput, "input", "", "read input from file")
}
