// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceOracleConfigurationConnectionDataArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceOracleConnectByServiceName")]
        public Input<Inputs.SourceOracleConfigurationConnectionDataSourceOracleConnectByServiceNameArgs>? SourceOracleConnectByServiceName { get; set; }

        [Input("sourceOracleConnectBySystemIdSid")]
        public Input<Inputs.SourceOracleConfigurationConnectionDataSourceOracleConnectBySystemIdSidArgs>? SourceOracleConnectBySystemIdSid { get; set; }

        [Input("sourceOracleUpdateConnectByServiceName")]
        public Input<Inputs.SourceOracleConfigurationConnectionDataSourceOracleUpdateConnectByServiceNameArgs>? SourceOracleUpdateConnectByServiceName { get; set; }

        [Input("sourceOracleUpdateConnectBySystemIdSid")]
        public Input<Inputs.SourceOracleConfigurationConnectionDataSourceOracleUpdateConnectBySystemIdSidArgs>? SourceOracleUpdateConnectBySystemIdSid { get; set; }

        public SourceOracleConfigurationConnectionDataArgs()
        {
        }
        public static new SourceOracleConfigurationConnectionDataArgs Empty => new SourceOracleConfigurationConnectionDataArgs();
    }
}
