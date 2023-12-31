// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceOrbitConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiToken", required: true)]
        public Input<string> ApiToken { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        [Input("workspace", required: true)]
        public Input<string> Workspace { get; set; } = null!;

        public SourceOrbitConfigurationArgs()
        {
        }
        public static new SourceOrbitConfigurationArgs Empty => new SourceOrbitConfigurationArgs();
    }
}
