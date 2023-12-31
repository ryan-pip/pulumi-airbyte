// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * SourcePostmarkapp DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourcePostmarkapp = airbyte.getSourcePostmarkapp({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourcePostmarkapp(args: GetSourcePostmarkappArgs, opts?: pulumi.InvokeOptions): Promise<GetSourcePostmarkappResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourcePostmarkapp:getSourcePostmarkapp", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourcePostmarkapp.
 */
export interface GetSourcePostmarkappArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourcePostmarkapp.
 */
export interface GetSourcePostmarkappResult {
    readonly configuration: outputs.GetSourcePostmarkappConfiguration;
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
 * SourcePostmarkapp DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourcePostmarkapp = airbyte.getSourcePostmarkapp({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourcePostmarkappOutput(args: GetSourcePostmarkappOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourcePostmarkappResult> {
    return pulumi.output(args).apply((a: any) => getSourcePostmarkapp(a, opts))
}

/**
 * A collection of arguments for invoking getSourcePostmarkapp.
 */
export interface GetSourcePostmarkappOutputArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
