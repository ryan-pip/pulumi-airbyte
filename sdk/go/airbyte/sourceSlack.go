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

// SourceSlack Resource
type SourceSlack struct {
	pulumi.CustomResourceState

	Configuration SourceSlackConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceSlack registers a new resource with the given unique name, arguments, and options.
func NewSourceSlack(ctx *pulumi.Context,
	name string, args *SourceSlackArgs, opts ...pulumi.ResourceOption) (*SourceSlack, error) {
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
	var resource SourceSlack
	err := ctx.RegisterResource("airbyte:index/sourceSlack:SourceSlack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSlack gets an existing SourceSlack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSlack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSlackState, opts ...pulumi.ResourceOption) (*SourceSlack, error) {
	var resource SourceSlack
	err := ctx.ReadResource("airbyte:index/sourceSlack:SourceSlack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSlack resources.
type sourceSlackState struct {
	Configuration *SourceSlackConfiguration `pulumi:"configuration"`
	Name          *string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceSlackState struct {
	Configuration SourceSlackConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceSlackState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSlackState)(nil)).Elem()
}

type sourceSlackArgs struct {
	Configuration SourceSlackConfiguration `pulumi:"configuration"`
	Name          string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSlack resource.
type SourceSlackArgs struct {
	Configuration SourceSlackConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSlackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSlackArgs)(nil)).Elem()
}

type SourceSlackInput interface {
	pulumi.Input

	ToSourceSlackOutput() SourceSlackOutput
	ToSourceSlackOutputWithContext(ctx context.Context) SourceSlackOutput
}

func (*SourceSlack) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSlack)(nil)).Elem()
}

func (i *SourceSlack) ToSourceSlackOutput() SourceSlackOutput {
	return i.ToSourceSlackOutputWithContext(context.Background())
}

func (i *SourceSlack) ToSourceSlackOutputWithContext(ctx context.Context) SourceSlackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSlackOutput)
}

// SourceSlackArrayInput is an input type that accepts SourceSlackArray and SourceSlackArrayOutput values.
// You can construct a concrete instance of `SourceSlackArrayInput` via:
//
//	SourceSlackArray{ SourceSlackArgs{...} }
type SourceSlackArrayInput interface {
	pulumi.Input

	ToSourceSlackArrayOutput() SourceSlackArrayOutput
	ToSourceSlackArrayOutputWithContext(context.Context) SourceSlackArrayOutput
}

type SourceSlackArray []SourceSlackInput

func (SourceSlackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSlack)(nil)).Elem()
}

func (i SourceSlackArray) ToSourceSlackArrayOutput() SourceSlackArrayOutput {
	return i.ToSourceSlackArrayOutputWithContext(context.Background())
}

func (i SourceSlackArray) ToSourceSlackArrayOutputWithContext(ctx context.Context) SourceSlackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSlackArrayOutput)
}

// SourceSlackMapInput is an input type that accepts SourceSlackMap and SourceSlackMapOutput values.
// You can construct a concrete instance of `SourceSlackMapInput` via:
//
//	SourceSlackMap{ "key": SourceSlackArgs{...} }
type SourceSlackMapInput interface {
	pulumi.Input

	ToSourceSlackMapOutput() SourceSlackMapOutput
	ToSourceSlackMapOutputWithContext(context.Context) SourceSlackMapOutput
}

type SourceSlackMap map[string]SourceSlackInput

func (SourceSlackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSlack)(nil)).Elem()
}

func (i SourceSlackMap) ToSourceSlackMapOutput() SourceSlackMapOutput {
	return i.ToSourceSlackMapOutputWithContext(context.Background())
}

func (i SourceSlackMap) ToSourceSlackMapOutputWithContext(ctx context.Context) SourceSlackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSlackMapOutput)
}

type SourceSlackOutput struct{ *pulumi.OutputState }

func (SourceSlackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSlack)(nil)).Elem()
}

func (o SourceSlackOutput) ToSourceSlackOutput() SourceSlackOutput {
	return o
}

func (o SourceSlackOutput) ToSourceSlackOutputWithContext(ctx context.Context) SourceSlackOutput {
	return o
}

func (o SourceSlackOutput) Configuration() SourceSlackConfigurationOutput {
	return o.ApplyT(func(v *SourceSlack) SourceSlackConfigurationOutput { return v.Configuration }).(SourceSlackConfigurationOutput)
}

func (o SourceSlackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSlack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceSlackOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSlack) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSlackOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSlack) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSlackOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSlack) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSlackOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSlack) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceSlackArrayOutput struct{ *pulumi.OutputState }

func (SourceSlackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSlack)(nil)).Elem()
}

func (o SourceSlackArrayOutput) ToSourceSlackArrayOutput() SourceSlackArrayOutput {
	return o
}

func (o SourceSlackArrayOutput) ToSourceSlackArrayOutputWithContext(ctx context.Context) SourceSlackArrayOutput {
	return o
}

func (o SourceSlackArrayOutput) Index(i pulumi.IntInput) SourceSlackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceSlack {
		return vs[0].([]*SourceSlack)[vs[1].(int)]
	}).(SourceSlackOutput)
}

type SourceSlackMapOutput struct{ *pulumi.OutputState }

func (SourceSlackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSlack)(nil)).Elem()
}

func (o SourceSlackMapOutput) ToSourceSlackMapOutput() SourceSlackMapOutput {
	return o
}

func (o SourceSlackMapOutput) ToSourceSlackMapOutputWithContext(ctx context.Context) SourceSlackMapOutput {
	return o
}

func (o SourceSlackMapOutput) MapIndex(k pulumi.StringInput) SourceSlackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceSlack {
		return vs[0].(map[string]*SourceSlack)[vs[1].(string)]
	}).(SourceSlackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSlackInput)(nil)).Elem(), &SourceSlack{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSlackArrayInput)(nil)).Elem(), SourceSlackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSlackMapInput)(nil)).Elem(), SourceSlackMap{})
	pulumi.RegisterOutputType(SourceSlackOutput{})
	pulumi.RegisterOutputType(SourceSlackArrayOutput{})
	pulumi.RegisterOutputType(SourceSlackMapOutput{})
}
