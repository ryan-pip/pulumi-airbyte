// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ConnectionConfigurations {
    streams: outputs.ConnectionConfigurationsStream[];
}

export interface ConnectionConfigurationsStream {
    cursorFields: string[];
    name: string;
    primaryKeys: string[][];
    syncMode: string;
}

export interface ConnectionSchedule {
    basicTiming: string;
    cronExpression: string;
    scheduleType: string;
}

export interface DestinationSnowflakeConfiguration {
    credentials?: outputs.DestinationSnowflakeConfigurationCredentials;
    database: string;
    destinationType: string;
    host: string;
    jdbcUrlParams?: string;
    rawDataSchema?: string;
    role: string;
    schema: string;
    use1s1tFormat?: boolean;
    username: string;
    warehouse: string;
}

export interface DestinationSnowflakeConfigurationCredentials {
    destinationSnowflakeAuthorizationMethodKeyPairAuthentication?: outputs.DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodKeyPairAuthentication;
    destinationSnowflakeAuthorizationMethodOAuth20?: outputs.DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodOAuth20;
    destinationSnowflakeAuthorizationMethodUsernameAndPassword?: outputs.DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodUsernameAndPassword;
    destinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication?: outputs.DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication;
    destinationSnowflakeUpdateAuthorizationMethodOAuth20?: outputs.DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodOAuth20;
    destinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword?: outputs.DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword;
}

export interface DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodKeyPairAuthentication {
    authType?: string;
    privateKey: string;
    privateKeyPassword?: string;
}

export interface DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodOAuth20 {
    accessToken: string;
    authType?: string;
    clientId?: string;
    clientSecret?: string;
    refreshToken: string;
}

export interface DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodUsernameAndPassword {
    authType?: string;
    password: string;
}

export interface DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication {
    authType?: string;
    privateKey: string;
    privateKeyPassword?: string;
}

export interface DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodOAuth20 {
    accessToken: string;
    authType?: string;
    clientId?: string;
    clientSecret?: string;
    refreshToken: string;
}

export interface DestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword {
    authType?: string;
    password: string;
}

export interface GetConnectionConfigurations {
    streams: outputs.GetConnectionConfigurationsStream[];
}

export interface GetConnectionConfigurationsStream {
    cursorFields: string[];
    name: string;
    primaryKeys: string[][];
    syncMode: string;
}

export interface GetConnectionSchedule {
    basicTiming: string;
    cronExpression: string;
    scheduleType: string;
}

export interface GetDestinationSnowflakeConfiguration {
    credentials: outputs.GetDestinationSnowflakeConfigurationCredentials;
    database: string;
    destinationType: string;
    host: string;
    jdbcUrlParams: string;
    rawDataSchema: string;
    role: string;
    schema: string;
    use1s1tFormat: boolean;
    username: string;
    warehouse: string;
}

export interface GetDestinationSnowflakeConfigurationCredentials {
    destinationSnowflakeAuthorizationMethodKeyPairAuthentication: outputs.GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodKeyPairAuthentication;
    destinationSnowflakeAuthorizationMethodOAuth20: outputs.GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodOAuth20;
    destinationSnowflakeAuthorizationMethodUsernameAndPassword: outputs.GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodUsernameAndPassword;
    destinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication: outputs.GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication;
    destinationSnowflakeUpdateAuthorizationMethodOAuth20: outputs.GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodOAuth20;
    destinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword: outputs.GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword;
}

export interface GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodKeyPairAuthentication {
    authType: string;
    privateKey: string;
    privateKeyPassword: string;
}

export interface GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodOAuth20 {
    accessToken: string;
    authType: string;
    clientId: string;
    clientSecret: string;
    refreshToken: string;
}

export interface GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeAuthorizationMethodUsernameAndPassword {
    authType: string;
    password: string;
}

export interface GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication {
    authType: string;
    privateKey: string;
    privateKeyPassword: string;
}

export interface GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodOAuth20 {
    accessToken: string;
    authType: string;
    clientId: string;
    clientSecret: string;
    refreshToken: string;
}

export interface GetDestinationSnowflakeConfigurationCredentialsDestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword {
    authType: string;
    password: string;
}

export interface GetSourceMixpanelConfiguration {
    attributionWindow: number;
    credentials: outputs.GetSourceMixpanelConfigurationCredentials;
    dateWindowSize: number;
    endDate: string;
    projectId: number;
    projectTimezone: string;
    region: string;
    selectPropertiesByDefault: boolean;
    sourceType: string;
    startDate: string;
}

export interface GetSourceMixpanelConfigurationCredentials {
    sourceMixpanelAuthenticationWildcardProjectSecret: outputs.GetSourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardProjectSecret;
    sourceMixpanelAuthenticationWildcardServiceAccount: outputs.GetSourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardServiceAccount;
    sourceMixpanelUpdateAuthenticationWildcardProjectSecret: outputs.GetSourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecret;
    sourceMixpanelUpdateAuthenticationWildcardServiceAccount: outputs.GetSourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccount;
}

export interface GetSourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardProjectSecret {
    apiSecret: string;
    optionTitle: string;
}

export interface GetSourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardServiceAccount {
    optionTitle: string;
    secret: string;
    username: string;
}

export interface GetSourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecret {
    apiSecret: string;
    optionTitle: string;
}

export interface GetSourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccount {
    optionTitle: string;
    secret: string;
    username: string;
}

export interface GetSourceSalesforceConfiguration {
    authType: string;
    clientId: string;
    clientSecret: string;
    forceUseBulkApi: boolean;
    isSandbox: boolean;
    refreshToken: string;
    sourceType: string;
    startDate: string;
    streamsCriterias: outputs.GetSourceSalesforceConfigurationStreamsCriteria[];
}

export interface GetSourceSalesforceConfigurationStreamsCriteria {
    criteria: string;
    value: string;
}

export interface SourceMixpanelConfiguration {
    attributionWindow?: number;
    credentials?: outputs.SourceMixpanelConfigurationCredentials;
    dateWindowSize?: number;
    endDate?: string;
    projectId?: number;
    projectTimezone?: string;
    region?: string;
    selectPropertiesByDefault?: boolean;
    sourceType?: string;
    startDate?: string;
}

export interface SourceMixpanelConfigurationCredentials {
    sourceMixpanelAuthenticationWildcardProjectSecret?: outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardProjectSecret;
    sourceMixpanelAuthenticationWildcardServiceAccount?: outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardServiceAccount;
    sourceMixpanelUpdateAuthenticationWildcardProjectSecret?: outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecret;
    sourceMixpanelUpdateAuthenticationWildcardServiceAccount?: outputs.SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccount;
}

export interface SourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardProjectSecret {
    apiSecret: string;
    optionTitle?: string;
}

export interface SourceMixpanelConfigurationCredentialsSourceMixpanelAuthenticationWildcardServiceAccount {
    optionTitle?: string;
    secret: string;
    username: string;
}

export interface SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardProjectSecret {
    apiSecret: string;
    optionTitle?: string;
}

export interface SourceMixpanelConfigurationCredentialsSourceMixpanelUpdateAuthenticationWildcardServiceAccount {
    optionTitle?: string;
    secret: string;
    username: string;
}

export interface SourceSalesforceConfiguration {
    authType?: string;
    clientId: string;
    clientSecret: string;
    forceUseBulkApi?: boolean;
    isSandbox?: boolean;
    refreshToken: string;
    sourceType: string;
    startDate?: string;
    streamsCriterias?: outputs.SourceSalesforceConfigurationStreamsCriteria[];
}

export interface SourceSalesforceConfigurationStreamsCriteria {
    criteria: string;
    value: string;
}
