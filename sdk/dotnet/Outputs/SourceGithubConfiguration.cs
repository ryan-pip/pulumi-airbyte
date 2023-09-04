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
    public sealed class SourceGithubConfiguration
    {
        public readonly string? Branch;
        public readonly Outputs.SourceGithubConfigurationCredentials? Credentials;
        public readonly string Repository;
        public readonly int? RequestsPerHour;
        public readonly string SourceType;
        public readonly string StartDate;

        [OutputConstructor]
        private SourceGithubConfiguration(
            string? branch,

            Outputs.SourceGithubConfigurationCredentials? credentials,

            string repository,

            int? requestsPerHour,

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