// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourcePosthog(args: GetSourcePosthogArgs, opts?: pulumi.InvokeOptions): Promise<GetSourcePosthogResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourcePosthog:getSourcePosthog", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourcePosthog.
 */
export interface GetSourcePosthogArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourcePosthog.
 */
export interface GetSourcePosthogResult {
    readonly configuration: outputs.GetSourcePosthogConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourcePosthogOutput(args: GetSourcePosthogOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourcePosthogResult> {
    return pulumi.output(args).apply((a: any) => getSourcePosthog(a, opts))
}

/**
 * A collection of arguments for invoking getSourcePosthog.
 */
export interface GetSourcePosthogOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}