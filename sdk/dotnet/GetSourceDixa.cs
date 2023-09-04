// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceDixa
    {
        public static Task<GetSourceDixaResult> InvokeAsync(GetSourceDixaArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceDixaResult>("airbyte:index/getSourceDixa:getSourceDixa", args ?? new GetSourceDixaArgs(), options.WithDefaults());

        public static Output<GetSourceDixaResult> Invoke(GetSourceDixaInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceDixaResult>("airbyte:index/getSourceDixa:getSourceDixa", args ?? new GetSourceDixaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceDixaArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceDixaArgs()
        {
        }
        public static new GetSourceDixaArgs Empty => new GetSourceDixaArgs();
    }

    public sealed class GetSourceDixaInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceDixaInvokeArgs()
        {
        }
        public static new GetSourceDixaInvokeArgs Empty => new GetSourceDixaInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceDixaResult
    {
        public readonly Outputs.GetSourceDixaConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceDixaResult(
            Outputs.GetSourceDixaConfiguration configuration,

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