// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportAuthenticationOAuth20GetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        public Input<string> AccessToken { get; set; } = null!;

        [Input("additionalProperties")]
        public Input<string>? AdditionalProperties { get; set; }

        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        public SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportAuthenticationOAuth20GetArgs()
        {
        }
        public static new SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportAuthenticationOAuth20GetArgs Empty => new SourceZendeskSupportConfigurationCredentialsSourceZendeskSupportAuthenticationOAuth20GetArgs();
    }
}