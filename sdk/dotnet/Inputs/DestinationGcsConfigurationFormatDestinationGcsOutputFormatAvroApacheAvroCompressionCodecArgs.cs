// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationGcsConfigurationFormatDestinationGcsOutputFormatAvroApacheAvroCompressionCodecArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationGcsOutputFormatAvroApacheAvroCompressionCodecBzip2")]
        public Input<Inputs.DestinationGcsConfigurationFormatDestinationGcsOutputFormatAvroApacheAvroCompressionCodecDestinationGcsOutputFormatAvroApacheAvroCompressionCodecBzip2Args>? DestinationGcsOutputFormatAvroApacheAvroCompressionCodecBzip2 { get; set; }

        [Input("destinationGcsOutputFormatAvroApacheAvroCompressionCodecDeflate")]
        public Input<Inputs.DestinationGcsConfigurationFormatDestinationGcsOutputFormatAvroApacheAvroCompressionCodecDestinationGcsOutputFormatAvroApacheAvroCompressionCodecDeflateArgs>? DestinationGcsOutputFormatAvroApacheAvroCompressionCodecDeflate { get; set; }

        [Input("destinationGcsOutputFormatAvroApacheAvroCompressionCodecNoCompression")]
        public Input<Inputs.DestinationGcsConfigurationFormatDestinationGcsOutputFormatAvroApacheAvroCompressionCodecDestinationGcsOutputFormatAvroApacheAvroCompressionCodecNoCompressionArgs>? DestinationGcsOutputFormatAvroApacheAvroCompressionCodecNoCompression { get; set; }

        [Input("destinationGcsOutputFormatAvroApacheAvroCompressionCodecSnappy")]
        public Input<Inputs.DestinationGcsConfigurationFormatDestinationGcsOutputFormatAvroApacheAvroCompressionCodecDestinationGcsOutputFormatAvroApacheAvroCompressionCodecSnappyArgs>? DestinationGcsOutputFormatAvroApacheAvroCompressionCodecSnappy { get; set; }

        [Input("destinationGcsOutputFormatAvroApacheAvroCompressionCodecXz")]
        public Input<Inputs.DestinationGcsConfigurationFormatDestinationGcsOutputFormatAvroApacheAvroCompressionCodecDestinationGcsOutputFormatAvroApacheAvroCompressionCodecXzArgs>? DestinationGcsOutputFormatAvroApacheAvroCompressionCodecXz { get; set; }

        [Input("destinationGcsOutputFormatAvroApacheAvroCompressionCodecZstandard")]
        public Input<Inputs.DestinationGcsConfigurationFormatDestinationGcsOutputFormatAvroApacheAvroCompressionCodecDestinationGcsOutputFormatAvroApacheAvroCompressionCodecZstandardArgs>? DestinationGcsOutputFormatAvroApacheAvroCompressionCodecZstandard { get; set; }

        public DestinationGcsConfigurationFormatDestinationGcsOutputFormatAvroApacheAvroCompressionCodecArgs()
        {
        }
        public static new DestinationGcsConfigurationFormatDestinationGcsOutputFormatAvroApacheAvroCompressionCodecArgs Empty => new DestinationGcsConfigurationFormatDestinationGcsOutputFormatAvroApacheAvroCompressionCodecArgs();
    }
}