// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceTwitterConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        public SourceTwitterConfigurationGetArgs()
        {
        }
        public static new SourceTwitterConfigurationGetArgs Empty => new SourceTwitterConfigurationGetArgs();
    }
}