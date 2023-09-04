// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class GetSourceAwsCloudtrailConfiguration
    {
        public readonly string AwsKeyId;
        public readonly string AwsRegionName;
        public readonly string AwsSecretKey;
        public readonly string SourceType;
        public readonly string StartDate;

        [OutputConstructor]
        private GetSourceAwsCloudtrailConfiguration(
            string awsKeyId,

            string awsRegionName,

            string awsSecretKey,

            string sourceType,

            string startDate)
        {
            AwsKeyId = awsKeyId;
            AwsRegionName = awsRegionName;
            AwsSecretKey = awsSecretKey;
            SourceType = sourceType;
            StartDate = startDate;
        }
    }
}