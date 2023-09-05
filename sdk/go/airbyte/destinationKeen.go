// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// DestinationKeen Resource
type DestinationKeen struct {
	pulumi.CustomResourceState

	Configuration   DestinationKeenConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                `pulumi:"destinationType"`
	Name            pulumi.StringOutput                `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                `pulumi:"workspaceId"`
}

// NewDestinationKeen registers a new resource with the given unique name, arguments, and options.
func NewDestinationKeen(ctx *pulumi.Context,
	name string, args *DestinationKeenArgs, opts ...pulumi.ResourceOption) (*DestinationKeen, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DestinationKeen
	err := ctx.RegisterResource("airbyte:index/destinationKeen:DestinationKeen", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationKeen gets an existing DestinationKeen resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationKeen(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationKeenState, opts ...pulumi.ResourceOption) (*DestinationKeen, error) {
	var resource DestinationKeen
	err := ctx.ReadResource("airbyte:index/destinationKeen:DestinationKeen", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationKeen resources.
type destinationKeenState struct {
	Configuration   *DestinationKeenConfiguration `pulumi:"configuration"`
	DestinationId   *string                       `pulumi:"destinationId"`
	DestinationType *string                       `pulumi:"destinationType"`
	Name            *string                       `pulumi:"name"`
	WorkspaceId     *string                       `pulumi:"workspaceId"`
}

type DestinationKeenState struct {
	Configuration   DestinationKeenConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationKeenState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationKeenState)(nil)).Elem()
}

type destinationKeenArgs struct {
	Configuration DestinationKeenConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	WorkspaceId   string                       `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationKeen resource.
type DestinationKeenArgs struct {
	Configuration DestinationKeenConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationKeenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationKeenArgs)(nil)).Elem()
}

type DestinationKeenInput interface {
	pulumi.Input

	ToDestinationKeenOutput() DestinationKeenOutput
	ToDestinationKeenOutputWithContext(ctx context.Context) DestinationKeenOutput
}

func (*DestinationKeen) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationKeen)(nil)).Elem()
}

func (i *DestinationKeen) ToDestinationKeenOutput() DestinationKeenOutput {
	return i.ToDestinationKeenOutputWithContext(context.Background())
}

func (i *DestinationKeen) ToDestinationKeenOutputWithContext(ctx context.Context) DestinationKeenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationKeenOutput)
}

type DestinationKeenOutput struct{ *pulumi.OutputState }

func (DestinationKeenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationKeen)(nil)).Elem()
}

func (o DestinationKeenOutput) ToDestinationKeenOutput() DestinationKeenOutput {
	return o
}

func (o DestinationKeenOutput) ToDestinationKeenOutputWithContext(ctx context.Context) DestinationKeenOutput {
	return o
}

func (o DestinationKeenOutput) Configuration() DestinationKeenConfigurationOutput {
	return o.ApplyT(func(v *DestinationKeen) DestinationKeenConfigurationOutput { return v.Configuration }).(DestinationKeenConfigurationOutput)
}

func (o DestinationKeenOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationKeen) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationKeenOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationKeen) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationKeenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationKeen) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationKeenOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationKeen) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationKeenInput)(nil)).Elem(), &DestinationKeen{})
	pulumi.RegisterOutputType(DestinationKeenOutput{})
}
