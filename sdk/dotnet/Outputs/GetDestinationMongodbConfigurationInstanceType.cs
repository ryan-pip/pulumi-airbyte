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
    public sealed class GetDestinationMongodbConfigurationInstanceType
    {
        public readonly Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbMongoDbInstanceTypeMongoDbAtlas DestinationMongodbMongoDbInstanceTypeMongoDbAtlas;
        public readonly Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbMongoDbInstanceTypeReplicaSet DestinationMongodbMongoDbInstanceTypeReplicaSet;
        public readonly Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance;
        public readonly Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbUpdateMongoDbInstanceTypeMongoDbAtlas DestinationMongodbUpdateMongoDbInstanceTypeMongoDbAtlas;
        public readonly Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbUpdateMongoDbInstanceTypeReplicaSet DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet;
        public readonly Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance;

        [OutputConstructor]
        private GetDestinationMongodbConfigurationInstanceType(
            Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbMongoDbInstanceTypeMongoDbAtlas destinationMongodbMongoDbInstanceTypeMongoDbAtlas,

            Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbMongoDbInstanceTypeReplicaSet destinationMongodbMongoDbInstanceTypeReplicaSet,

            Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance,

            Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbUpdateMongoDbInstanceTypeMongoDbAtlas destinationMongodbUpdateMongoDbInstanceTypeMongoDbAtlas,

            Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbUpdateMongoDbInstanceTypeReplicaSet destinationMongodbUpdateMongoDbInstanceTypeReplicaSet,

            Outputs.GetDestinationMongodbConfigurationInstanceTypeDestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance)
        {
            DestinationMongodbMongoDbInstanceTypeMongoDbAtlas = destinationMongodbMongoDbInstanceTypeMongoDbAtlas;
            DestinationMongodbMongoDbInstanceTypeReplicaSet = destinationMongodbMongoDbInstanceTypeReplicaSet;
            DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance = destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance;
            DestinationMongodbUpdateMongoDbInstanceTypeMongoDbAtlas = destinationMongodbUpdateMongoDbInstanceTypeMongoDbAtlas;
            DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet = destinationMongodbUpdateMongoDbInstanceTypeReplicaSet;
            DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance = destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance;
        }
    }
}
