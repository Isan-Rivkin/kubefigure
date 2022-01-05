package main

import (
	"github.com/isan-rivkin/kubefigure/sources/terraform"
	log "github.com/sirupsen/logrus"
)

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
func main() {
	//test1()
	//test2()
	test3()
	//test4()
}
