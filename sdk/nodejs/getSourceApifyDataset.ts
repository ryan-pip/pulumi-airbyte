// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * SourceApifyDataset DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceApifydataset = airbyte.getSourceApifyDataset({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceApifyDataset(args: GetSourceApifyDatasetArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceApifyDatasetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceApifyDataset:getSourceApifyDataset", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceApifyDataset.
 */
export interface GetSourceApifyDatasetArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceApifyDataset.
 */
export interface GetSourceApifyDatasetResult {
    readonly configuration: outputs.GetSourceApifyDatasetConfiguration;
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
 * SourceApifyDataset DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceApifydataset = airbyte.getSourceApifyDataset({
 *     secretId: "...my_secret_id...",
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceApifyDatasetOutput(args: GetSourceApifyDatasetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceApifyDatasetResult> {
    return pulumi.output(args).apply((a: any) => getSourceApifyDataset(a, opts))
}

/**
 * A collection of arguments for invoking getSourceApifyDataset.
 */
export interface GetSourceApifyDatasetOutputArgs {
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
