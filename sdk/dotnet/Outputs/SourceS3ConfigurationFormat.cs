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
    public sealed class SourceS3ConfigurationFormat
    {
        public readonly Outputs.SourceS3ConfigurationFormatSourceS3FileFormatAvro? SourceS3FileFormatAvro;
        public readonly Outputs.SourceS3ConfigurationFormatSourceS3FileFormatCsv? SourceS3FileFormatCsv;
        public readonly Outputs.SourceS3ConfigurationFormatSourceS3FileFormatJsonl? SourceS3FileFormatJsonl;
        public readonly Outputs.SourceS3ConfigurationFormatSourceS3FileFormatParquet? SourceS3FileFormatParquet;
        public readonly Outputs.SourceS3ConfigurationFormatSourceS3UpdateFileFormatAvro? SourceS3UpdateFileFormatAvro;
        public readonly Outputs.SourceS3ConfigurationFormatSourceS3UpdateFileFormatCsv? SourceS3UpdateFileFormatCsv;
        public readonly Outputs.SourceS3ConfigurationFormatSourceS3UpdateFileFormatJsonl? SourceS3UpdateFileFormatJsonl;
        public readonly Outputs.SourceS3ConfigurationFormatSourceS3UpdateFileFormatParquet? SourceS3UpdateFileFormatParquet;

        [OutputConstructor]
        private SourceS3ConfigurationFormat(
            Outputs.SourceS3ConfigurationFormatSourceS3FileFormatAvro? sourceS3FileFormatAvro,

            Outputs.SourceS3ConfigurationFormatSourceS3FileFormatCsv? sourceS3FileFormatCsv,

            Outputs.SourceS3ConfigurationFormatSourceS3FileFormatJsonl? sourceS3FileFormatJsonl,

            Outputs.SourceS3ConfigurationFormatSourceS3FileFormatParquet? sourceS3FileFormatParquet,

            Outputs.SourceS3ConfigurationFormatSourceS3UpdateFileFormatAvro? sourceS3UpdateFileFormatAvro,

            Outputs.SourceS3ConfigurationFormatSourceS3UpdateFileFormatCsv? sourceS3UpdateFileFormatCsv,

            Outputs.SourceS3ConfigurationFormatSourceS3UpdateFileFormatJsonl? sourceS3UpdateFileFormatJsonl,

            Outputs.SourceS3ConfigurationFormatSourceS3UpdateFileFormatParquet? sourceS3UpdateFileFormatParquet)
        {
            SourceS3FileFormatAvro = sourceS3FileFormatAvro;
            SourceS3FileFormatCsv = sourceS3FileFormatCsv;
            SourceS3FileFormatJsonl = sourceS3FileFormatJsonl;
            SourceS3FileFormatParquet = sourceS3FileFormatParquet;
            SourceS3UpdateFileFormatAvro = sourceS3UpdateFileFormatAvro;
            SourceS3UpdateFileFormatCsv = sourceS3UpdateFileFormatCsv;
            SourceS3UpdateFileFormatJsonl = sourceS3UpdateFileFormatJsonl;
            SourceS3UpdateFileFormatParquet = sourceS3UpdateFileFormatParquet;
        }
    }
}