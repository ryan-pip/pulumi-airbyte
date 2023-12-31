// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceGoogleSheets
    {
        public static Task<GetSourceGoogleSheetsResult> InvokeAsync(GetSourceGoogleSheetsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceGoogleSheetsResult>("airbyte:index/getSourceGoogleSheets:getSourceGoogleSheets", args ?? new GetSourceGoogleSheetsArgs(), options.WithDefaults());

        public static Output<GetSourceGoogleSheetsResult> Invoke(GetSourceGoogleSheetsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceGoogleSheetsResult>("airbyte:index/getSourceGoogleSheets:getSourceGoogleSheets", args ?? new GetSourceGoogleSheetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceGoogleSheetsArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public string? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceGoogleSheetsArgs()
        {
        }
        public static new GetSourceGoogleSheetsArgs Empty => new GetSourceGoogleSheetsArgs();
    }

    public sealed class GetSourceGoogleSheetsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceGoogleSheetsInvokeArgs()
        {
        }
        public static new GetSourceGoogleSheetsInvokeArgs Empty => new GetSourceGoogleSheetsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceGoogleSheetsResult
    {
        public readonly Outputs.GetSourceGoogleSheetsConfiguration Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? SecretId;
        public readonly string SourceId;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceGoogleSheetsResult(
            Outputs.GetSourceGoogleSheetsConfiguration configuration,

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
