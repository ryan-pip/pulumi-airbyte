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
    public sealed class GetSourceNotionConfigurationCredentialsSourceNotionAuthenticateUsingAccessToken
    {
        public readonly string AuthType;
        public readonly string Token;

        [OutputConstructor]
        private GetSourceNotionConfigurationCredentialsSourceNotionAuthenticateUsingAccessToken(
            string authType,

            string token)
        {
            AuthType = authType;
            Token = token;
        }
    }
}