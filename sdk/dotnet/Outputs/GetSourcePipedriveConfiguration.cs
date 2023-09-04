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
    public sealed class GetSourcePipedriveConfiguration
    {
        public readonly Outputs.GetSourcePipedriveConfigurationAuthorization Authorization;
        public readonly string ReplicationStartDate;
        public readonly string SourceType;

        [OutputConstructor]
        private GetSourcePipedriveConfiguration(
            Outputs.GetSourcePipedriveConfigurationAuthorization authorization,

            string replicationStartDate,

            string sourceType)
        {
            Authorization = authorization;
            ReplicationStartDate = replicationStartDate;
            SourceType = sourceType;
        }
    }
}