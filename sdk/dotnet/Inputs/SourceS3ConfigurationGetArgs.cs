// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceS3ConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dataset", required: true)]
        public Input<string> Dataset { get; set; } = null!;

        [Input("format")]
        public Input<Inputs.SourceS3ConfigurationFormatGetArgs>? Format { get; set; }

        [Input("pathPattern", required: true)]
        public Input<string> PathPattern { get; set; } = null!;

        [Input("provider", required: true)]
        public Input<Inputs.SourceS3ConfigurationProviderGetArgs> Provider { get; set; } = null!;

        [Input("schema")]
        public Input<string>? Schema { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        public SourceS3ConfigurationGetArgs()
        {
        }
        public static new SourceS3ConfigurationGetArgs Empty => new SourceS3ConfigurationGetArgs();
    }
}