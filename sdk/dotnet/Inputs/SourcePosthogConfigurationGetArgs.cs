// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourcePosthogConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public SourcePosthogConfigurationGetArgs()
        {
        }
        public static new SourcePosthogConfigurationGetArgs Empty => new SourcePosthogConfigurationGetArgs();
    }
}