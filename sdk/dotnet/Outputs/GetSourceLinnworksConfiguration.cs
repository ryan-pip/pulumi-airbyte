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
    public sealed class GetSourceLinnworksConfiguration
    {
        public readonly string ApplicationId;
        public readonly string ApplicationSecret;
        public readonly string SourceType;
        public readonly string StartDate;
        public readonly string Token;

        [OutputConstructor]
        private GetSourceLinnworksConfiguration(
            string applicationId,

            string applicationSecret,

            string sourceType,

            string startDate,

            string token)
        {
            ApplicationId = applicationId;
            ApplicationSecret = applicationSecret;
            SourceType = sourceType;
            StartDate = startDate;
            Token = token;
        }
    }
}
