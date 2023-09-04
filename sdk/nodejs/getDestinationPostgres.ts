// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDestinationPostgres(args: GetDestinationPostgresArgs, opts?: pulumi.InvokeOptions): Promise<GetDestinationPostgresResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getDestinationPostgres:getDestinationPostgres", {
        "destinationId": args.destinationId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDestinationPostgres.
 */
export interface GetDestinationPostgresArgs {
    destinationId: string;
}

/**
 * A collection of values returned by getDestinationPostgres.
 */
export interface GetDestinationPostgresResult {
    readonly configuration: outputs.GetDestinationPostgresConfiguration;
    readonly destinationId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly workspaceId: string;
}
export function getDestinationPostgresOutput(args: GetDestinationPostgresOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDestinationPostgresResult> {
    return pulumi.output(args).apply((a: any) => getDestinationPostgres(a, opts))
}

/**
 * A collection of arguments for invoking getDestinationPostgres.
 */
export interface GetDestinationPostgresOutputArgs {
    destinationId: pulumi.Input<string>;
}