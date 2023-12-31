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

// SourceAmplitude Resource
type SourceAmplitude struct {
	pulumi.CustomResourceState

	Configuration SourceAmplitudeConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceAmplitude registers a new resource with the given unique name, arguments, and options.
func NewSourceAmplitude(ctx *pulumi.Context,
	name string, args *SourceAmplitudeArgs, opts ...pulumi.ResourceOption) (*SourceAmplitude, error) {
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
	var resource SourceAmplitude
	err := ctx.RegisterResource("airbyte:index/sourceAmplitude:SourceAmplitude", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceAmplitude gets an existing SourceAmplitude resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceAmplitude(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceAmplitudeState, opts ...pulumi.ResourceOption) (*SourceAmplitude, error) {
	var resource SourceAmplitude
	err := ctx.ReadResource("airbyte:index/sourceAmplitude:SourceAmplitude", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceAmplitude resources.
type sourceAmplitudeState struct {
	Configuration *SourceAmplitudeConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceAmplitudeState struct {
	Configuration SourceAmplitudeConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceAmplitudeState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAmplitudeState)(nil)).Elem()
}

type sourceAmplitudeArgs struct {
	Configuration SourceAmplitudeConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceAmplitude resource.
type SourceAmplitudeArgs struct {
	Configuration SourceAmplitudeConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceAmplitudeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAmplitudeArgs)(nil)).Elem()
}

type SourceAmplitudeInput interface {
	pulumi.Input

	ToSourceAmplitudeOutput() SourceAmplitudeOutput
	ToSourceAmplitudeOutputWithContext(ctx context.Context) SourceAmplitudeOutput
}

func (*SourceAmplitude) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAmplitude)(nil)).Elem()
}

func (i *SourceAmplitude) ToSourceAmplitudeOutput() SourceAmplitudeOutput {
	return i.ToSourceAmplitudeOutputWithContext(context.Background())
}

func (i *SourceAmplitude) ToSourceAmplitudeOutputWithContext(ctx context.Context) SourceAmplitudeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAmplitudeOutput)
}

// SourceAmplitudeArrayInput is an input type that accepts SourceAmplitudeArray and SourceAmplitudeArrayOutput values.
// You can construct a concrete instance of `SourceAmplitudeArrayInput` via:
//
//	SourceAmplitudeArray{ SourceAmplitudeArgs{...} }
type SourceAmplitudeArrayInput interface {
	pulumi.Input

	ToSourceAmplitudeArrayOutput() SourceAmplitudeArrayOutput
	ToSourceAmplitudeArrayOutputWithContext(context.Context) SourceAmplitudeArrayOutput
}

type SourceAmplitudeArray []SourceAmplitudeInput

func (SourceAmplitudeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAmplitude)(nil)).Elem()
}

func (i SourceAmplitudeArray) ToSourceAmplitudeArrayOutput() SourceAmplitudeArrayOutput {
	return i.ToSourceAmplitudeArrayOutputWithContext(context.Background())
}

func (i SourceAmplitudeArray) ToSourceAmplitudeArrayOutputWithContext(ctx context.Context) SourceAmplitudeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAmplitudeArrayOutput)
}

// SourceAmplitudeMapInput is an input type that accepts SourceAmplitudeMap and SourceAmplitudeMapOutput values.
// You can construct a concrete instance of `SourceAmplitudeMapInput` via:
//
//	SourceAmplitudeMap{ "key": SourceAmplitudeArgs{...} }
type SourceAmplitudeMapInput interface {
	pulumi.Input

	ToSourceAmplitudeMapOutput() SourceAmplitudeMapOutput
	ToSourceAmplitudeMapOutputWithContext(context.Context) SourceAmplitudeMapOutput
}

type SourceAmplitudeMap map[string]SourceAmplitudeInput

func (SourceAmplitudeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAmplitude)(nil)).Elem()
}

func (i SourceAmplitudeMap) ToSourceAmplitudeMapOutput() SourceAmplitudeMapOutput {
	return i.ToSourceAmplitudeMapOutputWithContext(context.Background())
}

func (i SourceAmplitudeMap) ToSourceAmplitudeMapOutputWithContext(ctx context.Context) SourceAmplitudeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAmplitudeMapOutput)
}

type SourceAmplitudeOutput struct{ *pulumi.OutputState }

func (SourceAmplitudeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAmplitude)(nil)).Elem()
}

func (o SourceAmplitudeOutput) ToSourceAmplitudeOutput() SourceAmplitudeOutput {
	return o
}

func (o SourceAmplitudeOutput) ToSourceAmplitudeOutputWithContext(ctx context.Context) SourceAmplitudeOutput {
	return o
}

func (o SourceAmplitudeOutput) Configuration() SourceAmplitudeConfigurationOutput {
	return o.ApplyT(func(v *SourceAmplitude) SourceAmplitudeConfigurationOutput { return v.Configuration }).(SourceAmplitudeConfigurationOutput)
}

func (o SourceAmplitudeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmplitude) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceAmplitudeOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAmplitude) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceAmplitudeOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmplitude) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceAmplitudeOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmplitude) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceAmplitudeOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmplitude) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceAmplitudeArrayOutput struct{ *pulumi.OutputState }

func (SourceAmplitudeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAmplitude)(nil)).Elem()
}

func (o SourceAmplitudeArrayOutput) ToSourceAmplitudeArrayOutput() SourceAmplitudeArrayOutput {
	return o
}

func (o SourceAmplitudeArrayOutput) ToSourceAmplitudeArrayOutputWithContext(ctx context.Context) SourceAmplitudeArrayOutput {
	return o
}

func (o SourceAmplitudeArrayOutput) Index(i pulumi.IntInput) SourceAmplitudeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceAmplitude {
		return vs[0].([]*SourceAmplitude)[vs[1].(int)]
	}).(SourceAmplitudeOutput)
}

type SourceAmplitudeMapOutput struct{ *pulumi.OutputState }

func (SourceAmplitudeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAmplitude)(nil)).Elem()
}

func (o SourceAmplitudeMapOutput) ToSourceAmplitudeMapOutput() SourceAmplitudeMapOutput {
	return o
}

func (o SourceAmplitudeMapOutput) ToSourceAmplitudeMapOutputWithContext(ctx context.Context) SourceAmplitudeMapOutput {
	return o
}

func (o SourceAmplitudeMapOutput) MapIndex(k pulumi.StringInput) SourceAmplitudeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceAmplitude {
		return vs[0].(map[string]*SourceAmplitude)[vs[1].(string)]
	}).(SourceAmplitudeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAmplitudeInput)(nil)).Elem(), &SourceAmplitude{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAmplitudeArrayInput)(nil)).Elem(), SourceAmplitudeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAmplitudeMapInput)(nil)).Elem(), SourceAmplitudeMap{})
	pulumi.RegisterOutputType(SourceAmplitudeOutput{})
	pulumi.RegisterOutputType(SourceAmplitudeArrayOutput{})
	pulumi.RegisterOutputType(SourceAmplitudeMapOutput{})
}
