// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourcePinterestConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourcePinterestAuthorizationMethodAccessToken")]
        public Input<Inputs.SourcePinterestConfigurationCredentialsSourcePinterestAuthorizationMethodAccessTokenArgs>? SourcePinterestAuthorizationMethodAccessToken { get; set; }

        [Input("sourcePinterestAuthorizationMethodOAuth20")]
        public Input<Inputs.SourcePinterestConfigurationCredentialsSourcePinterestAuthorizationMethodOAuth20Args>? SourcePinterestAuthorizationMethodOAuth20 { get; set; }

        [Input("sourcePinterestUpdateAuthorizationMethodAccessToken")]
        public Input<Inputs.SourcePinterestConfigurationCredentialsSourcePinterestUpdateAuthorizationMethodAccessTokenArgs>? SourcePinterestUpdateAuthorizationMethodAccessToken { get; set; }

        [Input("sourcePinterestUpdateAuthorizationMethodOAuth20")]
        public Input<Inputs.SourcePinterestConfigurationCredentialsSourcePinterestUpdateAuthorizationMethodOAuth20Args>? SourcePinterestUpdateAuthorizationMethodOAuth20 { get; set; }

        public SourcePinterestConfigurationCredentialsArgs()
        {
        }
        public static new SourcePinterestConfigurationCredentialsArgs Empty => new SourcePinterestConfigurationCredentialsArgs();
    }
}
