// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    [AirbyteResourceType("airbyte:index/sourceBigcommerce:SourceBigcommerce")]
    public partial class SourceBigcommerce : global::Pulumi.CustomResource
    {
        [Output("configuration")]
        public Output<Outputs.SourceBigcommerceConfiguration> Configuration { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional secretID obtained through the public API OAuth redirect flow.
        /// </summary>
        [Output("secretId")]
        public Output<string?> SecretId { get; private set; } = null!;

        [Output("sourceId")]
        public Output<string> SourceId { get; private set; } = null!;

        [Output("sourceType")]
        public Output<string> SourceType { get; private set; } = null!;

        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a SourceBigcommerce resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SourceBigcommerce(string name, SourceBigcommerceArgs args, CustomResourceOptions? options = null)
            : base("airbyte:index/sourceBigcommerce:SourceBigcommerce", name, args ?? new SourceBigcommerceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SourceBigcommerce(string name, Input<string> id, SourceBigcommerceState? state = null, CustomResourceOptions? options = null)
            : base("airbyte:index/sourceBigcommerce:SourceBigcommerce", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SourceBigcommerce resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SourceBigcommerce Get(string name, Input<string> id, SourceBigcommerceState? state = null, CustomResourceOptions? options = null)
        {
            return new SourceBigcommerce(name, id, state, options);
        }
    }

    public sealed class SourceBigcommerceArgs : global::Pulumi.ResourceArgs
    {
        [Input("configuration", required: true)]
        public Input<Inputs.SourceBigcommerceConfigurationArgs> Configuration { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Optional secretID obtained through the public API OAuth redirect flow.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public SourceBigcommerceArgs()
        {
        }
        public static new SourceBigcommerceArgs Empty => new SourceBigcommerceArgs();
    }

    public sealed class SourceBigcommerceState : global::Pulumi.ResourceArgs
    {
        [Input("configuration")]
        public Input<Inputs.SourceBigcommerceConfigurationGetArgs>? Configuration { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional secretID obtained through the public API OAuth redirect flow.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public SourceBigcommerceState()
        {
        }
        public static new SourceBigcommerceState Empty => new SourceBigcommerceState();
    }
}