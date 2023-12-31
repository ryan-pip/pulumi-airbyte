// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceFreshdesk
    {
        public static Task<GetSourceFreshdeskResult> InvokeAsync(GetSourceFreshdeskArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceFreshdeskResult>("airbyte:index/getSourceFreshdesk:getSourceFreshdesk", args ?? new GetSourceFreshdeskArgs(), options.WithDefaults());

        public static Output<GetSourceFreshdeskResult> Invoke(GetSourceFreshdeskInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceFreshdeskResult>("airbyte:index/getSourceFreshdesk:getSourceFreshdesk", args ?? new GetSourceFreshdeskInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceFreshdeskArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceFreshdeskArgs()
        {
        }
        public static new GetSourceFreshdeskArgs Empty => new GetSourceFreshdeskArgs();
    }

    public sealed class GetSourceFreshdeskInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceFreshdeskInvokeArgs()
        {
        }
        public static new GetSourceFreshdeskInvokeArgs Empty => new GetSourceFreshdeskInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceFreshdeskResult
    {
        public readonly Outputs.GetSourceFreshdeskConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceFreshdeskResult(
            Outputs.GetSourceFreshdeskConfiguration configuration,

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
