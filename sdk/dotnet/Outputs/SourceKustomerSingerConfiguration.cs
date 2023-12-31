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
    public sealed class SourceKustomerSingerConfiguration
    {
        public readonly string ApiToken;
        public readonly string SourceType;
        public readonly string StartDate;

        [OutputConstructor]
        private SourceKustomerSingerConfiguration(
            string apiToken,

            string sourceType,

            string startDate)
        {
            ApiToken = apiToken;
            SourceType = sourceType;
            StartDate = startDate;
        }
    }
}
