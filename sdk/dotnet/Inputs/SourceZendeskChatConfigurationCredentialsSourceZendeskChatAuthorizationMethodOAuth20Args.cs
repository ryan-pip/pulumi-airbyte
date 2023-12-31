// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceZendeskChatConfigurationCredentialsSourceZendeskChatAuthorizationMethodOAuth20Args : global::Pulumi.ResourceArgs
    {
        [Input("accessToken")]
        public Input<string>? AccessToken { get; set; }

        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        [Input("credentials", required: true)]
        public Input<string> Credentials { get; set; } = null!;

        [Input("refreshToken")]
        public Input<string>? RefreshToken { get; set; }

        public SourceZendeskChatConfigurationCredentialsSourceZendeskChatAuthorizationMethodOAuth20Args()
        {
        }
        public static new SourceZendeskChatConfigurationCredentialsSourceZendeskChatAuthorizationMethodOAuth20Args Empty => new SourceZendeskChatConfigurationCredentialsSourceZendeskChatAuthorizationMethodOAuth20Args();
    }
}
