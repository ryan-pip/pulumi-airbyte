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
    public sealed class GetSourceQuickbooksConfigurationCredentials
    {
        public readonly Outputs.GetSourceQuickbooksConfigurationCredentialsSourceQuickbooksAuthorizationMethodOAuth20 SourceQuickbooksAuthorizationMethodOAuth20;
        public readonly Outputs.GetSourceQuickbooksConfigurationCredentialsSourceQuickbooksUpdateAuthorizationMethodOAuth20 SourceQuickbooksUpdateAuthorizationMethodOAuth20;

        [OutputConstructor]
        private GetSourceQuickbooksConfigurationCredentials(
            Outputs.GetSourceQuickbooksConfigurationCredentialsSourceQuickbooksAuthorizationMethodOAuth20 sourceQuickbooksAuthorizationMethodOAuth20,

            Outputs.GetSourceQuickbooksConfigurationCredentialsSourceQuickbooksUpdateAuthorizationMethodOAuth20 sourceQuickbooksUpdateAuthorizationMethodOAuth20)
        {
            SourceQuickbooksAuthorizationMethodOAuth20 = sourceQuickbooksAuthorizationMethodOAuth20;
            SourceQuickbooksUpdateAuthorizationMethodOAuth20 = sourceQuickbooksUpdateAuthorizationMethodOAuth20;
        }
    }
}