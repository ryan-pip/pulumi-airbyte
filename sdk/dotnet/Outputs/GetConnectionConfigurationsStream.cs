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
    public sealed class GetConnectionConfigurationsStream
    {
        public readonly ImmutableArray<string> CursorFields;
        public readonly string Name;
        public readonly ImmutableArray<ImmutableArray<string>> PrimaryKeys;
        public readonly string SyncMode;

        [OutputConstructor]
        private GetConnectionConfigurationsStream(
            ImmutableArray<string> cursorFields,

            string name,

            ImmutableArray<ImmutableArray<string>> primaryKeys,

            string syncMode)
        {
            CursorFields = cursorFields;
            Name = name;
            PrimaryKeys = primaryKeys;
            SyncMode = syncMode;
        }
    }
}
