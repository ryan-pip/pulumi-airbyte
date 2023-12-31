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
    public sealed class SourceMixpanelConfigurationCredentials
    {
        public readonly Outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardProjectSecret? SourceMixpanelAuthenticationWildcardProjectSecret;
        public readonly Outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardServiceAccount? SourceMixpanelAuthenticationWildcardServiceAccount;
        public readonly Outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecret? SourceMixpanelUpdateAuthenticationWildcardProjectSecret;
        public readonly Outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccount? SourceMixpanelUpdateAuthenticationWildcardServiceAccount;

        [OutputConstructor]
        private SourceMixpanelConfigurationCredentials(
            Outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardProjectSecret? sourceMixpanelAuthenticationWildcardProjectSecret,

            Outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardServiceAccount? sourceMixpanelAuthenticationWildcardServiceAccount,

            Outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecret? sourceMixpanelUpdateAuthenticationWildcardProjectSecret,

            Outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccount? sourceMixpanelUpdateAuthenticationWildcardServiceAccount)
        {
            SourceMixpanelAuthenticationWildcardProjectSecret = sourceMixpanelAuthenticationWildcardProjectSecret;
            SourceMixpanelAuthenticationWildcardServiceAccount = sourceMixpanelAuthenticationWildcardServiceAccount;
            SourceMixpanelUpdateAuthenticationWildcardProjectSecret = sourceMixpanelUpdateAuthenticationWildcardProjectSecret;
            SourceMixpanelUpdateAuthenticationWildcardServiceAccount = sourceMixpanelUpdateAuthenticationWildcardServiceAccount;
        }
    }
}
