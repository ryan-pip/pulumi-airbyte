// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationMongodbConfigurationAuthTypeDestinationMongodbAuthorizationTypeNoneGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorization", required: true)]
        public Input<string> Authorization { get; set; } = null!;

        public DestinationMongodbConfigurationAuthTypeDestinationMongodbAuthorizationTypeNoneGetArgs()
        {
        }
        public static new DestinationMongodbConfigurationAuthTypeDestinationMongodbAuthorizationTypeNoneGetArgs Empty => new DestinationMongodbConfigurationAuthTypeDestinationMongodbAuthorizationTypeNoneGetArgs();
    }
}
