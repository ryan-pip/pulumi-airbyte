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
    public sealed class SourcePostgresConfigurationTunnelMethod
    {
        public readonly Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresSshTunnelMethodNoTunnel? SourcePostgresSshTunnelMethodNoTunnel;
        public readonly Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresSshTunnelMethodPasswordAuthentication? SourcePostgresSshTunnelMethodPasswordAuthentication;
        public readonly Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresSshTunnelMethodSshKeyAuthentication? SourcePostgresSshTunnelMethodSshKeyAuthentication;
        public readonly Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresUpdateSshTunnelMethodNoTunnel? SourcePostgresUpdateSshTunnelMethodNoTunnel;
        public readonly Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresUpdateSshTunnelMethodPasswordAuthentication? SourcePostgresUpdateSshTunnelMethodPasswordAuthentication;
        public readonly Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresUpdateSshTunnelMethodSshKeyAuthentication? SourcePostgresUpdateSshTunnelMethodSshKeyAuthentication;

        [OutputConstructor]
        private SourcePostgresConfigurationTunnelMethod(
            Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresSshTunnelMethodNoTunnel? sourcePostgresSshTunnelMethodNoTunnel,

            Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresSshTunnelMethodPasswordAuthentication? sourcePostgresSshTunnelMethodPasswordAuthentication,

            Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresSshTunnelMethodSshKeyAuthentication? sourcePostgresSshTunnelMethodSshKeyAuthentication,

            Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresUpdateSshTunnelMethodNoTunnel? sourcePostgresUpdateSshTunnelMethodNoTunnel,

            Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresUpdateSshTunnelMethodPasswordAuthentication? sourcePostgresUpdateSshTunnelMethodPasswordAuthentication,

            Outputs.SourcePostgresConfigurationTunnelMethodSourcePostgresUpdateSshTunnelMethodSshKeyAuthentication? sourcePostgresUpdateSshTunnelMethodSshKeyAuthentication)
        {
            SourcePostgresSshTunnelMethodNoTunnel = sourcePostgresSshTunnelMethodNoTunnel;
            SourcePostgresSshTunnelMethodPasswordAuthentication = sourcePostgresSshTunnelMethodPasswordAuthentication;
            SourcePostgresSshTunnelMethodSshKeyAuthentication = sourcePostgresSshTunnelMethodSshKeyAuthentication;
            SourcePostgresUpdateSshTunnelMethodNoTunnel = sourcePostgresUpdateSshTunnelMethodNoTunnel;
            SourcePostgresUpdateSshTunnelMethodPasswordAuthentication = sourcePostgresUpdateSshTunnelMethodPasswordAuthentication;
            SourcePostgresUpdateSshTunnelMethodSshKeyAuthentication = sourcePostgresUpdateSshTunnelMethodSshKeyAuthentication;
        }
    }
}
