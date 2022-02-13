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
	"github.com/isan-rivkin/kubefigure/sources/vault"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type VaultInput struct {
	AppRole     string `yaml:"approle"`
	SecretPath  string `yaml:"path"`
	ValJsonPath string `yaml:"valpath"`
	VaultAddr   string `yaml:"address"`
}

var vltInput *VaultInput = &VaultInput{}

// vaultCmd represents the vault command
var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "read secrets from Vault",
	Long: `
$input vault --addr=vault.com --approle=123 --path=path/to/secret
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vault called")
		renderVaultInputFile()
		getVaultWithController()
		// vClient := vault.NewClientFromApprole(vltInput.VaultAddr, vltInput.AppRole)
		// secret, err := vClient.ReadAsVal(vltInput.SecretPath, "")

		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Println("secret => ", secret.JsonVal.Value.AsValueMap())
		// val, err := common.FindValInJson(vltInput.ValJsonPath, secret.JsonBytes)
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Println("json path secret val =>  ", val.AsString())
	},
}

func getVaultWithController() {
	sourcesController := sources.NewDataSourceController()
	ds, err := sourcesController.GetDataSource(sources.VaultSource, sources.DataSourceConfig{
		VaultConf: &vault.SourceInput{
			VaultAddr:     vltInput.VaultAddr,
			VaultAuthType: vault.ApproleAuthType,
			Approle: &vault.ApproleAuth{
				RoleID: vltInput.AppRole,
			},
			Secret: vault.SecretInfo{
				Path:            vltInput.SecretPath,
				SecretValuePath: vltInput.ValJsonPath,
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

}
func renderVaultInputFile() error {
	yamlFile, err := ioutil.ReadFile(fileInput)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, vltInput)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	return nil
}

func init() {
	inputCmd.AddCommand(vaultCmd)
	vaultCmd.PersistentFlags().StringVar(&vltInput.VaultAddr, "address", "", "vault server address")
	vaultCmd.PersistentFlags().StringVar(&vltInput.AppRole, "approle", "", "approle id for auth method")
	vaultCmd.PersistentFlags().StringVar(&vltInput.SecretPath, "path", "", "path to secret")
	vaultCmd.PersistentFlags().StringVar(&vltInput.ValJsonPath, "jsonpath", "", "path to value inside json secret based on json path")
	vaultCmd.PersistentFlags().StringVar(&fileInput, "input", "", "read input from file")
}
