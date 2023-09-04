// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSquareConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceSquareAuthenticationApiKey")]
        public Input<Inputs.SourceSquareConfigurationCredentialsSourceSquareAuthenticationApiKeyArgs>? SourceSquareAuthenticationApiKey { get; set; }

        [Input("sourceSquareAuthenticationOauthAuthentication")]
        public Input<Inputs.SourceSquareConfigurationCredentialsSourceSquareAuthenticationOauthAuthenticationArgs>? SourceSquareAuthenticationOauthAuthentication { get; set; }

        [Input("sourceSquareUpdateAuthenticationApiKey")]
        public Input<Inputs.SourceSquareConfigurationCredentialsSourceSquareUpdateAuthenticationApiKeyArgs>? SourceSquareUpdateAuthenticationApiKey { get; set; }

        [Input("sourceSquareUpdateAuthenticationOauthAuthentication")]
        public Input<Inputs.SourceSquareConfigurationCredentialsSourceSquareUpdateAuthenticationOauthAuthenticationArgs>? SourceSquareUpdateAuthenticationOauthAuthentication { get; set; }

        public SourceSquareConfigurationCredentialsArgs()
        {
        }
        public static new SourceSquareConfigurationCredentialsArgs Empty => new SourceSquareConfigurationCredentialsArgs();
    }
}