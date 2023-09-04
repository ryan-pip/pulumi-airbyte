// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationAzurBlobStorage
    {
        public static Task<GetDestinationAzurBlobStorageResult> InvokeAsync(GetDestinationAzurBlobStorageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationAzurBlobStorageResult>("airbyte:index/getDestinationAzurBlobStorage:getDestinationAzurBlobStorage", args ?? new GetDestinationAzurBlobStorageArgs(), options.WithDefaults());

        public static Output<GetDestinationAzurBlobStorageResult> Invoke(GetDestinationAzurBlobStorageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationAzurBlobStorageResult>("airbyte:index/getDestinationAzurBlobStorage:getDestinationAzurBlobStorage", args ?? new GetDestinationAzurBlobStorageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationAzurBlobStorageArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationAzurBlobStorageArgs()
        {
        }
        public static new GetDestinationAzurBlobStorageArgs Empty => new GetDestinationAzurBlobStorageArgs();
    }

    public sealed class GetDestinationAzurBlobStorageInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationAzurBlobStorageInvokeArgs()
        {
        }
        public static new GetDestinationAzurBlobStorageInvokeArgs Empty => new GetDestinationAzurBlobStorageInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationAzurBlobStorageResult
    {
        public readonly Outputs.GetDestinationAzurBlobStorageConfiguration Configuration;
        public readonly string DestinationId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetDestinationAzurBlobStorageResult(
            Outputs.GetDestinationAzurBlobStorageConfiguration configuration,

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