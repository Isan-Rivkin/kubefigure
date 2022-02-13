package vault

import (
	"gopkg.in/yaml.v3"
)

type AuthType string

const (
	ApproleAuthType AuthType = "approle"
)

type ApproleAuth struct {
	RoleID string `yaml:"approle"`
}

type SecretInfo struct {
	Path                  string `yaml:"path"`
	SecretValuePath       string `yaml:"valPath"`
	OptionalSecretVersion string `yaml:"secretVersion"`
}

type SourceInput struct {
	VaultAddr     string       `yaml:"address"`
	VaultAuthType AuthType     `yaml:"authType"`
	Approle       *ApproleAuth `yaml:"-"`
	Secret        SecretInfo   `yaml:"value"`
}

func (s *SourceInput) UnmarshalYAML(n *yaml.Node) error {
	var err error
	type S SourceInput

	type T struct {
		*S      `yaml:",inline"`
		Approle yaml.Node `yaml:"auth"`
	}
	obj := &T{S: (*S)(s)}

	if err := n.Decode(obj); err != nil {
		return err
	}

	switch s.VaultAuthType {
	case ApproleAuthType:
		s.Approle = new(ApproleAuth)
		obj.Approle.Decode(s.Approle)
	default:
		panic("kind unknown")
	}
	return err
}
