// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourcePokeapi(args: GetSourcePokeapiArgs, opts?: pulumi.InvokeOptions): Promise<GetSourcePokeapiResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourcePokeapi:getSourcePokeapi", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourcePokeapi.
 */
export interface GetSourcePokeapiArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourcePokeapi.
 */
export interface GetSourcePokeapiResult {
    readonly configuration: outputs.GetSourcePokeapiConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourcePokeapiOutput(args: GetSourcePokeapiOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourcePokeapiResult> {
    return pulumi.output(args).apply((a: any) => getSourcePokeapi(a, opts))
}

/**
 * A collection of arguments for invoking getSourcePokeapi.
 */
export interface GetSourcePokeapiOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}