// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationBigqueryDenormalizedConfigurationLoadingMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationBigqueryDenormalizedLoadingMethodGcsStaging")]
        public Input<Inputs.DestinationBigqueryDenormalizedConfigurationLoadingMethodDestinationBigqueryDenormalizedLoadingMethodGcsStagingArgs>? DestinationBigqueryDenormalizedLoadingMethodGcsStaging { get; set; }

        [Input("destinationBigqueryDenormalizedLoadingMethodStandardInserts")]
        public Input<Inputs.DestinationBigqueryDenormalizedConfigurationLoadingMethodDestinationBigqueryDenormalizedLoadingMethodStandardInsertsArgs>? DestinationBigqueryDenormalizedLoadingMethodStandardInserts { get; set; }

        [Input("destinationBigqueryDenormalizedUpdateLoadingMethodGcsStaging")]
        public Input<Inputs.DestinationBigqueryDenormalizedConfigurationLoadingMethodDestinationBigqueryDenormalizedUpdateLoadingMethodGcsStagingArgs>? DestinationBigqueryDenormalizedUpdateLoadingMethodGcsStaging { get; set; }

        [Input("destinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts")]
        public Input<Inputs.DestinationBigqueryDenormalizedConfigurationLoadingMethodDestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsArgs>? DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts { get; set; }

        public DestinationBigqueryDenormalizedConfigurationLoadingMethodArgs()
        {
        }
        public static new DestinationBigqueryDenormalizedConfigurationLoadingMethodArgs Empty => new DestinationBigqueryDenormalizedConfigurationLoadingMethodArgs();
    }
}
