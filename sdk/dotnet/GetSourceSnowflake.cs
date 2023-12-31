// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceSnowflake
    {
        public static Task<GetSourceSnowflakeResult> InvokeAsync(GetSourceSnowflakeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceSnowflakeResult>("airbyte:index/getSourceSnowflake:getSourceSnowflake", args ?? new GetSourceSnowflakeArgs(), options.WithDefaults());

        public static Output<GetSourceSnowflakeResult> Invoke(GetSourceSnowflakeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceSnowflakeResult>("airbyte:index/getSourceSnowflake:getSourceSnowflake", args ?? new GetSourceSnowflakeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceSnowflakeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceSnowflakeArgs()
        {
        }
        public static new GetSourceSnowflakeArgs Empty => new GetSourceSnowflakeArgs();
    }

    public sealed class GetSourceSnowflakeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceSnowflakeInvokeArgs()
        {
        }
        public static new GetSourceSnowflakeInvokeArgs Empty => new GetSourceSnowflakeInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceSnowflakeResult
    {
        public readonly Outputs.GetSourceSnowflakeConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceSnowflakeResult(
            Outputs.GetSourceSnowflakeConfiguration configuration,

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
