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
    public sealed class DestinationOracleConfigurationTunnelMethod
    {
        public readonly Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleSshTunnelMethodNoTunnel? DestinationOracleSshTunnelMethodNoTunnel;
        public readonly Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleSshTunnelMethodPasswordAuthentication? DestinationOracleSshTunnelMethodPasswordAuthentication;
        public readonly Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleSshTunnelMethodSshKeyAuthentication? DestinationOracleSshTunnelMethodSshKeyAuthentication;
        public readonly Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodNoTunnel? DestinationOracleUpdateSshTunnelMethodNoTunnel;
        public readonly Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodPasswordAuthentication? DestinationOracleUpdateSshTunnelMethodPasswordAuthentication;
        public readonly Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodSshKeyAuthentication? DestinationOracleUpdateSshTunnelMethodSshKeyAuthentication;

        [OutputConstructor]
        private DestinationOracleConfigurationTunnelMethod(
            Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleSshTunnelMethodNoTunnel? destinationOracleSshTunnelMethodNoTunnel,

            Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleSshTunnelMethodPasswordAuthentication? destinationOracleSshTunnelMethodPasswordAuthentication,

            Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleSshTunnelMethodSshKeyAuthentication? destinationOracleSshTunnelMethodSshKeyAuthentication,

            Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodNoTunnel? destinationOracleUpdateSshTunnelMethodNoTunnel,

            Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodPasswordAuthentication? destinationOracleUpdateSshTunnelMethodPasswordAuthentication,

            Outputs.DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodSshKeyAuthentication? destinationOracleUpdateSshTunnelMethodSshKeyAuthentication)
        {
            DestinationOracleSshTunnelMethodNoTunnel = destinationOracleSshTunnelMethodNoTunnel;
            DestinationOracleSshTunnelMethodPasswordAuthentication = destinationOracleSshTunnelMethodPasswordAuthentication;
            DestinationOracleSshTunnelMethodSshKeyAuthentication = destinationOracleSshTunnelMethodSshKeyAuthentication;
            DestinationOracleUpdateSshTunnelMethodNoTunnel = destinationOracleUpdateSshTunnelMethodNoTunnel;
            DestinationOracleUpdateSshTunnelMethodPasswordAuthentication = destinationOracleUpdateSshTunnelMethodPasswordAuthentication;
            DestinationOracleUpdateSshTunnelMethodSshKeyAuthentication = destinationOracleUpdateSshTunnelMethodSshKeyAuthentication;
        }
    }
}
