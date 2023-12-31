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
    public sealed class DestinationRedisConfiguration
    {
        public readonly string CacheType;
        public readonly string DestinationType;
        public readonly string Host;
        public readonly string? Password;
        public readonly int Port;
        public readonly bool? Ssl;
        public readonly Outputs.DestinationRedisConfigurationSslMode? SslMode;
        public readonly Outputs.DestinationRedisConfigurationTunnelMethod? TunnelMethod;
        public readonly string Username;

        [OutputConstructor]
        private DestinationRedisConfiguration(
            string cacheType,

            string destinationType,

            string host,

            string? password,

            int port,

            bool? ssl,

            Outputs.DestinationRedisConfigurationSslMode? sslMode,

            Outputs.DestinationRedisConfigurationTunnelMethod? tunnelMethod,

            string username)
        {
            CacheType = cacheType;
            DestinationType = destinationType;
            Host = host;
            Password = password;
            Port = port;
            Ssl = ssl;
            SslMode = sslMode;
            TunnelMethod = tunnelMethod;
            Username = username;
        }
    }
}
