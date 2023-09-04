// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDestinationGoogleSheets(args: GetDestinationGoogleSheetsArgs, opts?: pulumi.InvokeOptions): Promise<GetDestinationGoogleSheetsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getDestinationGoogleSheets:getDestinationGoogleSheets", {
        "destinationId": args.destinationId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDestinationGoogleSheets.
 */
export interface GetDestinationGoogleSheetsArgs {
    destinationId: string;
}

/**
 * A collection of values returned by getDestinationGoogleSheets.
 */
export interface GetDestinationGoogleSheetsResult {
    readonly configuration: outputs.GetDestinationGoogleSheetsConfiguration;
    readonly destinationId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly workspaceId: string;
}
export function getDestinationGoogleSheetsOutput(args: GetDestinationGoogleSheetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDestinationGoogleSheetsResult> {
    return pulumi.output(args).apply((a: any) => getDestinationGoogleSheets(a, opts))
}

/**
 * A collection of arguments for invoking getDestinationGoogleSheets.
 */
export interface GetDestinationGoogleSheetsOutputArgs {
    destinationId: pulumi.Input<string>;
}