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
    public sealed class GetSourceAzureBlobStorageConfigurationFormat
    {
        public readonly Outputs.GetSourceAzureBlobStorageConfigurationFormatSourceAzureBlobStorageInputFormatJsonLinesNewlineDelimitedJson SourceAzureBlobStorageInputFormatJsonLinesNewlineDelimitedJson;
        public readonly Outputs.GetSourceAzureBlobStorageConfigurationFormatSourceAzureBlobStorageUpdateInputFormatJsonLinesNewlineDelimitedJson SourceAzureBlobStorageUpdateInputFormatJsonLinesNewlineDelimitedJson;

        [OutputConstructor]
        private GetSourceAzureBlobStorageConfigurationFormat(
            Outputs.GetSourceAzureBlobStorageConfigurationFormatSourceAzureBlobStorageInputFormatJsonLinesNewlineDelimitedJson sourceAzureBlobStorageInputFormatJsonLinesNewlineDelimitedJson,

            Outputs.GetSourceAzureBlobStorageConfigurationFormatSourceAzureBlobStorageUpdateInputFormatJsonLinesNewlineDelimitedJson sourceAzureBlobStorageUpdateInputFormatJsonLinesNewlineDelimitedJson)
        {
            SourceAzureBlobStorageInputFormatJsonLinesNewlineDelimitedJson = sourceAzureBlobStorageInputFormatJsonLinesNewlineDelimitedJson;
            SourceAzureBlobStorageUpdateInputFormatJsonLinesNewlineDelimitedJson = sourceAzureBlobStorageUpdateInputFormatJsonLinesNewlineDelimitedJson;
        }
    }
}