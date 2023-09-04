// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SourceXero extends pulumi.CustomResource {
    /**
     * Get an existing SourceXero resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SourceXeroState, opts?: pulumi.CustomResourceOptions): SourceXero {
        return new SourceXero(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'airbyte:index/sourceXero:SourceXero';

    /**
     * Returns true if the given object is an instance of SourceXero.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SourceXero {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SourceXero.__pulumiType;
    }

    public readonly configuration!: pulumi.Output<outputs.SourceXeroConfiguration>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    public readonly secretId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly sourceId!: pulumi.Output<string>;
    public /*out*/ readonly sourceType!: pulumi.Output<string>;
    public readonly workspaceId!: pulumi.Output<string>;

    /**
     * Create a SourceXero resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SourceXeroArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SourceXeroArgs | SourceXeroState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SourceXeroState | undefined;
            resourceInputs["configuration"] = state ? state.configuration : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["secretId"] = state ? state.secretId : undefined;
            resourceInputs["sourceId"] = state ? state.sourceId : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
            resourceInputs["workspaceId"] = state ? state.workspaceId : undefined;
        } else {
            const args = argsOrState as SourceXeroArgs | undefined;
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
            resourceInputs["secretId"] = args ? args.secretId : undefined;
            resourceInputs["workspaceId"] = args ? args.workspaceId : undefined;
            resourceInputs["sourceId"] = undefined /*out*/;
            resourceInputs["sourceType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SourceXero.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SourceXero resources.
 */
export interface SourceXeroState {
    configuration?: pulumi.Input<inputs.SourceXeroConfiguration>;
    name?: pulumi.Input<string>;
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    sourceId?: pulumi.Input<string>;
    sourceType?: pulumi.Input<string>;
    workspaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SourceXero resource.
 */
export interface SourceXeroArgs {
    configuration: pulumi.Input<inputs.SourceXeroConfiguration>;
    name: pulumi.Input<string>;
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    workspaceId: pulumi.Input<string>;
}