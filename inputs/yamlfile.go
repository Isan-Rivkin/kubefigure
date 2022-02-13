package inputs

import (
	"io/ioutil"

	"github.com/isan-rivkin/kubefigure/sources/consuls"
	"github.com/isan-rivkin/kubefigure/sources/terraform"
	"github.com/isan-rivkin/kubefigure/sources/vault"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Source struct {
	SourceType string `yaml:"sourceType"`
	//Spec          interface{}            `yaml:"-"`
	TerraformConf *terraform.SourceInput `yaml:"terraform"`
	VaultConf     *vault.SourceInput     `yaml:"vault"`
	ConsulConf    *consuls.SourceInput   `yaml:"consul"`
}

type Input struct {
	Sources []*Source `yaml:"sources"`
}

func NewInputsFromYaml(path string) (*Input, error) {
	input := &Input{}
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, input)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}

	return input, nil
}
