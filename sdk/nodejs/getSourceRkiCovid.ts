// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourceRkiCovid(args: GetSourceRkiCovidArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceRkiCovidResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceRkiCovid:getSourceRkiCovid", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceRkiCovid.
 */
export interface GetSourceRkiCovidArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceRkiCovid.
 */
export interface GetSourceRkiCovidResult {
    readonly configuration: outputs.GetSourceRkiCovidConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceRkiCovidOutput(args: GetSourceRkiCovidOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceRkiCovidResult> {
    return pulumi.output(args).apply((a: any) => getSourceRkiCovid(a, opts))
}

/**
 * A collection of arguments for invoking getSourceRkiCovid.
 */
export interface GetSourceRkiCovidOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}