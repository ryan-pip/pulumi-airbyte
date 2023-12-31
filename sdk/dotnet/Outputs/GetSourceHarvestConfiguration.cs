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
    public sealed class GetSourceHarvestConfiguration
    {
        public readonly string AccountId;
        public readonly Outputs.GetSourceHarvestConfigurationCredentials Credentials;
        public readonly string ReplicationEndDate;
        public readonly string ReplicationStartDate;
        public readonly string SourceType;

        [OutputConstructor]
        private GetSourceHarvestConfiguration(
            string accountId,

            Outputs.GetSourceHarvestConfigurationCredentials credentials,

            string replicationEndDate,

            string replicationStartDate,

            string sourceType)
        {
            AccountId = accountId;
            Credentials = credentials;
            ReplicationEndDate = replicationEndDate;
            ReplicationStartDate = replicationStartDate;
            SourceType = sourceType;
        }
    }
}
