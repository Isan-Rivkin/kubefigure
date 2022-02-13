package consuls

import "gopkg.in/yaml.v3"

type AuthType string

const (
	HTTPAuthType AuthType = "http"
)

type HTTPAuth struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type QueryOpts struct {
	Namespace  string `yaml:"namespace"`
	Partition  string `yaml:"partition"`
	Datacenter string `yaml:"datacenter"`
}
type KVInfo struct {
	Path          string     `yaml:"key"`
	ValueJsonPath string     `yaml:"valPath"`
	Opts          *QueryOpts `yaml:"options"`
}

type SourceInput struct {
	ConsulAuthType AuthType  `yaml:"authType"`
	HttpAuth       *HTTPAuth `yaml:"-"`
	KV             KVInfo    `yaml:"value"`
}

func (s *SourceInput) UnmarshalYAML(n *yaml.Node) error {
	var err error
	type S SourceInput

	type T struct {
		*S       `yaml:",inline"`
		HttpAuth yaml.Node `yaml:"auth"`
	}
	obj := &T{S: (*S)(s)}

	if err := n.Decode(obj); err != nil {
		return err
	}

	switch s.ConsulAuthType {
	case HTTPAuthType:
		s.HttpAuth = new(HTTPAuth)
		obj.HttpAuth.Decode(s.HttpAuth)
	default:
		panic("kind unknown")
	}
	return err
}
