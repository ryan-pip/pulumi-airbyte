// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsV4ConfigurationCredentialsSourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthenticationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        [Input("credentialsJson", required: true)]
        public Input<string> CredentialsJson { get; set; } = null!;

        public SourceGoogleAnalyticsV4ConfigurationCredentialsSourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthenticationGetArgs()
        {
        }
        public static new SourceGoogleAnalyticsV4ConfigurationCredentialsSourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthenticationGetArgs Empty => new SourceGoogleAnalyticsV4ConfigurationCredentialsSourceGoogleAnalyticsV4UpdateCredentialsServiceAccountKeyAuthenticationGetArgs();
    }
}
