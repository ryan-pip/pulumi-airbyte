// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvroCompressionCodecDestinationS3OutputFormatAvroApacheAvroCompressionCodecXzArgs : global::Pulumi.ResourceArgs
    {
        [Input("codec", required: true)]
        public Input<string> Codec { get; set; } = null!;

        [Input("compressionLevel", required: true)]
        public Input<int> CompressionLevel { get; set; } = null!;

        public DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvroCompressionCodecDestinationS3OutputFormatAvroApacheAvroCompressionCodecXzArgs()
        {
        }
        public static new DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvroCompressionCodecDestinationS3OutputFormatAvroApacheAvroCompressionCodecXzArgs Empty => new DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvroCompressionCodecDestinationS3OutputFormatAvroApacheAvroCompressionCodecXzArgs();
    }
}
