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
    public sealed class GetSourceSurveySparrowConfigurationRegion
    {
        public readonly Outputs.GetSourceSurveySparrowConfigurationRegionSourceSurveySparrowBaseUrlEuBasedAccount SourceSurveySparrowBaseUrlEuBasedAccount;
        public readonly Outputs.GetSourceSurveySparrowConfigurationRegionSourceSurveySparrowBaseUrlGlobalAccount SourceSurveySparrowBaseUrlGlobalAccount;
        public readonly Outputs.GetSourceSurveySparrowConfigurationRegionSourceSurveySparrowUpdateBaseUrlEuBasedAccount SourceSurveySparrowUpdateBaseUrlEuBasedAccount;
        public readonly Outputs.GetSourceSurveySparrowConfigurationRegionSourceSurveySparrowUpdateBaseUrlGlobalAccount SourceSurveySparrowUpdateBaseUrlGlobalAccount;

        [OutputConstructor]
        private GetSourceSurveySparrowConfigurationRegion(
            Outputs.GetSourceSurveySparrowConfigurationRegionSourceSurveySparrowBaseUrlEuBasedAccount sourceSurveySparrowBaseUrlEuBasedAccount,

            Outputs.GetSourceSurveySparrowConfigurationRegionSourceSurveySparrowBaseUrlGlobalAccount sourceSurveySparrowBaseUrlGlobalAccount,

            Outputs.GetSourceSurveySparrowConfigurationRegionSourceSurveySparrowUpdateBaseUrlEuBasedAccount sourceSurveySparrowUpdateBaseUrlEuBasedAccount,

            Outputs.GetSourceSurveySparrowConfigurationRegionSourceSurveySparrowUpdateBaseUrlGlobalAccount sourceSurveySparrowUpdateBaseUrlGlobalAccount)
        {
            SourceSurveySparrowBaseUrlEuBasedAccount = sourceSurveySparrowBaseUrlEuBasedAccount;
            SourceSurveySparrowBaseUrlGlobalAccount = sourceSurveySparrowBaseUrlGlobalAccount;
            SourceSurveySparrowUpdateBaseUrlEuBasedAccount = sourceSurveySparrowUpdateBaseUrlEuBasedAccount;
            SourceSurveySparrowUpdateBaseUrlGlobalAccount = sourceSurveySparrowUpdateBaseUrlGlobalAccount;
        }
    }
}