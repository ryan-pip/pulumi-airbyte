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
    public sealed class GetSourceGoogleWorkspaceAdminReportsConfiguration
    {
        public readonly string CredentialsJson;
        public readonly string Email;
        public readonly int Lookback;
        public readonly string SourceType;

        [OutputConstructor]
        private GetSourceGoogleWorkspaceAdminReportsConfiguration(
            string credentialsJson,

            string email,

            int lookback,

            string sourceType)
        {
            CredentialsJson = credentialsJson;
            Email = email;
            Lookback = lookback;
            SourceType = sourceType;
        }
    }
}
