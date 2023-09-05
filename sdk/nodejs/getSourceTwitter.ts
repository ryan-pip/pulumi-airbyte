// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * SourceTwitter DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceTwitter = airbyte.getSourceTwitter({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceTwitter(args: GetSourceTwitterArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceTwitterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceTwitter:getSourceTwitter", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceTwitter.
 */
export interface GetSourceTwitterArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceTwitter.
 */
export interface GetSourceTwitterResult {
    readonly configuration: outputs.GetSourceTwitterConfiguration;
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
 * SourceTwitter DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceTwitter = airbyte.getSourceTwitter({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceTwitterOutput(args: GetSourceTwitterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceTwitterResult> {
    return pulumi.output(args).apply((a: any) => getSourceTwitter(a, opts))
}

/**
 * A collection of arguments for invoking getSourceTwitter.
 */
export interface GetSourceTwitterOutputArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
