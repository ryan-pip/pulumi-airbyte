// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationClickhouseConfigurationTunnelMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationClickhouseSshTunnelMethodNoTunnel")]
        public Input<Inputs.DestinationClickhouseConfigurationTunnelMethodDestinationClickhouseSshTunnelMethodNoTunnelArgs>? DestinationClickhouseSshTunnelMethodNoTunnel { get; set; }

        [Input("destinationClickhouseSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.DestinationClickhouseConfigurationTunnelMethodDestinationClickhouseSshTunnelMethodPasswordAuthenticationArgs>? DestinationClickhouseSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("destinationClickhouseSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.DestinationClickhouseConfigurationTunnelMethodDestinationClickhouseSshTunnelMethodSshKeyAuthenticationArgs>? DestinationClickhouseSshTunnelMethodSshKeyAuthentication { get; set; }

        [Input("destinationClickhouseUpdateSshTunnelMethodNoTunnel")]
        public Input<Inputs.DestinationClickhouseConfigurationTunnelMethodDestinationClickhouseUpdateSshTunnelMethodNoTunnelArgs>? DestinationClickhouseUpdateSshTunnelMethodNoTunnel { get; set; }

        [Input("destinationClickhouseUpdateSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.DestinationClickhouseConfigurationTunnelMethodDestinationClickhouseUpdateSshTunnelMethodPasswordAuthenticationArgs>? DestinationClickhouseUpdateSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("destinationClickhouseUpdateSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.DestinationClickhouseConfigurationTunnelMethodDestinationClickhouseUpdateSshTunnelMethodSshKeyAuthenticationArgs>? DestinationClickhouseUpdateSshTunnelMethodSshKeyAuthentication { get; set; }

        public DestinationClickhouseConfigurationTunnelMethodArgs()
        {
        }
        public static new DestinationClickhouseConfigurationTunnelMethodArgs Empty => new DestinationClickhouseConfigurationTunnelMethodArgs();
    }
}