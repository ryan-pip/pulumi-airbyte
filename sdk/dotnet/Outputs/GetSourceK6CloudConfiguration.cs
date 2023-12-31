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
    public sealed class GetSourceK6CloudConfiguration
    {
        public readonly string ApiToken;
        public readonly string SourceType;

        [OutputConstructor]
        private GetSourceK6CloudConfiguration(
            string apiToken,

            string sourceType)
        {
            ApiToken = apiToken;
            SourceType = sourceType;
        }
    }
}
