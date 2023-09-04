// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

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
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceApifyDatasetOutput(args: GetSourceApifyDatasetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceApifyDatasetResult> {
    return pulumi.output(args).apply((a: any) => getSourceApifyDataset(a, opts))
}

/**
 * A collection of arguments for invoking getSourceApifyDataset.
 */
export interface GetSourceApifyDatasetOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}