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
    public sealed class SourceOracleConfigurationEncryption
    {
        public readonly Outputs.SourceOracleConfigurationEncryptionSourceOracleEncryptionNativeNetworkEncryptionNne? SourceOracleEncryptionNativeNetworkEncryptionNne;
        public readonly Outputs.SourceOracleConfigurationEncryptionSourceOracleEncryptionTlsEncryptedVerifyCertificate? SourceOracleEncryptionTlsEncryptedVerifyCertificate;
        public readonly Outputs.SourceOracleConfigurationEncryptionSourceOracleUpdateEncryptionNativeNetworkEncryptionNne? SourceOracleUpdateEncryptionNativeNetworkEncryptionNne;
        public readonly Outputs.SourceOracleConfigurationEncryptionSourceOracleUpdateEncryptionTlsEncryptedVerifyCertificate? SourceOracleUpdateEncryptionTlsEncryptedVerifyCertificate;

        [OutputConstructor]
        private SourceOracleConfigurationEncryption(
            Outputs.SourceOracleConfigurationEncryptionSourceOracleEncryptionNativeNetworkEncryptionNne? sourceOracleEncryptionNativeNetworkEncryptionNne,

            Outputs.SourceOracleConfigurationEncryptionSourceOracleEncryptionTlsEncryptedVerifyCertificate? sourceOracleEncryptionTlsEncryptedVerifyCertificate,

            Outputs.SourceOracleConfigurationEncryptionSourceOracleUpdateEncryptionNativeNetworkEncryptionNne? sourceOracleUpdateEncryptionNativeNetworkEncryptionNne,

            Outputs.SourceOracleConfigurationEncryptionSourceOracleUpdateEncryptionTlsEncryptedVerifyCertificate? sourceOracleUpdateEncryptionTlsEncryptedVerifyCertificate)
        {
            SourceOracleEncryptionNativeNetworkEncryptionNne = sourceOracleEncryptionNativeNetworkEncryptionNne;
            SourceOracleEncryptionTlsEncryptedVerifyCertificate = sourceOracleEncryptionTlsEncryptedVerifyCertificate;
            SourceOracleUpdateEncryptionNativeNetworkEncryptionNne = sourceOracleUpdateEncryptionNativeNetworkEncryptionNne;
            SourceOracleUpdateEncryptionTlsEncryptedVerifyCertificate = sourceOracleUpdateEncryptionTlsEncryptedVerifyCertificate;
        }
    }
}
