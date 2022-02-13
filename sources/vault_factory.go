package sources

import (
	"fmt"

	"github.com/isan-rivkin/kubefigure/sources/vault"
)

type VaultFactory struct {
}

func (vf *VaultFactory) Create(input DataSourceConfig) (DataSource, error) {
	var client vault.Client
	authType := input.VaultConf.VaultAuthType

	switch authType {
	case vault.ApproleAuthType:
		client = vault.NewClientFromApprole(input.VaultConf.VaultAddr, input.VaultConf.Approle.RoleID)
	default:
		return nil, fmt.Errorf("no such auth method vault %s", authType)
	}

	c := NewVaultDataSource(client, input)

	return c, nil
}
