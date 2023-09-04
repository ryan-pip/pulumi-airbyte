// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDestinationOracle(args: GetDestinationOracleArgs, opts?: pulumi.InvokeOptions): Promise<GetDestinationOracleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getDestinationOracle:getDestinationOracle", {
        "destinationId": args.destinationId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDestinationOracle.
 */
export interface GetDestinationOracleArgs {
    destinationId: string;
}

/**
 * A collection of values returned by getDestinationOracle.
 */
export interface GetDestinationOracleResult {
    readonly configuration: outputs.GetDestinationOracleConfiguration;
    readonly destinationId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly workspaceId: string;
}
export function getDestinationOracleOutput(args: GetDestinationOracleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDestinationOracleResult> {
    return pulumi.output(args).apply((a: any) => getDestinationOracle(a, opts))
}

/**
 * A collection of arguments for invoking getDestinationOracle.
 */
export interface GetDestinationOracleOutputArgs {
    destinationId: pulumi.Input<string>;
}