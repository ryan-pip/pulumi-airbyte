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
    public sealed class GetDestinationAzurBlobStorageConfigurationFormat
    {
        public readonly Outputs.GetDestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageOutputFormatCsvCommaSeparatedValues DestinationAzureBlobStorageOutputFormatCsvCommaSeparatedValues;
        public readonly Outputs.GetDestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJson DestinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJson;
        public readonly Outputs.GetDestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageUpdateOutputFormatCsvCommaSeparatedValues DestinationAzureBlobStorageUpdateOutputFormatCsvCommaSeparatedValues;
        public readonly Outputs.GetDestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageUpdateOutputFormatJsonLinesNewlineDelimitedJson DestinationAzureBlobStorageUpdateOutputFormatJsonLinesNewlineDelimitedJson;

        [OutputConstructor]
        private GetDestinationAzurBlobStorageConfigurationFormat(
            Outputs.GetDestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageOutputFormatCsvCommaSeparatedValues destinationAzureBlobStorageOutputFormatCsvCommaSeparatedValues,

            Outputs.GetDestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJson destinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJson,

            Outputs.GetDestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageUpdateOutputFormatCsvCommaSeparatedValues destinationAzureBlobStorageUpdateOutputFormatCsvCommaSeparatedValues,

            Outputs.GetDestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageUpdateOutputFormatJsonLinesNewlineDelimitedJson destinationAzureBlobStorageUpdateOutputFormatJsonLinesNewlineDelimitedJson)
        {
            DestinationAzureBlobStorageOutputFormatCsvCommaSeparatedValues = destinationAzureBlobStorageOutputFormatCsvCommaSeparatedValues;
            DestinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJson = destinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJson;
            DestinationAzureBlobStorageUpdateOutputFormatCsvCommaSeparatedValues = destinationAzureBlobStorageUpdateOutputFormatCsvCommaSeparatedValues;
            DestinationAzureBlobStorageUpdateOutputFormatJsonLinesNewlineDelimitedJson = destinationAzureBlobStorageUpdateOutputFormatJsonLinesNewlineDelimitedJson;
        }
    }
}