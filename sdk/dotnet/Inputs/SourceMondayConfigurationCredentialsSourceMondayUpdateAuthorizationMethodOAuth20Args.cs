// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMondayConfigurationCredentialsSourceMondayUpdateAuthorizationMethodOAuth20Args : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        public Input<string> AccessToken { get; set; } = null!;

        [Input("authType", required: true)]
        public Input<string> AuthType { get; set; } = null!;

        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

        public SourceMondayConfigurationCredentialsSourceMondayUpdateAuthorizationMethodOAuth20Args()
        {
        }
        public static new SourceMondayConfigurationCredentialsSourceMondayUpdateAuthorizationMethodOAuth20Args Empty => new SourceMondayConfigurationCredentialsSourceMondayUpdateAuthorizationMethodOAuth20Args();
    }
}
