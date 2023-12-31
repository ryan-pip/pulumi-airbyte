// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceAzureBlobStorageConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("azureBlobStorageAccountKey", required: true)]
        public Input<string> AzureBlobStorageAccountKey { get; set; } = null!;

        [Input("azureBlobStorageAccountName", required: true)]
        public Input<string> AzureBlobStorageAccountName { get; set; } = null!;

        [Input("azureBlobStorageBlobsPrefix")]
        public Input<string>? AzureBlobStorageBlobsPrefix { get; set; }

        [Input("azureBlobStorageContainerName", required: true)]
        public Input<string> AzureBlobStorageContainerName { get; set; } = null!;

        [Input("azureBlobStorageEndpoint")]
        public Input<string>? AzureBlobStorageEndpoint { get; set; }

        [Input("azureBlobStorageSchemaInferenceLimit")]
        public Input<int>? AzureBlobStorageSchemaInferenceLimit { get; set; }

        [Input("format", required: true)]
        public Input<Inputs.SourceAzureBlobStorageConfigurationFormatArgs> Format { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        public SourceAzureBlobStorageConfigurationArgs()
        {
        }
        public static new SourceAzureBlobStorageConfigurationArgs Empty => new SourceAzureBlobStorageConfigurationArgs();
    }
}
