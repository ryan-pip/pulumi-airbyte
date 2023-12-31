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

// DestinationFirebolt Resource
type DestinationFirebolt struct {
	pulumi.CustomResourceState

	Configuration   DestinationFireboltConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                    `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                    `pulumi:"destinationType"`
	Name            pulumi.StringOutput                    `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                    `pulumi:"workspaceId"`
}

// NewDestinationFirebolt registers a new resource with the given unique name, arguments, and options.
func NewDestinationFirebolt(ctx *pulumi.Context,
	name string, args *DestinationFireboltArgs, opts ...pulumi.ResourceOption) (*DestinationFirebolt, error) {
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
	var resource DestinationFirebolt
	err := ctx.RegisterResource("airbyte:index/destinationFirebolt:DestinationFirebolt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationFirebolt gets an existing DestinationFirebolt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationFirebolt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationFireboltState, opts ...pulumi.ResourceOption) (*DestinationFirebolt, error) {
	var resource DestinationFirebolt
	err := ctx.ReadResource("airbyte:index/destinationFirebolt:DestinationFirebolt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationFirebolt resources.
type destinationFireboltState struct {
	Configuration   *DestinationFireboltConfiguration `pulumi:"configuration"`
	DestinationId   *string                           `pulumi:"destinationId"`
	DestinationType *string                           `pulumi:"destinationType"`
	Name            *string                           `pulumi:"name"`
	WorkspaceId     *string                           `pulumi:"workspaceId"`
}

type DestinationFireboltState struct {
	Configuration   DestinationFireboltConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationFireboltState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationFireboltState)(nil)).Elem()
}

type destinationFireboltArgs struct {
	Configuration DestinationFireboltConfiguration `pulumi:"configuration"`
	Name          string                           `pulumi:"name"`
	WorkspaceId   string                           `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationFirebolt resource.
type DestinationFireboltArgs struct {
	Configuration DestinationFireboltConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationFireboltArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationFireboltArgs)(nil)).Elem()
}

type DestinationFireboltInput interface {
	pulumi.Input

	ToDestinationFireboltOutput() DestinationFireboltOutput
	ToDestinationFireboltOutputWithContext(ctx context.Context) DestinationFireboltOutput
}

func (*DestinationFirebolt) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationFirebolt)(nil)).Elem()
}

func (i *DestinationFirebolt) ToDestinationFireboltOutput() DestinationFireboltOutput {
	return i.ToDestinationFireboltOutputWithContext(context.Background())
}

func (i *DestinationFirebolt) ToDestinationFireboltOutputWithContext(ctx context.Context) DestinationFireboltOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationFireboltOutput)
}

// DestinationFireboltArrayInput is an input type that accepts DestinationFireboltArray and DestinationFireboltArrayOutput values.
// You can construct a concrete instance of `DestinationFireboltArrayInput` via:
//
//	DestinationFireboltArray{ DestinationFireboltArgs{...} }
type DestinationFireboltArrayInput interface {
	pulumi.Input

	ToDestinationFireboltArrayOutput() DestinationFireboltArrayOutput
	ToDestinationFireboltArrayOutputWithContext(context.Context) DestinationFireboltArrayOutput
}

type DestinationFireboltArray []DestinationFireboltInput

func (DestinationFireboltArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationFirebolt)(nil)).Elem()
}

func (i DestinationFireboltArray) ToDestinationFireboltArrayOutput() DestinationFireboltArrayOutput {
	return i.ToDestinationFireboltArrayOutputWithContext(context.Background())
}

func (i DestinationFireboltArray) ToDestinationFireboltArrayOutputWithContext(ctx context.Context) DestinationFireboltArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationFireboltArrayOutput)
}

// DestinationFireboltMapInput is an input type that accepts DestinationFireboltMap and DestinationFireboltMapOutput values.
// You can construct a concrete instance of `DestinationFireboltMapInput` via:
//
//	DestinationFireboltMap{ "key": DestinationFireboltArgs{...} }
type DestinationFireboltMapInput interface {
	pulumi.Input

	ToDestinationFireboltMapOutput() DestinationFireboltMapOutput
	ToDestinationFireboltMapOutputWithContext(context.Context) DestinationFireboltMapOutput
}

type DestinationFireboltMap map[string]DestinationFireboltInput

func (DestinationFireboltMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationFirebolt)(nil)).Elem()
}

func (i DestinationFireboltMap) ToDestinationFireboltMapOutput() DestinationFireboltMapOutput {
	return i.ToDestinationFireboltMapOutputWithContext(context.Background())
}

func (i DestinationFireboltMap) ToDestinationFireboltMapOutputWithContext(ctx context.Context) DestinationFireboltMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationFireboltMapOutput)
}

type DestinationFireboltOutput struct{ *pulumi.OutputState }

func (DestinationFireboltOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationFirebolt)(nil)).Elem()
}

func (o DestinationFireboltOutput) ToDestinationFireboltOutput() DestinationFireboltOutput {
	return o
}

func (o DestinationFireboltOutput) ToDestinationFireboltOutputWithContext(ctx context.Context) DestinationFireboltOutput {
	return o
}

func (o DestinationFireboltOutput) Configuration() DestinationFireboltConfigurationOutput {
	return o.ApplyT(func(v *DestinationFirebolt) DestinationFireboltConfigurationOutput { return v.Configuration }).(DestinationFireboltConfigurationOutput)
}

func (o DestinationFireboltOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationFirebolt) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationFireboltOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationFirebolt) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationFireboltOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationFirebolt) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationFireboltOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationFirebolt) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type DestinationFireboltArrayOutput struct{ *pulumi.OutputState }

func (DestinationFireboltArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationFirebolt)(nil)).Elem()
}

func (o DestinationFireboltArrayOutput) ToDestinationFireboltArrayOutput() DestinationFireboltArrayOutput {
	return o
}

func (o DestinationFireboltArrayOutput) ToDestinationFireboltArrayOutputWithContext(ctx context.Context) DestinationFireboltArrayOutput {
	return o
}

func (o DestinationFireboltArrayOutput) Index(i pulumi.IntInput) DestinationFireboltOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationFirebolt {
		return vs[0].([]*DestinationFirebolt)[vs[1].(int)]
	}).(DestinationFireboltOutput)
}

type DestinationFireboltMapOutput struct{ *pulumi.OutputState }

func (DestinationFireboltMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationFirebolt)(nil)).Elem()
}

func (o DestinationFireboltMapOutput) ToDestinationFireboltMapOutput() DestinationFireboltMapOutput {
	return o
}

func (o DestinationFireboltMapOutput) ToDestinationFireboltMapOutputWithContext(ctx context.Context) DestinationFireboltMapOutput {
	return o
}

func (o DestinationFireboltMapOutput) MapIndex(k pulumi.StringInput) DestinationFireboltOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationFirebolt {
		return vs[0].(map[string]*DestinationFirebolt)[vs[1].(string)]
	}).(DestinationFireboltOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationFireboltInput)(nil)).Elem(), &DestinationFirebolt{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationFireboltArrayInput)(nil)).Elem(), DestinationFireboltArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationFireboltMapInput)(nil)).Elem(), DestinationFireboltMap{})
	pulumi.RegisterOutputType(DestinationFireboltOutput{})
	pulumi.RegisterOutputType(DestinationFireboltArrayOutput{})
	pulumi.RegisterOutputType(DestinationFireboltMapOutput{})
}
