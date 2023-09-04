// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMssqlConfigurationTunnelMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceMssqlSshTunnelMethodNoTunnel")]
        public Input<Inputs.SourceMssqlConfigurationTunnelMethodSourceMssqlSshTunnelMethodNoTunnelArgs>? SourceMssqlSshTunnelMethodNoTunnel { get; set; }

        [Input("sourceMssqlSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.SourceMssqlConfigurationTunnelMethodSourceMssqlSshTunnelMethodPasswordAuthenticationArgs>? SourceMssqlSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("sourceMssqlSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.SourceMssqlConfigurationTunnelMethodSourceMssqlSshTunnelMethodSshKeyAuthenticationArgs>? SourceMssqlSshTunnelMethodSshKeyAuthentication { get; set; }

        [Input("sourceMssqlUpdateSshTunnelMethodNoTunnel")]
        public Input<Inputs.SourceMssqlConfigurationTunnelMethodSourceMssqlUpdateSshTunnelMethodNoTunnelArgs>? SourceMssqlUpdateSshTunnelMethodNoTunnel { get; set; }

        [Input("sourceMssqlUpdateSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.SourceMssqlConfigurationTunnelMethodSourceMssqlUpdateSshTunnelMethodPasswordAuthenticationArgs>? SourceMssqlUpdateSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("sourceMssqlUpdateSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.SourceMssqlConfigurationTunnelMethodSourceMssqlUpdateSshTunnelMethodSshKeyAuthenticationArgs>? SourceMssqlUpdateSshTunnelMethodSshKeyAuthentication { get; set; }

        public SourceMssqlConfigurationTunnelMethodArgs()
        {
        }
        public static new SourceMssqlConfigurationTunnelMethodArgs Empty => new SourceMssqlConfigurationTunnelMethodArgs();
    }
}