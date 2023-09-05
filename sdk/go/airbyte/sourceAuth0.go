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

// SourceAuth0 Resource
type SourceAuth0 struct {
	pulumi.CustomResourceState

	Configuration SourceAuth0ConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceAuth0 registers a new resource with the given unique name, arguments, and options.
func NewSourceAuth0(ctx *pulumi.Context,
	name string, args *SourceAuth0Args, opts ...pulumi.ResourceOption) (*SourceAuth0, error) {
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
	var resource SourceAuth0
	err := ctx.RegisterResource("airbyte:index/sourceAuth0:SourceAuth0", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceAuth0 gets an existing SourceAuth0 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceAuth0(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceAuth0State, opts ...pulumi.ResourceOption) (*SourceAuth0, error) {
	var resource SourceAuth0
	err := ctx.ReadResource("airbyte:index/sourceAuth0:SourceAuth0", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceAuth0 resources.
type sourceAuth0State struct {
	Configuration *SourceAuth0Configuration `pulumi:"configuration"`
	Name          *string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceAuth0State struct {
	Configuration SourceAuth0ConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceAuth0State) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAuth0State)(nil)).Elem()
}

type sourceAuth0Args struct {
	Configuration SourceAuth0Configuration `pulumi:"configuration"`
	Name          string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceAuth0 resource.
type SourceAuth0Args struct {
	Configuration SourceAuth0ConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceAuth0Args) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAuth0Args)(nil)).Elem()
}

type SourceAuth0Input interface {
	pulumi.Input

	ToSourceAuth0Output() SourceAuth0Output
	ToSourceAuth0OutputWithContext(ctx context.Context) SourceAuth0Output
}

func (*SourceAuth0) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAuth0)(nil)).Elem()
}

func (i *SourceAuth0) ToSourceAuth0Output() SourceAuth0Output {
	return i.ToSourceAuth0OutputWithContext(context.Background())
}

func (i *SourceAuth0) ToSourceAuth0OutputWithContext(ctx context.Context) SourceAuth0Output {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAuth0Output)
}

// SourceAuth0ArrayInput is an input type that accepts SourceAuth0Array and SourceAuth0ArrayOutput values.
// You can construct a concrete instance of `SourceAuth0ArrayInput` via:
//
//	SourceAuth0Array{ SourceAuth0Args{...} }
type SourceAuth0ArrayInput interface {
	pulumi.Input

	ToSourceAuth0ArrayOutput() SourceAuth0ArrayOutput
	ToSourceAuth0ArrayOutputWithContext(context.Context) SourceAuth0ArrayOutput
}

type SourceAuth0Array []SourceAuth0Input

func (SourceAuth0Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAuth0)(nil)).Elem()
}

func (i SourceAuth0Array) ToSourceAuth0ArrayOutput() SourceAuth0ArrayOutput {
	return i.ToSourceAuth0ArrayOutputWithContext(context.Background())
}

func (i SourceAuth0Array) ToSourceAuth0ArrayOutputWithContext(ctx context.Context) SourceAuth0ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAuth0ArrayOutput)
}

// SourceAuth0MapInput is an input type that accepts SourceAuth0Map and SourceAuth0MapOutput values.
// You can construct a concrete instance of `SourceAuth0MapInput` via:
//
//	SourceAuth0Map{ "key": SourceAuth0Args{...} }
type SourceAuth0MapInput interface {
	pulumi.Input

	ToSourceAuth0MapOutput() SourceAuth0MapOutput
	ToSourceAuth0MapOutputWithContext(context.Context) SourceAuth0MapOutput
}

type SourceAuth0Map map[string]SourceAuth0Input

func (SourceAuth0Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAuth0)(nil)).Elem()
}

func (i SourceAuth0Map) ToSourceAuth0MapOutput() SourceAuth0MapOutput {
	return i.ToSourceAuth0MapOutputWithContext(context.Background())
}

func (i SourceAuth0Map) ToSourceAuth0MapOutputWithContext(ctx context.Context) SourceAuth0MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAuth0MapOutput)
}

type SourceAuth0Output struct{ *pulumi.OutputState }

func (SourceAuth0Output) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAuth0)(nil)).Elem()
}

func (o SourceAuth0Output) ToSourceAuth0Output() SourceAuth0Output {
	return o
}

func (o SourceAuth0Output) ToSourceAuth0OutputWithContext(ctx context.Context) SourceAuth0Output {
	return o
}

func (o SourceAuth0Output) Configuration() SourceAuth0ConfigurationOutput {
	return o.ApplyT(func(v *SourceAuth0) SourceAuth0ConfigurationOutput { return v.Configuration }).(SourceAuth0ConfigurationOutput)
}

func (o SourceAuth0Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAuth0) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceAuth0Output) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAuth0) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceAuth0Output) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAuth0) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceAuth0Output) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAuth0) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceAuth0Output) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAuth0) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceAuth0ArrayOutput struct{ *pulumi.OutputState }

func (SourceAuth0ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAuth0)(nil)).Elem()
}

func (o SourceAuth0ArrayOutput) ToSourceAuth0ArrayOutput() SourceAuth0ArrayOutput {
	return o
}

func (o SourceAuth0ArrayOutput) ToSourceAuth0ArrayOutputWithContext(ctx context.Context) SourceAuth0ArrayOutput {
	return o
}

func (o SourceAuth0ArrayOutput) Index(i pulumi.IntInput) SourceAuth0Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceAuth0 {
		return vs[0].([]*SourceAuth0)[vs[1].(int)]
	}).(SourceAuth0Output)
}

type SourceAuth0MapOutput struct{ *pulumi.OutputState }

func (SourceAuth0MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAuth0)(nil)).Elem()
}

func (o SourceAuth0MapOutput) ToSourceAuth0MapOutput() SourceAuth0MapOutput {
	return o
}

func (o SourceAuth0MapOutput) ToSourceAuth0MapOutputWithContext(ctx context.Context) SourceAuth0MapOutput {
	return o
}

func (o SourceAuth0MapOutput) MapIndex(k pulumi.StringInput) SourceAuth0Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceAuth0 {
		return vs[0].(map[string]*SourceAuth0)[vs[1].(string)]
	}).(SourceAuth0Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAuth0Input)(nil)).Elem(), &SourceAuth0{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAuth0ArrayInput)(nil)).Elem(), SourceAuth0Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAuth0MapInput)(nil)).Elem(), SourceAuth0Map{})
	pulumi.RegisterOutputType(SourceAuth0Output{})
	pulumi.RegisterOutputType(SourceAuth0ArrayOutput{})
	pulumi.RegisterOutputType(SourceAuth0MapOutput{})
}
