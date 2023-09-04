// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceNetsuiteConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("consumerKey", required: true)]
        public Input<string> ConsumerKey { get; set; } = null!;

        [Input("consumerSecret", required: true)]
        public Input<string> ConsumerSecret { get; set; } = null!;

        [Input("objectTypes")]
        private InputList<string>? _objectTypes;
        public InputList<string> ObjectTypes
        {
            get => _objectTypes ?? (_objectTypes = new InputList<string>());
            set => _objectTypes = value;
        }

        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        [Input("startDatetime", required: true)]
        public Input<string> StartDatetime { get; set; } = null!;

        [Input("tokenKey", required: true)]
        public Input<string> TokenKey { get; set; } = null!;

        [Input("tokenSecret", required: true)]
        public Input<string> TokenSecret { get; set; } = null!;

        [Input("windowInDays")]
        public Input<int>? WindowInDays { get; set; }

        public SourceNetsuiteConfigurationGetArgs()
        {
        }
        public static new SourceNetsuiteConfigurationGetArgs Empty => new SourceNetsuiteConfigurationGetArgs();
    }
}