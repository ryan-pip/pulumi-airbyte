package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/ryan-pip/terraform-provider-airbyte/internal/provider"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"
)


func NewProvider() tfpf.Provider {
	return provider.New(version)()
}