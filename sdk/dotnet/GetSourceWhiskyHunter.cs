// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceWhiskyHunter
    {
        public static Task<GetSourceWhiskyHunterResult> InvokeAsync(GetSourceWhiskyHunterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceWhiskyHunterResult>("airbyte:index/getSourceWhiskyHunter:getSourceWhiskyHunter", args ?? new GetSourceWhiskyHunterArgs(), options.WithDefaults());

        public static Output<GetSourceWhiskyHunterResult> Invoke(GetSourceWhiskyHunterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceWhiskyHunterResult>("airbyte:index/getSourceWhiskyHunter:getSourceWhiskyHunter", args ?? new GetSourceWhiskyHunterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceWhiskyHunterArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceWhiskyHunterArgs()
        {
        }
        public static new GetSourceWhiskyHunterArgs Empty => new GetSourceWhiskyHunterArgs();
    }

    public sealed class GetSourceWhiskyHunterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceWhiskyHunterInvokeArgs()
        {
        }
        public static new GetSourceWhiskyHunterInvokeArgs Empty => new GetSourceWhiskyHunterInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceWhiskyHunterResult
    {
        public readonly Outputs.GetSourceWhiskyHunterConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceWhiskyHunterResult(
            Outputs.GetSourceWhiskyHunterConfiguration configuration,

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