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
    public sealed class DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJson
    {
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompression? Compression;
        public readonly string? Flattening;
        public readonly string FormatType;

        [OutputConstructor]
        private DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJson(
            Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJsonCompression? compression,

            string? flattening,

            string formatType)
        {
            Compression = compression;
            Flattening = flattening;
            FormatType = formatType;
        }
    }
}