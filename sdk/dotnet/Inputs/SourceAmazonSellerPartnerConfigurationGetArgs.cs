// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceAmazonSellerPartnerConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("advancedStreamOptions")]
        public Input<string>? AdvancedStreamOptions { get; set; }

        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        [Input("awsAccessKey")]
        public Input<string>? AwsAccessKey { get; set; }

        [Input("awsEnvironment", required: true)]
        public Input<string> AwsEnvironment { get; set; } = null!;

        [Input("awsSecretKey")]
        public Input<string>? AwsSecretKey { get; set; }

        [Input("lwaAppId", required: true)]
        public Input<string> LwaAppId { get; set; } = null!;

        [Input("lwaClientSecret", required: true)]
        public Input<string> LwaClientSecret { get; set; } = null!;

        [Input("maxWaitSeconds")]
        public Input<int>? MaxWaitSeconds { get; set; }

        [Input("periodInDays")]
        public Input<int>? PeriodInDays { get; set; }

        [Input("refreshToken", required: true)]
        public Input<string> RefreshToken { get; set; } = null!;

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("replicationEndDate")]
        public Input<string>? ReplicationEndDate { get; set; }

        [Input("replicationStartDate", required: true)]
        public Input<string> ReplicationStartDate { get; set; } = null!;

        [Input("reportOptions")]
        public Input<string>? ReportOptions { get; set; }

        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        public SourceAmazonSellerPartnerConfigurationGetArgs()
        {
        }
        public static new SourceAmazonSellerPartnerConfigurationGetArgs Empty => new SourceAmazonSellerPartnerConfigurationGetArgs();
    }
}