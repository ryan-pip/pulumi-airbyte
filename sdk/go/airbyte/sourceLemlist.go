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

// SourceLemlist Resource
type SourceLemlist struct {
	pulumi.CustomResourceState

	Configuration SourceLemlistConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceLemlist registers a new resource with the given unique name, arguments, and options.
func NewSourceLemlist(ctx *pulumi.Context,
	name string, args *SourceLemlistArgs, opts ...pulumi.ResourceOption) (*SourceLemlist, error) {
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
	var resource SourceLemlist
	err := ctx.RegisterResource("airbyte:index/sourceLemlist:SourceLemlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceLemlist gets an existing SourceLemlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceLemlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceLemlistState, opts ...pulumi.ResourceOption) (*SourceLemlist, error) {
	var resource SourceLemlist
	err := ctx.ReadResource("airbyte:index/sourceLemlist:SourceLemlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceLemlist resources.
type sourceLemlistState struct {
	Configuration *SourceLemlistConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceLemlistState struct {
	Configuration SourceLemlistConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceLemlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLemlistState)(nil)).Elem()
}

type sourceLemlistArgs struct {
	Configuration SourceLemlistConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceLemlist resource.
type SourceLemlistArgs struct {
	Configuration SourceLemlistConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceLemlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLemlistArgs)(nil)).Elem()
}

type SourceLemlistInput interface {
	pulumi.Input

	ToSourceLemlistOutput() SourceLemlistOutput
	ToSourceLemlistOutputWithContext(ctx context.Context) SourceLemlistOutput
}

func (*SourceLemlist) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLemlist)(nil)).Elem()
}

func (i *SourceLemlist) ToSourceLemlistOutput() SourceLemlistOutput {
	return i.ToSourceLemlistOutputWithContext(context.Background())
}

func (i *SourceLemlist) ToSourceLemlistOutputWithContext(ctx context.Context) SourceLemlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLemlistOutput)
}

// SourceLemlistArrayInput is an input type that accepts SourceLemlistArray and SourceLemlistArrayOutput values.
// You can construct a concrete instance of `SourceLemlistArrayInput` via:
//
//	SourceLemlistArray{ SourceLemlistArgs{...} }
type SourceLemlistArrayInput interface {
	pulumi.Input

	ToSourceLemlistArrayOutput() SourceLemlistArrayOutput
	ToSourceLemlistArrayOutputWithContext(context.Context) SourceLemlistArrayOutput
}

type SourceLemlistArray []SourceLemlistInput

func (SourceLemlistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceLemlist)(nil)).Elem()
}

func (i SourceLemlistArray) ToSourceLemlistArrayOutput() SourceLemlistArrayOutput {
	return i.ToSourceLemlistArrayOutputWithContext(context.Background())
}

func (i SourceLemlistArray) ToSourceLemlistArrayOutputWithContext(ctx context.Context) SourceLemlistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLemlistArrayOutput)
}

// SourceLemlistMapInput is an input type that accepts SourceLemlistMap and SourceLemlistMapOutput values.
// You can construct a concrete instance of `SourceLemlistMapInput` via:
//
//	SourceLemlistMap{ "key": SourceLemlistArgs{...} }
type SourceLemlistMapInput interface {
	pulumi.Input

	ToSourceLemlistMapOutput() SourceLemlistMapOutput
	ToSourceLemlistMapOutputWithContext(context.Context) SourceLemlistMapOutput
}

type SourceLemlistMap map[string]SourceLemlistInput

func (SourceLemlistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceLemlist)(nil)).Elem()
}

func (i SourceLemlistMap) ToSourceLemlistMapOutput() SourceLemlistMapOutput {
	return i.ToSourceLemlistMapOutputWithContext(context.Background())
}

func (i SourceLemlistMap) ToSourceLemlistMapOutputWithContext(ctx context.Context) SourceLemlistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLemlistMapOutput)
}

type SourceLemlistOutput struct{ *pulumi.OutputState }

func (SourceLemlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLemlist)(nil)).Elem()
}

func (o SourceLemlistOutput) ToSourceLemlistOutput() SourceLemlistOutput {
	return o
}

func (o SourceLemlistOutput) ToSourceLemlistOutputWithContext(ctx context.Context) SourceLemlistOutput {
	return o
}

func (o SourceLemlistOutput) Configuration() SourceLemlistConfigurationOutput {
	return o.ApplyT(func(v *SourceLemlist) SourceLemlistConfigurationOutput { return v.Configuration }).(SourceLemlistConfigurationOutput)
}

func (o SourceLemlistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLemlist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceLemlistOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLemlist) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceLemlistOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLemlist) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceLemlistOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLemlist) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceLemlistOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLemlist) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceLemlistArrayOutput struct{ *pulumi.OutputState }

func (SourceLemlistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceLemlist)(nil)).Elem()
}

func (o SourceLemlistArrayOutput) ToSourceLemlistArrayOutput() SourceLemlistArrayOutput {
	return o
}

func (o SourceLemlistArrayOutput) ToSourceLemlistArrayOutputWithContext(ctx context.Context) SourceLemlistArrayOutput {
	return o
}

func (o SourceLemlistArrayOutput) Index(i pulumi.IntInput) SourceLemlistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceLemlist {
		return vs[0].([]*SourceLemlist)[vs[1].(int)]
	}).(SourceLemlistOutput)
}

type SourceLemlistMapOutput struct{ *pulumi.OutputState }

func (SourceLemlistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceLemlist)(nil)).Elem()
}

func (o SourceLemlistMapOutput) ToSourceLemlistMapOutput() SourceLemlistMapOutput {
	return o
}

func (o SourceLemlistMapOutput) ToSourceLemlistMapOutputWithContext(ctx context.Context) SourceLemlistMapOutput {
	return o
}

func (o SourceLemlistMapOutput) MapIndex(k pulumi.StringInput) SourceLemlistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceLemlist {
		return vs[0].(map[string]*SourceLemlist)[vs[1].(string)]
	}).(SourceLemlistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLemlistInput)(nil)).Elem(), &SourceLemlist{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLemlistArrayInput)(nil)).Elem(), SourceLemlistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLemlistMapInput)(nil)).Elem(), SourceLemlistMap{})
	pulumi.RegisterOutputType(SourceLemlistOutput{})
	pulumi.RegisterOutputType(SourceLemlistArrayOutput{})
	pulumi.RegisterOutputType(SourceLemlistMapOutput{})
}
