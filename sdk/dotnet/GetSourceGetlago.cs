// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceGetlago
    {
        public static Task<GetSourceGetlagoResult> InvokeAsync(GetSourceGetlagoArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceGetlagoResult>("airbyte:index/getSourceGetlago:getSourceGetlago", args ?? new GetSourceGetlagoArgs(), options.WithDefaults());

        public static Output<GetSourceGetlagoResult> Invoke(GetSourceGetlagoInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceGetlagoResult>("airbyte:index/getSourceGetlago:getSourceGetlago", args ?? new GetSourceGetlagoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceGetlagoArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceGetlagoArgs()
        {
        }
        public static new GetSourceGetlagoArgs Empty => new GetSourceGetlagoArgs();
    }

    public sealed class GetSourceGetlagoInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceGetlagoInvokeArgs()
        {
        }
        public static new GetSourceGetlagoInvokeArgs Empty => new GetSourceGetlagoInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceGetlagoResult
    {
        public readonly Outputs.GetSourceGetlagoConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceGetlagoResult(
            Outputs.GetSourceGetlagoConfiguration configuration,

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
