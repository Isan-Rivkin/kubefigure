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
	"github.com/isan-rivkin/kubefigure/sources/terraform"
	log "github.com/sirupsen/logrus"
)

func test1(bucket, key, region string) {

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
func test2(bucket, key, region string) {

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

func test3(bucket, key, region string) {

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

func test4(bucket, key, region string) {

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
	bucket, key, region := "", "", ""
	//test1(bucket, key, region )
	//test2(bucket, key, region )
	test3(bucket, key, region)
	//test4(bucket, key, region )
}
