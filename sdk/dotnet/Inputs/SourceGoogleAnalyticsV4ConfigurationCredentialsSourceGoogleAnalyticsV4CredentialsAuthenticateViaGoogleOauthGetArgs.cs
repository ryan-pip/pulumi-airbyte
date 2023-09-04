// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsV4ConfigurationCredentialsSourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauthGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken")]
        public Input<string>? AccessToken { get; set; }

        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        [Input("refreshToken", required: true)]
        public Input<string> RefreshToken { get; set; } = null!;

        public SourceGoogleAnalyticsV4ConfigurationCredentialsSourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauthGetArgs()
        {
        }
        public static new SourceGoogleAnalyticsV4ConfigurationCredentialsSourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauthGetArgs Empty => new SourceGoogleAnalyticsV4ConfigurationCredentialsSourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauthGetArgs();
    }
}