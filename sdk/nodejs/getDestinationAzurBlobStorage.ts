// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getDestinationAzurBlobStorage(args: GetDestinationAzurBlobStorageArgs, opts?: pulumi.InvokeOptions): Promise<GetDestinationAzurBlobStorageResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getDestinationAzurBlobStorage:getDestinationAzurBlobStorage", {
        "destinationId": args.destinationId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDestinationAzurBlobStorage.
 */
export interface GetDestinationAzurBlobStorageArgs {
    destinationId: string;
}

/**
 * A collection of values returned by getDestinationAzurBlobStorage.
 */
export interface GetDestinationAzurBlobStorageResult {
    readonly configuration: outputs.GetDestinationAzurBlobStorageConfiguration;
    readonly destinationId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly workspaceId: string;
}
export function getDestinationAzurBlobStorageOutput(args: GetDestinationAzurBlobStorageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDestinationAzurBlobStorageResult> {
    return pulumi.output(args).apply((a: any) => getDestinationAzurBlobStorage(a, opts))
}

/**
 * A collection of arguments for invoking getDestinationAzurBlobStorage.
 */
export interface GetDestinationAzurBlobStorageOutputArgs {
    destinationId: pulumi.Input<string>;
}