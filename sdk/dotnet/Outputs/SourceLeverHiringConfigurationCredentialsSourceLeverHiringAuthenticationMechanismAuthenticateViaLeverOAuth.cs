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
    public sealed class SourceLeverHiringConfigurationCredentialsSourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth
    {
        public readonly string? AuthType;
        public readonly string? ClientId;
        public readonly string? ClientSecret;
        public readonly string RefreshToken;

        [OutputConstructor]
        private SourceLeverHiringConfigurationCredentialsSourceLeverHiringAuthenticationMechanismAuthenticateViaLeverOAuth(
            string? authType,

            string? clientId,

            string? clientSecret,

            string refreshToken)
        {
            AuthType = authType;
            ClientId = clientId;
            ClientSecret = clientSecret;
            RefreshToken = refreshToken;
        }
    }
}
