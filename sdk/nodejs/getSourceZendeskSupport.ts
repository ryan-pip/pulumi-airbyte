// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * SourceZendeskSupport DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceZendesksupport = airbyte.getSourceZendeskSupport({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceZendeskSupport(args: GetSourceZendeskSupportArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceZendeskSupportResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceZendeskSupport:getSourceZendeskSupport", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceZendeskSupport.
 */
export interface GetSourceZendeskSupportArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceZendeskSupport.
 */
export interface GetSourceZendeskSupportResult {
    readonly configuration: outputs.GetSourceZendeskSupportConfiguration;
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
 * SourceZendeskSupport DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceZendesksupport = airbyte.getSourceZendeskSupport({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceZendeskSupportOutput(args: GetSourceZendeskSupportOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceZendeskSupportResult> {
    return pulumi.output(args).apply((a: any) => getSourceZendeskSupport(a, opts))
}

/**
 * A collection of arguments for invoking getSourceZendeskSupport.
 */
export interface GetSourceZendeskSupportOutputArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
