// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationElasticsearchConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authenticationMethod")]
        public Input<Inputs.DestinationElasticsearchConfigurationAuthenticationMethodGetArgs>? AuthenticationMethod { get; set; }

        [Input("caCertificate")]
        public Input<string>? CaCertificate { get; set; }

        [Input("destinationType", required: true)]
        public Input<string> DestinationType { get; set; } = null!;

        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        [Input("upsert")]
        public Input<bool>? Upsert { get; set; }

        public DestinationElasticsearchConfigurationGetArgs()
        {
        }
        public static new DestinationElasticsearchConfigurationGetArgs Empty => new DestinationElasticsearchConfigurationGetArgs();
    }
}
