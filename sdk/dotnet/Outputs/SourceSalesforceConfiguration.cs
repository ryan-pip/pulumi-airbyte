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
    public sealed class SourceSalesforceConfiguration
    {
        public readonly string? AuthType;
        public readonly string ClientId;
        public readonly string ClientSecret;
        public readonly bool? ForceUseBulkApi;
        public readonly bool? IsSandbox;
        public readonly string RefreshToken;
        public readonly string SourceType;
        public readonly string? StartDate;
        public readonly ImmutableArray<Outputs.SourceSalesforceConfigurationStreamsCriteria> StreamsCriterias;

        [OutputConstructor]
        private SourceSalesforceConfiguration(
            string? authType,

            string clientId,

            string clientSecret,

            bool? forceUseBulkApi,

            bool? isSandbox,

            string refreshToken,

            string sourceType,

            string? startDate,

            ImmutableArray<Outputs.SourceSalesforceConfigurationStreamsCriteria> streamsCriterias)
        {
            AuthType = authType;
            ClientId = clientId;
            ClientSecret = clientSecret;
            ForceUseBulkApi = forceUseBulkApi;
            IsSandbox = isSandbox;
            RefreshToken = refreshToken;
            SourceType = sourceType;
            StartDate = startDate;
            StreamsCriterias = streamsCriterias;
        }
    }
}