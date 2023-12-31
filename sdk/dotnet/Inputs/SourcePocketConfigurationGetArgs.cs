// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourcePocketConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        public Input<string> AccessToken { get; set; } = null!;

        [Input("consumerKey", required: true)]
        public Input<string> ConsumerKey { get; set; } = null!;

        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        [Input("detailType")]
        public Input<string>? DetailType { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("favorite")]
        public Input<bool>? Favorite { get; set; }

        [Input("search")]
        public Input<string>? Search { get; set; }

        [Input("since")]
        public Input<string>? Since { get; set; }

        [Input("sort")]
        public Input<string>? Sort { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public SourcePocketConfigurationGetArgs()
        {
        }
        public static new SourcePocketConfigurationGetArgs Empty => new SourcePocketConfigurationGetArgs();
    }
}
