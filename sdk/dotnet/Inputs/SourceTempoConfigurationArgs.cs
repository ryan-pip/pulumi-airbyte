// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceTempoConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiToken", required: true)]
        public Input<string> ApiToken { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        public SourceTempoConfigurationArgs()
        {
        }
        public static new SourceTempoConfigurationArgs Empty => new SourceTempoConfigurationArgs();
    }
}