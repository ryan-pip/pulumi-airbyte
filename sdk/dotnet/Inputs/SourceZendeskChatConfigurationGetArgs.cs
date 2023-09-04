// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceZendeskChatConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentials")]
        public Input<Inputs.SourceZendeskChatConfigurationCredentialsGetArgs>? Credentials { get; set; }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

        public SourceZendeskChatConfigurationGetArgs()
        {
        }
        public static new SourceZendeskChatConfigurationGetArgs Empty => new SourceZendeskChatConfigurationGetArgs();
    }
}