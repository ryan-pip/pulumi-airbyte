// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * SourceGoogleAnalyticsV4 Resource
 */
export class SourceGoogleAnalyticsV4 extends pulumi.CustomResource {
    /**
     * Get an existing SourceGoogleAnalyticsV4 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SourceGoogleAnalyticsV4State, opts?: pulumi.CustomResourceOptions): SourceGoogleAnalyticsV4 {
        return new SourceGoogleAnalyticsV4(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'airbyte:index/sourceGoogleAnalyticsV4:SourceGoogleAnalyticsV4';

    /**
     * Returns true if the given object is an instance of SourceGoogleAnalyticsV4.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SourceGoogleAnalyticsV4 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SourceGoogleAnalyticsV4.__pulumiType;
    }

    public readonly configuration!: pulumi.Output<outputs.SourceGoogleAnalyticsV4Configuration>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    public readonly secretId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly sourceId!: pulumi.Output<string>;
    public /*out*/ readonly sourceType!: pulumi.Output<string>;
    public readonly workspaceId!: pulumi.Output<string>;

    /**
     * Create a SourceGoogleAnalyticsV4 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SourceGoogleAnalyticsV4Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SourceGoogleAnalyticsV4Args | SourceGoogleAnalyticsV4State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SourceGoogleAnalyticsV4State | undefined;
            resourceInputs["configuration"] = state ? state.configuration : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["secretId"] = state ? state.secretId : undefined;
            resourceInputs["sourceId"] = state ? state.sourceId : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
            resourceInputs["workspaceId"] = state ? state.workspaceId : undefined;
        } else {
            const args = argsOrState as SourceGoogleAnalyticsV4Args | undefined;
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
        super(SourceGoogleAnalyticsV4.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SourceGoogleAnalyticsV4 resources.
 */
export interface SourceGoogleAnalyticsV4State {
    configuration?: pulumi.Input<inputs.SourceGoogleAnalyticsV4Configuration>;
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
 * The set of arguments for constructing a SourceGoogleAnalyticsV4 resource.
 */
export interface SourceGoogleAnalyticsV4Args {
    configuration: pulumi.Input<inputs.SourceGoogleAnalyticsV4Configuration>;
    name: pulumi.Input<string>;
    /**
     * Optional secretID obtained through the public API OAuth redirect flow.
     */
    secretId?: pulumi.Input<string>;
    workspaceId: pulumi.Input<string>;
}
