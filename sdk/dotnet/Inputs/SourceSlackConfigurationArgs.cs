// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSlackConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("channelFilters")]
        private InputList<string>? _channelFilters;
        public InputList<string> ChannelFilters
        {
            get => _channelFilters ?? (_channelFilters = new InputList<string>());
            set => _channelFilters = value;
        }

        [Input("credentials")]
        public Input<Inputs.SourceSlackConfigurationCredentialsArgs>? Credentials { get; set; }

        [Input("joinChannels", required: true)]
        public Input<bool> JoinChannels { get; set; } = null!;

        [Input("lookbackWindow", required: true)]
        public Input<int> LookbackWindow { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public SourceSlackConfigurationArgs()
        {
        }
        public static new SourceSlackConfigurationArgs Empty => new SourceSlackConfigurationArgs();
    }
}
