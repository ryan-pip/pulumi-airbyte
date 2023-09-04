// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceHarvestConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth")]
        public Input<Inputs.SourceHarvestConfigurationCredentialsSourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuthArgs>? SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth { get; set; }

        [Input("sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken")]
        public Input<Inputs.SourceHarvestConfigurationCredentialsSourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessTokenArgs>? SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken { get; set; }

        [Input("sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth")]
        public Input<Inputs.SourceHarvestConfigurationCredentialsSourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuthArgs>? SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth { get; set; }

        [Input("sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken")]
        public Input<Inputs.SourceHarvestConfigurationCredentialsSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenArgs>? SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken { get; set; }

        public SourceHarvestConfigurationCredentialsArgs()
        {
        }
        public static new SourceHarvestConfigurationCredentialsArgs Empty => new SourceHarvestConfigurationCredentialsArgs();
    }
}