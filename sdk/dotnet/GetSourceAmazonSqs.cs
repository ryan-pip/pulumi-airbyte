// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceAmazonSqs
    {
        public static Task<GetSourceAmazonSqsResult> InvokeAsync(GetSourceAmazonSqsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceAmazonSqsResult>("airbyte:index/getSourceAmazonSqs:getSourceAmazonSqs", args ?? new GetSourceAmazonSqsArgs(), options.WithDefaults());

        public static Output<GetSourceAmazonSqsResult> Invoke(GetSourceAmazonSqsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceAmazonSqsResult>("airbyte:index/getSourceAmazonSqs:getSourceAmazonSqs", args ?? new GetSourceAmazonSqsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceAmazonSqsArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceAmazonSqsArgs()
        {
        }
        public static new GetSourceAmazonSqsArgs Empty => new GetSourceAmazonSqsArgs();
    }

    public sealed class GetSourceAmazonSqsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceAmazonSqsInvokeArgs()
        {
        }
        public static new GetSourceAmazonSqsInvokeArgs Empty => new GetSourceAmazonSqsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceAmazonSqsResult
    {
        public readonly Outputs.GetSourceAmazonSqsConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceAmazonSqsResult(
            Outputs.GetSourceAmazonSqsConfiguration configuration,

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
