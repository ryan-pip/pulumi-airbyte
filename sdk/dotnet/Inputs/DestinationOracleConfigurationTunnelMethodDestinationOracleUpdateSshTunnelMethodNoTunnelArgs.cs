// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodNoTunnelArgs : global::Pulumi.ResourceArgs
    {
        [Input("tunnelMethod", required: true)]
        public Input<string> TunnelMethod { get; set; } = null!;

        public DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodNoTunnelArgs()
        {
        }
        public static new DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodNoTunnelArgs Empty => new DestinationOracleConfigurationTunnelMethodDestinationOracleUpdateSshTunnelMethodNoTunnelArgs();
    }
}