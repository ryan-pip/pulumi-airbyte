// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceAirtableConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceAirtableAuthenticationOAuth20")]
        public Input<Inputs.SourceAirtableConfigurationCredentialsSourceAirtableAuthenticationOAuth20Args>? SourceAirtableAuthenticationOAuth20 { get; set; }

        [Input("sourceAirtableAuthenticationPersonalAccessToken")]
        public Input<Inputs.SourceAirtableConfigurationCredentialsSourceAirtableAuthenticationPersonalAccessTokenArgs>? SourceAirtableAuthenticationPersonalAccessToken { get; set; }

        [Input("sourceAirtableUpdateAuthenticationOAuth20")]
        public Input<Inputs.SourceAirtableConfigurationCredentialsSourceAirtableUpdateAuthenticationOAuth20Args>? SourceAirtableUpdateAuthenticationOAuth20 { get; set; }

        [Input("sourceAirtableUpdateAuthenticationPersonalAccessToken")]
        public Input<Inputs.SourceAirtableConfigurationCredentialsSourceAirtableUpdateAuthenticationPersonalAccessTokenArgs>? SourceAirtableUpdateAuthenticationPersonalAccessToken { get; set; }

        public SourceAirtableConfigurationCredentialsArgs()
        {
        }
        public static new SourceAirtableConfigurationCredentialsArgs Empty => new SourceAirtableConfigurationCredentialsArgs();
    }
}