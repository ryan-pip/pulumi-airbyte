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
    public sealed class DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompression
    {
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionGzip? DestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionGzip;
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompression? DestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompression;

        [OutputConstructor]
        private DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompression(
            Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionGzip? destinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionGzip,

            Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompression? destinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompression)
        {
            DestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionGzip = destinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionGzip;
            DestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompression = destinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompression;
        }
    }
}
