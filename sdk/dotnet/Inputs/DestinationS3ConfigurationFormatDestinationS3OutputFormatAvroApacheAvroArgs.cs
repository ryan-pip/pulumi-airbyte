// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvroArgs : global::Pulumi.ResourceArgs
    {
        [Input("compressionCodec", required: true)]
        public Input<Inputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvroCompressionCodecArgs> CompressionCodec { get; set; } = null!;

        [Input("formatType", required: true)]
        public Input<string> FormatType { get; set; } = null!;

        public DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvroArgs()
        {
        }
        public static new DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvroArgs Empty => new DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvroArgs();
    }
}
