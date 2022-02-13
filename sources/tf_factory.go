package sources

import (
	"fmt"

	"github.com/isan-rivkin/kubefigure/sources/terraform"
)

type TerraformFactory struct {
}

func (tf *TerraformFactory) Create(input DataSourceConfig) (DataSource, error) {
	var connector terraform.RemoteStateConnector
	storageType := input.TerraformConf.StorageType

	switch storageType {
	case terraform.S3Storage:
		bucket := input.TerraformConf.S3Storage.Bucket
		region := input.TerraformConf.S3Storage.Region
		key := input.TerraformConf.Value.Key
		connector = terraform.NewDefaultS3RemoteStateConnector(bucket, key, region)
	default:
		return nil, fmt.Errorf("no such storage type for terraform %s", storageType)
	}

	ds := NewTerraformDataSource(connector, input)

	return ds, nil
}
