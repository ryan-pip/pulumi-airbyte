// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceUsCensus
    {
        public static Task<GetSourceUsCensusResult> InvokeAsync(GetSourceUsCensusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceUsCensusResult>("airbyte:index/getSourceUsCensus:getSourceUsCensus", args ?? new GetSourceUsCensusArgs(), options.WithDefaults());

        public static Output<GetSourceUsCensusResult> Invoke(GetSourceUsCensusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceUsCensusResult>("airbyte:index/getSourceUsCensus:getSourceUsCensus", args ?? new GetSourceUsCensusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceUsCensusArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceUsCensusArgs()
        {
        }
        public static new GetSourceUsCensusArgs Empty => new GetSourceUsCensusArgs();
    }

    public sealed class GetSourceUsCensusInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceUsCensusInvokeArgs()
        {
        }
        public static new GetSourceUsCensusInvokeArgs Empty => new GetSourceUsCensusInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceUsCensusResult
    {
        public readonly Outputs.GetSourceUsCensusConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceUsCensusResult(
            Outputs.GetSourceUsCensusConfiguration configuration,

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