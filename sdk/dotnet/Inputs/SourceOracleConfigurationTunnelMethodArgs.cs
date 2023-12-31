// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceOracleConfigurationTunnelMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceOracleSshTunnelMethodNoTunnel")]
        public Input<Inputs.SourceOracleConfigurationTunnelMethodSourceOracleSshTunnelMethodNoTunnelArgs>? SourceOracleSshTunnelMethodNoTunnel { get; set; }

        [Input("sourceOracleSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.SourceOracleConfigurationTunnelMethodSourceOracleSshTunnelMethodPasswordAuthenticationArgs>? SourceOracleSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("sourceOracleSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.SourceOracleConfigurationTunnelMethodSourceOracleSshTunnelMethodSshKeyAuthenticationArgs>? SourceOracleSshTunnelMethodSshKeyAuthentication { get; set; }

        [Input("sourceOracleUpdateSshTunnelMethodNoTunnel")]
        public Input<Inputs.SourceOracleConfigurationTunnelMethodSourceOracleUpdateSshTunnelMethodNoTunnelArgs>? SourceOracleUpdateSshTunnelMethodNoTunnel { get; set; }

        [Input("sourceOracleUpdateSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.SourceOracleConfigurationTunnelMethodSourceOracleUpdateSshTunnelMethodPasswordAuthenticationArgs>? SourceOracleUpdateSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("sourceOracleUpdateSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.SourceOracleConfigurationTunnelMethodSourceOracleUpdateSshTunnelMethodSshKeyAuthenticationArgs>? SourceOracleUpdateSshTunnelMethodSshKeyAuthentication { get; set; }

        public SourceOracleConfigurationTunnelMethodArgs()
        {
        }
        public static new SourceOracleConfigurationTunnelMethodArgs Empty => new SourceOracleConfigurationTunnelMethodArgs();
    }
}
