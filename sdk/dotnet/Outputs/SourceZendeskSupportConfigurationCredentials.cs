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
    public sealed class SourceZendeskSupportConfigurationCredentials
    {
        public readonly Outputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportAuthenticationApiToken? SourceZendeskSupportAuthenticationApiToken;
        public readonly Outputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportAuthenticationOAuth20? SourceZendeskSupportAuthenticationOAuth20;
        public readonly Outputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportUpdateAuthenticationApiToken? SourceZendeskSupportUpdateAuthenticationApiToken;
        public readonly Outputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportUpdateAuthenticationOAuth20? SourceZendeskSupportUpdateAuthenticationOAuth20;

        [OutputConstructor]
        private SourceZendeskSupportConfigurationCredentials(
            Outputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportAuthenticationApiToken? sourceZendeskSupportAuthenticationApiToken,

            Outputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportAuthenticationOAuth20? sourceZendeskSupportAuthenticationOAuth20,

            Outputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportUpdateAuthenticationApiToken? sourceZendeskSupportUpdateAuthenticationApiToken,

            Outputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportUpdateAuthenticationOAuth20? sourceZendeskSupportUpdateAuthenticationOAuth20)
        {
            SourceZendeskSupportAuthenticationApiToken = sourceZendeskSupportAuthenticationApiToken;
            SourceZendeskSupportAuthenticationOAuth20 = sourceZendeskSupportAuthenticationOAuth20;
            SourceZendeskSupportUpdateAuthenticationApiToken = sourceZendeskSupportUpdateAuthenticationApiToken;
            SourceZendeskSupportUpdateAuthenticationOAuth20 = sourceZendeskSupportUpdateAuthenticationOAuth20;
        }
    }
}
