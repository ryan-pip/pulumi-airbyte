// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        public SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceGetArgs()
        {
        }
        public static new SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceGetArgs Empty => new SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceGetArgs();
    }
}