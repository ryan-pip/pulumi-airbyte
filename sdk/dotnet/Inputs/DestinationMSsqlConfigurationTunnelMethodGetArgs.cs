// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationMSsqlConfigurationTunnelMethodGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationMssqlSshTunnelMethodNoTunnel")]
        public Input<Inputs.DestinationMSsqlConfigurationTunnelMethodDestinationMssqlSshTunnelMethodNoTunnelGetArgs>? DestinationMssqlSshTunnelMethodNoTunnel { get; set; }

        [Input("destinationMssqlSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.DestinationMSsqlConfigurationTunnelMethodDestinationMssqlSshTunnelMethodPasswordAuthenticationGetArgs>? DestinationMssqlSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("destinationMssqlSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.DestinationMSsqlConfigurationTunnelMethodDestinationMssqlSshTunnelMethodSshKeyAuthenticationGetArgs>? DestinationMssqlSshTunnelMethodSshKeyAuthentication { get; set; }

        [Input("destinationMssqlUpdateSshTunnelMethodNoTunnel")]
        public Input<Inputs.DestinationMSsqlConfigurationTunnelMethodDestinationMssqlUpdateSshTunnelMethodNoTunnelGetArgs>? DestinationMssqlUpdateSshTunnelMethodNoTunnel { get; set; }

        [Input("destinationMssqlUpdateSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.DestinationMSsqlConfigurationTunnelMethodDestinationMssqlUpdateSshTunnelMethodPasswordAuthenticationGetArgs>? DestinationMssqlUpdateSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("destinationMssqlUpdateSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.DestinationMSsqlConfigurationTunnelMethodDestinationMssqlUpdateSshTunnelMethodSshKeyAuthenticationGetArgs>? DestinationMssqlUpdateSshTunnelMethodSshKeyAuthentication { get; set; }

        public DestinationMSsqlConfigurationTunnelMethodGetArgs()
        {
        }
        public static new DestinationMSsqlConfigurationTunnelMethodGetArgs Empty => new DestinationMSsqlConfigurationTunnelMethodGetArgs();
    }
}