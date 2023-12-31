// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceFreshdeskConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("requestsPerMinute")]
        public Input<int>? RequestsPerMinute { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        public SourceFreshdeskConfigurationGetArgs()
        {
        }
        public static new SourceFreshdeskConfigurationGetArgs Empty => new SourceFreshdeskConfigurationGetArgs();
    }
}
