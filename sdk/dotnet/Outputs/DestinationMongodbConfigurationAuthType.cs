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
    public sealed class DestinationMongodbConfigurationAuthType
    {
        public readonly Outputs.DestinationMongodbConfigurationAuthTypeDestinationMongodbAuthorizationTypeLoginPassword? DestinationMongodbAuthorizationTypeLoginPassword;
        public readonly Outputs.DestinationMongodbConfigurationAuthTypeDestinationMongodbAuthorizationTypeNone? DestinationMongodbAuthorizationTypeNone;
        public readonly Outputs.DestinationMongodbConfigurationAuthTypeDestinationMongodbUpdateAuthorizationTypeLoginPassword? DestinationMongodbUpdateAuthorizationTypeLoginPassword;
        public readonly Outputs.DestinationMongodbConfigurationAuthTypeDestinationMongodbUpdateAuthorizationTypeNone? DestinationMongodbUpdateAuthorizationTypeNone;

        [OutputConstructor]
        private DestinationMongodbConfigurationAuthType(
            Outputs.DestinationMongodbConfigurationAuthTypeDestinationMongodbAuthorizationTypeLoginPassword? destinationMongodbAuthorizationTypeLoginPassword,

            Outputs.DestinationMongodbConfigurationAuthTypeDestinationMongodbAuthorizationTypeNone? destinationMongodbAuthorizationTypeNone,

            Outputs.DestinationMongodbConfigurationAuthTypeDestinationMongodbUpdateAuthorizationTypeLoginPassword? destinationMongodbUpdateAuthorizationTypeLoginPassword,

            Outputs.DestinationMongodbConfigurationAuthTypeDestinationMongodbUpdateAuthorizationTypeNone? destinationMongodbUpdateAuthorizationTypeNone)
        {
            DestinationMongodbAuthorizationTypeLoginPassword = destinationMongodbAuthorizationTypeLoginPassword;
            DestinationMongodbAuthorizationTypeNone = destinationMongodbAuthorizationTypeNone;
            DestinationMongodbUpdateAuthorizationTypeLoginPassword = destinationMongodbUpdateAuthorizationTypeLoginPassword;
            DestinationMongodbUpdateAuthorizationTypeNone = destinationMongodbUpdateAuthorizationTypeNone;
        }
    }
}
