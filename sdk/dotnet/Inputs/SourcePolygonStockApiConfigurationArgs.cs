// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourcePolygonStockApiConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("adjusted")]
        public Input<string>? Adjusted { get; set; }

        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        [Input("endDate", required: true)]
        public Input<string> EndDate { get; set; } = null!;

        [Input("limit")]
        public Input<int>? Limit { get; set; }

        [Input("multiplier", required: true)]
        public Input<int> Multiplier { get; set; } = null!;

        [Input("sort")]
        public Input<string>? Sort { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        [Input("stocksTicker", required: true)]
        public Input<string> StocksTicker { get; set; } = null!;

        [Input("timespan", required: true)]
        public Input<string> Timespan { get; set; } = null!;

        public SourcePolygonStockApiConfigurationArgs()
        {
        }
        public static new SourcePolygonStockApiConfigurationArgs Empty => new SourcePolygonStockApiConfigurationArgs();
    }
}
