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
    public sealed class GetSourceZendeskTalkConfigurationCredentials
    {
        public readonly Outputs.GetSourceZendeskTalkConfigurationCredentialsSourceZendeskTalkAuthenticationApiToken SourceZendeskTalkAuthenticationApiToken;
        public readonly Outputs.GetSourceZendeskTalkConfigurationCredentialsSourceZendeskTalkAuthenticationOAuth20 SourceZendeskTalkAuthenticationOAuth20;
        public readonly Outputs.GetSourceZendeskTalkConfigurationCredentialsSourceZendeskTalkUpdateAuthenticationApiToken SourceZendeskTalkUpdateAuthenticationApiToken;
        public readonly Outputs.GetSourceZendeskTalkConfigurationCredentialsSourceZendeskTalkUpdateAuthenticationOAuth20 SourceZendeskTalkUpdateAuthenticationOAuth20;

        [OutputConstructor]
        private GetSourceZendeskTalkConfigurationCredentials(
            Outputs.GetSourceZendeskTalkConfigurationCredentialsSourceZendeskTalkAuthenticationApiToken sourceZendeskTalkAuthenticationApiToken,

            Outputs.GetSourceZendeskTalkConfigurationCredentialsSourceZendeskTalkAuthenticationOAuth20 sourceZendeskTalkAuthenticationOAuth20,

            Outputs.GetSourceZendeskTalkConfigurationCredentialsSourceZendeskTalkUpdateAuthenticationApiToken sourceZendeskTalkUpdateAuthenticationApiToken,

            Outputs.GetSourceZendeskTalkConfigurationCredentialsSourceZendeskTalkUpdateAuthenticationOAuth20 sourceZendeskTalkUpdateAuthenticationOAuth20)
        {
            SourceZendeskTalkAuthenticationApiToken = sourceZendeskTalkAuthenticationApiToken;
            SourceZendeskTalkAuthenticationOAuth20 = sourceZendeskTalkAuthenticationOAuth20;
            SourceZendeskTalkUpdateAuthenticationApiToken = sourceZendeskTalkUpdateAuthenticationApiToken;
            SourceZendeskTalkUpdateAuthenticationOAuth20 = sourceZendeskTalkUpdateAuthenticationOAuth20;
        }
    }
}
