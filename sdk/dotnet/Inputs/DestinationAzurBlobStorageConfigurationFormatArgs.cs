// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationAzurBlobStorageConfigurationFormatArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationAzureBlobStorageOutputFormatCsvCommaSeparatedValues")]
        public Input<Inputs.DestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageOutputFormatCsvCommaSeparatedValuesArgs>? DestinationAzureBlobStorageOutputFormatCsvCommaSeparatedValues { get; set; }

        [Input("destinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJson")]
        public Input<Inputs.DestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJsonArgs>? DestinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJson { get; set; }

        [Input("destinationAzureBlobStorageUpdateOutputFormatCsvCommaSeparatedValues")]
        public Input<Inputs.DestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageUpdateOutputFormatCsvCommaSeparatedValuesArgs>? DestinationAzureBlobStorageUpdateOutputFormatCsvCommaSeparatedValues { get; set; }

        [Input("destinationAzureBlobStorageUpdateOutputFormatJsonLinesNewlineDelimitedJson")]
        public Input<Inputs.DestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageUpdateOutputFormatJsonLinesNewlineDelimitedJsonArgs>? DestinationAzureBlobStorageUpdateOutputFormatJsonLinesNewlineDelimitedJson { get; set; }

        public DestinationAzurBlobStorageConfigurationFormatArgs()
        {
        }
        public static new DestinationAzurBlobStorageConfigurationFormatArgs Empty => new DestinationAzurBlobStorageConfigurationFormatArgs();
    }
}