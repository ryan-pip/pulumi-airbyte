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
    public sealed class GetDestinationDevNullConfiguration
    {
        public readonly string DestinationType;
        public readonly Outputs.GetDestinationDevNullConfigurationTestDestination TestDestination;

        [OutputConstructor]
        private GetDestinationDevNullConfiguration(
            string destinationType,

            Outputs.GetDestinationDevNullConfigurationTestDestination testDestination)
        {
            DestinationType = destinationType;
            TestDestination = testDestination;
        }
    }
}
