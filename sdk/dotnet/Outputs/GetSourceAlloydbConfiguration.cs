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
    public sealed class GetSourceAlloydbConfiguration
    {
        public readonly string Database;
        public readonly string Host;
        public readonly string JdbcUrlParams;
        public readonly string Password;
        public readonly int Port;
        public readonly Outputs.GetSourceAlloydbConfigurationReplicationMethod ReplicationMethod;
        public readonly ImmutableArray<string> Schemas;
        public readonly string SourceType;
        public readonly Outputs.GetSourceAlloydbConfigurationSslMode SslMode;
        public readonly Outputs.GetSourceAlloydbConfigurationTunnelMethod TunnelMethod;
        public readonly string Username;

        [OutputConstructor]
        private GetSourceAlloydbConfiguration(
            string database,

            string host,

            string jdbcUrlParams,

            string password,

            int port,

            Outputs.GetSourceAlloydbConfigurationReplicationMethod replicationMethod,

            ImmutableArray<string> schemas,

            string sourceType,

            Outputs.GetSourceAlloydbConfigurationSslMode sslMode,

            Outputs.GetSourceAlloydbConfigurationTunnelMethod tunnelMethod,

            string username)
        {
            Database = database;
            Host = host;
            JdbcUrlParams = jdbcUrlParams;
            Password = password;
            Port = port;
            ReplicationMethod = replicationMethod;
            Schemas = schemas;
            SourceType = sourceType;
            SslMode = sslMode;
            TunnelMethod = tunnelMethod;
            Username = username;
        }
    }
}