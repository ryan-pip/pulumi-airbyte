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

// SourceSentry Resource
type SourceSentry struct {
	pulumi.CustomResourceState

	Configuration SourceSentryConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceSentry registers a new resource with the given unique name, arguments, and options.
func NewSourceSentry(ctx *pulumi.Context,
	name string, args *SourceSentryArgs, opts ...pulumi.ResourceOption) (*SourceSentry, error) {
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
	var resource SourceSentry
	err := ctx.RegisterResource("airbyte:index/sourceSentry:SourceSentry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSentry gets an existing SourceSentry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSentry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSentryState, opts ...pulumi.ResourceOption) (*SourceSentry, error) {
	var resource SourceSentry
	err := ctx.ReadResource("airbyte:index/sourceSentry:SourceSentry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSentry resources.
type sourceSentryState struct {
	Configuration *SourceSentryConfiguration `pulumi:"configuration"`
	Name          *string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceSentryState struct {
	Configuration SourceSentryConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceSentryState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSentryState)(nil)).Elem()
}

type sourceSentryArgs struct {
	Configuration SourceSentryConfiguration `pulumi:"configuration"`
	Name          string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSentry resource.
type SourceSentryArgs struct {
	Configuration SourceSentryConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSentryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSentryArgs)(nil)).Elem()
}

type SourceSentryInput interface {
	pulumi.Input

	ToSourceSentryOutput() SourceSentryOutput
	ToSourceSentryOutputWithContext(ctx context.Context) SourceSentryOutput
}

func (*SourceSentry) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSentry)(nil)).Elem()
}

func (i *SourceSentry) ToSourceSentryOutput() SourceSentryOutput {
	return i.ToSourceSentryOutputWithContext(context.Background())
}

func (i *SourceSentry) ToSourceSentryOutputWithContext(ctx context.Context) SourceSentryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSentryOutput)
}

// SourceSentryArrayInput is an input type that accepts SourceSentryArray and SourceSentryArrayOutput values.
// You can construct a concrete instance of `SourceSentryArrayInput` via:
//
//	SourceSentryArray{ SourceSentryArgs{...} }
type SourceSentryArrayInput interface {
	pulumi.Input

	ToSourceSentryArrayOutput() SourceSentryArrayOutput
	ToSourceSentryArrayOutputWithContext(context.Context) SourceSentryArrayOutput
}

type SourceSentryArray []SourceSentryInput

func (SourceSentryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSentry)(nil)).Elem()
}

func (i SourceSentryArray) ToSourceSentryArrayOutput() SourceSentryArrayOutput {
	return i.ToSourceSentryArrayOutputWithContext(context.Background())
}

func (i SourceSentryArray) ToSourceSentryArrayOutputWithContext(ctx context.Context) SourceSentryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSentryArrayOutput)
}

// SourceSentryMapInput is an input type that accepts SourceSentryMap and SourceSentryMapOutput values.
// You can construct a concrete instance of `SourceSentryMapInput` via:
//
//	SourceSentryMap{ "key": SourceSentryArgs{...} }
type SourceSentryMapInput interface {
	pulumi.Input

	ToSourceSentryMapOutput() SourceSentryMapOutput
	ToSourceSentryMapOutputWithContext(context.Context) SourceSentryMapOutput
}

type SourceSentryMap map[string]SourceSentryInput

func (SourceSentryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSentry)(nil)).Elem()
}

func (i SourceSentryMap) ToSourceSentryMapOutput() SourceSentryMapOutput {
	return i.ToSourceSentryMapOutputWithContext(context.Background())
}

func (i SourceSentryMap) ToSourceSentryMapOutputWithContext(ctx context.Context) SourceSentryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSentryMapOutput)
}

type SourceSentryOutput struct{ *pulumi.OutputState }

func (SourceSentryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSentry)(nil)).Elem()
}

func (o SourceSentryOutput) ToSourceSentryOutput() SourceSentryOutput {
	return o
}

func (o SourceSentryOutput) ToSourceSentryOutputWithContext(ctx context.Context) SourceSentryOutput {
	return o
}

func (o SourceSentryOutput) Configuration() SourceSentryConfigurationOutput {
	return o.ApplyT(func(v *SourceSentry) SourceSentryConfigurationOutput { return v.Configuration }).(SourceSentryConfigurationOutput)
}

func (o SourceSentryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSentry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceSentryOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSentry) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSentryOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSentry) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSentryOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSentry) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSentryOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSentry) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceSentryArrayOutput struct{ *pulumi.OutputState }

func (SourceSentryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSentry)(nil)).Elem()
}

func (o SourceSentryArrayOutput) ToSourceSentryArrayOutput() SourceSentryArrayOutput {
	return o
}

func (o SourceSentryArrayOutput) ToSourceSentryArrayOutputWithContext(ctx context.Context) SourceSentryArrayOutput {
	return o
}

func (o SourceSentryArrayOutput) Index(i pulumi.IntInput) SourceSentryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceSentry {
		return vs[0].([]*SourceSentry)[vs[1].(int)]
	}).(SourceSentryOutput)
}

type SourceSentryMapOutput struct{ *pulumi.OutputState }

func (SourceSentryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSentry)(nil)).Elem()
}

func (o SourceSentryMapOutput) ToSourceSentryMapOutput() SourceSentryMapOutput {
	return o
}

func (o SourceSentryMapOutput) ToSourceSentryMapOutputWithContext(ctx context.Context) SourceSentryMapOutput {
	return o
}

func (o SourceSentryMapOutput) MapIndex(k pulumi.StringInput) SourceSentryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceSentry {
		return vs[0].(map[string]*SourceSentry)[vs[1].(string)]
	}).(SourceSentryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSentryInput)(nil)).Elem(), &SourceSentry{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSentryArrayInput)(nil)).Elem(), SourceSentryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSentryMapInput)(nil)).Elem(), SourceSentryMap{})
	pulumi.RegisterOutputType(SourceSentryOutput{})
	pulumi.RegisterOutputType(SourceSentryArrayOutput{})
	pulumi.RegisterOutputType(SourceSentryMapOutput{})
}
