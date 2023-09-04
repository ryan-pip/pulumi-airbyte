// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDestinationXata(args: GetDestinationXataArgs, opts?: pulumi.InvokeOptions): Promise<GetDestinationXataResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getDestinationXata:getDestinationXata", {
        "destinationId": args.destinationId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDestinationXata.
 */
export interface GetDestinationXataArgs {
    destinationId: string;
}

/**
 * A collection of values returned by getDestinationXata.
 */
export interface GetDestinationXataResult {
    readonly configuration: outputs.GetDestinationXataConfiguration;
    readonly destinationId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly workspaceId: string;
}
export function getDestinationXataOutput(args: GetDestinationXataOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDestinationXataResult> {
    return pulumi.output(args).apply((a: any) => getDestinationXata(a, opts))
}

/**
 * A collection of arguments for invoking getDestinationXata.
 */
export interface GetDestinationXataOutputArgs {
    destinationId: pulumi.Input<string>;
}