// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationLangchainConfigurationIndexingArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationLangchainIndexingChromaLocalPersistance")]
        public Input<Inputs.DestinationLangchainConfigurationIndexingDestinationLangchainIndexingChromaLocalPersistanceArgs>? DestinationLangchainIndexingChromaLocalPersistance { get; set; }

        [Input("destinationLangchainIndexingDocArrayHnswSearch")]
        public Input<Inputs.DestinationLangchainConfigurationIndexingDestinationLangchainIndexingDocArrayHnswSearchArgs>? DestinationLangchainIndexingDocArrayHnswSearch { get; set; }

        [Input("destinationLangchainIndexingPinecone")]
        public Input<Inputs.DestinationLangchainConfigurationIndexingDestinationLangchainIndexingPineconeArgs>? DestinationLangchainIndexingPinecone { get; set; }

        [Input("destinationLangchainUpdateIndexingChromaLocalPersistance")]
        public Input<Inputs.DestinationLangchainConfigurationIndexingDestinationLangchainUpdateIndexingChromaLocalPersistanceArgs>? DestinationLangchainUpdateIndexingChromaLocalPersistance { get; set; }

        [Input("destinationLangchainUpdateIndexingDocArrayHnswSearch")]
        public Input<Inputs.DestinationLangchainConfigurationIndexingDestinationLangchainUpdateIndexingDocArrayHnswSearchArgs>? DestinationLangchainUpdateIndexingDocArrayHnswSearch { get; set; }

        [Input("destinationLangchainUpdateIndexingPinecone")]
        public Input<Inputs.DestinationLangchainConfigurationIndexingDestinationLangchainUpdateIndexingPineconeArgs>? DestinationLangchainUpdateIndexingPinecone { get; set; }

        public DestinationLangchainConfigurationIndexingArgs()
        {
        }
        public static new DestinationLangchainConfigurationIndexingArgs Empty => new DestinationLangchainConfigurationIndexingArgs();
    }
}
