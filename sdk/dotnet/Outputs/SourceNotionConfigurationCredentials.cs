// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class SourceNotionConfigurationCredentials
    {
        public readonly Outputs.SourceNotionConfigurationCredentialsSourceNotionAuthenticateUsingAccessToken? SourceNotionAuthenticateUsingAccessToken;
        public readonly Outputs.SourceNotionConfigurationCredentialsSourceNotionAuthenticateUsingOAuth20? SourceNotionAuthenticateUsingOAuth20;
        public readonly Outputs.SourceNotionConfigurationCredentialsSourceNotionUpdateAuthenticateUsingAccessToken? SourceNotionUpdateAuthenticateUsingAccessToken;
        public readonly Outputs.SourceNotionConfigurationCredentialsSourceNotionUpdateAuthenticateUsingOAuth20? SourceNotionUpdateAuthenticateUsingOAuth20;

        [OutputConstructor]
        private SourceNotionConfigurationCredentials(
            Outputs.SourceNotionConfigurationCredentialsSourceNotionAuthenticateUsingAccessToken? sourceNotionAuthenticateUsingAccessToken,

            Outputs.SourceNotionConfigurationCredentialsSourceNotionAuthenticateUsingOAuth20? sourceNotionAuthenticateUsingOAuth20,

            Outputs.SourceNotionConfigurationCredentialsSourceNotionUpdateAuthenticateUsingAccessToken? sourceNotionUpdateAuthenticateUsingAccessToken,

            Outputs.SourceNotionConfigurationCredentialsSourceNotionUpdateAuthenticateUsingOAuth20? sourceNotionUpdateAuthenticateUsingOAuth20)
        {
            SourceNotionAuthenticateUsingAccessToken = sourceNotionAuthenticateUsingAccessToken;
            SourceNotionAuthenticateUsingOAuth20 = sourceNotionAuthenticateUsingOAuth20;
            SourceNotionUpdateAuthenticateUsingAccessToken = sourceNotionUpdateAuthenticateUsingAccessToken;
            SourceNotionUpdateAuthenticateUsingOAuth20 = sourceNotionUpdateAuthenticateUsingOAuth20;
        }
    }
}