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

// SourceMonday Resource
type SourceMonday struct {
	pulumi.CustomResourceState

	Configuration SourceMondayConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceMonday registers a new resource with the given unique name, arguments, and options.
func NewSourceMonday(ctx *pulumi.Context,
	name string, args *SourceMondayArgs, opts ...pulumi.ResourceOption) (*SourceMonday, error) {
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
	var resource SourceMonday
	err := ctx.RegisterResource("airbyte:index/sourceMonday:SourceMonday", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceMonday gets an existing SourceMonday resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceMonday(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceMondayState, opts ...pulumi.ResourceOption) (*SourceMonday, error) {
	var resource SourceMonday
	err := ctx.ReadResource("airbyte:index/sourceMonday:SourceMonday", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceMonday resources.
type sourceMondayState struct {
	Configuration *SourceMondayConfiguration `pulumi:"configuration"`
	Name          *string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceMondayState struct {
	Configuration SourceMondayConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceMondayState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMondayState)(nil)).Elem()
}

type sourceMondayArgs struct {
	Configuration SourceMondayConfiguration `pulumi:"configuration"`
	Name          string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceMonday resource.
type SourceMondayArgs struct {
	Configuration SourceMondayConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceMondayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMondayArgs)(nil)).Elem()
}

type SourceMondayInput interface {
	pulumi.Input

	ToSourceMondayOutput() SourceMondayOutput
	ToSourceMondayOutputWithContext(ctx context.Context) SourceMondayOutput
}

func (*SourceMonday) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMonday)(nil)).Elem()
}

func (i *SourceMonday) ToSourceMondayOutput() SourceMondayOutput {
	return i.ToSourceMondayOutputWithContext(context.Background())
}

func (i *SourceMonday) ToSourceMondayOutputWithContext(ctx context.Context) SourceMondayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMondayOutput)
}

// SourceMondayArrayInput is an input type that accepts SourceMondayArray and SourceMondayArrayOutput values.
// You can construct a concrete instance of `SourceMondayArrayInput` via:
//
//	SourceMondayArray{ SourceMondayArgs{...} }
type SourceMondayArrayInput interface {
	pulumi.Input

	ToSourceMondayArrayOutput() SourceMondayArrayOutput
	ToSourceMondayArrayOutputWithContext(context.Context) SourceMondayArrayOutput
}

type SourceMondayArray []SourceMondayInput

func (SourceMondayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceMonday)(nil)).Elem()
}

func (i SourceMondayArray) ToSourceMondayArrayOutput() SourceMondayArrayOutput {
	return i.ToSourceMondayArrayOutputWithContext(context.Background())
}

func (i SourceMondayArray) ToSourceMondayArrayOutputWithContext(ctx context.Context) SourceMondayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMondayArrayOutput)
}

// SourceMondayMapInput is an input type that accepts SourceMondayMap and SourceMondayMapOutput values.
// You can construct a concrete instance of `SourceMondayMapInput` via:
//
//	SourceMondayMap{ "key": SourceMondayArgs{...} }
type SourceMondayMapInput interface {
	pulumi.Input

	ToSourceMondayMapOutput() SourceMondayMapOutput
	ToSourceMondayMapOutputWithContext(context.Context) SourceMondayMapOutput
}

type SourceMondayMap map[string]SourceMondayInput

func (SourceMondayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceMonday)(nil)).Elem()
}

func (i SourceMondayMap) ToSourceMondayMapOutput() SourceMondayMapOutput {
	return i.ToSourceMondayMapOutputWithContext(context.Background())
}

func (i SourceMondayMap) ToSourceMondayMapOutputWithContext(ctx context.Context) SourceMondayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMondayMapOutput)
}

type SourceMondayOutput struct{ *pulumi.OutputState }

func (SourceMondayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMonday)(nil)).Elem()
}

func (o SourceMondayOutput) ToSourceMondayOutput() SourceMondayOutput {
	return o
}

func (o SourceMondayOutput) ToSourceMondayOutputWithContext(ctx context.Context) SourceMondayOutput {
	return o
}

func (o SourceMondayOutput) Configuration() SourceMondayConfigurationOutput {
	return o.ApplyT(func(v *SourceMonday) SourceMondayConfigurationOutput { return v.Configuration }).(SourceMondayConfigurationOutput)
}

func (o SourceMondayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMonday) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceMondayOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceMonday) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceMondayOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMonday) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceMondayOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMonday) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceMondayOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMonday) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceMondayArrayOutput struct{ *pulumi.OutputState }

func (SourceMondayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceMonday)(nil)).Elem()
}

func (o SourceMondayArrayOutput) ToSourceMondayArrayOutput() SourceMondayArrayOutput {
	return o
}

func (o SourceMondayArrayOutput) ToSourceMondayArrayOutputWithContext(ctx context.Context) SourceMondayArrayOutput {
	return o
}

func (o SourceMondayArrayOutput) Index(i pulumi.IntInput) SourceMondayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceMonday {
		return vs[0].([]*SourceMonday)[vs[1].(int)]
	}).(SourceMondayOutput)
}

type SourceMondayMapOutput struct{ *pulumi.OutputState }

func (SourceMondayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceMonday)(nil)).Elem()
}

func (o SourceMondayMapOutput) ToSourceMondayMapOutput() SourceMondayMapOutput {
	return o
}

func (o SourceMondayMapOutput) ToSourceMondayMapOutputWithContext(ctx context.Context) SourceMondayMapOutput {
	return o
}

func (o SourceMondayMapOutput) MapIndex(k pulumi.StringInput) SourceMondayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceMonday {
		return vs[0].(map[string]*SourceMonday)[vs[1].(string)]
	}).(SourceMondayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMondayInput)(nil)).Elem(), &SourceMonday{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMondayArrayInput)(nil)).Elem(), SourceMondayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMondayMapInput)(nil)).Elem(), SourceMondayMap{})
	pulumi.RegisterOutputType(SourceMondayOutput{})
	pulumi.RegisterOutputType(SourceMondayArrayOutput{})
	pulumi.RegisterOutputType(SourceMondayMapOutput{})
}
