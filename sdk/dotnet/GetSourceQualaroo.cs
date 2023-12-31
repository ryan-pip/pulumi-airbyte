// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceQualaroo
    {
        public static Task<GetSourceQualarooResult> InvokeAsync(GetSourceQualarooArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceQualarooResult>("airbyte:index/getSourceQualaroo:getSourceQualaroo", args ?? new GetSourceQualarooArgs(), options.WithDefaults());

        public static Output<GetSourceQualarooResult> Invoke(GetSourceQualarooInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceQualarooResult>("airbyte:index/getSourceQualaroo:getSourceQualaroo", args ?? new GetSourceQualarooInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceQualarooArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceQualarooArgs()
        {
        }
        public static new GetSourceQualarooArgs Empty => new GetSourceQualarooArgs();
    }

    public sealed class GetSourceQualarooInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceQualarooInvokeArgs()
        {
        }
        public static new GetSourceQualarooInvokeArgs Empty => new GetSourceQualarooInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceQualarooResult
    {
        public readonly Outputs.GetSourceQualarooConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceQualarooResult(
            Outputs.GetSourceQualarooConfiguration configuration,

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
