// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceFaunaConfigurationCollectionDeletionsSourceFaunaCollectionDeletionModeDisabledGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("deletionMode", required: true)]
        public Input<string> DeletionMode { get; set; } = null!;

        public SourceFaunaConfigurationCollectionDeletionsSourceFaunaCollectionDeletionModeDisabledGetArgs()
        {
        }
        public static new SourceFaunaConfigurationCollectionDeletionsSourceFaunaCollectionDeletionModeDisabledGetArgs Empty => new SourceFaunaConfigurationCollectionDeletionsSourceFaunaCollectionDeletionModeDisabledGetArgs();
    }
}
