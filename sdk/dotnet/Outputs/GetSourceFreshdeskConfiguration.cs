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
    public sealed class GetSourceFreshdeskConfiguration
    {
        public readonly string ApiKey;
        public readonly string Domain;
        public readonly int RequestsPerMinute;
        public readonly string SourceType;
        public readonly string StartDate;

        [OutputConstructor]
        private GetSourceFreshdeskConfiguration(
            string apiKey,

            string domain,

            int requestsPerMinute,

            string sourceType,

            string startDate)
        {
            ApiKey = apiKey;
            Domain = domain;
            RequestsPerMinute = requestsPerMinute;
            SourceType = sourceType;
            StartDate = startDate;
        }
    }
}