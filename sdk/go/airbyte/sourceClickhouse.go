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

// SourceClickhouse Resource
type SourceClickhouse struct {
	pulumi.CustomResourceState

	Configuration SourceClickhouseConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                 `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceClickhouse registers a new resource with the given unique name, arguments, and options.
func NewSourceClickhouse(ctx *pulumi.Context,
	name string, args *SourceClickhouseArgs, opts ...pulumi.ResourceOption) (*SourceClickhouse, error) {
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
	var resource SourceClickhouse
	err := ctx.RegisterResource("airbyte:index/sourceClickhouse:SourceClickhouse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceClickhouse gets an existing SourceClickhouse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceClickhouse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceClickhouseState, opts ...pulumi.ResourceOption) (*SourceClickhouse, error) {
	var resource SourceClickhouse
	err := ctx.ReadResource("airbyte:index/sourceClickhouse:SourceClickhouse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceClickhouse resources.
type sourceClickhouseState struct {
	Configuration *SourceClickhouseConfiguration `pulumi:"configuration"`
	Name          *string                        `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceClickhouseState struct {
	Configuration SourceClickhouseConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceClickhouseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceClickhouseState)(nil)).Elem()
}

type sourceClickhouseArgs struct {
	Configuration SourceClickhouseConfiguration `pulumi:"configuration"`
	Name          string                        `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceClickhouse resource.
type SourceClickhouseArgs struct {
	Configuration SourceClickhouseConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceClickhouseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceClickhouseArgs)(nil)).Elem()
}

type SourceClickhouseInput interface {
	pulumi.Input

	ToSourceClickhouseOutput() SourceClickhouseOutput
	ToSourceClickhouseOutputWithContext(ctx context.Context) SourceClickhouseOutput
}

func (*SourceClickhouse) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceClickhouse)(nil)).Elem()
}

func (i *SourceClickhouse) ToSourceClickhouseOutput() SourceClickhouseOutput {
	return i.ToSourceClickhouseOutputWithContext(context.Background())
}

func (i *SourceClickhouse) ToSourceClickhouseOutputWithContext(ctx context.Context) SourceClickhouseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceClickhouseOutput)
}

// SourceClickhouseArrayInput is an input type that accepts SourceClickhouseArray and SourceClickhouseArrayOutput values.
// You can construct a concrete instance of `SourceClickhouseArrayInput` via:
//
//	SourceClickhouseArray{ SourceClickhouseArgs{...} }
type SourceClickhouseArrayInput interface {
	pulumi.Input

	ToSourceClickhouseArrayOutput() SourceClickhouseArrayOutput
	ToSourceClickhouseArrayOutputWithContext(context.Context) SourceClickhouseArrayOutput
}

type SourceClickhouseArray []SourceClickhouseInput

func (SourceClickhouseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceClickhouse)(nil)).Elem()
}

func (i SourceClickhouseArray) ToSourceClickhouseArrayOutput() SourceClickhouseArrayOutput {
	return i.ToSourceClickhouseArrayOutputWithContext(context.Background())
}

func (i SourceClickhouseArray) ToSourceClickhouseArrayOutputWithContext(ctx context.Context) SourceClickhouseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceClickhouseArrayOutput)
}

// SourceClickhouseMapInput is an input type that accepts SourceClickhouseMap and SourceClickhouseMapOutput values.
// You can construct a concrete instance of `SourceClickhouseMapInput` via:
//
//	SourceClickhouseMap{ "key": SourceClickhouseArgs{...} }
type SourceClickhouseMapInput interface {
	pulumi.Input

	ToSourceClickhouseMapOutput() SourceClickhouseMapOutput
	ToSourceClickhouseMapOutputWithContext(context.Context) SourceClickhouseMapOutput
}

type SourceClickhouseMap map[string]SourceClickhouseInput

func (SourceClickhouseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceClickhouse)(nil)).Elem()
}

func (i SourceClickhouseMap) ToSourceClickhouseMapOutput() SourceClickhouseMapOutput {
	return i.ToSourceClickhouseMapOutputWithContext(context.Background())
}

func (i SourceClickhouseMap) ToSourceClickhouseMapOutputWithContext(ctx context.Context) SourceClickhouseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceClickhouseMapOutput)
}

type SourceClickhouseOutput struct{ *pulumi.OutputState }

func (SourceClickhouseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceClickhouse)(nil)).Elem()
}

func (o SourceClickhouseOutput) ToSourceClickhouseOutput() SourceClickhouseOutput {
	return o
}

func (o SourceClickhouseOutput) ToSourceClickhouseOutputWithContext(ctx context.Context) SourceClickhouseOutput {
	return o
}

func (o SourceClickhouseOutput) Configuration() SourceClickhouseConfigurationOutput {
	return o.ApplyT(func(v *SourceClickhouse) SourceClickhouseConfigurationOutput { return v.Configuration }).(SourceClickhouseConfigurationOutput)
}

func (o SourceClickhouseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClickhouse) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceClickhouseOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceClickhouse) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceClickhouseOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClickhouse) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceClickhouseOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClickhouse) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceClickhouseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClickhouse) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceClickhouseArrayOutput struct{ *pulumi.OutputState }

func (SourceClickhouseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceClickhouse)(nil)).Elem()
}

func (o SourceClickhouseArrayOutput) ToSourceClickhouseArrayOutput() SourceClickhouseArrayOutput {
	return o
}

func (o SourceClickhouseArrayOutput) ToSourceClickhouseArrayOutputWithContext(ctx context.Context) SourceClickhouseArrayOutput {
	return o
}

func (o SourceClickhouseArrayOutput) Index(i pulumi.IntInput) SourceClickhouseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceClickhouse {
		return vs[0].([]*SourceClickhouse)[vs[1].(int)]
	}).(SourceClickhouseOutput)
}

type SourceClickhouseMapOutput struct{ *pulumi.OutputState }

func (SourceClickhouseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceClickhouse)(nil)).Elem()
}

func (o SourceClickhouseMapOutput) ToSourceClickhouseMapOutput() SourceClickhouseMapOutput {
	return o
}

func (o SourceClickhouseMapOutput) ToSourceClickhouseMapOutputWithContext(ctx context.Context) SourceClickhouseMapOutput {
	return o
}

func (o SourceClickhouseMapOutput) MapIndex(k pulumi.StringInput) SourceClickhouseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceClickhouse {
		return vs[0].(map[string]*SourceClickhouse)[vs[1].(string)]
	}).(SourceClickhouseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceClickhouseInput)(nil)).Elem(), &SourceClickhouse{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceClickhouseArrayInput)(nil)).Elem(), SourceClickhouseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceClickhouseMapInput)(nil)).Elem(), SourceClickhouseMap{})
	pulumi.RegisterOutputType(SourceClickhouseOutput{})
	pulumi.RegisterOutputType(SourceClickhouseArrayOutput{})
	pulumi.RegisterOutputType(SourceClickhouseMapOutput{})
}
