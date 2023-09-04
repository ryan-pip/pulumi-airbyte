// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourcePosthog
    {
        public static Task<GetSourcePosthogResult> InvokeAsync(GetSourcePosthogArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourcePosthogResult>("airbyte:index/getSourcePosthog:getSourcePosthog", args ?? new GetSourcePosthogArgs(), options.WithDefaults());

        public static Output<GetSourcePosthogResult> Invoke(GetSourcePosthogInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourcePosthogResult>("airbyte:index/getSourcePosthog:getSourcePosthog", args ?? new GetSourcePosthogInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourcePosthogArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourcePosthogArgs()
        {
        }
        public static new GetSourcePosthogArgs Empty => new GetSourcePosthogArgs();
    }

    public sealed class GetSourcePosthogInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourcePosthogInvokeArgs()
        {
        }
        public static new GetSourcePosthogInvokeArgs Empty => new GetSourcePosthogInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourcePosthogResult
    {
        public readonly Outputs.GetSourcePosthogConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourcePosthogResult(
            Outputs.GetSourcePosthogConfiguration configuration,

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