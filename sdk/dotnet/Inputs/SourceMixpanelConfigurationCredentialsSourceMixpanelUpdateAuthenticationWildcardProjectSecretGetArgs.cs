// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecretGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiSecret", required: true)]
        public Input<string> ApiSecret { get; set; } = null!;

        [Input("optionTitle")]
        public Input<string>? OptionTitle { get; set; }

        public SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecretGetArgs()
        {
        }
        public static new SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecretGetArgs Empty => new SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecretGetArgs();
    }
}
