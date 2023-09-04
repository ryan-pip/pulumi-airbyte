// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationDevNull
    {
        public static Task<GetDestinationDevNullResult> InvokeAsync(GetDestinationDevNullArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationDevNullResult>("airbyte:index/getDestinationDevNull:getDestinationDevNull", args ?? new GetDestinationDevNullArgs(), options.WithDefaults());

        public static Output<GetDestinationDevNullResult> Invoke(GetDestinationDevNullInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationDevNullResult>("airbyte:index/getDestinationDevNull:getDestinationDevNull", args ?? new GetDestinationDevNullInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationDevNullArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationDevNullArgs()
        {
        }
        public static new GetDestinationDevNullArgs Empty => new GetDestinationDevNullArgs();
    }

    public sealed class GetDestinationDevNullInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationDevNullInvokeArgs()
        {
        }
        public static new GetDestinationDevNullInvokeArgs Empty => new GetDestinationDevNullInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationDevNullResult
    {
        public readonly Outputs.GetDestinationDevNullConfiguration Configuration;
        public readonly string DestinationId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetDestinationDevNullResult(
            Outputs.GetDestinationDevNullConfiguration configuration,

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