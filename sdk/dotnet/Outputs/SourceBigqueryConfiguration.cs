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
    public sealed class SourceBigqueryConfiguration
    {
        public readonly string CredentialsJson;
        public readonly string? DatasetId;
        public readonly string ProjectId;
        public readonly string SourceType;

        [OutputConstructor]
        private SourceBigqueryConfiguration(
            string credentialsJson,

            string? datasetId,

            string projectId,

            string sourceType)
        {
            CredentialsJson = credentialsJson;
            DatasetId = datasetId;
            ProjectId = projectId;
            SourceType = sourceType;
        }
    }
}
