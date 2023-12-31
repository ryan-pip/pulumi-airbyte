// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMixpanelConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributionWindow")]
        public Input<int>? AttributionWindow { get; set; }

        [Input("credentials")]
        public Input<Inputs.SourceMixpanelConfigurationCredentialsArgs>? Credentials { get; set; }

        [Input("dateWindowSize")]
        public Input<int>? DateWindowSize { get; set; }

        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("projectTimezone")]
        public Input<string>? ProjectTimezone { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("selectPropertiesByDefault")]
        public Input<bool>? SelectPropertiesByDefault { get; set; }

        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        public SourceMixpanelConfigurationArgs()
        {
        }
        public static new SourceMixpanelConfigurationArgs Empty => new SourceMixpanelConfigurationArgs();
    }
}
