// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationMSsqlConfigurationSslMethodDestinationMssqlUpdateSslMethodEncryptedTrustServerCertificateGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("sslMethod", required: true)]
        public Input<string> SslMethod { get; set; } = null!;

        public DestinationMSsqlConfigurationSslMethodDestinationMssqlUpdateSslMethodEncryptedTrustServerCertificateGetArgs()
        {
        }
        public static new DestinationMSsqlConfigurationSslMethodDestinationMssqlUpdateSslMethodEncryptedTrustServerCertificateGetArgs Empty => new DestinationMSsqlConfigurationSslMethodDestinationMssqlUpdateSslMethodEncryptedTrustServerCertificateGetArgs();
    }
}
