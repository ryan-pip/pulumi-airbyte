// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationS3Glue
    {
        public static Task<GetDestinationS3GlueResult> InvokeAsync(GetDestinationS3GlueArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationS3GlueResult>("airbyte:index/getDestinationS3Glue:getDestinationS3Glue", args ?? new GetDestinationS3GlueArgs(), options.WithDefaults());

        public static Output<GetDestinationS3GlueResult> Invoke(GetDestinationS3GlueInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationS3GlueResult>("airbyte:index/getDestinationS3Glue:getDestinationS3Glue", args ?? new GetDestinationS3GlueInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationS3GlueArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationS3GlueArgs()
        {
        }
        public static new GetDestinationS3GlueArgs Empty => new GetDestinationS3GlueArgs();
    }

    public sealed class GetDestinationS3GlueInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationS3GlueInvokeArgs()
        {
        }
        public static new GetDestinationS3GlueInvokeArgs Empty => new GetDestinationS3GlueInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationS3GlueResult
    {
        public readonly Outputs.GetDestinationS3GlueConfiguration Configuration;
        public readonly string DestinationId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetDestinationS3GlueResult(
            Outputs.GetDestinationS3GlueConfiguration configuration,

            string destinationId,

            string id,

            string name,

            string workspaceId)
        {
            Configuration = configuration;
            DestinationId = destinationId;
            Id = id;
            Name = name;
            WorkspaceId = workspaceId;
        }
    }
}
