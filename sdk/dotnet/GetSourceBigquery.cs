// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceBigquery
    {
        public static Task<GetSourceBigqueryResult> InvokeAsync(GetSourceBigqueryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceBigqueryResult>("airbyte:index/getSourceBigquery:getSourceBigquery", args ?? new GetSourceBigqueryArgs(), options.WithDefaults());

        public static Output<GetSourceBigqueryResult> Invoke(GetSourceBigqueryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceBigqueryResult>("airbyte:index/getSourceBigquery:getSourceBigquery", args ?? new GetSourceBigqueryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceBigqueryArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceBigqueryArgs()
        {
        }
        public static new GetSourceBigqueryArgs Empty => new GetSourceBigqueryArgs();
    }

    public sealed class GetSourceBigqueryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceBigqueryInvokeArgs()
        {
        }
        public static new GetSourceBigqueryInvokeArgs Empty => new GetSourceBigqueryInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceBigqueryResult
    {
        public readonly Outputs.GetSourceBigqueryConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceBigqueryResult(
            Outputs.GetSourceBigqueryConfiguration configuration,

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
