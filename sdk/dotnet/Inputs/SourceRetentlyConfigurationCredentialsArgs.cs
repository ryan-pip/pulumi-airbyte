// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceRetentlyConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth")]
        public Input<Inputs.SourceRetentlyConfigurationCredentialsSourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuthArgs>? SourceRetentlyAuthenticationMechanismAuthenticateViaRetentlyOAuth { get; set; }

        [Input("sourceRetentlyAuthenticationMechanismAuthenticateWithApiToken")]
        public Input<Inputs.SourceRetentlyConfigurationCredentialsSourceRetentlyAuthenticationMechanismAuthenticateWithApiTokenArgs>? SourceRetentlyAuthenticationMechanismAuthenticateWithApiToken { get; set; }

        [Input("sourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth")]
        public Input<Inputs.SourceRetentlyConfigurationCredentialsSourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuthArgs>? SourceRetentlyUpdateAuthenticationMechanismAuthenticateViaRetentlyOAuth { get; set; }

        [Input("sourceRetentlyUpdateAuthenticationMechanismAuthenticateWithApiToken")]
        public Input<Inputs.SourceRetentlyConfigurationCredentialsSourceRetentlyUpdateAuthenticationMechanismAuthenticateWithApiTokenArgs>? SourceRetentlyUpdateAuthenticationMechanismAuthenticateWithApiToken { get; set; }

        public SourceRetentlyConfigurationCredentialsArgs()
        {
        }
        public static new SourceRetentlyConfigurationCredentialsArgs Empty => new SourceRetentlyConfigurationCredentialsArgs();
    }
}