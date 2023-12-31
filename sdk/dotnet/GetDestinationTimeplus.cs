// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationTimeplus
    {
        public static Task<GetDestinationTimeplusResult> InvokeAsync(GetDestinationTimeplusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationTimeplusResult>("airbyte:index/getDestinationTimeplus:getDestinationTimeplus", args ?? new GetDestinationTimeplusArgs(), options.WithDefaults());

        public static Output<GetDestinationTimeplusResult> Invoke(GetDestinationTimeplusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationTimeplusResult>("airbyte:index/getDestinationTimeplus:getDestinationTimeplus", args ?? new GetDestinationTimeplusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationTimeplusArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationTimeplusArgs()
        {
        }
        public static new GetDestinationTimeplusArgs Empty => new GetDestinationTimeplusArgs();
    }

    public sealed class GetDestinationTimeplusInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationTimeplusInvokeArgs()
        {
        }
        public static new GetDestinationTimeplusInvokeArgs Empty => new GetDestinationTimeplusInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationTimeplusResult
    {
        public readonly Outputs.GetDestinationTimeplusConfiguration Configuration;
        public readonly string DestinationId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetDestinationTimeplusResult(
            Outputs.GetDestinationTimeplusConfiguration configuration,

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
