// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceAsana
    {
        public static Task<GetSourceAsanaResult> InvokeAsync(GetSourceAsanaArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceAsanaResult>("airbyte:index/getSourceAsana:getSourceAsana", args ?? new GetSourceAsanaArgs(), options.WithDefaults());

        public static Output<GetSourceAsanaResult> Invoke(GetSourceAsanaInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceAsanaResult>("airbyte:index/getSourceAsana:getSourceAsana", args ?? new GetSourceAsanaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceAsanaArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceAsanaArgs()
        {
        }
        public static new GetSourceAsanaArgs Empty => new GetSourceAsanaArgs();
    }

    public sealed class GetSourceAsanaInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceAsanaInvokeArgs()
        {
        }
        public static new GetSourceAsanaInvokeArgs Empty => new GetSourceAsanaInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceAsanaResult
    {
        public readonly Outputs.GetSourceAsanaConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceAsanaResult(
            Outputs.GetSourceAsanaConfiguration configuration,

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
