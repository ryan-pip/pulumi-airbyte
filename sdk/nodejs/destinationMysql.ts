// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * DestinationMysql Resource
 */
export class DestinationMysql extends pulumi.CustomResource {
    /**
     * Get an existing DestinationMysql resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DestinationMysqlState, opts?: pulumi.CustomResourceOptions): DestinationMysql {
        return new DestinationMysql(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'airbyte:index/destinationMysql:DestinationMysql';

    /**
     * Returns true if the given object is an instance of DestinationMysql.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DestinationMysql {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DestinationMysql.__pulumiType;
    }

    public readonly configuration!: pulumi.Output<outputs.DestinationMysqlConfiguration>;
    public /*out*/ readonly destinationId!: pulumi.Output<string>;
    public /*out*/ readonly destinationType!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly workspaceId!: pulumi.Output<string>;

    /**
     * Create a DestinationMysql resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DestinationMysqlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DestinationMysqlArgs | DestinationMysqlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DestinationMysqlState | undefined;
            resourceInputs["configuration"] = state ? state.configuration : undefined;
            resourceInputs["destinationId"] = state ? state.destinationId : undefined;
            resourceInputs["destinationType"] = state ? state.destinationType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["workspaceId"] = state ? state.workspaceId : undefined;
        } else {
            const args = argsOrState as DestinationMysqlArgs | undefined;
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
        super(DestinationMysql.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DestinationMysql resources.
 */
export interface DestinationMysqlState {
    configuration?: pulumi.Input<inputs.DestinationMysqlConfiguration>;
    destinationId?: pulumi.Input<string>;
    destinationType?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    workspaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DestinationMysql resource.
 */
export interface DestinationMysqlArgs {
    configuration: pulumi.Input<inputs.DestinationMysqlConfiguration>;
    name: pulumi.Input<string>;
    workspaceId: pulumi.Input<string>;
}
