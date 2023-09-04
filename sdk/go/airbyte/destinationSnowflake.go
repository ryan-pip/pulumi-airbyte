// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DestinationSnowflake struct {
	pulumi.CustomResourceState

	Configuration   DestinationSnowflakeConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                     `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                     `pulumi:"destinationType"`
	Name            pulumi.StringOutput                     `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                     `pulumi:"workspaceId"`
}

// NewDestinationSnowflake registers a new resource with the given unique name, arguments, and options.
func NewDestinationSnowflake(ctx *pulumi.Context,
	name string, args *DestinationSnowflakeArgs, opts ...pulumi.ResourceOption) (*DestinationSnowflake, error) {
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
	opts = pkgResourceDefaultOpts(opts)
	var resource DestinationSnowflake
	err := ctx.RegisterResource("airbyte:index/destinationSnowflake:DestinationSnowflake", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationSnowflake gets an existing DestinationSnowflake resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationSnowflake(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationSnowflakeState, opts ...pulumi.ResourceOption) (*DestinationSnowflake, error) {
	var resource DestinationSnowflake
	err := ctx.ReadResource("airbyte:index/destinationSnowflake:DestinationSnowflake", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationSnowflake resources.
type destinationSnowflakeState struct {
	Configuration   *DestinationSnowflakeConfiguration `pulumi:"configuration"`
	DestinationId   *string                            `pulumi:"destinationId"`
	DestinationType *string                            `pulumi:"destinationType"`
	Name            *string                            `pulumi:"name"`
	WorkspaceId     *string                            `pulumi:"workspaceId"`
}

type DestinationSnowflakeState struct {
	Configuration   DestinationSnowflakeConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationSnowflakeState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationSnowflakeState)(nil)).Elem()
}

type destinationSnowflakeArgs struct {
	Configuration DestinationSnowflakeConfiguration `pulumi:"configuration"`
	Name          string                            `pulumi:"name"`
	WorkspaceId   string                            `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationSnowflake resource.
type DestinationSnowflakeArgs struct {
	Configuration DestinationSnowflakeConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationSnowflakeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationSnowflakeArgs)(nil)).Elem()
}

type DestinationSnowflakeInput interface {
	pulumi.Input

	ToDestinationSnowflakeOutput() DestinationSnowflakeOutput
	ToDestinationSnowflakeOutputWithContext(ctx context.Context) DestinationSnowflakeOutput
}

func (*DestinationSnowflake) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationSnowflake)(nil)).Elem()
}

func (i *DestinationSnowflake) ToDestinationSnowflakeOutput() DestinationSnowflakeOutput {
	return i.ToDestinationSnowflakeOutputWithContext(context.Background())
}

func (i *DestinationSnowflake) ToDestinationSnowflakeOutputWithContext(ctx context.Context) DestinationSnowflakeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationSnowflakeOutput)
}

type DestinationSnowflakeOutput struct{ *pulumi.OutputState }

func (DestinationSnowflakeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationSnowflake)(nil)).Elem()
}

func (o DestinationSnowflakeOutput) ToDestinationSnowflakeOutput() DestinationSnowflakeOutput {
	return o
}

func (o DestinationSnowflakeOutput) ToDestinationSnowflakeOutputWithContext(ctx context.Context) DestinationSnowflakeOutput {
	return o
}

func (o DestinationSnowflakeOutput) Configuration() DestinationSnowflakeConfigurationOutput {
	return o.ApplyT(func(v *DestinationSnowflake) DestinationSnowflakeConfigurationOutput { return v.Configuration }).(DestinationSnowflakeConfigurationOutput)
}

func (o DestinationSnowflakeOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationSnowflake) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationSnowflakeOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationSnowflake) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationSnowflakeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationSnowflake) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationSnowflakeOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationSnowflake) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationSnowflakeInput)(nil)).Elem(), &DestinationSnowflake{})
	pulumi.RegisterOutputType(DestinationSnowflakeOutput{})
}
