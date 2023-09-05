// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * SourceWhiskyHunter DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceWhiskyhunter = airbyte.getSourceWhiskyHunter({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceWhiskyHunter(args: GetSourceWhiskyHunterArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceWhiskyHunterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceWhiskyHunter:getSourceWhiskyHunter", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceWhiskyHunter.
 */
export interface GetSourceWhiskyHunterArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceWhiskyHunter.
 */
export interface GetSourceWhiskyHunterResult {
    readonly configuration: outputs.GetSourceWhiskyHunterConfiguration;
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
 * SourceWhiskyHunter DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceWhiskyhunter = airbyte.getSourceWhiskyHunter({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceWhiskyHunterOutput(args: GetSourceWhiskyHunterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceWhiskyHunterResult> {
    return pulumi.output(args).apply((a: any) => getSourceWhiskyHunter(a, opts))
}

/**
 * A collection of arguments for invoking getSourceWhiskyHunter.
 */
export interface GetSourceWhiskyHunterOutputArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
