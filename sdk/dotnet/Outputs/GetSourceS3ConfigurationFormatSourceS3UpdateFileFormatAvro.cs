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
    public sealed class GetSourceS3ConfigurationFormatSourceS3UpdateFileFormatAvro
    {
        public readonly string Filetype;

        [OutputConstructor]
        private GetSourceS3ConfigurationFormatSourceS3UpdateFileFormatAvro(string filetype)
        {
            Filetype = filetype;
        }
    }
}