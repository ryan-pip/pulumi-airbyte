// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceZendeskTalkConfigurationCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceZendeskTalkAuthenticationApiToken")]
        public Input<Inputs.SourceZendeskTalkConfigurationCredentialsSourceZendeskTalkAuthenticationApiTokenGetArgs>? SourceZendeskTalkAuthenticationApiToken { get; set; }

        [Input("sourceZendeskTalkAuthenticationOAuth20")]
        public Input<Inputs.SourceZendeskTalkConfigurationCredentialsSourceZendeskTalkAuthenticationOAuth20GetArgs>? SourceZendeskTalkAuthenticationOAuth20 { get; set; }

        [Input("sourceZendeskTalkUpdateAuthenticationApiToken")]
        public Input<Inputs.SourceZendeskTalkConfigurationCredentialsSourceZendeskTalkUpdateAuthenticationApiTokenGetArgs>? SourceZendeskTalkUpdateAuthenticationApiToken { get; set; }

        [Input("sourceZendeskTalkUpdateAuthenticationOAuth20")]
        public Input<Inputs.SourceZendeskTalkConfigurationCredentialsSourceZendeskTalkUpdateAuthenticationOAuth20GetArgs>? SourceZendeskTalkUpdateAuthenticationOAuth20 { get; set; }

        public SourceZendeskTalkConfigurationCredentialsGetArgs()
        {
        }
        public static new SourceZendeskTalkConfigurationCredentialsGetArgs Empty => new SourceZendeskTalkConfigurationCredentialsGetArgs();
    }
}
