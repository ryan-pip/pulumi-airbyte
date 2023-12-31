// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceAuth0
    {
        public static Task<GetSourceAuth0Result> InvokeAsync(GetSourceAuth0Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceAuth0Result>("airbyte:index/getSourceAuth0:getSourceAuth0", args ?? new GetSourceAuth0Args(), options.WithDefaults());

        public static Output<GetSourceAuth0Result> Invoke(GetSourceAuth0InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceAuth0Result>("airbyte:index/getSourceAuth0:getSourceAuth0", args ?? new GetSourceAuth0InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceAuth0Args : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceAuth0Args()
        {
        }
        public static new GetSourceAuth0Args Empty => new GetSourceAuth0Args();
    }

    public sealed class GetSourceAuth0InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceAuth0InvokeArgs()
        {
        }
        public static new GetSourceAuth0InvokeArgs Empty => new GetSourceAuth0InvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceAuth0Result
    {
        public readonly Outputs.GetSourceAuth0Configuration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceAuth0Result(
            Outputs.GetSourceAuth0Configuration configuration,

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
