// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationDatabricksConfigurationDataSourceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationDatabricksDataSourceAmazonS3")]
        public Input<Inputs.DestinationDatabricksConfigurationDataSourceDestinationDatabricksDataSourceAmazonS3GetArgs>? DestinationDatabricksDataSourceAmazonS3 { get; set; }

        [Input("destinationDatabricksDataSourceAzureBlobStorage")]
        public Input<Inputs.DestinationDatabricksConfigurationDataSourceDestinationDatabricksDataSourceAzureBlobStorageGetArgs>? DestinationDatabricksDataSourceAzureBlobStorage { get; set; }

        [Input("destinationDatabricksDataSourceRecommendedManagedTables")]
        public Input<Inputs.DestinationDatabricksConfigurationDataSourceDestinationDatabricksDataSourceRecommendedManagedTablesGetArgs>? DestinationDatabricksDataSourceRecommendedManagedTables { get; set; }

        [Input("destinationDatabricksUpdateDataSourceAmazonS3")]
        public Input<Inputs.DestinationDatabricksConfigurationDataSourceDestinationDatabricksUpdateDataSourceAmazonS3GetArgs>? DestinationDatabricksUpdateDataSourceAmazonS3 { get; set; }

        [Input("destinationDatabricksUpdateDataSourceAzureBlobStorage")]
        public Input<Inputs.DestinationDatabricksConfigurationDataSourceDestinationDatabricksUpdateDataSourceAzureBlobStorageGetArgs>? DestinationDatabricksUpdateDataSourceAzureBlobStorage { get; set; }

        [Input("destinationDatabricksUpdateDataSourceRecommendedManagedTables")]
        public Input<Inputs.DestinationDatabricksConfigurationDataSourceDestinationDatabricksUpdateDataSourceRecommendedManagedTablesGetArgs>? DestinationDatabricksUpdateDataSourceRecommendedManagedTables { get; set; }

        public DestinationDatabricksConfigurationDataSourceGetArgs()
        {
        }
        public static new DestinationDatabricksConfigurationDataSourceGetArgs Empty => new DestinationDatabricksConfigurationDataSourceGetArgs();
    }
}
