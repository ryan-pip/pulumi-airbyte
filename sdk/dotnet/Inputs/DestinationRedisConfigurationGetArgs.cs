// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationRedisConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("cacheType", required: true)]
        public Input<string> CacheType { get; set; } = null!;

        [Input("destinationType", required: true)]
        public Input<string> DestinationType { get; set; } = null!;

        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("ssl")]
        public Input<bool>? Ssl { get; set; }

        [Input("sslMode")]
        public Input<Inputs.DestinationRedisConfigurationSslModeGetArgs>? SslMode { get; set; }

        [Input("tunnelMethod")]
        public Input<Inputs.DestinationRedisConfigurationTunnelMethodGetArgs>? TunnelMethod { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public DestinationRedisConfigurationGetArgs()
        {
        }
        public static new DestinationRedisConfigurationGetArgs Empty => new DestinationRedisConfigurationGetArgs();
    }
}