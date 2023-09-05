package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/ryan-pip/terraform-provider-airbyte/internal/provider"
)

const (
	VersionDev  = "dev"
	VersionTest = "test"
)

func NewProvider() tfpf.Provider {
	return provider.New(VersionDev)
}
