// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationAWSDatalakeConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationAwsDatalakeAuthenticationModeIamRole")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationCredentialsDestinationAwsDatalakeAuthenticationModeIamRoleArgs>? DestinationAwsDatalakeAuthenticationModeIamRole { get; set; }

        [Input("destinationAwsDatalakeAuthenticationModeIamUser")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationCredentialsDestinationAwsDatalakeAuthenticationModeIamUserArgs>? DestinationAwsDatalakeAuthenticationModeIamUser { get; set; }

        [Input("destinationAwsDatalakeUpdateAuthenticationModeIamRole")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationCredentialsDestinationAwsDatalakeUpdateAuthenticationModeIamRoleArgs>? DestinationAwsDatalakeUpdateAuthenticationModeIamRole { get; set; }

        [Input("destinationAwsDatalakeUpdateAuthenticationModeIamUser")]
        public Input<Inputs.DestinationAWSDatalakeConfigurationCredentialsDestinationAwsDatalakeUpdateAuthenticationModeIamUserArgs>? DestinationAwsDatalakeUpdateAuthenticationModeIamUser { get; set; }

        public DestinationAWSDatalakeConfigurationCredentialsArgs()
        {
        }
        public static new DestinationAWSDatalakeConfigurationCredentialsArgs Empty => new DestinationAWSDatalakeConfigurationCredentialsArgs();
    }
}
