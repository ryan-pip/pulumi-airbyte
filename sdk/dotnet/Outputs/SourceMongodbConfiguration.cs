// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class SourceMongodbConfiguration
    {
        public readonly string? AuthSource;
        public readonly string Database;
        public readonly Outputs.SourceMongodbConfigurationInstanceType? InstanceType;
        public readonly string? Password;
        public readonly string SourceType;
        public readonly string? User;

        [OutputConstructor]
        private SourceMongodbConfiguration(
            string? authSource,

            string database,

            Outputs.SourceMongodbConfigurationInstanceType? instanceType,

            string? password,

            string sourceType,

            string? user)
        {
            AuthSource = authSource;
            Database = database;
            InstanceType = instanceType;
            Password = password;
            SourceType = sourceType;
            User = user;
        }
    }
}
