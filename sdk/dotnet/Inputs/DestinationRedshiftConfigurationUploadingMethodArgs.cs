// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationRedshiftConfigurationUploadingMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationRedshiftUpdateUploadingMethodS3Staging")]
        public Input<Inputs.DestinationRedshiftConfigurationUploadingMethodDestinationRedshiftUpdateUploadingMethodS3StagingArgs>? DestinationRedshiftUpdateUploadingMethodS3Staging { get; set; }

        [Input("destinationRedshiftUpdateUploadingMethodStandard")]
        public Input<Inputs.DestinationRedshiftConfigurationUploadingMethodDestinationRedshiftUpdateUploadingMethodStandardArgs>? DestinationRedshiftUpdateUploadingMethodStandard { get; set; }

        [Input("destinationRedshiftUploadingMethodS3Staging")]
        public Input<Inputs.DestinationRedshiftConfigurationUploadingMethodDestinationRedshiftUploadingMethodS3StagingArgs>? DestinationRedshiftUploadingMethodS3Staging { get; set; }

        [Input("destinationRedshiftUploadingMethodStandard")]
        public Input<Inputs.DestinationRedshiftConfigurationUploadingMethodDestinationRedshiftUploadingMethodStandardArgs>? DestinationRedshiftUploadingMethodStandard { get; set; }

        public DestinationRedshiftConfigurationUploadingMethodArgs()
        {
        }
        public static new DestinationRedshiftConfigurationUploadingMethodArgs Empty => new DestinationRedshiftConfigurationUploadingMethodArgs();
    }
}
