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
    public sealed class GetSourceSlackConfigurationCredentials
    {
        public readonly Outputs.GetSourceSlackConfigurationCredentialsSourceSlackAuthenticationMechanismApiToken SourceSlackAuthenticationMechanismApiToken;
        public readonly Outputs.GetSourceSlackConfigurationCredentialsSourceSlackAuthenticationMechanismSignInViaSlackOAuth SourceSlackAuthenticationMechanismSignInViaSlackOAuth;
        public readonly Outputs.GetSourceSlackConfigurationCredentialsSourceSlackUpdateAuthenticationMechanismApiToken SourceSlackUpdateAuthenticationMechanismApiToken;
        public readonly Outputs.GetSourceSlackConfigurationCredentialsSourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth;

        [OutputConstructor]
        private GetSourceSlackConfigurationCredentials(
            Outputs.GetSourceSlackConfigurationCredentialsSourceSlackAuthenticationMechanismApiToken sourceSlackAuthenticationMechanismApiToken,

            Outputs.GetSourceSlackConfigurationCredentialsSourceSlackAuthenticationMechanismSignInViaSlackOAuth sourceSlackAuthenticationMechanismSignInViaSlackOAuth,

            Outputs.GetSourceSlackConfigurationCredentialsSourceSlackUpdateAuthenticationMechanismApiToken sourceSlackUpdateAuthenticationMechanismApiToken,

            Outputs.GetSourceSlackConfigurationCredentialsSourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth sourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth)
        {
            SourceSlackAuthenticationMechanismApiToken = sourceSlackAuthenticationMechanismApiToken;
            SourceSlackAuthenticationMechanismSignInViaSlackOAuth = sourceSlackAuthenticationMechanismSignInViaSlackOAuth;
            SourceSlackUpdateAuthenticationMechanismApiToken = sourceSlackUpdateAuthenticationMechanismApiToken;
            SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth = sourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth;
        }
    }
}
