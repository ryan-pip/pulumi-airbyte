// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class SourceHarvestConfigurationCredentials
    {
        public readonly Outputs.SourceHarvestConfigurationCredentialsSourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth? SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth;
        public readonly Outputs.SourceHarvestConfigurationCredentialsSourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken? SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken;
        public readonly Outputs.SourceHarvestConfigurationCredentialsSourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth? SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth;
        public readonly Outputs.SourceHarvestConfigurationCredentialsSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken? SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken;

        [OutputConstructor]
        private SourceHarvestConfigurationCredentials(
            Outputs.SourceHarvestConfigurationCredentialsSourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth? sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth,

            Outputs.SourceHarvestConfigurationCredentialsSourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken? sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken,

            Outputs.SourceHarvestConfigurationCredentialsSourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth? sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth,

            Outputs.SourceHarvestConfigurationCredentialsSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken? sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken)
        {
            SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth = sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth;
            SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken = sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken;
            SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth = sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth;
            SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken = sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken;
        }
    }
}
