// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceZuora
    {
        public static Task<GetSourceZuoraResult> InvokeAsync(GetSourceZuoraArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceZuoraResult>("airbyte:index/getSourceZuora:getSourceZuora", args ?? new GetSourceZuoraArgs(), options.WithDefaults());

        public static Output<GetSourceZuoraResult> Invoke(GetSourceZuoraInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceZuoraResult>("airbyte:index/getSourceZuora:getSourceZuora", args ?? new GetSourceZuoraInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceZuoraArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceZuoraArgs()
        {
        }
        public static new GetSourceZuoraArgs Empty => new GetSourceZuoraArgs();
    }

    public sealed class GetSourceZuoraInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceZuoraInvokeArgs()
        {
        }
        public static new GetSourceZuoraInvokeArgs Empty => new GetSourceZuoraInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceZuoraResult
    {
        public readonly Outputs.GetSourceZuoraConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceZuoraResult(
            Outputs.GetSourceZuoraConfiguration configuration,

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
