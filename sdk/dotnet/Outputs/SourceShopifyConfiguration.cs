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
    public sealed class SourceShopifyConfiguration
    {
        public readonly Outputs.SourceShopifyConfigurationCredentials? Credentials;
        public readonly string Shop;
        public readonly string SourceType;
        public readonly string StartDate;

        [OutputConstructor]
        private SourceShopifyConfiguration(
            Outputs.SourceShopifyConfigurationCredentials? credentials,

            string shop,

            string sourceType,

            string startDate)
        {
            Credentials = credentials;
            Shop = shop;
            SourceType = sourceType;
            StartDate = startDate;
        }
    }
}
