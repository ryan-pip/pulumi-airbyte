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
    public sealed class GetDestinationVerticaConfiguration
    {
        public readonly string Database;
        public readonly string DestinationType;
        public readonly string Host;
        public readonly string JdbcUrlParams;
        public readonly string Password;
        public readonly int Port;
        public readonly string Schema;
        public readonly Outputs.GetDestinationVerticaConfigurationTunnelMethod TunnelMethod;
        public readonly string Username;

        [OutputConstructor]
        private GetDestinationVerticaConfiguration(
            string database,

            string destinationType,

            string host,

            string jdbcUrlParams,

            string password,

            int port,

            string schema,

            Outputs.GetDestinationVerticaConfigurationTunnelMethod tunnelMethod,

            string username)
        {
            Database = database;
            DestinationType = destinationType;
            Host = host;
            JdbcUrlParams = jdbcUrlParams;
            Password = password;
            Port = port;
            Schema = schema;
            TunnelMethod = tunnelMethod;
            Username = username;
        }
    }
}