// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourceAzureTable(args: GetSourceAzureTableArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceAzureTableResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceAzureTable:getSourceAzureTable", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceAzureTable.
 */
export interface GetSourceAzureTableArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceAzureTable.
 */
export interface GetSourceAzureTableResult {
    readonly configuration: outputs.GetSourceAzureTableConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceAzureTableOutput(args: GetSourceAzureTableOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceAzureTableResult> {
    return pulumi.output(args).apply((a: any) => getSourceAzureTable(a, opts))
}

/**
 * A collection of arguments for invoking getSourceAzureTable.
 */
export interface GetSourceAzureTableOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}