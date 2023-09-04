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
    public sealed class GetSourcePinterestConfigurationCredentialsSourcePinterestUpdateAuthorizationMethodOAuth20
    {
        public readonly string AuthMethod;
        public readonly string ClientId;
        public readonly string ClientSecret;
        public readonly string RefreshToken;

        [OutputConstructor]
        private GetSourcePinterestConfigurationCredentialsSourcePinterestUpdateAuthorizationMethodOAuth20(
            string authMethod,

            string clientId,

            string clientSecret,

            string refreshToken)
        {
            AuthMethod = authMethod;
            ClientId = clientId;
            ClientSecret = clientSecret;
            RefreshToken = refreshToken;
        }
    }
}