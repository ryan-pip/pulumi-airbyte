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
    public sealed class SourceGoogleAnalyticsDataApiConfigurationCredentialsSourceGoogleAnalyticsDataApiUpdateCredentialsServiceAccountKeyAuthentication
    {
        public readonly string? AuthType;
        public readonly string CredentialsJson;

        [OutputConstructor]
        private SourceGoogleAnalyticsDataApiConfigurationCredentialsSourceGoogleAnalyticsDataApiUpdateCredentialsServiceAccountKeyAuthentication(
            string? authType,

            string credentialsJson)
        {
            AuthType = authType;
            CredentialsJson = credentialsJson;
        }
    }
}