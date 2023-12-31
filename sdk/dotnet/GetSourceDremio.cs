// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceDremio
    {
        public static Task<GetSourceDremioResult> InvokeAsync(GetSourceDremioArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceDremioResult>("airbyte:index/getSourceDremio:getSourceDremio", args ?? new GetSourceDremioArgs(), options.WithDefaults());

        public static Output<GetSourceDremioResult> Invoke(GetSourceDremioInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceDremioResult>("airbyte:index/getSourceDremio:getSourceDremio", args ?? new GetSourceDremioInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceDremioArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceDremioArgs()
        {
        }
        public static new GetSourceDremioArgs Empty => new GetSourceDremioArgs();
    }

    public sealed class GetSourceDremioInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceDremioInvokeArgs()
        {
        }
        public static new GetSourceDremioInvokeArgs Empty => new GetSourceDremioInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceDremioResult
    {
        public readonly Outputs.GetSourceDremioConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceDremioResult(
            Outputs.GetSourceDremioConfiguration configuration,

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
