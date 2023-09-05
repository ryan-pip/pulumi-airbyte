// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * SourceSnapchatMarketing DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceSnapchatmarketing = airbyte.getSourceSnapchatMarketing({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceSnapchatMarketing(args: GetSourceSnapchatMarketingArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceSnapchatMarketingResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceSnapchatMarketing:getSourceSnapchatMarketing", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceSnapchatMarketing.
 */
export interface GetSourceSnapchatMarketingArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceSnapchatMarketing.
 */
export interface GetSourceSnapchatMarketingResult {
    readonly configuration: outputs.GetSourceSnapchatMarketingConfiguration;
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
 * SourceSnapchatMarketing DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceSnapchatmarketing = airbyte.getSourceSnapchatMarketing({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceSnapchatMarketingOutput(args: GetSourceSnapchatMarketingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceSnapchatMarketingResult> {
    return pulumi.output(args).apply((a: any) => getSourceSnapchatMarketing(a, opts))
}

/**
 * A collection of arguments for invoking getSourceSnapchatMarketing.
 */
export interface GetSourceSnapchatMarketingOutputArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
