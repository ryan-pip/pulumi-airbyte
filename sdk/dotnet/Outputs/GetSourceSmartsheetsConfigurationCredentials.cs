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
    public sealed class GetSourceSmartsheetsConfigurationCredentials
    {
        public readonly Outputs.GetSourceSmartsheetsConfigurationCredentialsSourceSmartsheetsAuthorizationMethodApiAccessToken SourceSmartsheetsAuthorizationMethodApiAccessToken;
        public readonly Outputs.GetSourceSmartsheetsConfigurationCredentialsSourceSmartsheetsAuthorizationMethodOAuth20 SourceSmartsheetsAuthorizationMethodOAuth20;
        public readonly Outputs.GetSourceSmartsheetsConfigurationCredentialsSourceSmartsheetsUpdateAuthorizationMethodApiAccessToken SourceSmartsheetsUpdateAuthorizationMethodApiAccessToken;
        public readonly Outputs.GetSourceSmartsheetsConfigurationCredentialsSourceSmartsheetsUpdateAuthorizationMethodOAuth20 SourceSmartsheetsUpdateAuthorizationMethodOAuth20;

        [OutputConstructor]
        private GetSourceSmartsheetsConfigurationCredentials(
            Outputs.GetSourceSmartsheetsConfigurationCredentialsSourceSmartsheetsAuthorizationMethodApiAccessToken sourceSmartsheetsAuthorizationMethodApiAccessToken,

            Outputs.GetSourceSmartsheetsConfigurationCredentialsSourceSmartsheetsAuthorizationMethodOAuth20 sourceSmartsheetsAuthorizationMethodOAuth20,

            Outputs.GetSourceSmartsheetsConfigurationCredentialsSourceSmartsheetsUpdateAuthorizationMethodApiAccessToken sourceSmartsheetsUpdateAuthorizationMethodApiAccessToken,

            Outputs.GetSourceSmartsheetsConfigurationCredentialsSourceSmartsheetsUpdateAuthorizationMethodOAuth20 sourceSmartsheetsUpdateAuthorizationMethodOAuth20)
        {
            SourceSmartsheetsAuthorizationMethodApiAccessToken = sourceSmartsheetsAuthorizationMethodApiAccessToken;
            SourceSmartsheetsAuthorizationMethodOAuth20 = sourceSmartsheetsAuthorizationMethodOAuth20;
            SourceSmartsheetsUpdateAuthorizationMethodApiAccessToken = sourceSmartsheetsUpdateAuthorizationMethodApiAccessToken;
            SourceSmartsheetsUpdateAuthorizationMethodOAuth20 = sourceSmartsheetsUpdateAuthorizationMethodOAuth20;
        }
    }
}