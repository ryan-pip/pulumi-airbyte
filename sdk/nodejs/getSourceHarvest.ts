// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * SourceHarvest DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceHarvest = airbyte.getSourceHarvest({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceHarvest(args: GetSourceHarvestArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceHarvestResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceHarvest:getSourceHarvest", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceHarvest.
 */
export interface GetSourceHarvestArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceHarvest.
 */
export interface GetSourceHarvestResult {
    readonly configuration: outputs.GetSourceHarvestConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
/**
 * SourceHarvest DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceHarvest = airbyte.getSourceHarvest({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceHarvestOutput(args: GetSourceHarvestOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceHarvestResult> {
    return pulumi.output(args).apply((a: any) => getSourceHarvest(a, opts))
}

/**
 * A collection of arguments for invoking getSourceHarvest.
 */
export interface GetSourceHarvestOutputArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
