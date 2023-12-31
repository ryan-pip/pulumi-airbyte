// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceRailz
    {
        public static Task<GetSourceRailzResult> InvokeAsync(GetSourceRailzArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceRailzResult>("airbyte:index/getSourceRailz:getSourceRailz", args ?? new GetSourceRailzArgs(), options.WithDefaults());

        public static Output<GetSourceRailzResult> Invoke(GetSourceRailzInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceRailzResult>("airbyte:index/getSourceRailz:getSourceRailz", args ?? new GetSourceRailzInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceRailzArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceRailzArgs()
        {
        }
        public static new GetSourceRailzArgs Empty => new GetSourceRailzArgs();
    }

    public sealed class GetSourceRailzInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceRailzInvokeArgs()
        {
        }
        public static new GetSourceRailzInvokeArgs Empty => new GetSourceRailzInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceRailzResult
    {
        public readonly Outputs.GetSourceRailzConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceRailzResult(
            Outputs.GetSourceRailzConfiguration configuration,

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
