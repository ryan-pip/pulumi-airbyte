// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceConfigcatConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SourceConfigcatConfigurationGetArgs()
        {
        }
        public static new SourceConfigcatConfigurationGetArgs Empty => new SourceConfigcatConfigurationGetArgs();
    }
}
