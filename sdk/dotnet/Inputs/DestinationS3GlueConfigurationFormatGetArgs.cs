// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationS3GlueConfigurationFormatGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationS3GlueOutputFormatJsonLinesNewlineDelimitedJson")]
        public Input<Inputs.DestinationS3GlueConfigurationFormatDestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJsonGetArgs>? DestinationS3GlueOutputFormatJsonLinesNewlineDelimitedJson { get; set; }

        [Input("destinationS3GlueUpdateOutputFormatJsonLinesNewlineDelimitedJson")]
        public Input<Inputs.DestinationS3GlueConfigurationFormatDestinationS3GlueUpdateOutputFormatJsonLinesNewlineDelimitedJsonGetArgs>? DestinationS3GlueUpdateOutputFormatJsonLinesNewlineDelimitedJson { get; set; }

        public DestinationS3GlueConfigurationFormatGetArgs()
        {
        }
        public static new DestinationS3GlueConfigurationFormatGetArgs Empty => new DestinationS3GlueConfigurationFormatGetArgs();
    }
}
