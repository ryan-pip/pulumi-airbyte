// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeMongoDbAtlasArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        public Input<string>? AdditionalProperties { get; set; }

        [Input("clusterUrl", required: true)]
        public Input<string> ClusterUrl { get; set; } = null!;

        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        public SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeMongoDbAtlasArgs()
        {
        }
        public static new SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeMongoDbAtlasArgs Empty => new SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeMongoDbAtlasArgs();
    }
}
