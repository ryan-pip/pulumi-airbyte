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
    public sealed class SourceMongodbConfigurationInstanceType
    {
        public readonly Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbMongoDbInstanceTypeMongoDbAtlas? SourceMongodbMongoDbInstanceTypeMongoDbAtlas;
        public readonly Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbMongoDbInstanceTypeReplicaSet? SourceMongodbMongoDbInstanceTypeReplicaSet;
        public readonly Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance? SourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance;
        public readonly Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeMongoDbAtlas? SourceMongodbUpdateMongoDbInstanceTypeMongoDbAtlas;
        public readonly Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeReplicaSet? SourceMongodbUpdateMongoDbInstanceTypeReplicaSet;
        public readonly Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance? SourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance;

        [OutputConstructor]
        private SourceMongodbConfigurationInstanceType(
            Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbMongoDbInstanceTypeMongoDbAtlas? sourceMongodbMongoDbInstanceTypeMongoDbAtlas,

            Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbMongoDbInstanceTypeReplicaSet? sourceMongodbMongoDbInstanceTypeReplicaSet,

            Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance? sourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance,

            Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeMongoDbAtlas? sourceMongodbUpdateMongoDbInstanceTypeMongoDbAtlas,

            Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeReplicaSet? sourceMongodbUpdateMongoDbInstanceTypeReplicaSet,

            Outputs.SourceMongodbConfigurationInstanceTypeSourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance? sourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance)
        {
            SourceMongodbMongoDbInstanceTypeMongoDbAtlas = sourceMongodbMongoDbInstanceTypeMongoDbAtlas;
            SourceMongodbMongoDbInstanceTypeReplicaSet = sourceMongodbMongoDbInstanceTypeReplicaSet;
            SourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance = sourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance;
            SourceMongodbUpdateMongoDbInstanceTypeMongoDbAtlas = sourceMongodbUpdateMongoDbInstanceTypeMongoDbAtlas;
            SourceMongodbUpdateMongoDbInstanceTypeReplicaSet = sourceMongodbUpdateMongoDbInstanceTypeReplicaSet;
            SourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance = sourceMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance;
        }
    }
}
