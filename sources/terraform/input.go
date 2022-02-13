package terraform

import "gopkg.in/yaml.v3"

type StateStorageType string

const (
	S3Storage StateStorageType = "s3"
)

type S3StateStorage struct {
	Bucket string `yaml:"bucket"`
	Region string `yaml:"region"`
}

type StateValueInput struct {
	Key            string `yaml:"key"`
	StateValuePath string `yaml:"stateValuePath"`
}

type SourceInput struct {
	StorageType StateStorageType `yaml:"storageType"`
	S3Storage   *S3StateStorage  `yaml:"-"`
	Value       StateValueInput  `yaml:"value"`
}

func (s *SourceInput) UnmarshalYAML(n *yaml.Node) error {
	var err error
	type S SourceInput

	type T struct {
		*S        `yaml:",inline"`
		S3Storage yaml.Node `yaml:"storage"`
	}
	obj := &T{S: (*S)(s)}

	if err := n.Decode(obj); err != nil {
		return err
	}

	switch s.StorageType {
	case S3Storage:
		s.S3Storage = new(S3StateStorage)
		obj.S3Storage.Decode(s.S3Storage)
	default:
		panic("kind unknown")
	}
	return err
}
