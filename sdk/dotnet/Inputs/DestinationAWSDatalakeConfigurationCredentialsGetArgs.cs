// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationAWSDatalakeConfigurationCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationAwsDatalakeAuthenticationModeIamRole")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationCredentialsDestinationAwsDatalakeAuthenticationModeIamRoleGetArgs>? DestinationAwsDatalakeAuthenticationModeIamRole { get; set; }

        [Input("destinationAwsDatalakeAuthenticationModeIamUser")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationCredentialsDestinationAwsDatalakeAuthenticationModeIamUserGetArgs>? DestinationAwsDatalakeAuthenticationModeIamUser { get; set; }

        [Input("destinationAwsDatalakeUpdateAuthenticationModeIamRole")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationCredentialsDestinationAwsDatalakeUpdateAuthenticationModeIamRoleGetArgs>? DestinationAwsDatalakeUpdateAuthenticationModeIamRole { get; set; }

        [Input("destinationAwsDatalakeUpdateAuthenticationModeIamUser")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationCredentialsDestinationAwsDatalakeUpdateAuthenticationModeIamUserGetArgs>? DestinationAwsDatalakeUpdateAuthenticationModeIamUser { get; set; }

        public DestinationAWSDatalakeConfigurationCredentialsGetArgs()
        {
        }
        public static new DestinationAWSDatalakeConfigurationCredentialsGetArgs Empty => new DestinationAWSDatalakeConfigurationCredentialsGetArgs();
    }
}
