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
    public sealed class SourceSurveySparrowConfigurationRegion
    {
        public readonly Outputs.SourceSurveySparrowConfigurationRegionSourceSurveySparrowBaseUrlEuBasedAccount? SourceSurveySparrowBaseUrlEuBasedAccount;
        public readonly Outputs.SourceSurveySparrowConfigurationRegionSourceSurveySparrowBaseUrlGlobalAccount? SourceSurveySparrowBaseUrlGlobalAccount;
        public readonly Outputs.SourceSurveySparrowConfigurationRegionSourceSurveySparrowUpdateBaseUrlEuBasedAccount? SourceSurveySparrowUpdateBaseUrlEuBasedAccount;
        public readonly Outputs.SourceSurveySparrowConfigurationRegionSourceSurveySparrowUpdateBaseUrlGlobalAccount? SourceSurveySparrowUpdateBaseUrlGlobalAccount;

        [OutputConstructor]
        private SourceSurveySparrowConfigurationRegion(
            Outputs.SourceSurveySparrowConfigurationRegionSourceSurveySparrowBaseUrlEuBasedAccount? sourceSurveySparrowBaseUrlEuBasedAccount,

            Outputs.SourceSurveySparrowConfigurationRegionSourceSurveySparrowBaseUrlGlobalAccount? sourceSurveySparrowBaseUrlGlobalAccount,

            Outputs.SourceSurveySparrowConfigurationRegionSourceSurveySparrowUpdateBaseUrlEuBasedAccount? sourceSurveySparrowUpdateBaseUrlEuBasedAccount,

            Outputs.SourceSurveySparrowConfigurationRegionSourceSurveySparrowUpdateBaseUrlGlobalAccount? sourceSurveySparrowUpdateBaseUrlGlobalAccount)
        {
            SourceSurveySparrowBaseUrlEuBasedAccount = sourceSurveySparrowBaseUrlEuBasedAccount;
            SourceSurveySparrowBaseUrlGlobalAccount = sourceSurveySparrowBaseUrlGlobalAccount;
            SourceSurveySparrowUpdateBaseUrlEuBasedAccount = sourceSurveySparrowUpdateBaseUrlEuBasedAccount;
            SourceSurveySparrowUpdateBaseUrlGlobalAccount = sourceSurveySparrowUpdateBaseUrlGlobalAccount;
        }
    }
}