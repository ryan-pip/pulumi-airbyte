// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceHarvestConfigurationCredentialsSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        public Input<string>? AdditionalProperties { get; set; }

        [Input("apiToken", required: true)]
        public Input<string> ApiToken { get; set; } = null!;

        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        public SourceHarvestConfigurationCredentialsSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenArgs()
        {
        }
        public static new SourceHarvestConfigurationCredentialsSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenArgs Empty => new SourceHarvestConfigurationCredentialsSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenArgs();
    }
}
