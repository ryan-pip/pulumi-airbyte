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
    public sealed class GetSourceMysqlConfiguration
    {
        public readonly string Database;
        public readonly string Host;
        public readonly string JdbcUrlParams;
        public readonly string Password;
        public readonly int Port;
        public readonly Outputs.GetSourceMysqlConfigurationReplicationMethod ReplicationMethod;
        public readonly string SourceType;
        public readonly Outputs.GetSourceMysqlConfigurationSslMode SslMode;
        public readonly Outputs.GetSourceMysqlConfigurationTunnelMethod TunnelMethod;
        public readonly string Username;

        [OutputConstructor]
        private GetSourceMysqlConfiguration(
            string database,

            string host,

            string jdbcUrlParams,

            string password,

            int port,

            Outputs.GetSourceMysqlConfigurationReplicationMethod replicationMethod,

            string sourceType,

            Outputs.GetSourceMysqlConfigurationSslMode sslMode,

            Outputs.GetSourceMysqlConfigurationTunnelMethod tunnelMethod,

            string username)
        {
            Database = database;
            Host = host;
            JdbcUrlParams = jdbcUrlParams;
            Password = password;
            Port = port;
            ReplicationMethod = replicationMethod;
            SourceType = sourceType;
            SslMode = sslMode;
            TunnelMethod = tunnelMethod;
            Username = username;
        }
    }
}