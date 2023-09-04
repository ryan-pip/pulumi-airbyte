// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourceMonday(args: GetSourceMondayArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceMondayResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceMonday:getSourceMonday", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceMonday.
 */
export interface GetSourceMondayArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceMonday.
 */
export interface GetSourceMondayResult {
    readonly configuration: outputs.GetSourceMondayConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceMondayOutput(args: GetSourceMondayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceMondayResult> {
    return pulumi.output(args).apply((a: any) => getSourceMonday(a, opts))
}

/**
 * A collection of arguments for invoking getSourceMonday.
 */
export interface GetSourceMondayOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}