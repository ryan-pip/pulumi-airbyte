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
    public sealed class DestinationPostgresConfigurationSslMode
    {
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesAllow? DestinationPostgresSslModesAllow;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesDisable? DestinationPostgresSslModesDisable;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesPrefer? DestinationPostgresSslModesPrefer;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesRequire? DestinationPostgresSslModesRequire;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesVerifyCa? DestinationPostgresSslModesVerifyCa;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesVerifyFull? DestinationPostgresSslModesVerifyFull;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesAllow? DestinationPostgresUpdateSslModesAllow;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesDisable? DestinationPostgresUpdateSslModesDisable;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesPrefer? DestinationPostgresUpdateSslModesPrefer;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesRequire? DestinationPostgresUpdateSslModesRequire;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesVerifyCa? DestinationPostgresUpdateSslModesVerifyCa;
        public readonly Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesVerifyFull? DestinationPostgresUpdateSslModesVerifyFull;

        [OutputConstructor]
        private DestinationPostgresConfigurationSslMode(
            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesAllow? destinationPostgresSslModesAllow,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesDisable? destinationPostgresSslModesDisable,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesPrefer? destinationPostgresSslModesPrefer,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesRequire? destinationPostgresSslModesRequire,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesVerifyCa? destinationPostgresSslModesVerifyCa,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresSslModesVerifyFull? destinationPostgresSslModesVerifyFull,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesAllow? destinationPostgresUpdateSslModesAllow,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesDisable? destinationPostgresUpdateSslModesDisable,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesPrefer? destinationPostgresUpdateSslModesPrefer,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesRequire? destinationPostgresUpdateSslModesRequire,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesVerifyCa? destinationPostgresUpdateSslModesVerifyCa,

            Outputs.DestinationPostgresConfigurationSslModeDestinationPostgresUpdateSslModesVerifyFull? destinationPostgresUpdateSslModesVerifyFull)
        {
            DestinationPostgresSslModesAllow = destinationPostgresSslModesAllow;
            DestinationPostgresSslModesDisable = destinationPostgresSslModesDisable;
            DestinationPostgresSslModesPrefer = destinationPostgresSslModesPrefer;
            DestinationPostgresSslModesRequire = destinationPostgresSslModesRequire;
            DestinationPostgresSslModesVerifyCa = destinationPostgresSslModesVerifyCa;
            DestinationPostgresSslModesVerifyFull = destinationPostgresSslModesVerifyFull;
            DestinationPostgresUpdateSslModesAllow = destinationPostgresUpdateSslModesAllow;
            DestinationPostgresUpdateSslModesDisable = destinationPostgresUpdateSslModesDisable;
            DestinationPostgresUpdateSslModesPrefer = destinationPostgresUpdateSslModesPrefer;
            DestinationPostgresUpdateSslModesRequire = destinationPostgresUpdateSslModesRequire;
            DestinationPostgresUpdateSslModesVerifyCa = destinationPostgresUpdateSslModesVerifyCa;
            DestinationPostgresUpdateSslModesVerifyFull = destinationPostgresUpdateSslModesVerifyFull;
        }
    }
}