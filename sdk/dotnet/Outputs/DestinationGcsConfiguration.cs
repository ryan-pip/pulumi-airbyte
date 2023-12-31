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
    public sealed class DestinationGcsConfiguration
    {
        public readonly Outputs.DestinationGcsConfigurationCredential Credential;
        public readonly string DestinationType;
        public readonly Outputs.DestinationGcsConfigurationFormat Format;
        public readonly string GcsBucketName;
        public readonly string GcsBucketPath;
        public readonly string? GcsBucketRegion;

        [OutputConstructor]
        private DestinationGcsConfiguration(
            Outputs.DestinationGcsConfigurationCredential credential,

            string destinationType,

            Outputs.DestinationGcsConfigurationFormat format,

            string gcsBucketName,

            string gcsBucketPath,

            string? gcsBucketRegion)
        {
            Credential = credential;
            DestinationType = destinationType;
            Format = format;
            GcsBucketName = gcsBucketName;
            GcsBucketPath = gcsBucketPath;
            GcsBucketRegion = gcsBucketRegion;
        }
    }
}
