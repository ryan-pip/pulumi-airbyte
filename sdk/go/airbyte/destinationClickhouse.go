// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationClickhouse Resource
type DestinationClickhouse struct {
	pulumi.CustomResourceState

	Configuration   DestinationClickhouseConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                      `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                      `pulumi:"destinationType"`
	Name            pulumi.StringOutput                      `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                      `pulumi:"workspaceId"`
}

// NewDestinationClickhouse registers a new resource with the given unique name, arguments, and options.
func NewDestinationClickhouse(ctx *pulumi.Context,
	name string, args *DestinationClickhouseArgs, opts ...pulumi.ResourceOption) (*DestinationClickhouse, error) {
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
	var resource DestinationClickhouse
	err := ctx.RegisterResource("airbyte:index/destinationClickhouse:DestinationClickhouse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationClickhouse gets an existing DestinationClickhouse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationClickhouse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationClickhouseState, opts ...pulumi.ResourceOption) (*DestinationClickhouse, error) {
	var resource DestinationClickhouse
	err := ctx.ReadResource("airbyte:index/destinationClickhouse:DestinationClickhouse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationClickhouse resources.
type destinationClickhouseState struct {
	Configuration   *DestinationClickhouseConfiguration `pulumi:"configuration"`
	DestinationId   *string                             `pulumi:"destinationId"`
	DestinationType *string                             `pulumi:"destinationType"`
	Name            *string                             `pulumi:"name"`
	WorkspaceId     *string                             `pulumi:"workspaceId"`
}

type DestinationClickhouseState struct {
	Configuration   DestinationClickhouseConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationClickhouseState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationClickhouseState)(nil)).Elem()
}

type destinationClickhouseArgs struct {
	Configuration DestinationClickhouseConfiguration `pulumi:"configuration"`
	Name          string                             `pulumi:"name"`
	WorkspaceId   string                             `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationClickhouse resource.
type DestinationClickhouseArgs struct {
	Configuration DestinationClickhouseConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationClickhouseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationClickhouseArgs)(nil)).Elem()
}

type DestinationClickhouseInput interface {
	pulumi.Input

	ToDestinationClickhouseOutput() DestinationClickhouseOutput
	ToDestinationClickhouseOutputWithContext(ctx context.Context) DestinationClickhouseOutput
}

func (*DestinationClickhouse) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationClickhouse)(nil)).Elem()
}

func (i *DestinationClickhouse) ToDestinationClickhouseOutput() DestinationClickhouseOutput {
	return i.ToDestinationClickhouseOutputWithContext(context.Background())
}

func (i *DestinationClickhouse) ToDestinationClickhouseOutputWithContext(ctx context.Context) DestinationClickhouseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationClickhouseOutput)
}

// DestinationClickhouseArrayInput is an input type that accepts DestinationClickhouseArray and DestinationClickhouseArrayOutput values.
// You can construct a concrete instance of `DestinationClickhouseArrayInput` via:
//
//	DestinationClickhouseArray{ DestinationClickhouseArgs{...} }
type DestinationClickhouseArrayInput interface {
	pulumi.Input

	ToDestinationClickhouseArrayOutput() DestinationClickhouseArrayOutput
	ToDestinationClickhouseArrayOutputWithContext(context.Context) DestinationClickhouseArrayOutput
}

type DestinationClickhouseArray []DestinationClickhouseInput

func (DestinationClickhouseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationClickhouse)(nil)).Elem()
}

func (i DestinationClickhouseArray) ToDestinationClickhouseArrayOutput() DestinationClickhouseArrayOutput {
	return i.ToDestinationClickhouseArrayOutputWithContext(context.Background())
}

func (i DestinationClickhouseArray) ToDestinationClickhouseArrayOutputWithContext(ctx context.Context) DestinationClickhouseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationClickhouseArrayOutput)
}

// DestinationClickhouseMapInput is an input type that accepts DestinationClickhouseMap and DestinationClickhouseMapOutput values.
// You can construct a concrete instance of `DestinationClickhouseMapInput` via:
//
//	DestinationClickhouseMap{ "key": DestinationClickhouseArgs{...} }
type DestinationClickhouseMapInput interface {
	pulumi.Input

	ToDestinationClickhouseMapOutput() DestinationClickhouseMapOutput
	ToDestinationClickhouseMapOutputWithContext(context.Context) DestinationClickhouseMapOutput
}

type DestinationClickhouseMap map[string]DestinationClickhouseInput

func (DestinationClickhouseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationClickhouse)(nil)).Elem()
}

func (i DestinationClickhouseMap) ToDestinationClickhouseMapOutput() DestinationClickhouseMapOutput {
	return i.ToDestinationClickhouseMapOutputWithContext(context.Background())
}

func (i DestinationClickhouseMap) ToDestinationClickhouseMapOutputWithContext(ctx context.Context) DestinationClickhouseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationClickhouseMapOutput)
}

type DestinationClickhouseOutput struct{ *pulumi.OutputState }

func (DestinationClickhouseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationClickhouse)(nil)).Elem()
}

func (o DestinationClickhouseOutput) ToDestinationClickhouseOutput() DestinationClickhouseOutput {
	return o
}

func (o DestinationClickhouseOutput) ToDestinationClickhouseOutputWithContext(ctx context.Context) DestinationClickhouseOutput {
	return o
}

func (o DestinationClickhouseOutput) Configuration() DestinationClickhouseConfigurationOutput {
	return o.ApplyT(func(v *DestinationClickhouse) DestinationClickhouseConfigurationOutput { return v.Configuration }).(DestinationClickhouseConfigurationOutput)
}

func (o DestinationClickhouseOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationClickhouse) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationClickhouseOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationClickhouse) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationClickhouseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationClickhouse) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationClickhouseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationClickhouse) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type DestinationClickhouseArrayOutput struct{ *pulumi.OutputState }

func (DestinationClickhouseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationClickhouse)(nil)).Elem()
}

func (o DestinationClickhouseArrayOutput) ToDestinationClickhouseArrayOutput() DestinationClickhouseArrayOutput {
	return o
}

func (o DestinationClickhouseArrayOutput) ToDestinationClickhouseArrayOutputWithContext(ctx context.Context) DestinationClickhouseArrayOutput {
	return o
}

func (o DestinationClickhouseArrayOutput) Index(i pulumi.IntInput) DestinationClickhouseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationClickhouse {
		return vs[0].([]*DestinationClickhouse)[vs[1].(int)]
	}).(DestinationClickhouseOutput)
}

type DestinationClickhouseMapOutput struct{ *pulumi.OutputState }

func (DestinationClickhouseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationClickhouse)(nil)).Elem()
}

func (o DestinationClickhouseMapOutput) ToDestinationClickhouseMapOutput() DestinationClickhouseMapOutput {
	return o
}

func (o DestinationClickhouseMapOutput) ToDestinationClickhouseMapOutputWithContext(ctx context.Context) DestinationClickhouseMapOutput {
	return o
}

func (o DestinationClickhouseMapOutput) MapIndex(k pulumi.StringInput) DestinationClickhouseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationClickhouse {
		return vs[0].(map[string]*DestinationClickhouse)[vs[1].(string)]
	}).(DestinationClickhouseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationClickhouseInput)(nil)).Elem(), &DestinationClickhouse{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationClickhouseArrayInput)(nil)).Elem(), DestinationClickhouseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationClickhouseMapInput)(nil)).Elem(), DestinationClickhouseMap{})
	pulumi.RegisterOutputType(DestinationClickhouseOutput{})
	pulumi.RegisterOutputType(DestinationClickhouseArrayOutput{})
	pulumi.RegisterOutputType(DestinationClickhouseMapOutput{})
}
