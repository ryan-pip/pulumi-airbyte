// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class GetDestinationVerticaConfigurationTunnelMethod
    {
        public readonly Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaSshTunnelMethodNoTunnel DestinationVerticaSshTunnelMethodNoTunnel;
        public readonly Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaSshTunnelMethodPasswordAuthentication DestinationVerticaSshTunnelMethodPasswordAuthentication;
        public readonly Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaSshTunnelMethodSshKeyAuthentication DestinationVerticaSshTunnelMethodSshKeyAuthentication;
        public readonly Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaUpdateSshTunnelMethodNoTunnel DestinationVerticaUpdateSshTunnelMethodNoTunnel;
        public readonly Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaUpdateSshTunnelMethodPasswordAuthentication DestinationVerticaUpdateSshTunnelMethodPasswordAuthentication;
        public readonly Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaUpdateSshTunnelMethodSshKeyAuthentication DestinationVerticaUpdateSshTunnelMethodSshKeyAuthentication;

        [OutputConstructor]
        private GetDestinationVerticaConfigurationTunnelMethod(
            Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaSshTunnelMethodNoTunnel destinationVerticaSshTunnelMethodNoTunnel,

            Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaSshTunnelMethodPasswordAuthentication destinationVerticaSshTunnelMethodPasswordAuthentication,

            Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaSshTunnelMethodSshKeyAuthentication destinationVerticaSshTunnelMethodSshKeyAuthentication,

            Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaUpdateSshTunnelMethodNoTunnel destinationVerticaUpdateSshTunnelMethodNoTunnel,

            Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaUpdateSshTunnelMethodPasswordAuthentication destinationVerticaUpdateSshTunnelMethodPasswordAuthentication,

            Outputs.GetDestinationVerticaConfigurationTunnelMethodDestinationVerticaUpdateSshTunnelMethodSshKeyAuthentication destinationVerticaUpdateSshTunnelMethodSshKeyAuthentication)
        {
            DestinationVerticaSshTunnelMethodNoTunnel = destinationVerticaSshTunnelMethodNoTunnel;
            DestinationVerticaSshTunnelMethodPasswordAuthentication = destinationVerticaSshTunnelMethodPasswordAuthentication;
            DestinationVerticaSshTunnelMethodSshKeyAuthentication = destinationVerticaSshTunnelMethodSshKeyAuthentication;
            DestinationVerticaUpdateSshTunnelMethodNoTunnel = destinationVerticaUpdateSshTunnelMethodNoTunnel;
            DestinationVerticaUpdateSshTunnelMethodPasswordAuthentication = destinationVerticaUpdateSshTunnelMethodPasswordAuthentication;
            DestinationVerticaUpdateSshTunnelMethodSshKeyAuthentication = destinationVerticaUpdateSshTunnelMethodSshKeyAuthentication;
        }
    }
}
