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
    public sealed class DestinationS3ConfigurationFormat
    {
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvro? DestinationS3OutputFormatAvroApacheAvro;
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatCsvCommaSeparatedValues? DestinationS3OutputFormatCsvCommaSeparatedValues;
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJson? DestinationS3OutputFormatJsonLinesNewlineDelimitedJson;
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatParquetColumnarStorage? DestinationS3OutputFormatParquetColumnarStorage;
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatAvroApacheAvro? DestinationS3UpdateOutputFormatAvroApacheAvro;
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatCsvCommaSeparatedValues? DestinationS3UpdateOutputFormatCsvCommaSeparatedValues;
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatJsonLinesNewlineDelimitedJson? DestinationS3UpdateOutputFormatJsonLinesNewlineDelimitedJson;
        public readonly Outputs.DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatParquetColumnarStorage? DestinationS3UpdateOutputFormatParquetColumnarStorage;

        [OutputConstructor]
        private DestinationS3ConfigurationFormat(
            Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatAvroApacheAvro? destinationS3OutputFormatAvroApacheAvro,

            Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatCsvCommaSeparatedValues? destinationS3OutputFormatCsvCommaSeparatedValues,

            Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatJsonLinesNewlineDelimitedJson? destinationS3OutputFormatJsonLinesNewlineDelimitedJson,

            Outputs.DestinationS3ConfigurationFormatDestinationS3OutputFormatParquetColumnarStorage? destinationS3OutputFormatParquetColumnarStorage,

            Outputs.DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatAvroApacheAvro? destinationS3UpdateOutputFormatAvroApacheAvro,

            Outputs.DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatCsvCommaSeparatedValues? destinationS3UpdateOutputFormatCsvCommaSeparatedValues,

            Outputs.DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatJsonLinesNewlineDelimitedJson? destinationS3UpdateOutputFormatJsonLinesNewlineDelimitedJson,

            Outputs.DestinationS3ConfigurationFormatDestinationS3UpdateOutputFormatParquetColumnarStorage? destinationS3UpdateOutputFormatParquetColumnarStorage)
        {
            DestinationS3OutputFormatAvroApacheAvro = destinationS3OutputFormatAvroApacheAvro;
            DestinationS3OutputFormatCsvCommaSeparatedValues = destinationS3OutputFormatCsvCommaSeparatedValues;
            DestinationS3OutputFormatJsonLinesNewlineDelimitedJson = destinationS3OutputFormatJsonLinesNewlineDelimitedJson;
            DestinationS3OutputFormatParquetColumnarStorage = destinationS3OutputFormatParquetColumnarStorage;
            DestinationS3UpdateOutputFormatAvroApacheAvro = destinationS3UpdateOutputFormatAvroApacheAvro;
            DestinationS3UpdateOutputFormatCsvCommaSeparatedValues = destinationS3UpdateOutputFormatCsvCommaSeparatedValues;
            DestinationS3UpdateOutputFormatJsonLinesNewlineDelimitedJson = destinationS3UpdateOutputFormatJsonLinesNewlineDelimitedJson;
            DestinationS3UpdateOutputFormatParquetColumnarStorage = destinationS3UpdateOutputFormatParquetColumnarStorage;
        }
    }
}
