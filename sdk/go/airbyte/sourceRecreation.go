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

// SourceRecreation Resource
type SourceRecreation struct {
	pulumi.CustomResourceState

	Configuration SourceRecreationConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                 `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceRecreation registers a new resource with the given unique name, arguments, and options.
func NewSourceRecreation(ctx *pulumi.Context,
	name string, args *SourceRecreationArgs, opts ...pulumi.ResourceOption) (*SourceRecreation, error) {
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
	var resource SourceRecreation
	err := ctx.RegisterResource("airbyte:index/sourceRecreation:SourceRecreation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceRecreation gets an existing SourceRecreation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceRecreation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceRecreationState, opts ...pulumi.ResourceOption) (*SourceRecreation, error) {
	var resource SourceRecreation
	err := ctx.ReadResource("airbyte:index/sourceRecreation:SourceRecreation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceRecreation resources.
type sourceRecreationState struct {
	Configuration *SourceRecreationConfiguration `pulumi:"configuration"`
	Name          *string                        `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceRecreationState struct {
	Configuration SourceRecreationConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceRecreationState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRecreationState)(nil)).Elem()
}

type sourceRecreationArgs struct {
	Configuration SourceRecreationConfiguration `pulumi:"configuration"`
	Name          string                        `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceRecreation resource.
type SourceRecreationArgs struct {
	Configuration SourceRecreationConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceRecreationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRecreationArgs)(nil)).Elem()
}

type SourceRecreationInput interface {
	pulumi.Input

	ToSourceRecreationOutput() SourceRecreationOutput
	ToSourceRecreationOutputWithContext(ctx context.Context) SourceRecreationOutput
}

func (*SourceRecreation) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRecreation)(nil)).Elem()
}

func (i *SourceRecreation) ToSourceRecreationOutput() SourceRecreationOutput {
	return i.ToSourceRecreationOutputWithContext(context.Background())
}

func (i *SourceRecreation) ToSourceRecreationOutputWithContext(ctx context.Context) SourceRecreationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRecreationOutput)
}

// SourceRecreationArrayInput is an input type that accepts SourceRecreationArray and SourceRecreationArrayOutput values.
// You can construct a concrete instance of `SourceRecreationArrayInput` via:
//
//	SourceRecreationArray{ SourceRecreationArgs{...} }
type SourceRecreationArrayInput interface {
	pulumi.Input

	ToSourceRecreationArrayOutput() SourceRecreationArrayOutput
	ToSourceRecreationArrayOutputWithContext(context.Context) SourceRecreationArrayOutput
}

type SourceRecreationArray []SourceRecreationInput

func (SourceRecreationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceRecreation)(nil)).Elem()
}

func (i SourceRecreationArray) ToSourceRecreationArrayOutput() SourceRecreationArrayOutput {
	return i.ToSourceRecreationArrayOutputWithContext(context.Background())
}

func (i SourceRecreationArray) ToSourceRecreationArrayOutputWithContext(ctx context.Context) SourceRecreationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRecreationArrayOutput)
}

// SourceRecreationMapInput is an input type that accepts SourceRecreationMap and SourceRecreationMapOutput values.
// You can construct a concrete instance of `SourceRecreationMapInput` via:
//
//	SourceRecreationMap{ "key": SourceRecreationArgs{...} }
type SourceRecreationMapInput interface {
	pulumi.Input

	ToSourceRecreationMapOutput() SourceRecreationMapOutput
	ToSourceRecreationMapOutputWithContext(context.Context) SourceRecreationMapOutput
}

type SourceRecreationMap map[string]SourceRecreationInput

func (SourceRecreationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceRecreation)(nil)).Elem()
}

func (i SourceRecreationMap) ToSourceRecreationMapOutput() SourceRecreationMapOutput {
	return i.ToSourceRecreationMapOutputWithContext(context.Background())
}

func (i SourceRecreationMap) ToSourceRecreationMapOutputWithContext(ctx context.Context) SourceRecreationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRecreationMapOutput)
}

type SourceRecreationOutput struct{ *pulumi.OutputState }

func (SourceRecreationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRecreation)(nil)).Elem()
}

func (o SourceRecreationOutput) ToSourceRecreationOutput() SourceRecreationOutput {
	return o
}

func (o SourceRecreationOutput) ToSourceRecreationOutputWithContext(ctx context.Context) SourceRecreationOutput {
	return o
}

func (o SourceRecreationOutput) Configuration() SourceRecreationConfigurationOutput {
	return o.ApplyT(func(v *SourceRecreation) SourceRecreationConfigurationOutput { return v.Configuration }).(SourceRecreationConfigurationOutput)
}

func (o SourceRecreationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRecreation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceRecreationOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceRecreation) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceRecreationOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRecreation) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceRecreationOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRecreation) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceRecreationOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRecreation) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceRecreationArrayOutput struct{ *pulumi.OutputState }

func (SourceRecreationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceRecreation)(nil)).Elem()
}

func (o SourceRecreationArrayOutput) ToSourceRecreationArrayOutput() SourceRecreationArrayOutput {
	return o
}

func (o SourceRecreationArrayOutput) ToSourceRecreationArrayOutputWithContext(ctx context.Context) SourceRecreationArrayOutput {
	return o
}

func (o SourceRecreationArrayOutput) Index(i pulumi.IntInput) SourceRecreationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceRecreation {
		return vs[0].([]*SourceRecreation)[vs[1].(int)]
	}).(SourceRecreationOutput)
}

type SourceRecreationMapOutput struct{ *pulumi.OutputState }

func (SourceRecreationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceRecreation)(nil)).Elem()
}

func (o SourceRecreationMapOutput) ToSourceRecreationMapOutput() SourceRecreationMapOutput {
	return o
}

func (o SourceRecreationMapOutput) ToSourceRecreationMapOutputWithContext(ctx context.Context) SourceRecreationMapOutput {
	return o
}

func (o SourceRecreationMapOutput) MapIndex(k pulumi.StringInput) SourceRecreationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceRecreation {
		return vs[0].(map[string]*SourceRecreation)[vs[1].(string)]
	}).(SourceRecreationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRecreationInput)(nil)).Elem(), &SourceRecreation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRecreationArrayInput)(nil)).Elem(), SourceRecreationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRecreationMapInput)(nil)).Elem(), SourceRecreationMap{})
	pulumi.RegisterOutputType(SourceRecreationOutput{})
	pulumi.RegisterOutputType(SourceRecreationArrayOutput{})
	pulumi.RegisterOutputType(SourceRecreationMapOutput{})
}
