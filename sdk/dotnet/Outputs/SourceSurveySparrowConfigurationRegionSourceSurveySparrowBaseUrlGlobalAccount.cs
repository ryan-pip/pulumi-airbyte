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
    public sealed class SourceSurveySparrowConfigurationRegionSourceSurveySparrowBaseUrlGlobalAccount
    {
        public readonly string? UrlBase;

        [OutputConstructor]
        private SourceSurveySparrowConfigurationRegionSourceSurveySparrowBaseUrlGlobalAccount(string? urlBase)
        {
            UrlBase = urlBase;
        }
    }
}