// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceFileSecureConfigurationProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceFileSecureStorageProviderAzBlobAzureBlobStorage")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderAzBlobAzureBlobStorageArgs>? SourceFileSecureStorageProviderAzBlobAzureBlobStorage { get; set; }

        [Input("sourceFileSecureStorageProviderGcsGoogleCloudStorage")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderGcsGoogleCloudStorageArgs>? SourceFileSecureStorageProviderGcsGoogleCloudStorage { get; set; }

        [Input("sourceFileSecureStorageProviderHttpsPublicWeb")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderHttpsPublicWebArgs>? SourceFileSecureStorageProviderHttpsPublicWeb { get; set; }

        [Input("sourceFileSecureStorageProviderS3AmazonWebServices")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderS3AmazonWebServicesArgs>? SourceFileSecureStorageProviderS3AmazonWebServices { get; set; }

        [Input("sourceFileSecureStorageProviderScpSecureCopyProtocol")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderScpSecureCopyProtocolArgs>? SourceFileSecureStorageProviderScpSecureCopyProtocol { get; set; }

        [Input("sourceFileSecureStorageProviderSftpSecureFileTransferProtocol")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderSftpSecureFileTransferProtocolArgs>? SourceFileSecureStorageProviderSftpSecureFileTransferProtocol { get; set; }

        [Input("sourceFileSecureStorageProviderSshSecureShell")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureStorageProviderSshSecureShellArgs>? SourceFileSecureStorageProviderSshSecureShell { get; set; }

        [Input("sourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorageArgs>? SourceFileSecureUpdateStorageProviderAzBlobAzureBlobStorage { get; set; }

        [Input("sourceFileSecureUpdateStorageProviderGcsGoogleCloudStorage")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureUpdateStorageProviderGcsGoogleCloudStorageArgs>? SourceFileSecureUpdateStorageProviderGcsGoogleCloudStorage { get; set; }

        [Input("sourceFileSecureUpdateStorageProviderHttpsPublicWeb")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureUpdateStorageProviderHttpsPublicWebArgs>? SourceFileSecureUpdateStorageProviderHttpsPublicWeb { get; set; }

        [Input("sourceFileSecureUpdateStorageProviderS3AmazonWebServices")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureUpdateStorageProviderS3AmazonWebServicesArgs>? SourceFileSecureUpdateStorageProviderS3AmazonWebServices { get; set; }

        [Input("sourceFileSecureUpdateStorageProviderScpSecureCopyProtocol")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureUpdateStorageProviderScpSecureCopyProtocolArgs>? SourceFileSecureUpdateStorageProviderScpSecureCopyProtocol { get; set; }

        [Input("sourceFileSecureUpdateStorageProviderSftpSecureFileTransferProtocol")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureUpdateStorageProviderSftpSecureFileTransferProtocolArgs>? SourceFileSecureUpdateStorageProviderSftpSecureFileTransferProtocol { get; set; }

        [Input("sourceFileSecureUpdateStorageProviderSshSecureShell")]
        public Input<Inputs.SourceFileSecureConfigurationProviderSourceFileSecureUpdateStorageProviderSshSecureShellArgs>? SourceFileSecureUpdateStorageProviderSshSecureShell { get; set; }

        public SourceFileSecureConfigurationProviderArgs()
        {
        }
        public static new SourceFileSecureConfigurationProviderArgs Empty => new SourceFileSecureConfigurationProviderArgs();
    }
}
