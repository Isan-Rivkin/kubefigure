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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type TerraformInput struct {
	BackendType    string
	BackendName    string
	BackendKey     string
	RemoteRegion   string
	StateValuePath string
	StatusOnly     bool
}

var tfInput *TerraformInput = &TerraformInput{}

// terraformCmd represents the terraform command
var terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "fetch terraform inputs",
	Long:  `$input terraform --backed-type=s3 --backend-name=bucket-name --backend-key=some/key/state --backend-region=us-east-1 --state-value-path=outputs.alb_sg --status-only`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("terraform called")
		log.Info("called")
	},
}

func init() {
	inputCmd.AddCommand(terraformCmd)
	// input := struct {
	// 	BackendType    string
	// 	BucketName     string
	// 	BucketKey      string
	// 	StateRegion    string
	// 	StateValuePath string
	// 	StatusOnly     bool
	// }

	terraformCmd.PersistentFlags().StringVar(&tfInput.BackendType, "backend-type", "", "--backend-type s3")
	terraformCmd.PersistentFlags().StringVar(&tfInput.BackendName, "backend-name", "", "--backend-name my-s3-bucket")
	terraformCmd.PersistentFlags().StringVar(&tfInput.BackendKey, "backend-key", "", "--backend-key key/to/state/in/bucket")
	terraformCmd.PersistentFlags().StringVar(&tfInput.RemoteRegion, "remote-region", "us-east-1", "--remote-region us-east-1")
	terraformCmd.PersistentFlags().StringVar(&tfInput.StateValuePath, "state-value-path", "", "--state-value-path outputs.elb_sg")
	terraformCmd.PersistentFlags().BoolVar(&tfInput.StatusOnly, "status-only", true, "--status-only will print the remote value info")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// terraformCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
