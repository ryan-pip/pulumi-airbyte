// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourceOmnisend(args: GetSourceOmnisendArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceOmnisendResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceOmnisend:getSourceOmnisend", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceOmnisend.
 */
export interface GetSourceOmnisendArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceOmnisend.
 */
export interface GetSourceOmnisendResult {
    readonly configuration: outputs.GetSourceOmnisendConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceOmnisendOutput(args: GetSourceOmnisendOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceOmnisendResult> {
    return pulumi.output(args).apply((a: any) => getSourceOmnisend(a, opts))
}

/**
 * A collection of arguments for invoking getSourceOmnisend.
 */
export interface GetSourceOmnisendOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}