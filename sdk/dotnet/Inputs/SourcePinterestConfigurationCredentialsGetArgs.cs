// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourcePinterestConfigurationCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourcePinterestAuthorizationMethodAccessToken")]
        public Input<Inputs.SourcePinterestConfigurationCredentialsSourcePinterestAuthorizationMethodAccessTokenGetArgs>? SourcePinterestAuthorizationMethodAccessToken { get; set; }

        [Input("sourcePinterestAuthorizationMethodOAuth20")]
        public Input<Inputs.SourcePinterestConfigurationCredentialsSourcePinterestAuthorizationMethodOAuth20GetArgs>? SourcePinterestAuthorizationMethodOAuth20 { get; set; }

        [Input("sourcePinterestUpdateAuthorizationMethodAccessToken")]
        public Input<Inputs.SourcePinterestConfigurationCredentialsSourcePinterestUpdateAuthorizationMethodAccessTokenGetArgs>? SourcePinterestUpdateAuthorizationMethodAccessToken { get; set; }

        [Input("sourcePinterestUpdateAuthorizationMethodOAuth20")]
        public Input<Inputs.SourcePinterestConfigurationCredentialsSourcePinterestUpdateAuthorizationMethodOAuth20GetArgs>? SourcePinterestUpdateAuthorizationMethodOAuth20 { get; set; }

        public SourcePinterestConfigurationCredentialsGetArgs()
        {
        }
        public static new SourcePinterestConfigurationCredentialsGetArgs Empty => new SourcePinterestConfigurationCredentialsGetArgs();
    }
}
