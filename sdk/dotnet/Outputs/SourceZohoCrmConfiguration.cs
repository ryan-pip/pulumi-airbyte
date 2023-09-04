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
    public sealed class SourceZohoCrmConfiguration
    {
        public readonly string ClientId;
        public readonly string ClientSecret;
        public readonly string DcRegion;
        public readonly string Edition;
        public readonly string Environment;
        public readonly string RefreshToken;
        public readonly string SourceType;
        public readonly string? StartDatetime;

        [OutputConstructor]
        private SourceZohoCrmConfiguration(
            string clientId,

            string clientSecret,

            string dcRegion,

            string edition,

            string environment,

            string refreshToken,

            string sourceType,

            string? startDatetime)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            DcRegion = dcRegion;
            Edition = edition;
            Environment = environment;
            RefreshToken = refreshToken;
            SourceType = sourceType;
            StartDatetime = startDatetime;
        }
    }
}