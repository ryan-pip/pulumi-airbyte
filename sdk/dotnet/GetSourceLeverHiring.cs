// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceLeverHiring
    {
        public static Task<GetSourceLeverHiringResult> InvokeAsync(GetSourceLeverHiringArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceLeverHiringResult>("airbyte:index/getSourceLeverHiring:getSourceLeverHiring", args ?? new GetSourceLeverHiringArgs(), options.WithDefaults());

        public static Output<GetSourceLeverHiringResult> Invoke(GetSourceLeverHiringInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceLeverHiringResult>("airbyte:index/getSourceLeverHiring:getSourceLeverHiring", args ?? new GetSourceLeverHiringInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceLeverHiringArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceLeverHiringArgs()
        {
        }
        public static new GetSourceLeverHiringArgs Empty => new GetSourceLeverHiringArgs();
    }

    public sealed class GetSourceLeverHiringInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceLeverHiringInvokeArgs()
        {
        }
        public static new GetSourceLeverHiringInvokeArgs Empty => new GetSourceLeverHiringInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceLeverHiringResult
    {
        public readonly Outputs.GetSourceLeverHiringConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceLeverHiringResult(
            Outputs.GetSourceLeverHiringConfiguration configuration,

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
