// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSmartsheetsConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceSmartsheetsAuthorizationMethodApiAccessToken")]
        public Input<Inputs.SourceSmartsheetsConfigurationCredentialsSourceSmartsheetsAuthorizationMethodApiAccessTokenArgs>? SourceSmartsheetsAuthorizationMethodApiAccessToken { get; set; }

        [Input("sourceSmartsheetsAuthorizationMethodOAuth20")]
        public Input<Inputs.SourceSmartsheetsConfigurationCredentialsSourceSmartsheetsAuthorizationMethodOAuth20Args>? SourceSmartsheetsAuthorizationMethodOAuth20 { get; set; }

        [Input("sourceSmartsheetsUpdateAuthorizationMethodApiAccessToken")]
        public Input<Inputs.SourceSmartsheetsConfigurationCredentialsSourceSmartsheetsUpdateAuthorizationMethodApiAccessTokenArgs>? SourceSmartsheetsUpdateAuthorizationMethodApiAccessToken { get; set; }

        [Input("sourceSmartsheetsUpdateAuthorizationMethodOAuth20")]
        public Input<Inputs.SourceSmartsheetsConfigurationCredentialsSourceSmartsheetsUpdateAuthorizationMethodOAuth20Args>? SourceSmartsheetsUpdateAuthorizationMethodOAuth20 { get; set; }

        public SourceSmartsheetsConfigurationCredentialsArgs()
        {
        }
        public static new SourceSmartsheetsConfigurationCredentialsArgs Empty => new SourceSmartsheetsConfigurationCredentialsArgs();
    }
}
