// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationSftpJsonConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationPath", required: true)]
        public Input<string> DestinationPath { get; set; } = null!;

        [Input("destinationType", required: true)]
        public Input<string> DestinationType { get; set; } = null!;

        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public DestinationSftpJsonConfigurationArgs()
        {
        }
        public static new DestinationSftpJsonConfigurationArgs Empty => new DestinationSftpJsonConfigurationArgs();
    }
}