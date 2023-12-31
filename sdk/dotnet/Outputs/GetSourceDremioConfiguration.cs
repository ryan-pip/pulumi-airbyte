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
    public sealed class GetSourceDremioConfiguration
    {
        public readonly string ApiKey;
        public readonly string BaseUrl;
        public readonly string SourceType;

        [OutputConstructor]
        private GetSourceDremioConfiguration(
            string apiKey,

            string baseUrl,

            string sourceType)
        {
            ApiKey = apiKey;
            BaseUrl = baseUrl;
            SourceType = sourceType;
        }
    }
}
