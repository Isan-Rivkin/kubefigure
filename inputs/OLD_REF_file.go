package inputs

// import (
// 	"io/ioutil"

// 	"github.com/isan-rivkin/kubefigure/sources"
// 	"github.com/isan-rivkin/kubefigure/sources/consuls"
// 	"github.com/isan-rivkin/kubefigure/sources/terraform"
// 	"github.com/isan-rivkin/kubefigure/sources/vault"
// 	log "github.com/sirupsen/logrus"
// 	"gopkg.in/yaml.v3"
// )

// type Source struct {
// 	SourceType    string                 `yaml:"sourceType"`
// 	Spec          interface{}            `yaml:"-"`
// 	TerraformConf *terraform.SourceInput `yaml:"-"`
// 	VaultConf     *vault.SourceInput     `yaml:"-"`
// 	ConsulConf    *consuls.SourceInput   `yaml:"-"`
// }

// func (s *Source) UnmarshalYAML(n *yaml.Node) error {
// 	type S Source
// 	type T struct {
// 		*S   `yaml:",inline"`
// 		Spec yaml.Node `yaml:"spec"`
// 	}

// 	obj := &T{S: (*S)(s)}

// 	if err := n.Decode(obj); err != nil {
// 		return err
// 	}

// 	switch sources.SourceType(s.SourceType) {
// 	case sources.VaultSource:
// 		s.Spec = new(vault.SourceInput)
// 	case sources.TerraformSource:
// 		s.Spec = new(terraform.SourceInput)
// 	case sources.ConsulSource:
// 		s.Spec = new(consuls.SourceInput)
// 	default:
// 		panic("kind unknown")
// 	}
// 	return obj.Spec.Decode(s.Spec)
// }

// type Input struct {
// 	Sources []*Source `yaml:"sources"`
// }

// func NewInputsFromYaml(path string) (*Input, error) {
// 	input := &Input{}
// 	yamlFile, err := ioutil.ReadFile(path)

// 	if err != nil {
// 		log.Printf("yamlFile.Get err   #%v ", err)
// 		return nil, err
// 	}

// 	err = yaml.Unmarshal(yamlFile, input)

// 	if err != nil {
// 		log.Fatalf("Unmarshal: %v", err)
// 		return nil, err
// 	}

// 	return input, nil
// }
