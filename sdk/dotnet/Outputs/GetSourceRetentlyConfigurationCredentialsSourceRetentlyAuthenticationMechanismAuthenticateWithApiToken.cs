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
    public sealed class GetSourceRetentlyConfigurationCredentialsSourceRetentlyAuthenticationMechanismAuthenticateWithApiToken
    {
        public readonly string? AdditionalProperties;
        public readonly string ApiKey;
        public readonly string AuthType;

        [OutputConstructor]
        private GetSourceRetentlyConfigurationCredentialsSourceRetentlyAuthenticationMechanismAuthenticateWithApiToken(
            string? additionalProperties,

            string apiKey,

            string authType)
        {
            AdditionalProperties = additionalProperties;
            ApiKey = apiKey;
            AuthType = authType;
        }
    }
}
