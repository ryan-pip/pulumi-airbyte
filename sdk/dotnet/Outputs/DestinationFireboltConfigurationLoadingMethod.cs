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
    public sealed class DestinationFireboltConfigurationLoadingMethod
    {
        public readonly Outputs.DestinationFireboltConfigurationLoadingMethodDestinationFireboltLoadingMethodExternalTableViaS3? DestinationFireboltLoadingMethodExternalTableViaS3;
        public readonly Outputs.DestinationFireboltConfigurationLoadingMethodDestinationFireboltLoadingMethodSqlInserts? DestinationFireboltLoadingMethodSqlInserts;
        public readonly Outputs.DestinationFireboltConfigurationLoadingMethodDestinationFireboltUpdateLoadingMethodExternalTableViaS3? DestinationFireboltUpdateLoadingMethodExternalTableViaS3;
        public readonly Outputs.DestinationFireboltConfigurationLoadingMethodDestinationFireboltUpdateLoadingMethodSqlInserts? DestinationFireboltUpdateLoadingMethodSqlInserts;

        [OutputConstructor]
        private DestinationFireboltConfigurationLoadingMethod(
            Outputs.DestinationFireboltConfigurationLoadingMethodDestinationFireboltLoadingMethodExternalTableViaS3? destinationFireboltLoadingMethodExternalTableViaS3,

            Outputs.DestinationFireboltConfigurationLoadingMethodDestinationFireboltLoadingMethodSqlInserts? destinationFireboltLoadingMethodSqlInserts,

            Outputs.DestinationFireboltConfigurationLoadingMethodDestinationFireboltUpdateLoadingMethodExternalTableViaS3? destinationFireboltUpdateLoadingMethodExternalTableViaS3,

            Outputs.DestinationFireboltConfigurationLoadingMethodDestinationFireboltUpdateLoadingMethodSqlInserts? destinationFireboltUpdateLoadingMethodSqlInserts)
        {
            DestinationFireboltLoadingMethodExternalTableViaS3 = destinationFireboltLoadingMethodExternalTableViaS3;
            DestinationFireboltLoadingMethodSqlInserts = destinationFireboltLoadingMethodSqlInserts;
            DestinationFireboltUpdateLoadingMethodExternalTableViaS3 = destinationFireboltUpdateLoadingMethodExternalTableViaS3;
            DestinationFireboltUpdateLoadingMethodSqlInserts = destinationFireboltUpdateLoadingMethodSqlInserts;
        }
    }
}