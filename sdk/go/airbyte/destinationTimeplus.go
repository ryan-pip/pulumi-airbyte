// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DestinationTimeplus struct {
	pulumi.CustomResourceState

	Configuration   DestinationTimeplusConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                    `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                    `pulumi:"destinationType"`
	Name            pulumi.StringOutput                    `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                    `pulumi:"workspaceId"`
}

// NewDestinationTimeplus registers a new resource with the given unique name, arguments, and options.
func NewDestinationTimeplus(ctx *pulumi.Context,
	name string, args *DestinationTimeplusArgs, opts ...pulumi.ResourceOption) (*DestinationTimeplus, error) {
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
	var resource DestinationTimeplus
	err := ctx.RegisterResource("airbyte:index/destinationTimeplus:DestinationTimeplus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationTimeplus gets an existing DestinationTimeplus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationTimeplus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationTimeplusState, opts ...pulumi.ResourceOption) (*DestinationTimeplus, error) {
	var resource DestinationTimeplus
	err := ctx.ReadResource("airbyte:index/destinationTimeplus:DestinationTimeplus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationTimeplus resources.
type destinationTimeplusState struct {
	Configuration   *DestinationTimeplusConfiguration `pulumi:"configuration"`
	DestinationId   *string                           `pulumi:"destinationId"`
	DestinationType *string                           `pulumi:"destinationType"`
	Name            *string                           `pulumi:"name"`
	WorkspaceId     *string                           `pulumi:"workspaceId"`
}

type DestinationTimeplusState struct {
	Configuration   DestinationTimeplusConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationTimeplusState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationTimeplusState)(nil)).Elem()
}

type destinationTimeplusArgs struct {
	Configuration DestinationTimeplusConfiguration `pulumi:"configuration"`
	Name          string                           `pulumi:"name"`
	WorkspaceId   string                           `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationTimeplus resource.
type DestinationTimeplusArgs struct {
	Configuration DestinationTimeplusConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationTimeplusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationTimeplusArgs)(nil)).Elem()
}

type DestinationTimeplusInput interface {
	pulumi.Input

	ToDestinationTimeplusOutput() DestinationTimeplusOutput
	ToDestinationTimeplusOutputWithContext(ctx context.Context) DestinationTimeplusOutput
}

func (*DestinationTimeplus) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationTimeplus)(nil)).Elem()
}

func (i *DestinationTimeplus) ToDestinationTimeplusOutput() DestinationTimeplusOutput {
	return i.ToDestinationTimeplusOutputWithContext(context.Background())
}

func (i *DestinationTimeplus) ToDestinationTimeplusOutputWithContext(ctx context.Context) DestinationTimeplusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationTimeplusOutput)
}

type DestinationTimeplusOutput struct{ *pulumi.OutputState }

func (DestinationTimeplusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationTimeplus)(nil)).Elem()
}

func (o DestinationTimeplusOutput) ToDestinationTimeplusOutput() DestinationTimeplusOutput {
	return o
}

func (o DestinationTimeplusOutput) ToDestinationTimeplusOutputWithContext(ctx context.Context) DestinationTimeplusOutput {
	return o
}

func (o DestinationTimeplusOutput) Configuration() DestinationTimeplusConfigurationOutput {
	return o.ApplyT(func(v *DestinationTimeplus) DestinationTimeplusConfigurationOutput { return v.Configuration }).(DestinationTimeplusConfigurationOutput)
}

func (o DestinationTimeplusOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationTimeplus) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationTimeplusOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationTimeplus) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationTimeplusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationTimeplus) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationTimeplusOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationTimeplus) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationTimeplusInput)(nil)).Elem(), &DestinationTimeplus{})
	pulumi.RegisterOutputType(DestinationTimeplusOutput{})
}