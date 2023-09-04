// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationBigqueryConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bigQueryClientBufferSizeMb")]
        public Input<int>? BigQueryClientBufferSizeMb { get; set; }

        [Input("credentialsJson")]
        public Input<string>? CredentialsJson { get; set; }

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("datasetLocation", required: true)]
        public Input<string> DatasetLocation { get; set; } = null!;

        [Input("destinationType", required: true)]
        public Input<string> DestinationType { get; set; } = null!;

        [Input("loadingMethod")]
        public Input<Inputs.DestinationBigqueryConfigurationLoadingMethodGetArgs>? LoadingMethod { get; set; }

        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("rawDataDataset")]
        public Input<string>? RawDataDataset { get; set; }

        [Input("transformationPriority")]
        public Input<string>? TransformationPriority { get; set; }

        [Input("use1s1tFormat")]
        public Input<bool>? Use1s1tFormat { get; set; }

        public DestinationBigqueryConfigurationGetArgs()
        {
        }
        public static new DestinationBigqueryConfigurationGetArgs Empty => new DestinationBigqueryConfigurationGetArgs();
    }
}