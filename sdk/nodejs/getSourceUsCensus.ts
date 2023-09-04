// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourceUsCensus(args: GetSourceUsCensusArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceUsCensusResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceUsCensus:getSourceUsCensus", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceUsCensus.
 */
export interface GetSourceUsCensusArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceUsCensus.
 */
export interface GetSourceUsCensusResult {
    readonly configuration: outputs.GetSourceUsCensusConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceUsCensusOutput(args: GetSourceUsCensusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceUsCensusResult> {
    return pulumi.output(args).apply((a: any) => getSourceUsCensus(a, opts))
}

/**
 * A collection of arguments for invoking getSourceUsCensus.
 */
export interface GetSourceUsCensusOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}