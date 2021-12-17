package main

import (
	"github.com/isan-rivkin/kubefigure/sources/awsclient"
	"github.com/isan-rivkin/kubefigure/sources/terraform"
	log "github.com/sirupsen/logrus"
)

func test1() {
	bucket := "similarweb-terraform-remote-states"
	key := "private/terraform/route53-zones/int-similarweb-io/terraform.tfstate"
	region := "us-east-1"
	s3Api := awsclient.NewS3API(awsclient.CreateNewSession("", "", "", region))
	connector := terraform.NewS3RemoteStateConnector(s3Api, bucket, key, region)

	file, err := connector.DownloadAsStatefile()

	if err != nil {
		log.WithError(err).Error("failed getting s3")
	}

	outputs, err := terraform.GetOutputsFromStateFile(file)

	if err != nil {
		log.WithError(err).Error("failed getting s3")
		return
	}

	for _, o := range outputs {
		log.Info(o.Key + " : " + o.Value)
	}

}
func test2() {
	bucket := "similarweb-terraform-remote-states"
	key := "private/terraform/route53-zones/int-similarweb-io/terraform.tfstate"
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

func main() {
	test2()
}
