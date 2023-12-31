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
    public sealed class GetSourceGithubConfiguration
    {
        public readonly string Branch;
        public readonly Outputs.GetSourceGithubConfigurationCredentials Credentials;
        public readonly string Repository;
        public readonly int RequestsPerHour;
        public readonly string SourceType;
        public readonly string StartDate;

        [OutputConstructor]
        private GetSourceGithubConfiguration(
            string branch,

            Outputs.GetSourceGithubConfigurationCredentials credentials,

            string repository,

            int requestsPerHour,

            string sourceType,

            string startDate)
        {
            Branch = branch;
            Credentials = credentials;
            Repository = repository;
            RequestsPerHour = requestsPerHour;
            SourceType = sourceType;
            StartDate = startDate;
        }
    }
}
