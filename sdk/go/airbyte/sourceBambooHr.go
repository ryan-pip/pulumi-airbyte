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

// SourceBambooHr Resource
type SourceBambooHr struct {
	pulumi.CustomResourceState

	Configuration SourceBambooHrConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceBambooHr registers a new resource with the given unique name, arguments, and options.
func NewSourceBambooHr(ctx *pulumi.Context,
	name string, args *SourceBambooHrArgs, opts ...pulumi.ResourceOption) (*SourceBambooHr, error) {
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
	var resource SourceBambooHr
	err := ctx.RegisterResource("airbyte:index/sourceBambooHr:SourceBambooHr", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceBambooHr gets an existing SourceBambooHr resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceBambooHr(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceBambooHrState, opts ...pulumi.ResourceOption) (*SourceBambooHr, error) {
	var resource SourceBambooHr
	err := ctx.ReadResource("airbyte:index/sourceBambooHr:SourceBambooHr", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceBambooHr resources.
type sourceBambooHrState struct {
	Configuration *SourceBambooHrConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceBambooHrState struct {
	Configuration SourceBambooHrConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceBambooHrState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceBambooHrState)(nil)).Elem()
}

type sourceBambooHrArgs struct {
	Configuration SourceBambooHrConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceBambooHr resource.
type SourceBambooHrArgs struct {
	Configuration SourceBambooHrConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceBambooHrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceBambooHrArgs)(nil)).Elem()
}

type SourceBambooHrInput interface {
	pulumi.Input

	ToSourceBambooHrOutput() SourceBambooHrOutput
	ToSourceBambooHrOutputWithContext(ctx context.Context) SourceBambooHrOutput
}

func (*SourceBambooHr) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceBambooHr)(nil)).Elem()
}

func (i *SourceBambooHr) ToSourceBambooHrOutput() SourceBambooHrOutput {
	return i.ToSourceBambooHrOutputWithContext(context.Background())
}

func (i *SourceBambooHr) ToSourceBambooHrOutputWithContext(ctx context.Context) SourceBambooHrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceBambooHrOutput)
}

// SourceBambooHrArrayInput is an input type that accepts SourceBambooHrArray and SourceBambooHrArrayOutput values.
// You can construct a concrete instance of `SourceBambooHrArrayInput` via:
//
//	SourceBambooHrArray{ SourceBambooHrArgs{...} }
type SourceBambooHrArrayInput interface {
	pulumi.Input

	ToSourceBambooHrArrayOutput() SourceBambooHrArrayOutput
	ToSourceBambooHrArrayOutputWithContext(context.Context) SourceBambooHrArrayOutput
}

type SourceBambooHrArray []SourceBambooHrInput

func (SourceBambooHrArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceBambooHr)(nil)).Elem()
}

func (i SourceBambooHrArray) ToSourceBambooHrArrayOutput() SourceBambooHrArrayOutput {
	return i.ToSourceBambooHrArrayOutputWithContext(context.Background())
}

func (i SourceBambooHrArray) ToSourceBambooHrArrayOutputWithContext(ctx context.Context) SourceBambooHrArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceBambooHrArrayOutput)
}

// SourceBambooHrMapInput is an input type that accepts SourceBambooHrMap and SourceBambooHrMapOutput values.
// You can construct a concrete instance of `SourceBambooHrMapInput` via:
//
//	SourceBambooHrMap{ "key": SourceBambooHrArgs{...} }
type SourceBambooHrMapInput interface {
	pulumi.Input

	ToSourceBambooHrMapOutput() SourceBambooHrMapOutput
	ToSourceBambooHrMapOutputWithContext(context.Context) SourceBambooHrMapOutput
}

type SourceBambooHrMap map[string]SourceBambooHrInput

func (SourceBambooHrMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceBambooHr)(nil)).Elem()
}

func (i SourceBambooHrMap) ToSourceBambooHrMapOutput() SourceBambooHrMapOutput {
	return i.ToSourceBambooHrMapOutputWithContext(context.Background())
}

func (i SourceBambooHrMap) ToSourceBambooHrMapOutputWithContext(ctx context.Context) SourceBambooHrMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceBambooHrMapOutput)
}

type SourceBambooHrOutput struct{ *pulumi.OutputState }

func (SourceBambooHrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceBambooHr)(nil)).Elem()
}

func (o SourceBambooHrOutput) ToSourceBambooHrOutput() SourceBambooHrOutput {
	return o
}

func (o SourceBambooHrOutput) ToSourceBambooHrOutputWithContext(ctx context.Context) SourceBambooHrOutput {
	return o
}

func (o SourceBambooHrOutput) Configuration() SourceBambooHrConfigurationOutput {
	return o.ApplyT(func(v *SourceBambooHr) SourceBambooHrConfigurationOutput { return v.Configuration }).(SourceBambooHrConfigurationOutput)
}

func (o SourceBambooHrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBambooHr) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceBambooHrOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceBambooHr) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceBambooHrOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBambooHr) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceBambooHrOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBambooHr) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceBambooHrOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBambooHr) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceBambooHrArrayOutput struct{ *pulumi.OutputState }

func (SourceBambooHrArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceBambooHr)(nil)).Elem()
}

func (o SourceBambooHrArrayOutput) ToSourceBambooHrArrayOutput() SourceBambooHrArrayOutput {
	return o
}

func (o SourceBambooHrArrayOutput) ToSourceBambooHrArrayOutputWithContext(ctx context.Context) SourceBambooHrArrayOutput {
	return o
}

func (o SourceBambooHrArrayOutput) Index(i pulumi.IntInput) SourceBambooHrOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceBambooHr {
		return vs[0].([]*SourceBambooHr)[vs[1].(int)]
	}).(SourceBambooHrOutput)
}

type SourceBambooHrMapOutput struct{ *pulumi.OutputState }

func (SourceBambooHrMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceBambooHr)(nil)).Elem()
}

func (o SourceBambooHrMapOutput) ToSourceBambooHrMapOutput() SourceBambooHrMapOutput {
	return o
}

func (o SourceBambooHrMapOutput) ToSourceBambooHrMapOutputWithContext(ctx context.Context) SourceBambooHrMapOutput {
	return o
}

func (o SourceBambooHrMapOutput) MapIndex(k pulumi.StringInput) SourceBambooHrOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceBambooHr {
		return vs[0].(map[string]*SourceBambooHr)[vs[1].(string)]
	}).(SourceBambooHrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceBambooHrInput)(nil)).Elem(), &SourceBambooHr{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceBambooHrArrayInput)(nil)).Elem(), SourceBambooHrArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceBambooHrMapInput)(nil)).Elem(), SourceBambooHrMap{})
	pulumi.RegisterOutputType(SourceBambooHrOutput{})
	pulumi.RegisterOutputType(SourceBambooHrArrayOutput{})
	pulumi.RegisterOutputType(SourceBambooHrMapOutput{})
}
