// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceFacebookMarketingConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        public Input<string> AccessToken { get; set; } = null!;

        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        [Input("actionBreakdownsAllowEmpty")]
        public Input<bool>? ActionBreakdownsAllowEmpty { get; set; }

        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        [Input("customInsights")]
        private InputList<Inputs.SourceFacebookMarketingConfigurationCustomInsightGetArgs>? _customInsights;
        public InputList<Inputs.SourceFacebookMarketingConfigurationCustomInsightGetArgs> CustomInsights
        {
            get => _customInsights ?? (_customInsights = new InputList<Inputs.SourceFacebookMarketingConfigurationCustomInsightGetArgs>());
            set => _customInsights = value;
        }

        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        [Input("fetchThumbnailImages")]
        public Input<bool>? FetchThumbnailImages { get; set; }

        [Input("includeDeleted")]
        public Input<bool>? IncludeDeleted { get; set; }

        [Input("insightsLookbackWindow")]
        public Input<int>? InsightsLookbackWindow { get; set; }

        [Input("maxBatchSize")]
        public Input<int>? MaxBatchSize { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public SourceFacebookMarketingConfigurationGetArgs()
        {
        }
        public static new SourceFacebookMarketingConfigurationGetArgs Empty => new SourceFacebookMarketingConfigurationGetArgs();
    }
}