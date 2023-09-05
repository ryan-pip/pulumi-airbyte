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

// DestinationTimeplus Resource
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
	opts = internal.PkgResourceDefaultOpts(opts)
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

// DestinationTimeplusArrayInput is an input type that accepts DestinationTimeplusArray and DestinationTimeplusArrayOutput values.
// You can construct a concrete instance of `DestinationTimeplusArrayInput` via:
//
//	DestinationTimeplusArray{ DestinationTimeplusArgs{...} }
type DestinationTimeplusArrayInput interface {
	pulumi.Input

	ToDestinationTimeplusArrayOutput() DestinationTimeplusArrayOutput
	ToDestinationTimeplusArrayOutputWithContext(context.Context) DestinationTimeplusArrayOutput
}

type DestinationTimeplusArray []DestinationTimeplusInput

func (DestinationTimeplusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationTimeplus)(nil)).Elem()
}

func (i DestinationTimeplusArray) ToDestinationTimeplusArrayOutput() DestinationTimeplusArrayOutput {
	return i.ToDestinationTimeplusArrayOutputWithContext(context.Background())
}

func (i DestinationTimeplusArray) ToDestinationTimeplusArrayOutputWithContext(ctx context.Context) DestinationTimeplusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationTimeplusArrayOutput)
}

// DestinationTimeplusMapInput is an input type that accepts DestinationTimeplusMap and DestinationTimeplusMapOutput values.
// You can construct a concrete instance of `DestinationTimeplusMapInput` via:
//
//	DestinationTimeplusMap{ "key": DestinationTimeplusArgs{...} }
type DestinationTimeplusMapInput interface {
	pulumi.Input

	ToDestinationTimeplusMapOutput() DestinationTimeplusMapOutput
	ToDestinationTimeplusMapOutputWithContext(context.Context) DestinationTimeplusMapOutput
}

type DestinationTimeplusMap map[string]DestinationTimeplusInput

func (DestinationTimeplusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationTimeplus)(nil)).Elem()
}

func (i DestinationTimeplusMap) ToDestinationTimeplusMapOutput() DestinationTimeplusMapOutput {
	return i.ToDestinationTimeplusMapOutputWithContext(context.Background())
}

func (i DestinationTimeplusMap) ToDestinationTimeplusMapOutputWithContext(ctx context.Context) DestinationTimeplusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationTimeplusMapOutput)
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

type DestinationTimeplusArrayOutput struct{ *pulumi.OutputState }

func (DestinationTimeplusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationTimeplus)(nil)).Elem()
}

func (o DestinationTimeplusArrayOutput) ToDestinationTimeplusArrayOutput() DestinationTimeplusArrayOutput {
	return o
}

func (o DestinationTimeplusArrayOutput) ToDestinationTimeplusArrayOutputWithContext(ctx context.Context) DestinationTimeplusArrayOutput {
	return o
}

func (o DestinationTimeplusArrayOutput) Index(i pulumi.IntInput) DestinationTimeplusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationTimeplus {
		return vs[0].([]*DestinationTimeplus)[vs[1].(int)]
	}).(DestinationTimeplusOutput)
}

type DestinationTimeplusMapOutput struct{ *pulumi.OutputState }

func (DestinationTimeplusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationTimeplus)(nil)).Elem()
}

func (o DestinationTimeplusMapOutput) ToDestinationTimeplusMapOutput() DestinationTimeplusMapOutput {
	return o
}

func (o DestinationTimeplusMapOutput) ToDestinationTimeplusMapOutputWithContext(ctx context.Context) DestinationTimeplusMapOutput {
	return o
}

func (o DestinationTimeplusMapOutput) MapIndex(k pulumi.StringInput) DestinationTimeplusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationTimeplus {
		return vs[0].(map[string]*DestinationTimeplus)[vs[1].(string)]
	}).(DestinationTimeplusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationTimeplusInput)(nil)).Elem(), &DestinationTimeplus{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationTimeplusArrayInput)(nil)).Elem(), DestinationTimeplusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationTimeplusMapInput)(nil)).Elem(), DestinationTimeplusMap{})
	pulumi.RegisterOutputType(DestinationTimeplusOutput{})
	pulumi.RegisterOutputType(DestinationTimeplusArrayOutput{})
	pulumi.RegisterOutputType(DestinationTimeplusMapOutput{})
}
