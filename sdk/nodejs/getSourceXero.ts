// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourceXero(args: GetSourceXeroArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceXeroResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceXero:getSourceXero", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceXero.
 */
export interface GetSourceXeroArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceXero.
 */
export interface GetSourceXeroResult {
    readonly configuration: outputs.GetSourceXeroConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceXeroOutput(args: GetSourceXeroOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceXeroResult> {
    return pulumi.output(args).apply((a: any) => getSourceXero(a, opts))
}

/**
 * A collection of arguments for invoking getSourceXero.
 */
export interface GetSourceXeroOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}