// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationGcsConfigurationFormatDestinationGcsUpdateOutputFormatAvroApacheAvroCompressionCodecDestinationGcsUpdateOutputFormatAvroApacheAvroCompressionCodecZstandardGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("codec", required: true)]
        public Input<string> Codec { get; set; } = null!;

        [Input("compressionLevel")]
        public Input<int>? CompressionLevel { get; set; }

        [Input("includeChecksum")]
        public Input<bool>? IncludeChecksum { get; set; }

        public DestinationGcsConfigurationFormatDestinationGcsUpdateOutputFormatAvroApacheAvroCompressionCodecDestinationGcsUpdateOutputFormatAvroApacheAvroCompressionCodecZstandardGetArgs()
        {
        }
        public static new DestinationGcsConfigurationFormatDestinationGcsUpdateOutputFormatAvroApacheAvroCompressionCodecDestinationGcsUpdateOutputFormatAvroApacheAvroCompressionCodecZstandardGetArgs Empty => new DestinationGcsConfigurationFormatDestinationGcsUpdateOutputFormatAvroApacheAvroCompressionCodecDestinationGcsUpdateOutputFormatAvroApacheAvroCompressionCodecZstandardGetArgs();
    }
}