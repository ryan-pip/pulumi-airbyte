// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGithubConfigurationCredentialsSourceGithubUpdateAuthenticationPersonalAccessTokenGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("optionTitle")]
        public Input<string>? OptionTitle { get; set; }

        [Input("personalAccessToken", required: true)]
        public Input<string> PersonalAccessToken { get; set; } = null!;

        public SourceGithubConfigurationCredentialsSourceGithubUpdateAuthenticationPersonalAccessTokenGetArgs()
        {
        }
        public static new SourceGithubConfigurationCredentialsSourceGithubUpdateAuthenticationPersonalAccessTokenGetArgs Empty => new SourceGithubConfigurationCredentialsSourceGithubUpdateAuthenticationPersonalAccessTokenGetArgs();
    }
}