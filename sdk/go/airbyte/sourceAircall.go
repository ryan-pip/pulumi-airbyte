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

// SourceAircall Resource
type SourceAircall struct {
	pulumi.CustomResourceState

	Configuration SourceAircallConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceAircall registers a new resource with the given unique name, arguments, and options.
func NewSourceAircall(ctx *pulumi.Context,
	name string, args *SourceAircallArgs, opts ...pulumi.ResourceOption) (*SourceAircall, error) {
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
	var resource SourceAircall
	err := ctx.RegisterResource("airbyte:index/sourceAircall:SourceAircall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceAircall gets an existing SourceAircall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceAircall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceAircallState, opts ...pulumi.ResourceOption) (*SourceAircall, error) {
	var resource SourceAircall
	err := ctx.ReadResource("airbyte:index/sourceAircall:SourceAircall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceAircall resources.
type sourceAircallState struct {
	Configuration *SourceAircallConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceAircallState struct {
	Configuration SourceAircallConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceAircallState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAircallState)(nil)).Elem()
}

type sourceAircallArgs struct {
	Configuration SourceAircallConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceAircall resource.
type SourceAircallArgs struct {
	Configuration SourceAircallConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceAircallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAircallArgs)(nil)).Elem()
}

type SourceAircallInput interface {
	pulumi.Input

	ToSourceAircallOutput() SourceAircallOutput
	ToSourceAircallOutputWithContext(ctx context.Context) SourceAircallOutput
}

func (*SourceAircall) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAircall)(nil)).Elem()
}

func (i *SourceAircall) ToSourceAircallOutput() SourceAircallOutput {
	return i.ToSourceAircallOutputWithContext(context.Background())
}

func (i *SourceAircall) ToSourceAircallOutputWithContext(ctx context.Context) SourceAircallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAircallOutput)
}

// SourceAircallArrayInput is an input type that accepts SourceAircallArray and SourceAircallArrayOutput values.
// You can construct a concrete instance of `SourceAircallArrayInput` via:
//
//	SourceAircallArray{ SourceAircallArgs{...} }
type SourceAircallArrayInput interface {
	pulumi.Input

	ToSourceAircallArrayOutput() SourceAircallArrayOutput
	ToSourceAircallArrayOutputWithContext(context.Context) SourceAircallArrayOutput
}

type SourceAircallArray []SourceAircallInput

func (SourceAircallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAircall)(nil)).Elem()
}

func (i SourceAircallArray) ToSourceAircallArrayOutput() SourceAircallArrayOutput {
	return i.ToSourceAircallArrayOutputWithContext(context.Background())
}

func (i SourceAircallArray) ToSourceAircallArrayOutputWithContext(ctx context.Context) SourceAircallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAircallArrayOutput)
}

// SourceAircallMapInput is an input type that accepts SourceAircallMap and SourceAircallMapOutput values.
// You can construct a concrete instance of `SourceAircallMapInput` via:
//
//	SourceAircallMap{ "key": SourceAircallArgs{...} }
type SourceAircallMapInput interface {
	pulumi.Input

	ToSourceAircallMapOutput() SourceAircallMapOutput
	ToSourceAircallMapOutputWithContext(context.Context) SourceAircallMapOutput
}

type SourceAircallMap map[string]SourceAircallInput

func (SourceAircallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAircall)(nil)).Elem()
}

func (i SourceAircallMap) ToSourceAircallMapOutput() SourceAircallMapOutput {
	return i.ToSourceAircallMapOutputWithContext(context.Background())
}

func (i SourceAircallMap) ToSourceAircallMapOutputWithContext(ctx context.Context) SourceAircallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAircallMapOutput)
}

type SourceAircallOutput struct{ *pulumi.OutputState }

func (SourceAircallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAircall)(nil)).Elem()
}

func (o SourceAircallOutput) ToSourceAircallOutput() SourceAircallOutput {
	return o
}

func (o SourceAircallOutput) ToSourceAircallOutputWithContext(ctx context.Context) SourceAircallOutput {
	return o
}

func (o SourceAircallOutput) Configuration() SourceAircallConfigurationOutput {
	return o.ApplyT(func(v *SourceAircall) SourceAircallConfigurationOutput { return v.Configuration }).(SourceAircallConfigurationOutput)
}

func (o SourceAircallOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAircall) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceAircallOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAircall) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceAircallOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAircall) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceAircallOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAircall) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceAircallOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAircall) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceAircallArrayOutput struct{ *pulumi.OutputState }

func (SourceAircallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAircall)(nil)).Elem()
}

func (o SourceAircallArrayOutput) ToSourceAircallArrayOutput() SourceAircallArrayOutput {
	return o
}

func (o SourceAircallArrayOutput) ToSourceAircallArrayOutputWithContext(ctx context.Context) SourceAircallArrayOutput {
	return o
}

func (o SourceAircallArrayOutput) Index(i pulumi.IntInput) SourceAircallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceAircall {
		return vs[0].([]*SourceAircall)[vs[1].(int)]
	}).(SourceAircallOutput)
}

type SourceAircallMapOutput struct{ *pulumi.OutputState }

func (SourceAircallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAircall)(nil)).Elem()
}

func (o SourceAircallMapOutput) ToSourceAircallMapOutput() SourceAircallMapOutput {
	return o
}

func (o SourceAircallMapOutput) ToSourceAircallMapOutputWithContext(ctx context.Context) SourceAircallMapOutput {
	return o
}

func (o SourceAircallMapOutput) MapIndex(k pulumi.StringInput) SourceAircallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceAircall {
		return vs[0].(map[string]*SourceAircall)[vs[1].(string)]
	}).(SourceAircallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAircallInput)(nil)).Elem(), &SourceAircall{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAircallArrayInput)(nil)).Elem(), SourceAircallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAircallMapInput)(nil)).Elem(), SourceAircallMap{})
	pulumi.RegisterOutputType(SourceAircallOutput{})
	pulumi.RegisterOutputType(SourceAircallArrayOutput{})
	pulumi.RegisterOutputType(SourceAircallMapOutput{})
}
