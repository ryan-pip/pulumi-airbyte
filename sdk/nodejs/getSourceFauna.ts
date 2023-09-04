// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourceFauna(args: GetSourceFaunaArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceFaunaResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceFauna:getSourceFauna", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceFauna.
 */
export interface GetSourceFaunaArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceFauna.
 */
export interface GetSourceFaunaResult {
    readonly configuration: outputs.GetSourceFaunaConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceFaunaOutput(args: GetSourceFaunaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceFaunaResult> {
    return pulumi.output(args).apply((a: any) => getSourceFauna(a, opts))
}

/**
 * A collection of arguments for invoking getSourceFauna.
 */
export interface GetSourceFaunaOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}