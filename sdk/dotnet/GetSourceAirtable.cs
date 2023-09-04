// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceAirtable
    {
        public static Task<GetSourceAirtableResult> InvokeAsync(GetSourceAirtableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceAirtableResult>("airbyte:index/getSourceAirtable:getSourceAirtable", args ?? new GetSourceAirtableArgs(), options.WithDefaults());

        public static Output<GetSourceAirtableResult> Invoke(GetSourceAirtableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceAirtableResult>("airbyte:index/getSourceAirtable:getSourceAirtable", args ?? new GetSourceAirtableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceAirtableArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceAirtableArgs()
        {
        }
        public static new GetSourceAirtableArgs Empty => new GetSourceAirtableArgs();
    }

    public sealed class GetSourceAirtableInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceAirtableInvokeArgs()
        {
        }
        public static new GetSourceAirtableInvokeArgs Empty => new GetSourceAirtableInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceAirtableResult
    {
        public readonly Outputs.GetSourceAirtableConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceAirtableResult(
            Outputs.GetSourceAirtableConfiguration configuration,

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