// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationOracleConfigurationTunnelMethodGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationOracleSshTunnelMethodNoTunnel")]
        public Input<Inputs.DestinationOracleConfigurationTunnelMethodDestinationOracleSshTunnelMethodNoTunnelGetArgs>? DestinationOracleSshTunnelMethodNoTunnel { get; set; }

        [Input("destinationOracleSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.DestinationOracleConfigurationTunnelMethodDestinationOracleSshTunnelMethodPasswordAuthenticationGetArgs>? DestinationOracleSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("destinationOracleSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.DestinationOracleConfigurationTunnelMethodDestinationOracleSshTunnelMethodSshKeyAuthenticationGetArgs>? DestinationOracleSshTunnelMethodSshKeyAuthentication { get; set; }

        [Input("destinationOracleUpdateSshTunnelMethodNoTunnel")]
        public Input<Inputs.DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodNoTunnelGetArgs>? DestinationOracleUpdateSshTunnelMethodNoTunnel { get; set; }

        [Input("destinationOracleUpdateSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodPasswordAuthenticationGetArgs>? DestinationOracleUpdateSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("destinationOracleUpdateSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodSshKeyAuthenticationGetArgs>? DestinationOracleUpdateSshTunnelMethodSshKeyAuthentication { get; set; }

        public DestinationOracleConfigurationTunnelMethodGetArgs()
        {
        }
        public static new DestinationOracleConfigurationTunnelMethodGetArgs Empty => new DestinationOracleConfigurationTunnelMethodGetArgs();
    }
}