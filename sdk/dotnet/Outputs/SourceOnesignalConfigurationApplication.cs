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
    public sealed class SourceOnesignalConfigurationApplication
    {
        public readonly string AppApiKey;
        public readonly string AppId;
        public readonly string? AppName;

        [OutputConstructor]
        private SourceOnesignalConfigurationApplication(
            string appApiKey,

            string appId,

            string? appName)
        {
            AppApiKey = appApiKey;
            AppId = appId;
            AppName = appName;
        }
    }
}
