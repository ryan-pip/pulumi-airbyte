// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJsonArgs : global::Pulumi.ResourceArgs
    {
        [Input("formatType", required: true)]
        public Input<string> FormatType { get; set; } = null!;

        public DestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJsonArgs()
        {
        }
        public static new DestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJsonArgs Empty => new DestinationAzurBlobStorageConfigurationFormatDestinationAzureBlobStorageOutputFormatJsonLinesNewlineDelimitedJsonArgs();
    }
}