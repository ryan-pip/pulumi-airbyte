// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceZendeskSupportConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceZendeskSupportAuthenticationApiToken")]
        public Input<Inputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportAuthenticationApiTokenArgs>? SourceZendeskSupportAuthenticationApiToken { get; set; }

        [Input("sourceZendeskSupportAuthenticationOAuth20")]
        public Input<Inputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportAuthenticationOAuth20Args>? SourceZendeskSupportAuthenticationOAuth20 { get; set; }

        [Input("sourceZendeskSupportUpdateAuthenticationApiToken")]
        public Input<Inputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportUpdateAuthenticationApiTokenArgs>? SourceZendeskSupportUpdateAuthenticationApiToken { get; set; }

        [Input("sourceZendeskSupportUpdateAuthenticationOAuth20")]
        public Input<Inputs.SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportUpdateAuthenticationOAuth20Args>? SourceZendeskSupportUpdateAuthenticationOAuth20 { get; set; }

        public SourceZendeskSupportConfigurationCredentialsArgs()
        {
        }
        public static new SourceZendeskSupportConfigurationCredentialsArgs Empty => new SourceZendeskSupportConfigurationCredentialsArgs();
    }
}
