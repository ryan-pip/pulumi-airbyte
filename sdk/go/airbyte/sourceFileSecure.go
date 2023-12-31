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

// SourceFileSecure Resource
type SourceFileSecure struct {
	pulumi.CustomResourceState

	Configuration SourceFileSecureConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                 `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceFileSecure registers a new resource with the given unique name, arguments, and options.
func NewSourceFileSecure(ctx *pulumi.Context,
	name string, args *SourceFileSecureArgs, opts ...pulumi.ResourceOption) (*SourceFileSecure, error) {
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
	var resource SourceFileSecure
	err := ctx.RegisterResource("airbyte:index/sourceFileSecure:SourceFileSecure", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceFileSecure gets an existing SourceFileSecure resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceFileSecure(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceFileSecureState, opts ...pulumi.ResourceOption) (*SourceFileSecure, error) {
	var resource SourceFileSecure
	err := ctx.ReadResource("airbyte:index/sourceFileSecure:SourceFileSecure", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceFileSecure resources.
type sourceFileSecureState struct {
	Configuration *SourceFileSecureConfiguration `pulumi:"configuration"`
	Name          *string                        `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceFileSecureState struct {
	Configuration SourceFileSecureConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceFileSecureState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFileSecureState)(nil)).Elem()
}

type sourceFileSecureArgs struct {
	Configuration SourceFileSecureConfiguration `pulumi:"configuration"`
	Name          string                        `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceFileSecure resource.
type SourceFileSecureArgs struct {
	Configuration SourceFileSecureConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceFileSecureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFileSecureArgs)(nil)).Elem()
}

type SourceFileSecureInput interface {
	pulumi.Input

	ToSourceFileSecureOutput() SourceFileSecureOutput
	ToSourceFileSecureOutputWithContext(ctx context.Context) SourceFileSecureOutput
}

func (*SourceFileSecure) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFileSecure)(nil)).Elem()
}

func (i *SourceFileSecure) ToSourceFileSecureOutput() SourceFileSecureOutput {
	return i.ToSourceFileSecureOutputWithContext(context.Background())
}

func (i *SourceFileSecure) ToSourceFileSecureOutputWithContext(ctx context.Context) SourceFileSecureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFileSecureOutput)
}

// SourceFileSecureArrayInput is an input type that accepts SourceFileSecureArray and SourceFileSecureArrayOutput values.
// You can construct a concrete instance of `SourceFileSecureArrayInput` via:
//
//	SourceFileSecureArray{ SourceFileSecureArgs{...} }
type SourceFileSecureArrayInput interface {
	pulumi.Input

	ToSourceFileSecureArrayOutput() SourceFileSecureArrayOutput
	ToSourceFileSecureArrayOutputWithContext(context.Context) SourceFileSecureArrayOutput
}

type SourceFileSecureArray []SourceFileSecureInput

func (SourceFileSecureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceFileSecure)(nil)).Elem()
}

func (i SourceFileSecureArray) ToSourceFileSecureArrayOutput() SourceFileSecureArrayOutput {
	return i.ToSourceFileSecureArrayOutputWithContext(context.Background())
}

func (i SourceFileSecureArray) ToSourceFileSecureArrayOutputWithContext(ctx context.Context) SourceFileSecureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFileSecureArrayOutput)
}

// SourceFileSecureMapInput is an input type that accepts SourceFileSecureMap and SourceFileSecureMapOutput values.
// You can construct a concrete instance of `SourceFileSecureMapInput` via:
//
//	SourceFileSecureMap{ "key": SourceFileSecureArgs{...} }
type SourceFileSecureMapInput interface {
	pulumi.Input

	ToSourceFileSecureMapOutput() SourceFileSecureMapOutput
	ToSourceFileSecureMapOutputWithContext(context.Context) SourceFileSecureMapOutput
}

type SourceFileSecureMap map[string]SourceFileSecureInput

func (SourceFileSecureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceFileSecure)(nil)).Elem()
}

func (i SourceFileSecureMap) ToSourceFileSecureMapOutput() SourceFileSecureMapOutput {
	return i.ToSourceFileSecureMapOutputWithContext(context.Background())
}

func (i SourceFileSecureMap) ToSourceFileSecureMapOutputWithContext(ctx context.Context) SourceFileSecureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFileSecureMapOutput)
}

type SourceFileSecureOutput struct{ *pulumi.OutputState }

func (SourceFileSecureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFileSecure)(nil)).Elem()
}

func (o SourceFileSecureOutput) ToSourceFileSecureOutput() SourceFileSecureOutput {
	return o
}

func (o SourceFileSecureOutput) ToSourceFileSecureOutputWithContext(ctx context.Context) SourceFileSecureOutput {
	return o
}

func (o SourceFileSecureOutput) Configuration() SourceFileSecureConfigurationOutput {
	return o.ApplyT(func(v *SourceFileSecure) SourceFileSecureConfigurationOutput { return v.Configuration }).(SourceFileSecureConfigurationOutput)
}

func (o SourceFileSecureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFileSecure) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceFileSecureOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceFileSecure) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceFileSecureOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFileSecure) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceFileSecureOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFileSecure) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceFileSecureOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFileSecure) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceFileSecureArrayOutput struct{ *pulumi.OutputState }

func (SourceFileSecureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceFileSecure)(nil)).Elem()
}

func (o SourceFileSecureArrayOutput) ToSourceFileSecureArrayOutput() SourceFileSecureArrayOutput {
	return o
}

func (o SourceFileSecureArrayOutput) ToSourceFileSecureArrayOutputWithContext(ctx context.Context) SourceFileSecureArrayOutput {
	return o
}

func (o SourceFileSecureArrayOutput) Index(i pulumi.IntInput) SourceFileSecureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceFileSecure {
		return vs[0].([]*SourceFileSecure)[vs[1].(int)]
	}).(SourceFileSecureOutput)
}

type SourceFileSecureMapOutput struct{ *pulumi.OutputState }

func (SourceFileSecureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceFileSecure)(nil)).Elem()
}

func (o SourceFileSecureMapOutput) ToSourceFileSecureMapOutput() SourceFileSecureMapOutput {
	return o
}

func (o SourceFileSecureMapOutput) ToSourceFileSecureMapOutputWithContext(ctx context.Context) SourceFileSecureMapOutput {
	return o
}

func (o SourceFileSecureMapOutput) MapIndex(k pulumi.StringInput) SourceFileSecureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceFileSecure {
		return vs[0].(map[string]*SourceFileSecure)[vs[1].(string)]
	}).(SourceFileSecureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFileSecureInput)(nil)).Elem(), &SourceFileSecure{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFileSecureArrayInput)(nil)).Elem(), SourceFileSecureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFileSecureMapInput)(nil)).Elem(), SourceFileSecureMap{})
	pulumi.RegisterOutputType(SourceFileSecureOutput{})
	pulumi.RegisterOutputType(SourceFileSecureArrayOutput{})
	pulumi.RegisterOutputType(SourceFileSecureMapOutput{})
}
