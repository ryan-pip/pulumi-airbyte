// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationDevNullConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationType", required: true)]
        public Input<string> DestinationType { get; set; } = null!;

        [Input("testDestination", required: true)]
        public Input<Inputs.DestinationDevNullConfigurationTestDestinationArgs> TestDestination { get; set; } = null!;

        public DestinationDevNullConfigurationArgs()
        {
        }
        public static new DestinationDevNullConfigurationArgs Empty => new DestinationDevNullConfigurationArgs();
    }
}
