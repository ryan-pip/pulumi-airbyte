// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleSearchConsoleConfigurationAuthorizationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceGoogleSearchConsoleAuthenticationTypeOAuth")]
        public Input<Inputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleAuthenticationTypeOAuthGetArgs>? SourceGoogleSearchConsoleAuthenticationTypeOAuth { get; set; }

        [Input("sourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication")]
        public Input<Inputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthenticationGetArgs>? SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication { get; set; }

        [Input("sourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth")]
        public Input<Inputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthGetArgs>? SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth { get; set; }

        [Input("sourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication")]
        public Input<Inputs.SourceGoogleSearchConsoleConfigurationAuthorizationSourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationGetArgs>? SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication { get; set; }

        public SourceGoogleSearchConsoleConfigurationAuthorizationGetArgs()
        {
        }
        public static new SourceGoogleSearchConsoleConfigurationAuthorizationGetArgs Empty => new SourceGoogleSearchConsoleConfigurationAuthorizationGetArgs();
    }
}
