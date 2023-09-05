// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * DestinationGoogleSheets Resource
 */
export class DestinationGoogleSheets extends pulumi.CustomResource {
    /**
     * Get an existing DestinationGoogleSheets resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DestinationGoogleSheetsState, opts?: pulumi.CustomResourceOptions): DestinationGoogleSheets {
        return new DestinationGoogleSheets(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'airbyte:index/destinationGoogleSheets:DestinationGoogleSheets';

    /**
     * Returns true if the given object is an instance of DestinationGoogleSheets.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DestinationGoogleSheets {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DestinationGoogleSheets.__pulumiType;
    }

    public readonly configuration!: pulumi.Output<outputs.DestinationGoogleSheetsConfiguration>;
    public /*out*/ readonly destinationId!: pulumi.Output<string>;
    public /*out*/ readonly destinationType!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly workspaceId!: pulumi.Output<string>;

    /**
     * Create a DestinationGoogleSheets resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DestinationGoogleSheetsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DestinationGoogleSheetsArgs | DestinationGoogleSheetsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DestinationGoogleSheetsState | undefined;
            resourceInputs["configuration"] = state ? state.configuration : undefined;
            resourceInputs["destinationId"] = state ? state.destinationId : undefined;
            resourceInputs["destinationType"] = state ? state.destinationType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["workspaceId"] = state ? state.workspaceId : undefined;
        } else {
            const args = argsOrState as DestinationGoogleSheetsArgs | undefined;
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.workspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceId'");
            }
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["workspaceId"] = args ? args.workspaceId : undefined;
            resourceInputs["destinationId"] = undefined /*out*/;
            resourceInputs["destinationType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DestinationGoogleSheets.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DestinationGoogleSheets resources.
 */
export interface DestinationGoogleSheetsState {
    configuration?: pulumi.Input<inputs.DestinationGoogleSheetsConfiguration>;
    destinationId?: pulumi.Input<string>;
    destinationType?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    workspaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DestinationGoogleSheets resource.
 */
export interface DestinationGoogleSheetsArgs {
    configuration: pulumi.Input<inputs.DestinationGoogleSheetsConfiguration>;
    name: pulumi.Input<string>;
    workspaceId: pulumi.Input<string>;
}
