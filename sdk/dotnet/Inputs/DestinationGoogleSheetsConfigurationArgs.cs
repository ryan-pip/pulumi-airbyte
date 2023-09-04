// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationGoogleSheetsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentials", required: true)]
        public Input<Inputs.DestinationGoogleSheetsConfigurationCredentialsArgs> Credentials { get; set; } = null!;

        [Input("destinationType", required: true)]
        public Input<string> DestinationType { get; set; } = null!;

        [Input("spreadsheetId", required: true)]
        public Input<string> SpreadsheetId { get; set; } = null!;

        public DestinationGoogleSheetsConfigurationArgs()
        {
        }
        public static new DestinationGoogleSheetsConfigurationArgs Empty => new DestinationGoogleSheetsConfigurationArgs();
    }
}