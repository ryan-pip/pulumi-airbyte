// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceOutbrainAmplifyConfigurationCredentialsSourceOutbrainAmplifyAuthenticationMethodAccessTokenArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        public Input<string> AccessToken { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SourceOutbrainAmplifyConfigurationCredentialsSourceOutbrainAmplifyAuthenticationMethodAccessTokenArgs()
        {
        }
        public static new SourceOutbrainAmplifyConfigurationCredentialsSourceOutbrainAmplifyAuthenticationMethodAccessTokenArgs Empty => new SourceOutbrainAmplifyConfigurationCredentialsSourceOutbrainAmplifyAuthenticationMethodAccessTokenArgs();
    }
}
