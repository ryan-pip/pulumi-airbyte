// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationPostgresConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("destinationType", required: true)]
        public Input<string> DestinationType { get; set; } = null!;

        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("jdbcUrlParams")]
        public Input<string>? JdbcUrlParams { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        [Input("sslMode")]
        public Input<Inputs.DestinationPostgresConfigurationSslModeArgs>? SslMode { get; set; }

        [Input("tunnelMethod")]
        public Input<Inputs.DestinationPostgresConfigurationTunnelMethodArgs>? TunnelMethod { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public DestinationPostgresConfigurationArgs()
        {
        }
        public static new DestinationPostgresConfigurationArgs Empty => new DestinationPostgresConfigurationArgs();
    }
}