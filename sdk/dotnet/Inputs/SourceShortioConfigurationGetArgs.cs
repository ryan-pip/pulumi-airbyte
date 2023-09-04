// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceShortioConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        [Input("secretKey", required: true)]
        public Input<string> SecretKey { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public SourceShortioConfigurationGetArgs()
        {
        }
        public static new SourceShortioConfigurationGetArgs Empty => new SourceShortioConfigurationGetArgs();
    }
}