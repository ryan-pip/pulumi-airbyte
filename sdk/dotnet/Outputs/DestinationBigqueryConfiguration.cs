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
    public sealed class DestinationBigqueryConfiguration
    {
        public readonly int? BigQueryClientBufferSizeMb;
        public readonly string? CredentialsJson;
        public readonly string DatasetId;
        public readonly string DatasetLocation;
        public readonly string DestinationType;
        public readonly Outputs.DestinationBigqueryConfigurationLoadingMethod? LoadingMethod;
        public readonly string ProjectId;
        public readonly string? RawDataDataset;
        public readonly string? TransformationPriority;
        public readonly bool? Use1s1tFormat;

        [OutputConstructor]
        private DestinationBigqueryConfiguration(
            int? bigQueryClientBufferSizeMb,

            string? credentialsJson,

            string datasetId,

            string datasetLocation,

            string destinationType,

            Outputs.DestinationBigqueryConfigurationLoadingMethod? loadingMethod,

            string projectId,

            string? rawDataDataset,

            string? transformationPriority,

            bool? use1s1tFormat)
        {
            BigQueryClientBufferSizeMb = bigQueryClientBufferSizeMb;
            CredentialsJson = credentialsJson;
            DatasetId = datasetId;
            DatasetLocation = datasetLocation;
            DestinationType = destinationType;
            LoadingMethod = loadingMethod;
            ProjectId = projectId;
            RawDataDataset = rawDataDataset;
            TransformationPriority = transformationPriority;
            Use1s1tFormat = use1s1tFormat;
        }
    }
}
