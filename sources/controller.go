package sources

import "fmt"

type SourceType string

type Controller interface {
	GetDataSource(ty SourceType, input DataSourceConfig) (DataSource, error)
}

type DataSourceController struct {
	factories map[SourceType]DataSourceFactory
}

func NewDataSourceController() Controller {
	return &DataSourceController{
		factories: map[SourceType]DataSourceFactory{
			TerraformSource: &TerraformFactory{},
			VaultSource:     &VaultFactory{},
			ConsulSource:    &ConsulFactory{},
		},
	}
}

func (c *DataSourceController) GetDataSource(ty SourceType, input DataSourceConfig) (DataSource, error) {
	fact, found := c.factories[ty]

	if !found {
		return nil, fmt.Errorf("no such data source factory %s", ty)
	}

	ds, err := fact.Create(input)

	if err != nil {
		return nil, fmt.Errorf("factory failed creating datasource %s", err)
	}

	return ds, nil
}
