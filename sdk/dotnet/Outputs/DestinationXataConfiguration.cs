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
    public sealed class DestinationXataConfiguration
    {
        public readonly string ApiKey;
        public readonly string DbUrl;
        public readonly string DestinationType;

        [OutputConstructor]
        private DestinationXataConfiguration(
            string apiKey,

            string dbUrl,

            string destinationType)
        {
            ApiKey = apiKey;
            DbUrl = dbUrl;
            DestinationType = destinationType;
        }
    }
}