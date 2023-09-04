// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMssqlConfigurationTunnelMethodSourceMssqlSshTunnelMethodSshKeyAuthenticationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("sshKey", required: true)]
        public Input<string> SshKey { get; set; } = null!;

        [Input("tunnelHost", required: true)]
        public Input<string> TunnelHost { get; set; } = null!;

        [Input("tunnelMethod", required: true)]
        public Input<string> TunnelMethod { get; set; } = null!;

        [Input("tunnelPort", required: true)]
        public Input<int> TunnelPort { get; set; } = null!;

        [Input("tunnelUser", required: true)]
        public Input<string> TunnelUser { get; set; } = null!;

        public SourceMssqlConfigurationTunnelMethodSourceMssqlSshTunnelMethodSshKeyAuthenticationGetArgs()
        {
        }
        public static new SourceMssqlConfigurationTunnelMethodSourceMssqlSshTunnelMethodSshKeyAuthenticationGetArgs Empty => new SourceMssqlConfigurationTunnelMethodSourceMssqlSshTunnelMethodSshKeyAuthenticationGetArgs();
    }
}