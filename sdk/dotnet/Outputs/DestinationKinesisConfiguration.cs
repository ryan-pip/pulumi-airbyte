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
    public sealed class DestinationKinesisConfiguration
    {
        public readonly string AccessKey;
        public readonly int BufferSize;
        public readonly string DestinationType;
        public readonly string Endpoint;
        public readonly string PrivateKey;
        public readonly string Region;
        public readonly int ShardCount;

        [OutputConstructor]
        private DestinationKinesisConfiguration(
            string accessKey,

            int bufferSize,

            string destinationType,

            string endpoint,

            string privateKey,

            string region,

            int shardCount)
        {
            AccessKey = accessKey;
            BufferSize = bufferSize;
            DestinationType = destinationType;
            Endpoint = endpoint;
            PrivateKey = privateKey;
            Region = region;
            ShardCount = shardCount;
        }
    }
}
