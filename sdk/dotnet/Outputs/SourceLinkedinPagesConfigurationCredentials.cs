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
    public sealed class SourceLinkedinPagesConfigurationCredentials
    {
        public readonly Outputs.SourceLinkedinPagesConfigurationCredentialsSourceLinkedinPagesAuthenticationAccessToken? SourceLinkedinPagesAuthenticationAccessToken;
        public readonly Outputs.SourceLinkedinPagesConfigurationCredentialsSourceLinkedinPagesAuthenticationOAuth20? SourceLinkedinPagesAuthenticationOAuth20;
        public readonly Outputs.SourceLinkedinPagesConfigurationCredentialsSourceLinkedinPagesUpdateAuthenticationAccessToken? SourceLinkedinPagesUpdateAuthenticationAccessToken;
        public readonly Outputs.SourceLinkedinPagesConfigurationCredentialsSourceLinkedinPagesUpdateAuthenticationOAuth20? SourceLinkedinPagesUpdateAuthenticationOAuth20;

        [OutputConstructor]
        private SourceLinkedinPagesConfigurationCredentials(
            Outputs.SourceLinkedinPagesConfigurationCredentialsSourceLinkedinPagesAuthenticationAccessToken? sourceLinkedinPagesAuthenticationAccessToken,

            Outputs.SourceLinkedinPagesConfigurationCredentialsSourceLinkedinPagesAuthenticationOAuth20? sourceLinkedinPagesAuthenticationOAuth20,

            Outputs.SourceLinkedinPagesConfigurationCredentialsSourceLinkedinPagesUpdateAuthenticationAccessToken? sourceLinkedinPagesUpdateAuthenticationAccessToken,

            Outputs.SourceLinkedinPagesConfigurationCredentialsSourceLinkedinPagesUpdateAuthenticationOAuth20? sourceLinkedinPagesUpdateAuthenticationOAuth20)
        {
            SourceLinkedinPagesAuthenticationAccessToken = sourceLinkedinPagesAuthenticationAccessToken;
            SourceLinkedinPagesAuthenticationOAuth20 = sourceLinkedinPagesAuthenticationOAuth20;
            SourceLinkedinPagesUpdateAuthenticationAccessToken = sourceLinkedinPagesUpdateAuthenticationAccessToken;
            SourceLinkedinPagesUpdateAuthenticationOAuth20 = sourceLinkedinPagesUpdateAuthenticationOAuth20;
        }
    }
}
