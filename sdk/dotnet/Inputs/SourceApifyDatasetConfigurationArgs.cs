// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceApifyDatasetConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("clean")]
        public Input<bool>? Clean { get; set; }

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        public SourceApifyDatasetConfigurationArgs()
        {
        }
        public static new SourceApifyDatasetConfigurationArgs Empty => new SourceApifyDatasetConfigurationArgs();
    }
}