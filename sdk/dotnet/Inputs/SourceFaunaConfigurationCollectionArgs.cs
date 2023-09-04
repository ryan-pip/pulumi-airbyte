// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceFaunaConfigurationCollectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("deletions", required: true)]
        public Input<Inputs.SourceFaunaConfigurationCollectionDeletionsArgs> Deletions { get; set; } = null!;

        [Input("pageSize", required: true)]
        public Input<int> PageSize { get; set; } = null!;

        public SourceFaunaConfigurationCollectionArgs()
        {
        }
        public static new SourceFaunaConfigurationCollectionArgs Empty => new SourceFaunaConfigurationCollectionArgs();
    }
}