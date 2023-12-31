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
    public sealed class SourceAmazonSqsConfiguration
    {
        public readonly string? AccessKey;
        public readonly string? AttributesToReturn;
        public readonly bool DeleteMessages;
        public readonly int? MaxBatchSize;
        public readonly int? MaxWaitTime;
        public readonly string QueueUrl;
        public readonly string Region;
        public readonly string? SecretKey;
        public readonly string SourceType;
        public readonly int? VisibilityTimeout;

        [OutputConstructor]
        private SourceAmazonSqsConfiguration(
            string? accessKey,

            string? attributesToReturn,

            bool deleteMessages,

            int? maxBatchSize,

            int? maxWaitTime,

            string queueUrl,

            string region,

            string? secretKey,

            string sourceType,

            int? visibilityTimeout)
        {
            AccessKey = accessKey;
            AttributesToReturn = attributesToReturn;
            DeleteMessages = deleteMessages;
            MaxBatchSize = maxBatchSize;
            MaxWaitTime = maxWaitTime;
            QueueUrl = queueUrl;
            Region = region;
            SecretKey = secretKey;
            SourceType = sourceType;
            VisibilityTimeout = visibilityTimeout;
        }
    }
}
