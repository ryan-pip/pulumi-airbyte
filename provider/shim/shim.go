package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/ryan-pip/pulumi-airbyte/provider/pkg/version"
	"github.com/ryan-pip/terraform-provider-airbyte/internal/provider"
)

func NewProvider() tfpf.Provider {
	return provider.New(version.Version)()
}
