// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatCsvCommaSeparatedValuesArgs : global::Pulumi.ResourceArgs
    {
        [Input("compression")]
        public Input<Inputs.DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatCsvCommaSeparatedValuesCompressionArgs>? Compression { get; set; }

        [Input("flattening", required: true)]
        public Input<string> Flattening { get; set; } = null!;

        [Input("formatType", required: true)]
        public Input<string> FormatType { get; set; } = null!;

        public DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatCsvCommaSeparatedValuesArgs()
        {
        }
        public static new DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatCsvCommaSeparatedValuesArgs Empty => new DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatCsvCommaSeparatedValuesArgs();
    }
}