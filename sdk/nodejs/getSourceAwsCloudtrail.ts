// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourceAwsCloudtrail(args: GetSourceAwsCloudtrailArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceAwsCloudtrailResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceAwsCloudtrail:getSourceAwsCloudtrail", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceAwsCloudtrail.
 */
export interface GetSourceAwsCloudtrailArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceAwsCloudtrail.
 */
export interface GetSourceAwsCloudtrailResult {
    readonly configuration: outputs.GetSourceAwsCloudtrailConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceAwsCloudtrailOutput(args: GetSourceAwsCloudtrailOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceAwsCloudtrailResult> {
    return pulumi.output(args).apply((a: any) => getSourceAwsCloudtrail(a, opts))
}

/**
 * A collection of arguments for invoking getSourceAwsCloudtrail.
 */
export interface GetSourceAwsCloudtrailOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}