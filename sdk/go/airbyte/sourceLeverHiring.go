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

// SourceLeverHiring Resource
type SourceLeverHiring struct {
	pulumi.CustomResourceState

	Configuration SourceLeverHiringConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceLeverHiring registers a new resource with the given unique name, arguments, and options.
func NewSourceLeverHiring(ctx *pulumi.Context,
	name string, args *SourceLeverHiringArgs, opts ...pulumi.ResourceOption) (*SourceLeverHiring, error) {
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
	var resource SourceLeverHiring
	err := ctx.RegisterResource("airbyte:index/sourceLeverHiring:SourceLeverHiring", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceLeverHiring gets an existing SourceLeverHiring resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceLeverHiring(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceLeverHiringState, opts ...pulumi.ResourceOption) (*SourceLeverHiring, error) {
	var resource SourceLeverHiring
	err := ctx.ReadResource("airbyte:index/sourceLeverHiring:SourceLeverHiring", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceLeverHiring resources.
type sourceLeverHiringState struct {
	Configuration *SourceLeverHiringConfiguration `pulumi:"configuration"`
	Name          *string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceLeverHiringState struct {
	Configuration SourceLeverHiringConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceLeverHiringState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLeverHiringState)(nil)).Elem()
}

type sourceLeverHiringArgs struct {
	Configuration SourceLeverHiringConfiguration `pulumi:"configuration"`
	Name          string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceLeverHiring resource.
type SourceLeverHiringArgs struct {
	Configuration SourceLeverHiringConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceLeverHiringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLeverHiringArgs)(nil)).Elem()
}

type SourceLeverHiringInput interface {
	pulumi.Input

	ToSourceLeverHiringOutput() SourceLeverHiringOutput
	ToSourceLeverHiringOutputWithContext(ctx context.Context) SourceLeverHiringOutput
}

func (*SourceLeverHiring) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLeverHiring)(nil)).Elem()
}

func (i *SourceLeverHiring) ToSourceLeverHiringOutput() SourceLeverHiringOutput {
	return i.ToSourceLeverHiringOutputWithContext(context.Background())
}

func (i *SourceLeverHiring) ToSourceLeverHiringOutputWithContext(ctx context.Context) SourceLeverHiringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLeverHiringOutput)
}

// SourceLeverHiringArrayInput is an input type that accepts SourceLeverHiringArray and SourceLeverHiringArrayOutput values.
// You can construct a concrete instance of `SourceLeverHiringArrayInput` via:
//
//	SourceLeverHiringArray{ SourceLeverHiringArgs{...} }
type SourceLeverHiringArrayInput interface {
	pulumi.Input

	ToSourceLeverHiringArrayOutput() SourceLeverHiringArrayOutput
	ToSourceLeverHiringArrayOutputWithContext(context.Context) SourceLeverHiringArrayOutput
}

type SourceLeverHiringArray []SourceLeverHiringInput

func (SourceLeverHiringArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceLeverHiring)(nil)).Elem()
}

func (i SourceLeverHiringArray) ToSourceLeverHiringArrayOutput() SourceLeverHiringArrayOutput {
	return i.ToSourceLeverHiringArrayOutputWithContext(context.Background())
}

func (i SourceLeverHiringArray) ToSourceLeverHiringArrayOutputWithContext(ctx context.Context) SourceLeverHiringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLeverHiringArrayOutput)
}

// SourceLeverHiringMapInput is an input type that accepts SourceLeverHiringMap and SourceLeverHiringMapOutput values.
// You can construct a concrete instance of `SourceLeverHiringMapInput` via:
//
//	SourceLeverHiringMap{ "key": SourceLeverHiringArgs{...} }
type SourceLeverHiringMapInput interface {
	pulumi.Input

	ToSourceLeverHiringMapOutput() SourceLeverHiringMapOutput
	ToSourceLeverHiringMapOutputWithContext(context.Context) SourceLeverHiringMapOutput
}

type SourceLeverHiringMap map[string]SourceLeverHiringInput

func (SourceLeverHiringMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceLeverHiring)(nil)).Elem()
}

func (i SourceLeverHiringMap) ToSourceLeverHiringMapOutput() SourceLeverHiringMapOutput {
	return i.ToSourceLeverHiringMapOutputWithContext(context.Background())
}

func (i SourceLeverHiringMap) ToSourceLeverHiringMapOutputWithContext(ctx context.Context) SourceLeverHiringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLeverHiringMapOutput)
}

type SourceLeverHiringOutput struct{ *pulumi.OutputState }

func (SourceLeverHiringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLeverHiring)(nil)).Elem()
}

func (o SourceLeverHiringOutput) ToSourceLeverHiringOutput() SourceLeverHiringOutput {
	return o
}

func (o SourceLeverHiringOutput) ToSourceLeverHiringOutputWithContext(ctx context.Context) SourceLeverHiringOutput {
	return o
}

func (o SourceLeverHiringOutput) Configuration() SourceLeverHiringConfigurationOutput {
	return o.ApplyT(func(v *SourceLeverHiring) SourceLeverHiringConfigurationOutput { return v.Configuration }).(SourceLeverHiringConfigurationOutput)
}

func (o SourceLeverHiringOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLeverHiring) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceLeverHiringOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLeverHiring) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceLeverHiringOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLeverHiring) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceLeverHiringOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLeverHiring) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceLeverHiringOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLeverHiring) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceLeverHiringArrayOutput struct{ *pulumi.OutputState }

func (SourceLeverHiringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceLeverHiring)(nil)).Elem()
}

func (o SourceLeverHiringArrayOutput) ToSourceLeverHiringArrayOutput() SourceLeverHiringArrayOutput {
	return o
}

func (o SourceLeverHiringArrayOutput) ToSourceLeverHiringArrayOutputWithContext(ctx context.Context) SourceLeverHiringArrayOutput {
	return o
}

func (o SourceLeverHiringArrayOutput) Index(i pulumi.IntInput) SourceLeverHiringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceLeverHiring {
		return vs[0].([]*SourceLeverHiring)[vs[1].(int)]
	}).(SourceLeverHiringOutput)
}

type SourceLeverHiringMapOutput struct{ *pulumi.OutputState }

func (SourceLeverHiringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceLeverHiring)(nil)).Elem()
}

func (o SourceLeverHiringMapOutput) ToSourceLeverHiringMapOutput() SourceLeverHiringMapOutput {
	return o
}

func (o SourceLeverHiringMapOutput) ToSourceLeverHiringMapOutputWithContext(ctx context.Context) SourceLeverHiringMapOutput {
	return o
}

func (o SourceLeverHiringMapOutput) MapIndex(k pulumi.StringInput) SourceLeverHiringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceLeverHiring {
		return vs[0].(map[string]*SourceLeverHiring)[vs[1].(string)]
	}).(SourceLeverHiringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLeverHiringInput)(nil)).Elem(), &SourceLeverHiring{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLeverHiringArrayInput)(nil)).Elem(), SourceLeverHiringArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLeverHiringMapInput)(nil)).Elem(), SourceLeverHiringMap{})
	pulumi.RegisterOutputType(SourceLeverHiringOutput{})
	pulumi.RegisterOutputType(SourceLeverHiringArrayOutput{})
	pulumi.RegisterOutputType(SourceLeverHiringMapOutput{})
}
