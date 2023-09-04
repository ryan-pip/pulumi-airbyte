// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourceAircall(args: GetSourceAircallArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceAircallResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceAircall:getSourceAircall", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceAircall.
 */
export interface GetSourceAircallArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceAircall.
 */
export interface GetSourceAircallResult {
    readonly configuration: outputs.GetSourceAircallConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceAircallOutput(args: GetSourceAircallOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceAircallResult> {
    return pulumi.output(args).apply((a: any) => getSourceAircall(a, opts))
}

/**
 * A collection of arguments for invoking getSourceAircall.
 */
export interface GetSourceAircallOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}