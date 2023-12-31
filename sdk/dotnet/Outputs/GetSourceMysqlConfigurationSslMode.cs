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
    public sealed class GetSourceMysqlConfigurationSslMode
    {
        public readonly Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlSslModesPreferred SourceMysqlSslModesPreferred;
        public readonly Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlSslModesRequired SourceMysqlSslModesRequired;
        public readonly Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlSslModesVerifyCa SourceMysqlSslModesVerifyCa;
        public readonly Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlSslModesVerifyIdentity SourceMysqlSslModesVerifyIdentity;
        public readonly Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlUpdateSslModesPreferred SourceMysqlUpdateSslModesPreferred;
        public readonly Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlUpdateSslModesRequired SourceMysqlUpdateSslModesRequired;
        public readonly Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlUpdateSslModesVerifyCa SourceMysqlUpdateSslModesVerifyCa;
        public readonly Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlUpdateSslModesVerifyIdentity SourceMysqlUpdateSslModesVerifyIdentity;

        [OutputConstructor]
        private GetSourceMysqlConfigurationSslMode(
            Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlSslModesPreferred sourceMysqlSslModesPreferred,

            Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlSslModesRequired sourceMysqlSslModesRequired,

            Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlSslModesVerifyCa sourceMysqlSslModesVerifyCa,

            Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlSslModesVerifyIdentity sourceMysqlSslModesVerifyIdentity,

            Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlUpdateSslModesPreferred sourceMysqlUpdateSslModesPreferred,

            Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlUpdateSslModesRequired sourceMysqlUpdateSslModesRequired,

            Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlUpdateSslModesVerifyCa sourceMysqlUpdateSslModesVerifyCa,

            Outputs.GetSourceMysqlConfigurationSslModeSourceMysqlUpdateSslModesVerifyIdentity sourceMysqlUpdateSslModesVerifyIdentity)
        {
            SourceMysqlSslModesPreferred = sourceMysqlSslModesPreferred;
            SourceMysqlSslModesRequired = sourceMysqlSslModesRequired;
            SourceMysqlSslModesVerifyCa = sourceMysqlSslModesVerifyCa;
            SourceMysqlSslModesVerifyIdentity = sourceMysqlSslModesVerifyIdentity;
            SourceMysqlUpdateSslModesPreferred = sourceMysqlUpdateSslModesPreferred;
            SourceMysqlUpdateSslModesRequired = sourceMysqlUpdateSslModesRequired;
            SourceMysqlUpdateSslModesVerifyCa = sourceMysqlUpdateSslModesVerifyCa;
            SourceMysqlUpdateSslModesVerifyIdentity = sourceMysqlUpdateSslModesVerifyIdentity;
        }
    }
}
