// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceXkcd
    {
        public static Task<GetSourceXkcdResult> InvokeAsync(GetSourceXkcdArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceXkcdResult>("airbyte:index/getSourceXkcd:getSourceXkcd", args ?? new GetSourceXkcdArgs(), options.WithDefaults());

        public static Output<GetSourceXkcdResult> Invoke(GetSourceXkcdInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceXkcdResult>("airbyte:index/getSourceXkcd:getSourceXkcd", args ?? new GetSourceXkcdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceXkcdArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceXkcdArgs()
        {
        }
        public static new GetSourceXkcdArgs Empty => new GetSourceXkcdArgs();
    }

    public sealed class GetSourceXkcdInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceXkcdInvokeArgs()
        {
        }
        public static new GetSourceXkcdInvokeArgs Empty => new GetSourceXkcdInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceXkcdResult
    {
        public readonly Outputs.GetSourceXkcdConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceXkcdResult(
            Outputs.GetSourceXkcdConfiguration configuration,

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
