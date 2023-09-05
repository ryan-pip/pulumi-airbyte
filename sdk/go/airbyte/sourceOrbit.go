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

// SourceOrbit Resource
type SourceOrbit struct {
	pulumi.CustomResourceState

	Configuration SourceOrbitConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceOrbit registers a new resource with the given unique name, arguments, and options.
func NewSourceOrbit(ctx *pulumi.Context,
	name string, args *SourceOrbitArgs, opts ...pulumi.ResourceOption) (*SourceOrbit, error) {
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
	var resource SourceOrbit
	err := ctx.RegisterResource("airbyte:index/sourceOrbit:SourceOrbit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceOrbit gets an existing SourceOrbit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceOrbit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceOrbitState, opts ...pulumi.ResourceOption) (*SourceOrbit, error) {
	var resource SourceOrbit
	err := ctx.ReadResource("airbyte:index/sourceOrbit:SourceOrbit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceOrbit resources.
type sourceOrbitState struct {
	Configuration *SourceOrbitConfiguration `pulumi:"configuration"`
	Name          *string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceOrbitState struct {
	Configuration SourceOrbitConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceOrbitState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOrbitState)(nil)).Elem()
}

type sourceOrbitArgs struct {
	Configuration SourceOrbitConfiguration `pulumi:"configuration"`
	Name          string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceOrbit resource.
type SourceOrbitArgs struct {
	Configuration SourceOrbitConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceOrbitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOrbitArgs)(nil)).Elem()
}

type SourceOrbitInput interface {
	pulumi.Input

	ToSourceOrbitOutput() SourceOrbitOutput
	ToSourceOrbitOutputWithContext(ctx context.Context) SourceOrbitOutput
}

func (*SourceOrbit) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOrbit)(nil)).Elem()
}

func (i *SourceOrbit) ToSourceOrbitOutput() SourceOrbitOutput {
	return i.ToSourceOrbitOutputWithContext(context.Background())
}

func (i *SourceOrbit) ToSourceOrbitOutputWithContext(ctx context.Context) SourceOrbitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOrbitOutput)
}

// SourceOrbitArrayInput is an input type that accepts SourceOrbitArray and SourceOrbitArrayOutput values.
// You can construct a concrete instance of `SourceOrbitArrayInput` via:
//
//	SourceOrbitArray{ SourceOrbitArgs{...} }
type SourceOrbitArrayInput interface {
	pulumi.Input

	ToSourceOrbitArrayOutput() SourceOrbitArrayOutput
	ToSourceOrbitArrayOutputWithContext(context.Context) SourceOrbitArrayOutput
}

type SourceOrbitArray []SourceOrbitInput

func (SourceOrbitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceOrbit)(nil)).Elem()
}

func (i SourceOrbitArray) ToSourceOrbitArrayOutput() SourceOrbitArrayOutput {
	return i.ToSourceOrbitArrayOutputWithContext(context.Background())
}

func (i SourceOrbitArray) ToSourceOrbitArrayOutputWithContext(ctx context.Context) SourceOrbitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOrbitArrayOutput)
}

// SourceOrbitMapInput is an input type that accepts SourceOrbitMap and SourceOrbitMapOutput values.
// You can construct a concrete instance of `SourceOrbitMapInput` via:
//
//	SourceOrbitMap{ "key": SourceOrbitArgs{...} }
type SourceOrbitMapInput interface {
	pulumi.Input

	ToSourceOrbitMapOutput() SourceOrbitMapOutput
	ToSourceOrbitMapOutputWithContext(context.Context) SourceOrbitMapOutput
}

type SourceOrbitMap map[string]SourceOrbitInput

func (SourceOrbitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceOrbit)(nil)).Elem()
}

func (i SourceOrbitMap) ToSourceOrbitMapOutput() SourceOrbitMapOutput {
	return i.ToSourceOrbitMapOutputWithContext(context.Background())
}

func (i SourceOrbitMap) ToSourceOrbitMapOutputWithContext(ctx context.Context) SourceOrbitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOrbitMapOutput)
}

type SourceOrbitOutput struct{ *pulumi.OutputState }

func (SourceOrbitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOrbit)(nil)).Elem()
}

func (o SourceOrbitOutput) ToSourceOrbitOutput() SourceOrbitOutput {
	return o
}

func (o SourceOrbitOutput) ToSourceOrbitOutputWithContext(ctx context.Context) SourceOrbitOutput {
	return o
}

func (o SourceOrbitOutput) Configuration() SourceOrbitConfigurationOutput {
	return o.ApplyT(func(v *SourceOrbit) SourceOrbitConfigurationOutput { return v.Configuration }).(SourceOrbitConfigurationOutput)
}

func (o SourceOrbitOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOrbit) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceOrbitOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOrbit) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceOrbitOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOrbit) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceOrbitOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOrbit) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceOrbitOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOrbit) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceOrbitArrayOutput struct{ *pulumi.OutputState }

func (SourceOrbitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceOrbit)(nil)).Elem()
}

func (o SourceOrbitArrayOutput) ToSourceOrbitArrayOutput() SourceOrbitArrayOutput {
	return o
}

func (o SourceOrbitArrayOutput) ToSourceOrbitArrayOutputWithContext(ctx context.Context) SourceOrbitArrayOutput {
	return o
}

func (o SourceOrbitArrayOutput) Index(i pulumi.IntInput) SourceOrbitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceOrbit {
		return vs[0].([]*SourceOrbit)[vs[1].(int)]
	}).(SourceOrbitOutput)
}

type SourceOrbitMapOutput struct{ *pulumi.OutputState }

func (SourceOrbitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceOrbit)(nil)).Elem()
}

func (o SourceOrbitMapOutput) ToSourceOrbitMapOutput() SourceOrbitMapOutput {
	return o
}

func (o SourceOrbitMapOutput) ToSourceOrbitMapOutputWithContext(ctx context.Context) SourceOrbitMapOutput {
	return o
}

func (o SourceOrbitMapOutput) MapIndex(k pulumi.StringInput) SourceOrbitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceOrbit {
		return vs[0].(map[string]*SourceOrbit)[vs[1].(string)]
	}).(SourceOrbitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOrbitInput)(nil)).Elem(), &SourceOrbit{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOrbitArrayInput)(nil)).Elem(), SourceOrbitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOrbitMapInput)(nil)).Elem(), SourceOrbitMap{})
	pulumi.RegisterOutputType(SourceOrbitOutput{})
	pulumi.RegisterOutputType(SourceOrbitArrayOutput{})
	pulumi.RegisterOutputType(SourceOrbitMapOutput{})
}
