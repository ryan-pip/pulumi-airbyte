// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSalesloftConfigurationCredentialsSourceSalesloftCredentialsAuthenticateViaApiKeyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        [Input("authType", required: true)]
        public Input<string> AuthType { get; set; } = null!;

        public SourceSalesloftConfigurationCredentialsSourceSalesloftCredentialsAuthenticateViaApiKeyGetArgs()
        {
        }
        public static new SourceSalesloftConfigurationCredentialsSourceSalesloftCredentialsAuthenticateViaApiKeyGetArgs Empty => new SourceSalesloftConfigurationCredentialsSourceSalesloftCredentialsAuthenticateViaApiKeyGetArgs();
    }
}