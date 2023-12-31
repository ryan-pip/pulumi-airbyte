// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceZenloopConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiToken", required: true)]
        public Input<string> ApiToken { get; set; } = null!;

        [Input("dateFrom")]
        public Input<string>? DateFrom { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("surveyGroupId")]
        public Input<string>? SurveyGroupId { get; set; }

        [Input("surveyId")]
        public Input<string>? SurveyId { get; set; }

        public SourceZenloopConfigurationArgs()
        {
        }
        public static new SourceZenloopConfigurationArgs Empty => new SourceZenloopConfigurationArgs();
    }
}
