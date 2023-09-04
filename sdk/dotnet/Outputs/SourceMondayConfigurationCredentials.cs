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
    public sealed class SourceMondayConfigurationCredentials
    {
        public readonly Outputs.SourceMondayConfigurationCredentialsSourceMondayAuthorizationMethodApiToken? SourceMondayAuthorizationMethodApiToken;
        public readonly Outputs.SourceMondayConfigurationCredentialsSourceMondayAuthorizationMethodOAuth20? SourceMondayAuthorizationMethodOAuth20;
        public readonly Outputs.SourceMondayConfigurationCredentialsSourceMondayUpdateAuthorizationMethodApiToken? SourceMondayUpdateAuthorizationMethodApiToken;
        public readonly Outputs.SourceMondayConfigurationCredentialsSourceMondayUpdateAuthorizationMethodOAuth20? SourceMondayUpdateAuthorizationMethodOAuth20;

        [OutputConstructor]
        private SourceMondayConfigurationCredentials(
            Outputs.SourceMondayConfigurationCredentialsSourceMondayAuthorizationMethodApiToken? sourceMondayAuthorizationMethodApiToken,

            Outputs.SourceMondayConfigurationCredentialsSourceMondayAuthorizationMethodOAuth20? sourceMondayAuthorizationMethodOAuth20,

            Outputs.SourceMondayConfigurationCredentialsSourceMondayUpdateAuthorizationMethodApiToken? sourceMondayUpdateAuthorizationMethodApiToken,

            Outputs.SourceMondayConfigurationCredentialsSourceMondayUpdateAuthorizationMethodOAuth20? sourceMondayUpdateAuthorizationMethodOAuth20)
        {
            SourceMondayAuthorizationMethodApiToken = sourceMondayAuthorizationMethodApiToken;
            SourceMondayAuthorizationMethodOAuth20 = sourceMondayAuthorizationMethodOAuth20;
            SourceMondayUpdateAuthorizationMethodApiToken = sourceMondayUpdateAuthorizationMethodApiToken;
            SourceMondayUpdateAuthorizationMethodOAuth20 = sourceMondayUpdateAuthorizationMethodOAuth20;
        }
    }
}