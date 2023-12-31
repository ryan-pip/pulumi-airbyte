// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceTiktokMarketing
    {
        public static Task<GetSourceTiktokMarketingResult> InvokeAsync(GetSourceTiktokMarketingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceTiktokMarketingResult>("airbyte:index/getSourceTiktokMarketing:getSourceTiktokMarketing", args ?? new GetSourceTiktokMarketingArgs(), options.WithDefaults());

        public static Output<GetSourceTiktokMarketingResult> Invoke(GetSourceTiktokMarketingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceTiktokMarketingResult>("airbyte:index/getSourceTiktokMarketing:getSourceTiktokMarketing", args ?? new GetSourceTiktokMarketingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceTiktokMarketingArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceTiktokMarketingArgs()
        {
        }
        public static new GetSourceTiktokMarketingArgs Empty => new GetSourceTiktokMarketingArgs();
    }

    public sealed class GetSourceTiktokMarketingInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceTiktokMarketingInvokeArgs()
        {
        }
        public static new GetSourceTiktokMarketingInvokeArgs Empty => new GetSourceTiktokMarketingInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceTiktokMarketingResult
    {
        public readonly Outputs.GetSourceTiktokMarketingConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceTiktokMarketingResult(
            Outputs.GetSourceTiktokMarketingConfiguration configuration,

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
