// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceOracleConfigurationEncryptionSourceOracleUpdateEncryptionNativeNetworkEncryptionNneGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        [Input("encryptionMethod", required: true)]
        public Input<string> EncryptionMethod { get; set; } = null!;

        public SourceOracleConfigurationEncryptionSourceOracleUpdateEncryptionNativeNetworkEncryptionNneGetArgs()
        {
        }
        public static new SourceOracleConfigurationEncryptionSourceOracleUpdateEncryptionNativeNetworkEncryptionNneGetArgs Empty => new SourceOracleConfigurationEncryptionSourceOracleUpdateEncryptionNativeNetworkEncryptionNneGetArgs();
    }
}