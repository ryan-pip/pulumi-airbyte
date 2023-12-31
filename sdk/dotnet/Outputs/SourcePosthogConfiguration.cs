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
    public sealed class SourcePosthogConfiguration
    {
        public readonly string ApiKey;
        public readonly string? BaseUrl;
        public readonly string SourceType;
        public readonly string StartDate;

        [OutputConstructor]
        private SourcePosthogConfiguration(
            string apiKey,

            string? baseUrl,

            string sourceType,

            string startDate)
        {
            ApiKey = apiKey;
            BaseUrl = baseUrl;
            SourceType = sourceType;
            StartDate = startDate;
        }
    }
}
