// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderScpSecureCopyProtocolArgs : global::Pulumi.ResourceArgs
    {
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("port")]
        public Input<string>? Port { get; set; }

        [Input("storage", required: true)]
        public Input<string> Storage { get; set; } = null!;

        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderScpSecureCopyProtocolArgs()
        {
        }
        public static new SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderScpSecureCopyProtocolArgs Empty => new SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderScpSecureCopyProtocolArgs();
    }
}
