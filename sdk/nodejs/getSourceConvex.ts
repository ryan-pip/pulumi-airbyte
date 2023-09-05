// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * SourceConvex DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceConvex = airbyte.getSourceConvex({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceConvex(args: GetSourceConvexArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceConvexResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceConvex:getSourceConvex", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceConvex.
 */
export interface GetSourceConvexArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceConvex.
 */
export interface GetSourceConvexResult {
    readonly configuration: outputs.GetSourceConvexConfiguration;
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
 * SourceConvex DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceConvex = airbyte.getSourceConvex({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceConvexOutput(args: GetSourceConvexOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceConvexResult> {
    return pulumi.output(args).apply((a: any) => getSourceConvex(a, opts))
}

/**
 * A collection of arguments for invoking getSourceConvex.
 */
export interface GetSourceConvexOutputArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
