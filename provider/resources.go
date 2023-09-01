// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package airbyte

import (
	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"
	"unicode"

    // "github.com/hashicorp/terraform-plugin-framework/provider"
    pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
    "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
    "github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
    "github.com/ryan-pip/terraform-provider-airbyte/shim"
)

// all of the airbyte token components used below.
const (
	airbytePkg = "airbyte"
	airbyteMod = "index"
)

// airbyteMember manufactures a type token for the airbyte package and the given module and type.
func airbyteMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(airbytePkg + ":" + mod + ":" + mem)
}

// airbyteType manufactures a type token for the airbyte package and the given module and type.
func airbyteType(mod string, typ string) tokens.Type {
	return tokens.Type(airbyteMember(mod, typ))
}

func airbyteDataSource(mod string, res string) tokens.ModuleMember {
    fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
    return airbyteMember(mod+"/"+fn, res)
}

// airbyteResource manufactures a standard resource token given a module and resource name.  It automatically uses the
// airbyte package and names the file by simply lower casing the resource's first character.
func airbyteResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return airbyteType(mod+"/"+fn, res)
}
//go:embed cmd/pulumi-resource-airbyte/bridge-metadata.json
var bridgeMetadata []byte

func Provider() tfbridge.ProviderInfo {
    info := tfbridge.ProviderInfo{
        P:       pf.ShimProvider(shim.NewProvider()),
        Name:    "airbyte",
        Version: "1.2.3",
        // GitHubOrg: "ryan-pip",
        Resources: map[string]*tfbridge.ResourceInfo{
            "airbyte_connection":              {Tok: airbyteResource(airbyteMod, "Connection")},
            "airbyte_workspace":               {Tok: airbyteResource(airbyteMod, "Workspace")},
            "airbyte_destination_snowflake":   {Tok: airbyteResource(airbyteMod, "DestinationSnowflake")},
            "airbyte_source_salesforce":       {Tok: airbyteResource(airbyteMod, "SourceSalesforce")},
            "airbyte_source_mixpanel":         {Tok: airbyteResource(airbyteMod, "SourceMixpanel")},
        },
        DataSources: map[string]*tfbridge.DataSourceInfo{
            "airbyte_connection":              {Tok: airbyteDataSource(airbyteMod, "getConnection")},
            "airbyte_workspace":               {Tok: airbyteDataSource(airbyteMod, "getWorkspace")},
            "airbyte_destination_snowflake":   {Tok: airbyteDataSource(airbyteMod, "getDestinationSnowflake")},
            "airbyte_source_salesforce":       {Tok: airbyteDataSource(airbyteMod, "getSourceSalesforce")},
            "airbyte_source_mixpanel":         {Tok: airbyteDataSource(airbyteMod, "getSourceMixpanel")},
        },
        MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),
    }
    return info
}