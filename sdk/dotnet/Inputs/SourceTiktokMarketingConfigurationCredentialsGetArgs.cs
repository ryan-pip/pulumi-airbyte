// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceTiktokMarketingConfigurationCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceTiktokMarketingAuthenticationMethodOAuth20")]
        public Input<Inputs.SourceTiktokMarketingConfigurationCredentialsSourceTiktokMarketingAuthenticationMethodOAuth20GetArgs>? SourceTiktokMarketingAuthenticationMethodOAuth20 { get; set; }

        [Input("sourceTiktokMarketingAuthenticationMethodSandboxAccessToken")]
        public Input<Inputs.SourceTiktokMarketingConfigurationCredentialsSourceTiktokMarketingAuthenticationMethodSandboxAccessTokenGetArgs>? SourceTiktokMarketingAuthenticationMethodSandboxAccessToken { get; set; }

        [Input("sourceTiktokMarketingUpdateAuthenticationMethodOAuth20")]
        public Input<Inputs.SourceTiktokMarketingConfigurationCredentialsSourceTiktokMarketingUpdateAuthenticationMethodOAuth20GetArgs>? SourceTiktokMarketingUpdateAuthenticationMethodOAuth20 { get; set; }

        [Input("sourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken")]
        public Input<Inputs.SourceTiktokMarketingConfigurationCredentialsSourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessTokenGetArgs>? SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken { get; set; }

        public SourceTiktokMarketingConfigurationCredentialsGetArgs()
        {
        }
        public static new SourceTiktokMarketingConfigurationCredentialsGetArgs Empty => new SourceTiktokMarketingConfigurationCredentialsGetArgs();
    }
}