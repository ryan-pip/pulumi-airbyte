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
    public sealed class GetDestinationRedisConfigurationSslMode
    {
        public readonly Outputs.GetDestinationRedisConfigurationSslModeDestinationRedisSslModesDisable DestinationRedisSslModesDisable;
        public readonly Outputs.GetDestinationRedisConfigurationSslModeDestinationRedisSslModesVerifyFull DestinationRedisSslModesVerifyFull;
        public readonly Outputs.GetDestinationRedisConfigurationSslModeDestinationRedisUpdateSslModesDisable DestinationRedisUpdateSslModesDisable;
        public readonly Outputs.GetDestinationRedisConfigurationSslModeDestinationRedisUpdateSslModesVerifyFull DestinationRedisUpdateSslModesVerifyFull;

        [OutputConstructor]
        private GetDestinationRedisConfigurationSslMode(
            Outputs.GetDestinationRedisConfigurationSslModeDestinationRedisSslModesDisable destinationRedisSslModesDisable,

            Outputs.GetDestinationRedisConfigurationSslModeDestinationRedisSslModesVerifyFull destinationRedisSslModesVerifyFull,

            Outputs.GetDestinationRedisConfigurationSslModeDestinationRedisUpdateSslModesDisable destinationRedisUpdateSslModesDisable,

            Outputs.GetDestinationRedisConfigurationSslModeDestinationRedisUpdateSslModesVerifyFull destinationRedisUpdateSslModesVerifyFull)
        {
            DestinationRedisSslModesDisable = destinationRedisSslModesDisable;
            DestinationRedisSslModesVerifyFull = destinationRedisSslModesVerifyFull;
            DestinationRedisUpdateSslModesDisable = destinationRedisUpdateSslModesDisable;
            DestinationRedisUpdateSslModesVerifyFull = destinationRedisUpdateSslModesVerifyFull;
        }
    }
}