// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class DestinationLangchainConfiguration
    {
        public readonly string DestinationType;
        public readonly Outputs.DestinationLangchainConfigurationEmbedding Embedding;
        public readonly Outputs.DestinationLangchainConfigurationIndexing Indexing;
        public readonly Outputs.DestinationLangchainConfigurationProcessing Processing;

        [OutputConstructor]
        private DestinationLangchainConfiguration(
            string destinationType,

            Outputs.DestinationLangchainConfigurationEmbedding embedding,

            Outputs.DestinationLangchainConfigurationIndexing indexing,

            Outputs.DestinationLangchainConfigurationProcessing processing)
        {
            DestinationType = destinationType;
            Embedding = embedding;
            Indexing = indexing;
            Processing = processing;
        }
    }
}