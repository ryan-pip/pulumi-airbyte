// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourcePipedrive Resource
type SourcePipedrive struct {
	pulumi.CustomResourceState

	Configuration SourcePipedriveConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourcePipedrive registers a new resource with the given unique name, arguments, and options.
func NewSourcePipedrive(ctx *pulumi.Context,
	name string, args *SourcePipedriveArgs, opts ...pulumi.ResourceOption) (*SourcePipedrive, error) {
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
	var resource SourcePipedrive
	err := ctx.RegisterResource("airbyte:index/sourcePipedrive:SourcePipedrive", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourcePipedrive gets an existing SourcePipedrive resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourcePipedrive(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourcePipedriveState, opts ...pulumi.ResourceOption) (*SourcePipedrive, error) {
	var resource SourcePipedrive
	err := ctx.ReadResource("airbyte:index/sourcePipedrive:SourcePipedrive", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourcePipedrive resources.
type sourcePipedriveState struct {
	Configuration *SourcePipedriveConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourcePipedriveState struct {
	Configuration SourcePipedriveConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourcePipedriveState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePipedriveState)(nil)).Elem()
}

type sourcePipedriveArgs struct {
	Configuration SourcePipedriveConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourcePipedrive resource.
type SourcePipedriveArgs struct {
	Configuration SourcePipedriveConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourcePipedriveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePipedriveArgs)(nil)).Elem()
}

type SourcePipedriveInput interface {
	pulumi.Input

	ToSourcePipedriveOutput() SourcePipedriveOutput
	ToSourcePipedriveOutputWithContext(ctx context.Context) SourcePipedriveOutput
}

func (*SourcePipedrive) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePipedrive)(nil)).Elem()
}

func (i *SourcePipedrive) ToSourcePipedriveOutput() SourcePipedriveOutput {
	return i.ToSourcePipedriveOutputWithContext(context.Background())
}

func (i *SourcePipedrive) ToSourcePipedriveOutputWithContext(ctx context.Context) SourcePipedriveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePipedriveOutput)
}

// SourcePipedriveArrayInput is an input type that accepts SourcePipedriveArray and SourcePipedriveArrayOutput values.
// You can construct a concrete instance of `SourcePipedriveArrayInput` via:
//
//	SourcePipedriveArray{ SourcePipedriveArgs{...} }
type SourcePipedriveArrayInput interface {
	pulumi.Input

	ToSourcePipedriveArrayOutput() SourcePipedriveArrayOutput
	ToSourcePipedriveArrayOutputWithContext(context.Context) SourcePipedriveArrayOutput
}

type SourcePipedriveArray []SourcePipedriveInput

func (SourcePipedriveArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourcePipedrive)(nil)).Elem()
}

func (i SourcePipedriveArray) ToSourcePipedriveArrayOutput() SourcePipedriveArrayOutput {
	return i.ToSourcePipedriveArrayOutputWithContext(context.Background())
}

func (i SourcePipedriveArray) ToSourcePipedriveArrayOutputWithContext(ctx context.Context) SourcePipedriveArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePipedriveArrayOutput)
}

// SourcePipedriveMapInput is an input type that accepts SourcePipedriveMap and SourcePipedriveMapOutput values.
// You can construct a concrete instance of `SourcePipedriveMapInput` via:
//
//	SourcePipedriveMap{ "key": SourcePipedriveArgs{...} }
type SourcePipedriveMapInput interface {
	pulumi.Input

	ToSourcePipedriveMapOutput() SourcePipedriveMapOutput
	ToSourcePipedriveMapOutputWithContext(context.Context) SourcePipedriveMapOutput
}

type SourcePipedriveMap map[string]SourcePipedriveInput

func (SourcePipedriveMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourcePipedrive)(nil)).Elem()
}

func (i SourcePipedriveMap) ToSourcePipedriveMapOutput() SourcePipedriveMapOutput {
	return i.ToSourcePipedriveMapOutputWithContext(context.Background())
}

func (i SourcePipedriveMap) ToSourcePipedriveMapOutputWithContext(ctx context.Context) SourcePipedriveMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePipedriveMapOutput)
}

type SourcePipedriveOutput struct{ *pulumi.OutputState }

func (SourcePipedriveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePipedrive)(nil)).Elem()
}

func (o SourcePipedriveOutput) ToSourcePipedriveOutput() SourcePipedriveOutput {
	return o
}

func (o SourcePipedriveOutput) ToSourcePipedriveOutputWithContext(ctx context.Context) SourcePipedriveOutput {
	return o
}

func (o SourcePipedriveOutput) Configuration() SourcePipedriveConfigurationOutput {
	return o.ApplyT(func(v *SourcePipedrive) SourcePipedriveConfigurationOutput { return v.Configuration }).(SourcePipedriveConfigurationOutput)
}

func (o SourcePipedriveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePipedrive) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourcePipedriveOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePipedrive) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourcePipedriveOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePipedrive) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourcePipedriveOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePipedrive) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourcePipedriveOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePipedrive) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourcePipedriveArrayOutput struct{ *pulumi.OutputState }

func (SourcePipedriveArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourcePipedrive)(nil)).Elem()
}

func (o SourcePipedriveArrayOutput) ToSourcePipedriveArrayOutput() SourcePipedriveArrayOutput {
	return o
}

func (o SourcePipedriveArrayOutput) ToSourcePipedriveArrayOutputWithContext(ctx context.Context) SourcePipedriveArrayOutput {
	return o
}

func (o SourcePipedriveArrayOutput) Index(i pulumi.IntInput) SourcePipedriveOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourcePipedrive {
		return vs[0].([]*SourcePipedrive)[vs[1].(int)]
	}).(SourcePipedriveOutput)
}

type SourcePipedriveMapOutput struct{ *pulumi.OutputState }

func (SourcePipedriveMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourcePipedrive)(nil)).Elem()
}

func (o SourcePipedriveMapOutput) ToSourcePipedriveMapOutput() SourcePipedriveMapOutput {
	return o
}

func (o SourcePipedriveMapOutput) ToSourcePipedriveMapOutputWithContext(ctx context.Context) SourcePipedriveMapOutput {
	return o
}

func (o SourcePipedriveMapOutput) MapIndex(k pulumi.StringInput) SourcePipedriveOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourcePipedrive {
		return vs[0].(map[string]*SourcePipedrive)[vs[1].(string)]
	}).(SourcePipedriveOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePipedriveInput)(nil)).Elem(), &SourcePipedrive{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePipedriveArrayInput)(nil)).Elem(), SourcePipedriveArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePipedriveMapInput)(nil)).Elem(), SourcePipedriveMap{})
	pulumi.RegisterOutputType(SourcePipedriveOutput{})
	pulumi.RegisterOutputType(SourcePipedriveArrayOutput{})
	pulumi.RegisterOutputType(SourcePipedriveMapOutput{})
}
