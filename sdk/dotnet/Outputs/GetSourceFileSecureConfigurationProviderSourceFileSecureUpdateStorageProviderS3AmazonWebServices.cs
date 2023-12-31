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
    public sealed class GetSourceFileSecureConfigurationProviderSourceFileSecureUpdateStorageProviderS3AmazonWebServices
    {
        public readonly string AwsAccessKeyId;
        public readonly string AwsSecretAccessKey;
        public readonly string Storage;

        [OutputConstructor]
        private GetSourceFileSecureConfigurationProviderSourceFileSecureUpdateStorageProviderS3AmazonWebServices(
            string awsAccessKeyId,

            string awsSecretAccessKey,

            string storage)
        {
            AwsAccessKeyId = awsAccessKeyId;
            AwsSecretAccessKey = awsSecretAccessKey;
            Storage = storage;
        }
    }
}
