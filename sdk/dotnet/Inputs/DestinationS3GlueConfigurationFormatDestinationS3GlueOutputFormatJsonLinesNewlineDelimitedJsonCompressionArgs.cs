// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationS3GlueConfigurationFormatDestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionGzip")]
        public Input<Inputs.DestinationS3GlueConfigurationFormatDestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionDestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionGzipArgs>? DestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionGzip { get; set; }

        [Input("destinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompression")]
        public Input<Inputs.DestinationS3GlueConfigurationFormatDestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionDestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompressionArgs>? DestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionNoCompression { get; set; }

        public DestinationS3GlueConfigurationFormatDestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionArgs()
        {
        }
        public static new DestinationS3GlueConfigurationFormatDestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionArgs Empty => new DestinationS3GlueConfigurationFormatDestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonCompressionArgs();
    }
}