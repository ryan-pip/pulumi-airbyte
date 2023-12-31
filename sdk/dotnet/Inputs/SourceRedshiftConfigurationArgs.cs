// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceRedshiftConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("jdbcUrlParams")]
        public Input<string>? JdbcUrlParams { get; set; }

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("schemas")]
        private InputList<string>? _schemas;
        public InputList<string> Schemas
        {
            get => _schemas ?? (_schemas = new InputList<string>());
            set => _schemas = value;
        }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SourceRedshiftConfigurationArgs()
        {
        }
        public static new SourceRedshiftConfigurationArgs Empty => new SourceRedshiftConfigurationArgs();
    }
}
