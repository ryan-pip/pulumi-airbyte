// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationElasticsearchConfigurationAuthenticationMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationElasticsearchAuthenticationMethodApiKeySecret")]
        public Input<Inputs.DestinationElasticsearchConfigurationAuthenticationMethodDestinationElasticsearchAuthenticationMethodApiKeySecretArgs>? DestinationElasticsearchAuthenticationMethodApiKeySecret { get; set; }

        [Input("destinationElasticsearchAuthenticationMethodUsernamePassword")]
        public Input<Inputs.DestinationElasticsearchConfigurationAuthenticationMethodDestinationElasticsearchAuthenticationMethodUsernamePasswordArgs>? DestinationElasticsearchAuthenticationMethodUsernamePassword { get; set; }

        [Input("destinationElasticsearchUpdateAuthenticationMethodApiKeySecret")]
        public Input<Inputs.DestinationElasticsearchConfigurationAuthenticationMethodDestinationElasticsearchUpdateAuthenticationMethodApiKeySecretArgs>? DestinationElasticsearchUpdateAuthenticationMethodApiKeySecret { get; set; }

        [Input("destinationElasticsearchUpdateAuthenticationMethodUsernamePassword")]
        public Input<Inputs.DestinationElasticsearchConfigurationAuthenticationMethodDestinationElasticsearchUpdateAuthenticationMethodUsernamePasswordArgs>? DestinationElasticsearchUpdateAuthenticationMethodUsernamePassword { get; set; }

        public DestinationElasticsearchConfigurationAuthenticationMethodArgs()
        {
        }
        public static new DestinationElasticsearchConfigurationAuthenticationMethodArgs Empty => new DestinationElasticsearchConfigurationAuthenticationMethodArgs();
    }
}
