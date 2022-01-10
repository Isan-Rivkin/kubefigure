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
	"fmt"
	"io/ioutil"

	"github.com/isan-rivkin/kubefigure/sources/vault"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type VaultInput struct {
	AppRole    string `yaml:"approle"`
	SecretPath string `yaml:"path"`
	VaultAddr  string `yaml:"address"`
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
		vClient := vault.NewClientFromApprole(vltInput.VaultAddr, vltInput.AppRole)
		secret, err := vClient.Read(vltInput.SecretPath, "")

		if err != nil {
			panic(err)
		}

		fmt.Println("secret => ", secret)
	},
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
	terraformCmd.PersistentFlags().StringVar(&vltInput.VaultAddr, "address", "", "vault server address")
	terraformCmd.PersistentFlags().StringVar(&vltInput.AppRole, "approle", "", "approle id for auth method")
	terraformCmd.PersistentFlags().StringVar(&vltInput.SecretPath, "path", "", "path to secret")
	vaultCmd.PersistentFlags().StringVar(&fileInput, "input", "", "read input from file")
}
