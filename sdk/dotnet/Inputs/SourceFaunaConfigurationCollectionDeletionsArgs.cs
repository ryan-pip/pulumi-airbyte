// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceFaunaConfigurationCollectionDeletionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceFaunaCollectionDeletionModeDisabled")]
        public Input<Inputs.SourceFaunaConfigurationCollectionDeletionsSourceFaunaCollectionDeletionModeDisabledArgs>? SourceFaunaCollectionDeletionModeDisabled { get; set; }

        [Input("sourceFaunaCollectionDeletionModeEnabled")]
        public Input<Inputs.SourceFaunaConfigurationCollectionDeletionsSourceFaunaCollectionDeletionModeEnabledArgs>? SourceFaunaCollectionDeletionModeEnabled { get; set; }

        [Input("sourceFaunaUpdateCollectionDeletionModeDisabled")]
        public Input<Inputs.SourceFaunaConfigurationCollectionDeletionsSourceFaunaUpdateCollectionDeletionModeDisabledArgs>? SourceFaunaUpdateCollectionDeletionModeDisabled { get; set; }

        [Input("sourceFaunaUpdateCollectionDeletionModeEnabled")]
        public Input<Inputs.SourceFaunaConfigurationCollectionDeletionsSourceFaunaUpdateCollectionDeletionModeEnabledArgs>? SourceFaunaUpdateCollectionDeletionModeEnabled { get; set; }

        public SourceFaunaConfigurationCollectionDeletionsArgs()
        {
        }
        public static new SourceFaunaConfigurationCollectionDeletionsArgs Empty => new SourceFaunaConfigurationCollectionDeletionsArgs();
    }
}