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
    public sealed class SourceBambooHrConfiguration
    {
        public readonly string ApiKey;
        public readonly string? CustomReportsFields;
        public readonly bool? CustomReportsIncludeDefaultFields;
        public readonly string SourceType;
        public readonly string Subdomain;

        [OutputConstructor]
        private SourceBambooHrConfiguration(
            string apiKey,

            string? customReportsFields,

            bool? customReportsIncludeDefaultFields,

            string sourceType,

            string subdomain)
        {
            ApiKey = apiKey;
            CustomReportsFields = customReportsFields;
            CustomReportsIncludeDefaultFields = customReportsIncludeDefaultFields;
            SourceType = sourceType;
            Subdomain = subdomain;
        }
    }
}
