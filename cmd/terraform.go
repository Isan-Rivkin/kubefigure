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

	"github.com/isan-rivkin/kubefigure/sources/terraform"
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
	},
}

func test1() {
	bucket := "similarweb-terraform-remote-states"
	//key := "private/terraform/route53-zones/int-similarweb-io/terraform.tfstate"
	key := "private/terraform/nomad_agent/us-east-1_production_datacollection/terraform.tfstate"
	region := "us-east-1"

	connector := terraform.NewDefaultS3RemoteStateConnector(bucket, key, region)

	tfClient := terraform.NewTerraformClient()
	state, err := tfClient.GetState(connector)

	if err != nil {
		log.Error("failed getting state")
		panic(err)
	}

	outputs, err := state.OutputsStatus()

	if err != nil {
		log.Error("failed getting state outputs")
		panic(err)
	}

	for _, o := range outputs {
		log.Info(o.Key + " : " + o.Value)
	}

	vals := terraform.GetOutputsValuesFromStateFile(state.File)
	log.Info("@@@@@@@@@@")
	for key, val := range vals {
		log.Info("key " + key)
		log.Info(val.Value.Type().GoString())
		log.Info(val.Value)
	}

}
func test2() {
	bucket := "similarweb-terraform-remote-states"
	//key := "private/terraform/route53-zones/int-similarweb-io/terraform.tfstate"
	key := "private/terraform/nomad_agent/us-east-1_production_datacollection/terraform.tfstate"
	region := "us-east-1"

	connector := terraform.NewDefaultS3RemoteStateConnector(bucket, key, region)

	tfClient := terraform.NewTerraformClient()
	state, err := tfClient.GetState(connector)

	if err != nil {
		log.Error("failed getting state")
		panic(err)
	}

	outputs, err := state.OutputsStatus()

	if err != nil {
		log.Error("failed getting state outputs")
		panic(err)
	}

	for _, o := range outputs {
		log.Info(o.Key + " : " + o.Value)
	}

	vals := terraform.GetOutputsValuesFromStateFile(state.File)
	log.Info("@@@@@@@@@@")
	for key, val := range vals {
		log.Info("key " + key)
		log.Info(val.Value.Type().GoString())
		log.Info(val.Value)
	}

	log.Info("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@!!!!")
	for n, m := range state.File.State.Modules {
		log.Info(n)
		log.Info(m.Addr)
		for k, r := range m.Resources {
			log.WithFields(log.Fields{
				"resource":      k,
				"resource_addr": r.Addr,
			}).Info("A resource!")

			for i, d := range r.Instances {
				log.WithFields(log.Fields{
					"instances": i,
					"data":      string(d.Current.AttrsJSON),
				}).Info("An instances!")
				for _, dep := range d.Current.Dependencies {
					log.WithFields(log.Fields{
						"addr": dep.String(),
					}).Info("Depends on!")
				}
			}
		}
	}
}

func test3() {
	bucket := "similarweb-terraform-remote-states"
	//key := "private/terraform/route53-zones/int-similarweb-io/terraform.tfstate"
	key := "private/terraform/nomad_agent/us-east-1_production_datacollection/terraform.tfstate"
	region := "us-east-1"

	connector := terraform.NewDefaultS3RemoteStateConnector(bucket, key, region)

	tfClient := terraform.NewTerraformClient()
	state, err := tfClient.GetState(connector)

	if err != nil {
		log.Error("failed getting state")
		panic(err)
	}

	outputs, err := state.OutputsStatus()

	if err != nil {
		log.Error("failed getting state outputs")
		panic(err)
	}

	for _, o := range outputs {
		log.Info(o.Key + " : " + o.Value)
	}

	vals := terraform.GetOutputsValuesFromStateFile(state.File)
	log.Info("@@@@@@@@@@")
	for key, val := range vals {
		log.Info("key " + key)
		log.Info(val.Value.Type().GoString())
		log.Info(val.Value)
	}

	log.Info("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@!!!!")
	state.File.State.RootModule()
	for n, m := range state.File.State.Modules {
		log.Info(n)
		log.Info(m.Addr)
		if n == "" {
			log.Info("EMPRTY ADDD!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			for resName := range m.Resources {
				log.Info(resName)
			}
			log.Info("END EMPRTY ADDD!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		}
		// for k, r := range m.Resources {
		// 	log.WithFields(log.Fields{
		// 		"resource":      k,
		// 		"resource_addr": r.Addr,
		// 	}).Info("A resource!")

		// 	for i, d := range r.Instances {
		// 		log.WithFields(log.Fields{
		// 			"instances": i,
		// 			"data":      string(d.Current.AttrsJSON),
		// 		}).Info("An instances!")
		// 		for _, dep := range d.Current.Dependencies {
		// 			log.WithFields(log.Fields{
		// 				"addr": dep.String(),
		// 			}).Info("Depends on!")
		// 		}
		// 	}
		// }
	}
}

func test4() {
	bucket := "similarweb-terraform-remote-states"
	//key := "private/terraform/route53-zones/int-similarweb-io/terraform.tfstate"
	key := "private/terraform/nomad_agent/us-east-1_production_datacollection/terraform.tfstate"
	region := "us-east-1"

	connector := terraform.NewDefaultS3RemoteStateConnector(bucket, key, region)

	tfClient := terraform.NewTerraformClient()
	state, err := tfClient.GetState(connector)

	if err != nil {
		log.Error("failed getting state")
		panic(err)
	}

	for _, addr := range state.ListModulesAddrs() {
		log.WithField("address", addr).Info("a module!")
	}

	for _, addr := range state.ListAllResources() {
		log.WithField("address", addr).Info("a resources!")
	}

}
func runTestsTF() {
	//test1()
	//test2()
	test3()
	//test4()
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
