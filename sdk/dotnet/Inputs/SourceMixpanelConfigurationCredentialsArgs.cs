// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMixpanelConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceMixpanelAuthenticationWildcardProjectSecret")]
        public Input<Inputs.SourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardProjectSecretArgs>? SourceMixpanelAuthenticationWildcardProjectSecret { get; set; }

        [Input("sourceMixpanelAuthenticationWildcardServiceAccount")]
        public Input<Inputs.SourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardServiceAccountArgs>? SourceMixpanelAuthenticationWildcardServiceAccount { get; set; }

        [Input("sourceMixpanelUpdateAuthenticationWildcardProjectSecret")]
        public Input<Inputs.SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecretArgs>? SourceMixpanelUpdateAuthenticationWildcardProjectSecret { get; set; }

        [Input("sourceMixpanelUpdateAuthenticationWildcardServiceAccount")]
        public Input<Inputs.SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccountArgs>? SourceMixpanelUpdateAuthenticationWildcardServiceAccount { get; set; }

        public SourceMixpanelConfigurationCredentialsArgs()
        {
        }
        public static new SourceMixpanelConfigurationCredentialsArgs Empty => new SourceMixpanelConfigurationCredentialsArgs();
    }
}