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
    public sealed class SourceSlackConfigurationCredentials
    {
        public readonly Outputs.SourceSlackConfigurationCredentialsSourceSlackAuthenticationMechanismApiToken? SourceSlackAuthenticationMechanismApiToken;
        public readonly Outputs.SourceSlackConfigurationCredentialsSourceSlackAuthenticationMechanismSignInViaSlackOAuth? SourceSlackAuthenticationMechanismSignInViaSlackOAuth;
        public readonly Outputs.SourceSlackConfigurationCredentialsSourceSlackUpdateAuthenticationMechanismApiToken? SourceSlackUpdateAuthenticationMechanismApiToken;
        public readonly Outputs.SourceSlackConfigurationCredentialsSourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth? SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth;

        [OutputConstructor]
        private SourceSlackConfigurationCredentials(
            Outputs.SourceSlackConfigurationCredentialsSourceSlackAuthenticationMechanismApiToken? sourceSlackAuthenticationMechanismApiToken,

            Outputs.SourceSlackConfigurationCredentialsSourceSlackAuthenticationMechanismSignInViaSlackOAuth? sourceSlackAuthenticationMechanismSignInViaSlackOAuth,

            Outputs.SourceSlackConfigurationCredentialsSourceSlackUpdateAuthenticationMechanismApiToken? sourceSlackUpdateAuthenticationMechanismApiToken,

            Outputs.SourceSlackConfigurationCredentialsSourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth? sourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth)
        {
            SourceSlackAuthenticationMechanismApiToken = sourceSlackAuthenticationMechanismApiToken;
            SourceSlackAuthenticationMechanismSignInViaSlackOAuth = sourceSlackAuthenticationMechanismSignInViaSlackOAuth;
            SourceSlackUpdateAuthenticationMechanismApiToken = sourceSlackUpdateAuthenticationMechanismApiToken;
            SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth = sourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth;
        }
    }
}
