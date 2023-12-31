// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceFreshcaller
    {
        public static Task<GetSourceFreshcallerResult> InvokeAsync(GetSourceFreshcallerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceFreshcallerResult>("airbyte:index/getSourceFreshcaller:getSourceFreshcaller", args ?? new GetSourceFreshcallerArgs(), options.WithDefaults());

        public static Output<GetSourceFreshcallerResult> Invoke(GetSourceFreshcallerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceFreshcallerResult>("airbyte:index/getSourceFreshcaller:getSourceFreshcaller", args ?? new GetSourceFreshcallerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceFreshcallerArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceFreshcallerArgs()
        {
        }
        public static new GetSourceFreshcallerArgs Empty => new GetSourceFreshcallerArgs();
    }

    public sealed class GetSourceFreshcallerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceFreshcallerInvokeArgs()
        {
        }
        public static new GetSourceFreshcallerInvokeArgs Empty => new GetSourceFreshcallerInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceFreshcallerResult
    {
        public readonly Outputs.GetSourceFreshcallerConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceFreshcallerResult(
            Outputs.GetSourceFreshcallerConfiguration configuration,

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
