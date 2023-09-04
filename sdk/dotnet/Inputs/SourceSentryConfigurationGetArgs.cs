// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSentryConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authToken", required: true)]
        public Input<string> AuthToken { get; set; } = null!;

        [Input("discoverFields")]
        private InputList<string>? _discoverFields;
        public InputList<string> DiscoverFields
        {
            get => _discoverFields ?? (_discoverFields = new InputList<string>());
            set => _discoverFields = value;
        }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        public SourceSentryConfigurationGetArgs()
        {
        }
        public static new SourceSentryConfigurationGetArgs Empty => new SourceSentryConfigurationGetArgs();
    }
}