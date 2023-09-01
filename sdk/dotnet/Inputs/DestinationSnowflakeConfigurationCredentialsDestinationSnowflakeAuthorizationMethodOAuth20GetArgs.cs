// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodOAuth20GetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        public Input<string> AccessToken { get; set; } = null!;

        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        [Input("refreshToken", required: true)]
        public Input<string> RefreshToken { get; set; } = null!;

        public DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodOAuth20GetArgs()
        {
        }
        public static new DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodOAuth20GetArgs Empty => new DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodOAuth20GetArgs();
    }
}