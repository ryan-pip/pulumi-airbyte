// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMysqlConfigurationTunnelMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceMysqlSshTunnelMethodNoTunnel")]
        public Input<Inputs.SourceMysqlConfigurationTunnelMethodSourceMysqlSshTunnelMethodNoTunnelArgs>? SourceMysqlSshTunnelMethodNoTunnel { get; set; }

        [Input("sourceMysqlSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.SourceMysqlConfigurationTunnelMethodSourceMysqlSshTunnelMethodPasswordAuthenticationArgs>? SourceMysqlSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("sourceMysqlSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.SourceMysqlConfigurationTunnelMethodSourceMysqlSshTunnelMethodSshKeyAuthenticationArgs>? SourceMysqlSshTunnelMethodSshKeyAuthentication { get; set; }

        [Input("sourceMysqlUpdateSshTunnelMethodNoTunnel")]
        public Input<Inputs.SourceMysqlConfigurationTunnelMethodSourceMysqlUpdateSshTunnelMethodNoTunnelArgs>? SourceMysqlUpdateSshTunnelMethodNoTunnel { get; set; }

        [Input("sourceMysqlUpdateSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.SourceMysqlConfigurationTunnelMethodSourceMysqlUpdateSshTunnelMethodPasswordAuthenticationArgs>? SourceMysqlUpdateSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("sourceMysqlUpdateSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.SourceMysqlConfigurationTunnelMethodSourceMysqlUpdateSshTunnelMethodSshKeyAuthenticationArgs>? SourceMysqlUpdateSshTunnelMethodSshKeyAuthentication { get; set; }

        public SourceMysqlConfigurationTunnelMethodArgs()
        {
        }
        public static new SourceMysqlConfigurationTunnelMethodArgs Empty => new SourceMysqlConfigurationTunnelMethodArgs();
    }
}
