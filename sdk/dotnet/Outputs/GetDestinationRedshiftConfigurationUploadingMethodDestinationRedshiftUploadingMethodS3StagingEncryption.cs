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
    public sealed class GetDestinationRedshiftConfigurationUploadingMethodDestinationRedshiftUploadingMethodS3StagingEncryption
    {
        public readonly Outputs.GetDestinationRedshiftConfigurationUploadingMethodDestinationRedshiftUploadingMethodS3StagingEncryptionDestinationRedshiftUploadingMethodS3StagingEncryptionAesCbcEnvelopeEncryption DestinationRedshiftUploadingMethodS3StagingEncryptionAesCbcEnvelopeEncryption;
        public readonly Outputs.GetDestinationRedshiftConfigurationUploadingMethodDestinationRedshiftUploadingMethodS3StagingEncryptionDestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption DestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption;

        [OutputConstructor]
        private GetDestinationRedshiftConfigurationUploadingMethodDestinationRedshiftUploadingMethodS3StagingEncryption(
            Outputs.GetDestinationRedshiftConfigurationUploadingMethodDestinationRedshiftUploadingMethodS3StagingEncryptionDestinationRedshiftUploadingMethodS3StagingEncryptionAesCbcEnvelopeEncryption destinationRedshiftUploadingMethodS3StagingEncryptionAesCbcEnvelopeEncryption,

            Outputs.GetDestinationRedshiftConfigurationUploadingMethodDestinationRedshiftUploadingMethodS3StagingEncryptionDestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption destinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption)
        {
            DestinationRedshiftUploadingMethodS3StagingEncryptionAesCbcEnvelopeEncryption = destinationRedshiftUploadingMethodS3StagingEncryptionAesCbcEnvelopeEncryption;
            DestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption = destinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption;
        }
    }
}