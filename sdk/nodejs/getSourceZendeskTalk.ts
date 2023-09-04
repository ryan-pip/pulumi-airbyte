// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSourceZendeskTalk(args: GetSourceZendeskTalkArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceZendeskTalkResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceZendeskTalk:getSourceZendeskTalk", {
        "secretId": args.secretId,
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceZendeskTalk.
 */
export interface GetSourceZendeskTalkArgs {
    secretId?: string;
    sourceId: string;
}

/**
 * A collection of values returned by getSourceZendeskTalk.
 */
export interface GetSourceZendeskTalkResult {
    readonly configuration: outputs.GetSourceZendeskTalkConfiguration;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly secretId?: string;
    readonly sourceId: string;
    readonly workspaceId: string;
}
export function getSourceZendeskTalkOutput(args: GetSourceZendeskTalkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceZendeskTalkResult> {
    return pulumi.output(args).apply((a: any) => getSourceZendeskTalk(a, opts))
}

/**
 * A collection of arguments for invoking getSourceZendeskTalk.
 */
export interface GetSourceZendeskTalkOutputArgs {
    secretId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}