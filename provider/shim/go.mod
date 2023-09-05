module github.com/airbytehq/terraform-provider-airbyte/shim

go 1.21

toolchain go1.21.0

require (
	github.com/hashicorp/terraform-plugin-framework v1.3.5
	github.com/ryan-pip/pulumi-airbyte/provider v0.0.0-20230904223513-fe62fe13c1a5
	github.com/ryan-pip/terraform-provider-airbyte v0.3.7
)

require (
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.10.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.18.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
)

replace github.com/ryan-pip/terraform-provider-airbyte => github.com/ryan-pip/terraform-provider-airbyte v0.3.7
