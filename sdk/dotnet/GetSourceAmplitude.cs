// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceAmplitude
    {
        public static Task<GetSourceAmplitudeResult> InvokeAsync(GetSourceAmplitudeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceAmplitudeResult>("airbyte:index/getSourceAmplitude:getSourceAmplitude", args ?? new GetSourceAmplitudeArgs(), options.WithDefaults());

        public static Output<GetSourceAmplitudeResult> Invoke(GetSourceAmplitudeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceAmplitudeResult>("airbyte:index/getSourceAmplitude:getSourceAmplitude", args ?? new GetSourceAmplitudeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceAmplitudeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceAmplitudeArgs()
        {
        }
        public static new GetSourceAmplitudeArgs Empty => new GetSourceAmplitudeArgs();
    }

    public sealed class GetSourceAmplitudeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceAmplitudeInvokeArgs()
        {
        }
        public static new GetSourceAmplitudeInvokeArgs Empty => new GetSourceAmplitudeInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceAmplitudeResult
    {
        public readonly Outputs.GetSourceAmplitudeConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceAmplitudeResult(
            Outputs.GetSourceAmplitudeConfiguration configuration,

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
