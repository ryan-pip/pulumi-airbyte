// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationAWSDatalakeConfigurationFormatArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationAwsDatalakeOutputFormatWildcardJsonLinesNewlineDelimitedJson")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationFormatDestinationAwsDatalakeOutputFormatWildcardJsonLinesNewlineDelimitedJsonArgs>? DestinationAwsDatalakeOutputFormatWildcardJsonLinesNewlineDelimitedJson { get; set; }

        [Input("destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationFormatDestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageArgs>? DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage { get; set; }

        [Input("destinationAwsDatalakeUpdateOutputFormatWildcardJsonLinesNewlineDelimitedJson")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationFormatDestinationAwsDatalakeUpdateOutputFormatWildcardJsonLinesNewlineDelimitedJsonArgs>? DestinationAwsDatalakeUpdateOutputFormatWildcardJsonLinesNewlineDelimitedJson { get; set; }

        [Input("destinationAwsDatalakeUpdateOutputFormatWildcardParquetColumnarStorage")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationFormatDestinationAwsDatalakeUpdateOutputFormatWildcardParquetColumnarStorageArgs>? DestinationAwsDatalakeUpdateOutputFormatWildcardParquetColumnarStorage { get; set; }

        public DestinationAWSDatalakeConfigurationFormatArgs()
        {
        }
        public static new DestinationAWSDatalakeConfigurationFormatArgs Empty => new DestinationAWSDatalakeConfigurationFormatArgs();
    }
}
