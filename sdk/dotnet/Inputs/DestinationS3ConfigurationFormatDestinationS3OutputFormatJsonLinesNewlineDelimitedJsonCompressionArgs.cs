// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionGzip")]
        public Input<Inputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionGzipArgs>? DestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionGzip { get; set; }

        [Input("destinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompression")]
        public Input<Inputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompressionArgs>? DestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompression { get; set; }

        public DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionArgs()
        {
        }
        public static new DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionArgs Empty => new DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionArgs();
    }
}
