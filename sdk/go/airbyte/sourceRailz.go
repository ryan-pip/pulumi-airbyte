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

// SourceRailz Resource
type SourceRailz struct {
	pulumi.CustomResourceState

	Configuration SourceRailzConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceRailz registers a new resource with the given unique name, arguments, and options.
func NewSourceRailz(ctx *pulumi.Context,
	name string, args *SourceRailzArgs, opts ...pulumi.ResourceOption) (*SourceRailz, error) {
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
	var resource SourceRailz
	err := ctx.RegisterResource("airbyte:index/sourceRailz:SourceRailz", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceRailz gets an existing SourceRailz resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceRailz(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceRailzState, opts ...pulumi.ResourceOption) (*SourceRailz, error) {
	var resource SourceRailz
	err := ctx.ReadResource("airbyte:index/sourceRailz:SourceRailz", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceRailz resources.
type sourceRailzState struct {
	Configuration *SourceRailzConfiguration `pulumi:"configuration"`
	Name          *string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceRailzState struct {
	Configuration SourceRailzConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceRailzState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRailzState)(nil)).Elem()
}

type sourceRailzArgs struct {
	Configuration SourceRailzConfiguration `pulumi:"configuration"`
	Name          string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceRailz resource.
type SourceRailzArgs struct {
	Configuration SourceRailzConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceRailzArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRailzArgs)(nil)).Elem()
}

type SourceRailzInput interface {
	pulumi.Input

	ToSourceRailzOutput() SourceRailzOutput
	ToSourceRailzOutputWithContext(ctx context.Context) SourceRailzOutput
}

func (*SourceRailz) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRailz)(nil)).Elem()
}

func (i *SourceRailz) ToSourceRailzOutput() SourceRailzOutput {
	return i.ToSourceRailzOutputWithContext(context.Background())
}

func (i *SourceRailz) ToSourceRailzOutputWithContext(ctx context.Context) SourceRailzOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRailzOutput)
}

// SourceRailzArrayInput is an input type that accepts SourceRailzArray and SourceRailzArrayOutput values.
// You can construct a concrete instance of `SourceRailzArrayInput` via:
//
//	SourceRailzArray{ SourceRailzArgs{...} }
type SourceRailzArrayInput interface {
	pulumi.Input

	ToSourceRailzArrayOutput() SourceRailzArrayOutput
	ToSourceRailzArrayOutputWithContext(context.Context) SourceRailzArrayOutput
}

type SourceRailzArray []SourceRailzInput

func (SourceRailzArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceRailz)(nil)).Elem()
}

func (i SourceRailzArray) ToSourceRailzArrayOutput() SourceRailzArrayOutput {
	return i.ToSourceRailzArrayOutputWithContext(context.Background())
}

func (i SourceRailzArray) ToSourceRailzArrayOutputWithContext(ctx context.Context) SourceRailzArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRailzArrayOutput)
}

// SourceRailzMapInput is an input type that accepts SourceRailzMap and SourceRailzMapOutput values.
// You can construct a concrete instance of `SourceRailzMapInput` via:
//
//	SourceRailzMap{ "key": SourceRailzArgs{...} }
type SourceRailzMapInput interface {
	pulumi.Input

	ToSourceRailzMapOutput() SourceRailzMapOutput
	ToSourceRailzMapOutputWithContext(context.Context) SourceRailzMapOutput
}

type SourceRailzMap map[string]SourceRailzInput

func (SourceRailzMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceRailz)(nil)).Elem()
}

func (i SourceRailzMap) ToSourceRailzMapOutput() SourceRailzMapOutput {
	return i.ToSourceRailzMapOutputWithContext(context.Background())
}

func (i SourceRailzMap) ToSourceRailzMapOutputWithContext(ctx context.Context) SourceRailzMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRailzMapOutput)
}

type SourceRailzOutput struct{ *pulumi.OutputState }

func (SourceRailzOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRailz)(nil)).Elem()
}

func (o SourceRailzOutput) ToSourceRailzOutput() SourceRailzOutput {
	return o
}

func (o SourceRailzOutput) ToSourceRailzOutputWithContext(ctx context.Context) SourceRailzOutput {
	return o
}

func (o SourceRailzOutput) Configuration() SourceRailzConfigurationOutput {
	return o.ApplyT(func(v *SourceRailz) SourceRailzConfigurationOutput { return v.Configuration }).(SourceRailzConfigurationOutput)
}

func (o SourceRailzOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRailz) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceRailzOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceRailz) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceRailzOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRailz) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceRailzOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRailz) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceRailzOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRailz) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceRailzArrayOutput struct{ *pulumi.OutputState }

func (SourceRailzArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceRailz)(nil)).Elem()
}

func (o SourceRailzArrayOutput) ToSourceRailzArrayOutput() SourceRailzArrayOutput {
	return o
}

func (o SourceRailzArrayOutput) ToSourceRailzArrayOutputWithContext(ctx context.Context) SourceRailzArrayOutput {
	return o
}

func (o SourceRailzArrayOutput) Index(i pulumi.IntInput) SourceRailzOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceRailz {
		return vs[0].([]*SourceRailz)[vs[1].(int)]
	}).(SourceRailzOutput)
}

type SourceRailzMapOutput struct{ *pulumi.OutputState }

func (SourceRailzMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceRailz)(nil)).Elem()
}

func (o SourceRailzMapOutput) ToSourceRailzMapOutput() SourceRailzMapOutput {
	return o
}

func (o SourceRailzMapOutput) ToSourceRailzMapOutputWithContext(ctx context.Context) SourceRailzMapOutput {
	return o
}

func (o SourceRailzMapOutput) MapIndex(k pulumi.StringInput) SourceRailzOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceRailz {
		return vs[0].(map[string]*SourceRailz)[vs[1].(string)]
	}).(SourceRailzOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRailzInput)(nil)).Elem(), &SourceRailz{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRailzArrayInput)(nil)).Elem(), SourceRailzArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRailzMapInput)(nil)).Elem(), SourceRailzMap{})
	pulumi.RegisterOutputType(SourceRailzOutput{})
	pulumi.RegisterOutputType(SourceRailzArrayOutput{})
	pulumi.RegisterOutputType(SourceRailzMapOutput{})
}
