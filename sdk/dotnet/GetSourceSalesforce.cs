// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceSalesforce
    {
        public static Task<GetSourceSalesforceResult> InvokeAsync(GetSourceSalesforceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceSalesforceResult>("airbyte:index/getSourceSalesforce:getSourceSalesforce", args ?? new GetSourceSalesforceArgs(), options.WithDefaults());

        public static Output<GetSourceSalesforceResult> Invoke(GetSourceSalesforceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceSalesforceResult>("airbyte:index/getSourceSalesforce:getSourceSalesforce", args ?? new GetSourceSalesforceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceSalesforceArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceSalesforceArgs()
        {
        }
        public static new GetSourceSalesforceArgs Empty => new GetSourceSalesforceArgs();
    }

    public sealed class GetSourceSalesforceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceSalesforceInvokeArgs()
        {
        }
        public static new GetSourceSalesforceInvokeArgs Empty => new GetSourceSalesforceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceSalesforceResult
    {
        public readonly Outputs.GetSourceSalesforceConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceSalesforceResult(
            Outputs.GetSourceSalesforceConfiguration configuration,

            string id,

            string name,

            string? secretId,

            string sourceId,

            string workspaceId)
        {
            Configuration = configuration;
            Id = id;
            Name = name;
            SecretId = secretId;
            SourceId = sourceId;
            WorkspaceId = workspaceId;
        }
    }
}