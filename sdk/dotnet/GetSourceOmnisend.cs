// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceOmnisend
    {
        public static Task<GetSourceOmnisendResult> InvokeAsync(GetSourceOmnisendArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceOmnisendResult>("airbyte:index/getSourceOmnisend:getSourceOmnisend", args ?? new GetSourceOmnisendArgs(), options.WithDefaults());

        public static Output<GetSourceOmnisendResult> Invoke(GetSourceOmnisendInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceOmnisendResult>("airbyte:index/getSourceOmnisend:getSourceOmnisend", args ?? new GetSourceOmnisendInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceOmnisendArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceOmnisendArgs()
        {
        }
        public static new GetSourceOmnisendArgs Empty => new GetSourceOmnisendArgs();
    }

    public sealed class GetSourceOmnisendInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceOmnisendInvokeArgs()
        {
        }
        public static new GetSourceOmnisendInvokeArgs Empty => new GetSourceOmnisendInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceOmnisendResult
    {
        public readonly Outputs.GetSourceOmnisendConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceOmnisendResult(
            Outputs.GetSourceOmnisendConfiguration configuration,

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