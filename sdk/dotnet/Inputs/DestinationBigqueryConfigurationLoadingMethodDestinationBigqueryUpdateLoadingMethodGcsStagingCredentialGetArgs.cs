// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationBigqueryConfigurationLoadingMethodDestinationBigqueryUpdateLoadingMethodGcsStagingCredentialGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationBigqueryUpdateLoadingMethodGcsStagingCredentialHmacKey")]
        public Input<Inputs.DestinationBigqueryConfigurationLoadingMethodDestinationBigqueryUpdateLoadingMethodGcsStagingCredentialDestinationBigqueryUpdateLoadingMethodGcsStagingCredentialHmacKeyGetArgs>? DestinationBigqueryUpdateLoadingMethodGcsStagingCredentialHmacKey { get; set; }

        public DestinationBigqueryConfigurationLoadingMethodDestinationBigqueryUpdateLoadingMethodGcsStagingCredentialGetArgs()
        {
        }
        public static new DestinationBigqueryConfigurationLoadingMethodDestinationBigqueryUpdateLoadingMethodGcsStagingCredentialGetArgs Empty => new DestinationBigqueryConfigurationLoadingMethodDestinationBigqueryUpdateLoadingMethodGcsStagingCredentialGetArgs();
    }
}