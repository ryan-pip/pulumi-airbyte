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
    public sealed class SourceGoogleSearchConsoleConfigurationAuthorization
    {
        public readonly Outputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleAuthenticationTypeOAuth? SourceGoogleSearchConsoleAuthenticationTypeOAuth;
        public readonly Outputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication? SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication;
        public readonly Outputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth? SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth;
        public readonly Outputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication? SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication;

        [OutputConstructor]
        private SourceGoogleSearchConsoleConfigurationAuthorization(
            Outputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleAuthenticationTypeOAuth? sourceGoogleSearchConsoleAuthenticationTypeOAuth,

            Outputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication? sourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication,

            Outputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth? sourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth,

            Outputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication? sourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication)
        {
            SourceGoogleSearchConsoleAuthenticationTypeOAuth = sourceGoogleSearchConsoleAuthenticationTypeOAuth;
            SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication = sourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication;
            SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth = sourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth;
            SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication = sourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication;
        }
    }
}
