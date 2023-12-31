// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * SourceRecruitee DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceRecruitee = airbyte.getSourceRecruitee({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceRecruitee(args: GetSourceRecruiteeArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceRecruiteeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceRecruitee:getSourceRecruitee", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceRecruitee.
 */
export interface GetSourceRecruiteeArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceRecruitee.
 */
export interface GetSourceRecruiteeResult {
    readonly configuration: outputs.GetSourceRecruiteeConfiguration;
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
 * SourceRecruitee DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceRecruitee = airbyte.getSourceRecruitee({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceRecruiteeOutput(args: GetSourceRecruiteeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceRecruiteeResult> {
    return pulumi.output(args).apply((a: any) => getSourceRecruitee(a, opts))
}

/**
 * A collection of arguments for invoking getSourceRecruitee.
 */
export interface GetSourceRecruiteeOutputArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
