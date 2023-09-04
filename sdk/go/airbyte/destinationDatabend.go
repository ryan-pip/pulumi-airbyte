// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DestinationDatabend struct {
	pulumi.CustomResourceState

	Configuration   DestinationDatabendConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                    `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                    `pulumi:"destinationType"`
	Name            pulumi.StringOutput                    `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                    `pulumi:"workspaceId"`
}

// NewDestinationDatabend registers a new resource with the given unique name, arguments, and options.
func NewDestinationDatabend(ctx *pulumi.Context,
	name string, args *DestinationDatabendArgs, opts ...pulumi.ResourceOption) (*DestinationDatabend, error) {
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
	var resource DestinationDatabend
	err := ctx.RegisterResource("airbyte:index/destinationDatabend:DestinationDatabend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationDatabend gets an existing DestinationDatabend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationDatabend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationDatabendState, opts ...pulumi.ResourceOption) (*DestinationDatabend, error) {
	var resource DestinationDatabend
	err := ctx.ReadResource("airbyte:index/destinationDatabend:DestinationDatabend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationDatabend resources.
type destinationDatabendState struct {
	Configuration   *DestinationDatabendConfiguration `pulumi:"configuration"`
	DestinationId   *string                           `pulumi:"destinationId"`
	DestinationType *string                           `pulumi:"destinationType"`
	Name            *string                           `pulumi:"name"`
	WorkspaceId     *string                           `pulumi:"workspaceId"`
}

type DestinationDatabendState struct {
	Configuration   DestinationDatabendConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationDatabendState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationDatabendState)(nil)).Elem()
}

type destinationDatabendArgs struct {
	Configuration DestinationDatabendConfiguration `pulumi:"configuration"`
	Name          string                           `pulumi:"name"`
	WorkspaceId   string                           `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationDatabend resource.
type DestinationDatabendArgs struct {
	Configuration DestinationDatabendConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationDatabendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationDatabendArgs)(nil)).Elem()
}

type DestinationDatabendInput interface {
	pulumi.Input

	ToDestinationDatabendOutput() DestinationDatabendOutput
	ToDestinationDatabendOutputWithContext(ctx context.Context) DestinationDatabendOutput
}

func (*DestinationDatabend) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationDatabend)(nil)).Elem()
}

func (i *DestinationDatabend) ToDestinationDatabendOutput() DestinationDatabendOutput {
	return i.ToDestinationDatabendOutputWithContext(context.Background())
}

func (i *DestinationDatabend) ToDestinationDatabendOutputWithContext(ctx context.Context) DestinationDatabendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationDatabendOutput)
}

type DestinationDatabendOutput struct{ *pulumi.OutputState }

func (DestinationDatabendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationDatabend)(nil)).Elem()
}

func (o DestinationDatabendOutput) ToDestinationDatabendOutput() DestinationDatabendOutput {
	return o
}

func (o DestinationDatabendOutput) ToDestinationDatabendOutputWithContext(ctx context.Context) DestinationDatabendOutput {
	return o
}

func (o DestinationDatabendOutput) Configuration() DestinationDatabendConfigurationOutput {
	return o.ApplyT(func(v *DestinationDatabend) DestinationDatabendConfigurationOutput { return v.Configuration }).(DestinationDatabendConfigurationOutput)
}

func (o DestinationDatabendOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationDatabend) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationDatabendOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationDatabend) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationDatabendOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationDatabend) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationDatabendOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationDatabend) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationDatabendInput)(nil)).Elem(), &DestinationDatabend{})
	pulumi.RegisterOutputType(DestinationDatabendOutput{})
}