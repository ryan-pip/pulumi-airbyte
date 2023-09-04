// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceHubplanner
    {
        public static Task<GetSourceHubplannerResult> InvokeAsync(GetSourceHubplannerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceHubplannerResult>("airbyte:index/getSourceHubplanner:getSourceHubplanner", args ?? new GetSourceHubplannerArgs(), options.WithDefaults());

        public static Output<GetSourceHubplannerResult> Invoke(GetSourceHubplannerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceHubplannerResult>("airbyte:index/getSourceHubplanner:getSourceHubplanner", args ?? new GetSourceHubplannerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceHubplannerArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceHubplannerArgs()
        {
        }
        public static new GetSourceHubplannerArgs Empty => new GetSourceHubplannerArgs();
    }

    public sealed class GetSourceHubplannerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceHubplannerInvokeArgs()
        {
        }
        public static new GetSourceHubplannerInvokeArgs Empty => new GetSourceHubplannerInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceHubplannerResult
    {
        public readonly Outputs.GetSourceHubplannerConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceHubplannerResult(
            Outputs.GetSourceHubplannerConfiguration configuration,

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