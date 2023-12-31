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
    public sealed class GetSourceChargebeeConfiguration
    {
        public readonly string ProductCatalog;
        public readonly string Site;
        public readonly string SiteApiKey;
        public readonly string SourceType;
        public readonly string StartDate;

        [OutputConstructor]
        private GetSourceChargebeeConfiguration(
            string productCatalog,

            string site,

            string siteApiKey,

            string sourceType,

            string startDate)
        {
            ProductCatalog = productCatalog;
            Site = site;
            SiteApiKey = siteApiKey;
            SourceType = sourceType;
            StartDate = startDate;
        }
    }
}
