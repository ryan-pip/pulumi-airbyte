// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceGithub
    {
        public static Task<GetSourceGithubResult> InvokeAsync(GetSourceGithubArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceGithubResult>("airbyte:index/getSourceGithub:getSourceGithub", args ?? new GetSourceGithubArgs(), options.WithDefaults());

        public static Output<GetSourceGithubResult> Invoke(GetSourceGithubInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceGithubResult>("airbyte:index/getSourceGithub:getSourceGithub", args ?? new GetSourceGithubInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceGithubArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceGithubArgs()
        {
        }
        public static new GetSourceGithubArgs Empty => new GetSourceGithubArgs();
    }

    public sealed class GetSourceGithubInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceGithubInvokeArgs()
        {
        }
        public static new GetSourceGithubInvokeArgs Empty => new GetSourceGithubInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceGithubResult
    {
        public readonly Outputs.GetSourceGithubConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceGithubResult(
            Outputs.GetSourceGithubConfiguration configuration,

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
