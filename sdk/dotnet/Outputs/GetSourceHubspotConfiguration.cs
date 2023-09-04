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
    public sealed class GetSourceHubspotConfiguration
    {
        public readonly Outputs.GetSourceHubspotConfigurationCredentials Credentials;
        public readonly string SourceType;
        public readonly string StartDate;

        [OutputConstructor]
        private GetSourceHubspotConfiguration(
            Outputs.GetSourceHubspotConfigurationCredentials credentials,

            string sourceType,

            string startDate)
        {
            Credentials = credentials;
            SourceType = sourceType;
            StartDate = startDate;
        }
    }
}