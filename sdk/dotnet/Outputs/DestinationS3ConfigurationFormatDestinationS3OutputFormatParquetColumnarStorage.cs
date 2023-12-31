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
    public sealed class DestinationS3ConfigurationFormatDestinationS3OutputFormatParquetColumnarStorage
    {
        public readonly int? BlockSizeMb;
        public readonly string? CompressionCodec;
        public readonly bool? DictionaryEncoding;
        public readonly int? DictionaryPageSizeKb;
        public readonly string FormatType;
        public readonly int? MaxPaddingSizeMb;
        public readonly int? PageSizeKb;

        [OutputConstructor]
        private DestinationS3ConfigurationFormatDestinationS3OutputFormatParquetColumnarStorage(
            int? blockSizeMb,

            string? compressionCodec,

            bool? dictionaryEncoding,

            int? dictionaryPageSizeKb,

            string formatType,

            int? maxPaddingSizeMb,

            int? pageSizeKb)
        {
            BlockSizeMb = blockSizeMb;
            CompressionCodec = compressionCodec;
            DictionaryEncoding = dictionaryEncoding;
            DictionaryPageSizeKb = dictionaryPageSizeKb;
            FormatType = formatType;
            MaxPaddingSizeMb = maxPaddingSizeMb;
            PageSizeKb = pageSizeKb;
        }
    }
}
