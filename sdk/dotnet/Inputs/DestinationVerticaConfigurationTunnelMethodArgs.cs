// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationVerticaConfigurationTunnelMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationVerticaSshTunnelMethodNoTunnel")]
        public Input<Inputs.DestinationVerticaConfigurationTunnelMethodDestinationVerticaSshTunnelMethodNoTunnelArgs>? DestinationVerticaSshTunnelMethodNoTunnel { get; set; }

        [Input("destinationVerticaSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.DestinationVerticaConfigurationTunnelMethodDestinationVerticaSshTunnelMethodPasswordAuthenticationArgs>? DestinationVerticaSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("destinationVerticaSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.DestinationVerticaConfigurationTunnelMethodDestinationVerticaSshTunnelMethodSshKeyAuthenticationArgs>? DestinationVerticaSshTunnelMethodSshKeyAuthentication { get; set; }

        [Input("destinationVerticaUpdateSshTunnelMethodNoTunnel")]
        public Input<Inputs.DestinationVerticaConfigurationTunnelMethodDestinationVerticaUpdateSshTunnelMethodNoTunnelArgs>? DestinationVerticaUpdateSshTunnelMethodNoTunnel { get; set; }

        [Input("destinationVerticaUpdateSshTunnelMethodPasswordAuthentication")]
        public Input<Inputs.DestinationVerticaConfigurationTunnelMethodDestinationVerticaUpdateSshTunnelMethodPasswordAuthenticationArgs>? DestinationVerticaUpdateSshTunnelMethodPasswordAuthentication { get; set; }

        [Input("destinationVerticaUpdateSshTunnelMethodSshKeyAuthentication")]
        public Input<Inputs.DestinationVerticaConfigurationTunnelMethodDestinationVerticaUpdateSshTunnelMethodSshKeyAuthenticationArgs>? DestinationVerticaUpdateSshTunnelMethodSshKeyAuthentication { get; set; }

        public DestinationVerticaConfigurationTunnelMethodArgs()
        {
        }
        public static new DestinationVerticaConfigurationTunnelMethodArgs Empty => new DestinationVerticaConfigurationTunnelMethodArgs();
    }
}