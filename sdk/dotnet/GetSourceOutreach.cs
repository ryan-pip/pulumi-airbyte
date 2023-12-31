// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceOutreach
    {
        public static Task<GetSourceOutreachResult> InvokeAsync(GetSourceOutreachArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceOutreachResult>("airbyte:index/getSourceOutreach:getSourceOutreach", args ?? new GetSourceOutreachArgs(), options.WithDefaults());

        public static Output<GetSourceOutreachResult> Invoke(GetSourceOutreachInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceOutreachResult>("airbyte:index/getSourceOutreach:getSourceOutreach", args ?? new GetSourceOutreachInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceOutreachArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceOutreachArgs()
        {
        }
        public static new GetSourceOutreachArgs Empty => new GetSourceOutreachArgs();
    }

    public sealed class GetSourceOutreachInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceOutreachInvokeArgs()
        {
        }
        public static new GetSourceOutreachInvokeArgs Empty => new GetSourceOutreachInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceOutreachResult
    {
        public readonly Outputs.GetSourceOutreachConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceOutreachResult(
            Outputs.GetSourceOutreachConfiguration configuration,

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
