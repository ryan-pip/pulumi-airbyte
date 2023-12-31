// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceClickupApiConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiToken", required: true)]
        public Input<string> ApiToken { get; set; } = null!;

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("includeClosedTasks")]
        public Input<bool>? IncludeClosedTasks { get; set; }

        [Input("listId")]
        public Input<string>? ListId { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public SourceClickupApiConfigurationGetArgs()
        {
        }
        public static new SourceClickupApiConfigurationGetArgs Empty => new SourceClickupApiConfigurationGetArgs();
    }
}
