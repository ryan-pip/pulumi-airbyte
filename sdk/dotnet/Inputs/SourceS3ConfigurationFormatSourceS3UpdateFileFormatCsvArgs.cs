// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceS3ConfigurationFormatSourceS3UpdateFileFormatCsvArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalReaderOptions")]
        public Input<string>? AdditionalReaderOptions { get; set; }

        [Input("advancedOptions")]
        public Input<string>? AdvancedOptions { get; set; }

        [Input("blockSize")]
        public Input<int>? BlockSize { get; set; }

        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        [Input("doubleQuote")]
        public Input<bool>? DoubleQuote { get; set; }

        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        [Input("escapeChar")]
        public Input<string>? EscapeChar { get; set; }

        [Input("filetype")]
        public Input<string>? Filetype { get; set; }

        [Input("inferDatatypes")]
        public Input<bool>? InferDatatypes { get; set; }

        [Input("newlinesInValues")]
        public Input<bool>? NewlinesInValues { get; set; }

        [Input("quoteChar")]
        public Input<string>? QuoteChar { get; set; }

        public SourceS3ConfigurationFormatSourceS3UpdateFileFormatCsvArgs()
        {
        }
        public static new SourceS3ConfigurationFormatSourceS3UpdateFileFormatCsvArgs Empty => new SourceS3ConfigurationFormatSourceS3UpdateFileFormatCsvArgs();
    }
}
