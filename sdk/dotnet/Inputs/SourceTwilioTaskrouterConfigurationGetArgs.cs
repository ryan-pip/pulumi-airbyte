// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceTwilioTaskrouterConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountSid", required: true)]
        public Input<string> AccountSid { get; set; } = null!;

        [Input("authToken", required: true)]
        public Input<string> AuthToken { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        public SourceTwilioTaskrouterConfigurationGetArgs()
        {
        }
        public static new SourceTwilioTaskrouterConfigurationGetArgs Empty => new SourceTwilioTaskrouterConfigurationGetArgs();
    }
}