// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceAppfollowConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiSecret")]
        public Input<string>? ApiSecret { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        public SourceAppfollowConfigurationGetArgs()
        {
        }
        public static new SourceAppfollowConfigurationGetArgs Empty => new SourceAppfollowConfigurationGetArgs();
    }
}
