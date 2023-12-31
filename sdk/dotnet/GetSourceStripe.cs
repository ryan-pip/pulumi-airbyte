// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceStripe
    {
        public static Task<GetSourceStripeResult> InvokeAsync(GetSourceStripeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceStripeResult>("airbyte:index/getSourceStripe:getSourceStripe", args ?? new GetSourceStripeArgs(), options.WithDefaults());

        public static Output<GetSourceStripeResult> Invoke(GetSourceStripeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceStripeResult>("airbyte:index/getSourceStripe:getSourceStripe", args ?? new GetSourceStripeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceStripeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceStripeArgs()
        {
        }
        public static new GetSourceStripeArgs Empty => new GetSourceStripeArgs();
    }

    public sealed class GetSourceStripeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceStripeInvokeArgs()
        {
        }
        public static new GetSourceStripeInvokeArgs Empty => new GetSourceStripeInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceStripeResult
    {
        public readonly Outputs.GetSourceStripeConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceStripeResult(
            Outputs.GetSourceStripeConfiguration configuration,

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
