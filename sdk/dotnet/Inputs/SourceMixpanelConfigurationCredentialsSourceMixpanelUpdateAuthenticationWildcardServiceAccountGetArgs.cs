// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccountGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("optionTitle")]
        public Input<string>? OptionTitle { get; set; }

        [Input("secret", required: true)]
        public Input<string> Secret { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccountGetArgs()
        {
        }
        public static new SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccountGetArgs Empty => new SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccountGetArgs();
    }
}
