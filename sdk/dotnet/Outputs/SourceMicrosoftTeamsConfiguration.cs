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
    public sealed class SourceMicrosoftTeamsConfiguration
    {
        public readonly Outputs.SourceMicrosoftTeamsConfigurationCredentials? Credentials;
        public readonly string Period;
        public readonly string SourceType;

        [OutputConstructor]
        private SourceMicrosoftTeamsConfiguration(
            Outputs.SourceMicrosoftTeamsConfigurationCredentials? credentials,

            string period,

            string sourceType)
        {
            Credentials = credentials;
            Period = period;
            SourceType = sourceType;
        }
    }
}
