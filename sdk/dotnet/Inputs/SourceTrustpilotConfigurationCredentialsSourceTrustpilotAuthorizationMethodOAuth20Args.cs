// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceTrustpilotConfigurationCredentialsSourceTrustpilotAuthorizationMethodOAuth20Args : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        public Input<string> AccessToken { get; set; } = null!;

        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        [Input("refreshToken", required: true)]
        public Input<string> RefreshToken { get; set; } = null!;

        [Input("tokenExpiryDate", required: true)]
        public Input<string> TokenExpiryDate { get; set; } = null!;

        public SourceTrustpilotConfigurationCredentialsSourceTrustpilotAuthorizationMethodOAuth20Args()
        {
        }
        public static new SourceTrustpilotConfigurationCredentialsSourceTrustpilotAuthorizationMethodOAuth20Args Empty => new SourceTrustpilotConfigurationCredentialsSourceTrustpilotAuthorizationMethodOAuth20Args();
    }
}