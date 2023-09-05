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

// SourceBraintree Resource
type SourceBraintree struct {
	pulumi.CustomResourceState

	Configuration SourceBraintreeConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceBraintree registers a new resource with the given unique name, arguments, and options.
func NewSourceBraintree(ctx *pulumi.Context,
	name string, args *SourceBraintreeArgs, opts ...pulumi.ResourceOption) (*SourceBraintree, error) {
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
	var resource SourceBraintree
	err := ctx.RegisterResource("airbyte:index/sourceBraintree:SourceBraintree", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceBraintree gets an existing SourceBraintree resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceBraintree(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceBraintreeState, opts ...pulumi.ResourceOption) (*SourceBraintree, error) {
	var resource SourceBraintree
	err := ctx.ReadResource("airbyte:index/sourceBraintree:SourceBraintree", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceBraintree resources.
type sourceBraintreeState struct {
	Configuration *SourceBraintreeConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceBraintreeState struct {
	Configuration SourceBraintreeConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceBraintreeState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceBraintreeState)(nil)).Elem()
}

type sourceBraintreeArgs struct {
	Configuration SourceBraintreeConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceBraintree resource.
type SourceBraintreeArgs struct {
	Configuration SourceBraintreeConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceBraintreeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceBraintreeArgs)(nil)).Elem()
}

type SourceBraintreeInput interface {
	pulumi.Input

	ToSourceBraintreeOutput() SourceBraintreeOutput
	ToSourceBraintreeOutputWithContext(ctx context.Context) SourceBraintreeOutput
}

func (*SourceBraintree) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceBraintree)(nil)).Elem()
}

func (i *SourceBraintree) ToSourceBraintreeOutput() SourceBraintreeOutput {
	return i.ToSourceBraintreeOutputWithContext(context.Background())
}

func (i *SourceBraintree) ToSourceBraintreeOutputWithContext(ctx context.Context) SourceBraintreeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceBraintreeOutput)
}

// SourceBraintreeArrayInput is an input type that accepts SourceBraintreeArray and SourceBraintreeArrayOutput values.
// You can construct a concrete instance of `SourceBraintreeArrayInput` via:
//
//	SourceBraintreeArray{ SourceBraintreeArgs{...} }
type SourceBraintreeArrayInput interface {
	pulumi.Input

	ToSourceBraintreeArrayOutput() SourceBraintreeArrayOutput
	ToSourceBraintreeArrayOutputWithContext(context.Context) SourceBraintreeArrayOutput
}

type SourceBraintreeArray []SourceBraintreeInput

func (SourceBraintreeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceBraintree)(nil)).Elem()
}

func (i SourceBraintreeArray) ToSourceBraintreeArrayOutput() SourceBraintreeArrayOutput {
	return i.ToSourceBraintreeArrayOutputWithContext(context.Background())
}

func (i SourceBraintreeArray) ToSourceBraintreeArrayOutputWithContext(ctx context.Context) SourceBraintreeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceBraintreeArrayOutput)
}

// SourceBraintreeMapInput is an input type that accepts SourceBraintreeMap and SourceBraintreeMapOutput values.
// You can construct a concrete instance of `SourceBraintreeMapInput` via:
//
//	SourceBraintreeMap{ "key": SourceBraintreeArgs{...} }
type SourceBraintreeMapInput interface {
	pulumi.Input

	ToSourceBraintreeMapOutput() SourceBraintreeMapOutput
	ToSourceBraintreeMapOutputWithContext(context.Context) SourceBraintreeMapOutput
}

type SourceBraintreeMap map[string]SourceBraintreeInput

func (SourceBraintreeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceBraintree)(nil)).Elem()
}

func (i SourceBraintreeMap) ToSourceBraintreeMapOutput() SourceBraintreeMapOutput {
	return i.ToSourceBraintreeMapOutputWithContext(context.Background())
}

func (i SourceBraintreeMap) ToSourceBraintreeMapOutputWithContext(ctx context.Context) SourceBraintreeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceBraintreeMapOutput)
}

type SourceBraintreeOutput struct{ *pulumi.OutputState }

func (SourceBraintreeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceBraintree)(nil)).Elem()
}

func (o SourceBraintreeOutput) ToSourceBraintreeOutput() SourceBraintreeOutput {
	return o
}

func (o SourceBraintreeOutput) ToSourceBraintreeOutputWithContext(ctx context.Context) SourceBraintreeOutput {
	return o
}

func (o SourceBraintreeOutput) Configuration() SourceBraintreeConfigurationOutput {
	return o.ApplyT(func(v *SourceBraintree) SourceBraintreeConfigurationOutput { return v.Configuration }).(SourceBraintreeConfigurationOutput)
}

func (o SourceBraintreeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBraintree) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceBraintreeOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceBraintree) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceBraintreeOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBraintree) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceBraintreeOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBraintree) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceBraintreeOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBraintree) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceBraintreeArrayOutput struct{ *pulumi.OutputState }

func (SourceBraintreeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceBraintree)(nil)).Elem()
}

func (o SourceBraintreeArrayOutput) ToSourceBraintreeArrayOutput() SourceBraintreeArrayOutput {
	return o
}

func (o SourceBraintreeArrayOutput) ToSourceBraintreeArrayOutputWithContext(ctx context.Context) SourceBraintreeArrayOutput {
	return o
}

func (o SourceBraintreeArrayOutput) Index(i pulumi.IntInput) SourceBraintreeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceBraintree {
		return vs[0].([]*SourceBraintree)[vs[1].(int)]
	}).(SourceBraintreeOutput)
}

type SourceBraintreeMapOutput struct{ *pulumi.OutputState }

func (SourceBraintreeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceBraintree)(nil)).Elem()
}

func (o SourceBraintreeMapOutput) ToSourceBraintreeMapOutput() SourceBraintreeMapOutput {
	return o
}

func (o SourceBraintreeMapOutput) ToSourceBraintreeMapOutputWithContext(ctx context.Context) SourceBraintreeMapOutput {
	return o
}

func (o SourceBraintreeMapOutput) MapIndex(k pulumi.StringInput) SourceBraintreeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceBraintree {
		return vs[0].(map[string]*SourceBraintree)[vs[1].(string)]
	}).(SourceBraintreeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceBraintreeInput)(nil)).Elem(), &SourceBraintree{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceBraintreeArrayInput)(nil)).Elem(), SourceBraintreeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceBraintreeMapInput)(nil)).Elem(), SourceBraintreeMap{})
	pulumi.RegisterOutputType(SourceBraintreeOutput{})
	pulumi.RegisterOutputType(SourceBraintreeArrayOutput{})
	pulumi.RegisterOutputType(SourceBraintreeMapOutput{})
}
