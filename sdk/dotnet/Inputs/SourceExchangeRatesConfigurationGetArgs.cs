// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceExchangeRatesConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        [Input("base")]
        public Input<string>? Base { get; set; }

        [Input("ignoreWeekends")]
        public Input<bool>? IgnoreWeekends { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public SourceExchangeRatesConfigurationGetArgs()
        {
        }
        public static new SourceExchangeRatesConfigurationGetArgs Empty => new SourceExchangeRatesConfigurationGetArgs();
    }
}