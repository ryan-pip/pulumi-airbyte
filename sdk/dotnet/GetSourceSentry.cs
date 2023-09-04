// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceSentry
    {
        public static Task<GetSourceSentryResult> InvokeAsync(GetSourceSentryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceSentryResult>("airbyte:index/getSourceSentry:getSourceSentry", args ?? new GetSourceSentryArgs(), options.WithDefaults());

        public static Output<GetSourceSentryResult> Invoke(GetSourceSentryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceSentryResult>("airbyte:index/getSourceSentry:getSourceSentry", args ?? new GetSourceSentryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceSentryArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceSentryArgs()
        {
        }
        public static new GetSourceSentryArgs Empty => new GetSourceSentryArgs();
    }

    public sealed class GetSourceSentryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceSentryInvokeArgs()
        {
        }
        public static new GetSourceSentryInvokeArgs Empty => new GetSourceSentryInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceSentryResult
    {
        public readonly Outputs.GetSourceSentryConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceSentryResult(
            Outputs.GetSourceSentryConfiguration configuration,

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