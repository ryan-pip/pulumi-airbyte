// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceAirtableConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentials")]
        public Input<Inputs.SourceAirtableConfigurationCredentialsArgs>? Credentials { get; set; }

        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        public SourceAirtableConfigurationArgs()
        {
        }
        public static new SourceAirtableConfigurationArgs Empty => new SourceAirtableConfigurationArgs();
    }
}