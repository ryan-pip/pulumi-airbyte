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

// SourceGnews Resource
type SourceGnews struct {
	pulumi.CustomResourceState

	Configuration SourceGnewsConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceGnews registers a new resource with the given unique name, arguments, and options.
func NewSourceGnews(ctx *pulumi.Context,
	name string, args *SourceGnewsArgs, opts ...pulumi.ResourceOption) (*SourceGnews, error) {
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
	var resource SourceGnews
	err := ctx.RegisterResource("airbyte:index/sourceGnews:SourceGnews", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceGnews gets an existing SourceGnews resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceGnews(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceGnewsState, opts ...pulumi.ResourceOption) (*SourceGnews, error) {
	var resource SourceGnews
	err := ctx.ReadResource("airbyte:index/sourceGnews:SourceGnews", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceGnews resources.
type sourceGnewsState struct {
	Configuration *SourceGnewsConfiguration `pulumi:"configuration"`
	Name          *string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceGnewsState struct {
	Configuration SourceGnewsConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceGnewsState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGnewsState)(nil)).Elem()
}

type sourceGnewsArgs struct {
	Configuration SourceGnewsConfiguration `pulumi:"configuration"`
	Name          string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceGnews resource.
type SourceGnewsArgs struct {
	Configuration SourceGnewsConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceGnewsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGnewsArgs)(nil)).Elem()
}

type SourceGnewsInput interface {
	pulumi.Input

	ToSourceGnewsOutput() SourceGnewsOutput
	ToSourceGnewsOutputWithContext(ctx context.Context) SourceGnewsOutput
}

func (*SourceGnews) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGnews)(nil)).Elem()
}

func (i *SourceGnews) ToSourceGnewsOutput() SourceGnewsOutput {
	return i.ToSourceGnewsOutputWithContext(context.Background())
}

func (i *SourceGnews) ToSourceGnewsOutputWithContext(ctx context.Context) SourceGnewsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGnewsOutput)
}

// SourceGnewsArrayInput is an input type that accepts SourceGnewsArray and SourceGnewsArrayOutput values.
// You can construct a concrete instance of `SourceGnewsArrayInput` via:
//
//	SourceGnewsArray{ SourceGnewsArgs{...} }
type SourceGnewsArrayInput interface {
	pulumi.Input

	ToSourceGnewsArrayOutput() SourceGnewsArrayOutput
	ToSourceGnewsArrayOutputWithContext(context.Context) SourceGnewsArrayOutput
}

type SourceGnewsArray []SourceGnewsInput

func (SourceGnewsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGnews)(nil)).Elem()
}

func (i SourceGnewsArray) ToSourceGnewsArrayOutput() SourceGnewsArrayOutput {
	return i.ToSourceGnewsArrayOutputWithContext(context.Background())
}

func (i SourceGnewsArray) ToSourceGnewsArrayOutputWithContext(ctx context.Context) SourceGnewsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGnewsArrayOutput)
}

// SourceGnewsMapInput is an input type that accepts SourceGnewsMap and SourceGnewsMapOutput values.
// You can construct a concrete instance of `SourceGnewsMapInput` via:
//
//	SourceGnewsMap{ "key": SourceGnewsArgs{...} }
type SourceGnewsMapInput interface {
	pulumi.Input

	ToSourceGnewsMapOutput() SourceGnewsMapOutput
	ToSourceGnewsMapOutputWithContext(context.Context) SourceGnewsMapOutput
}

type SourceGnewsMap map[string]SourceGnewsInput

func (SourceGnewsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGnews)(nil)).Elem()
}

func (i SourceGnewsMap) ToSourceGnewsMapOutput() SourceGnewsMapOutput {
	return i.ToSourceGnewsMapOutputWithContext(context.Background())
}

func (i SourceGnewsMap) ToSourceGnewsMapOutputWithContext(ctx context.Context) SourceGnewsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGnewsMapOutput)
}

type SourceGnewsOutput struct{ *pulumi.OutputState }

func (SourceGnewsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGnews)(nil)).Elem()
}

func (o SourceGnewsOutput) ToSourceGnewsOutput() SourceGnewsOutput {
	return o
}

func (o SourceGnewsOutput) ToSourceGnewsOutputWithContext(ctx context.Context) SourceGnewsOutput {
	return o
}

func (o SourceGnewsOutput) Configuration() SourceGnewsConfigurationOutput {
	return o.ApplyT(func(v *SourceGnews) SourceGnewsConfigurationOutput { return v.Configuration }).(SourceGnewsConfigurationOutput)
}

func (o SourceGnewsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGnews) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceGnewsOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGnews) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceGnewsOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGnews) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceGnewsOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGnews) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceGnewsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGnews) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceGnewsArrayOutput struct{ *pulumi.OutputState }

func (SourceGnewsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGnews)(nil)).Elem()
}

func (o SourceGnewsArrayOutput) ToSourceGnewsArrayOutput() SourceGnewsArrayOutput {
	return o
}

func (o SourceGnewsArrayOutput) ToSourceGnewsArrayOutputWithContext(ctx context.Context) SourceGnewsArrayOutput {
	return o
}

func (o SourceGnewsArrayOutput) Index(i pulumi.IntInput) SourceGnewsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceGnews {
		return vs[0].([]*SourceGnews)[vs[1].(int)]
	}).(SourceGnewsOutput)
}

type SourceGnewsMapOutput struct{ *pulumi.OutputState }

func (SourceGnewsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGnews)(nil)).Elem()
}

func (o SourceGnewsMapOutput) ToSourceGnewsMapOutput() SourceGnewsMapOutput {
	return o
}

func (o SourceGnewsMapOutput) ToSourceGnewsMapOutputWithContext(ctx context.Context) SourceGnewsMapOutput {
	return o
}

func (o SourceGnewsMapOutput) MapIndex(k pulumi.StringInput) SourceGnewsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceGnews {
		return vs[0].(map[string]*SourceGnews)[vs[1].(string)]
	}).(SourceGnewsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGnewsInput)(nil)).Elem(), &SourceGnews{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGnewsArrayInput)(nil)).Elem(), SourceGnewsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGnewsMapInput)(nil)).Elem(), SourceGnewsMap{})
	pulumi.RegisterOutputType(SourceGnewsOutput{})
	pulumi.RegisterOutputType(SourceGnewsArrayOutput{})
	pulumi.RegisterOutputType(SourceGnewsMapOutput{})
}
