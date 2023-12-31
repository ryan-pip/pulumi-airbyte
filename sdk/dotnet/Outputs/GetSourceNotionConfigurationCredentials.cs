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
    public sealed class GetSourceNotionConfigurationCredentials
    {
        public readonly Outputs.GetSourceNotionConfigurationCredentialsSourceNotionAuthenticateUsingAccessToken SourceNotionAuthenticateUsingAccessToken;
        public readonly Outputs.GetSourceNotionConfigurationCredentialsSourceNotionAuthenticateUsingOAuth20 SourceNotionAuthenticateUsingOAuth20;
        public readonly Outputs.GetSourceNotionConfigurationCredentialsSourceNotionUpdateAuthenticateUsingAccessToken SourceNotionUpdateAuthenticateUsingAccessToken;
        public readonly Outputs.GetSourceNotionConfigurationCredentialsSourceNotionUpdateAuthenticateUsingOAuth20 SourceNotionUpdateAuthenticateUsingOAuth20;

        [OutputConstructor]
        private GetSourceNotionConfigurationCredentials(
            Outputs.GetSourceNotionConfigurationCredentialsSourceNotionAuthenticateUsingAccessToken sourceNotionAuthenticateUsingAccessToken,

            Outputs.GetSourceNotionConfigurationCredentialsSourceNotionAuthenticateUsingOAuth20 sourceNotionAuthenticateUsingOAuth20,

            Outputs.GetSourceNotionConfigurationCredentialsSourceNotionUpdateAuthenticateUsingAccessToken sourceNotionUpdateAuthenticateUsingAccessToken,

            Outputs.GetSourceNotionConfigurationCredentialsSourceNotionUpdateAuthenticateUsingOAuth20 sourceNotionUpdateAuthenticateUsingOAuth20)
        {
            SourceNotionAuthenticateUsingAccessToken = sourceNotionAuthenticateUsingAccessToken;
            SourceNotionAuthenticateUsingOAuth20 = sourceNotionAuthenticateUsingOAuth20;
            SourceNotionUpdateAuthenticateUsingAccessToken = sourceNotionUpdateAuthenticateUsingAccessToken;
            SourceNotionUpdateAuthenticateUsingOAuth20 = sourceNotionUpdateAuthenticateUsingOAuth20;
        }
    }
}
