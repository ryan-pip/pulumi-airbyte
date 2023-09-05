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

// SourceYandexMetrica Resource
type SourceYandexMetrica struct {
	pulumi.CustomResourceState

	Configuration SourceYandexMetricaConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceYandexMetrica registers a new resource with the given unique name, arguments, and options.
func NewSourceYandexMetrica(ctx *pulumi.Context,
	name string, args *SourceYandexMetricaArgs, opts ...pulumi.ResourceOption) (*SourceYandexMetrica, error) {
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
	var resource SourceYandexMetrica
	err := ctx.RegisterResource("airbyte:index/sourceYandexMetrica:SourceYandexMetrica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceYandexMetrica gets an existing SourceYandexMetrica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceYandexMetrica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceYandexMetricaState, opts ...pulumi.ResourceOption) (*SourceYandexMetrica, error) {
	var resource SourceYandexMetrica
	err := ctx.ReadResource("airbyte:index/sourceYandexMetrica:SourceYandexMetrica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceYandexMetrica resources.
type sourceYandexMetricaState struct {
	Configuration *SourceYandexMetricaConfiguration `pulumi:"configuration"`
	Name          *string                           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceYandexMetricaState struct {
	Configuration SourceYandexMetricaConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceYandexMetricaState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceYandexMetricaState)(nil)).Elem()
}

type sourceYandexMetricaArgs struct {
	Configuration SourceYandexMetricaConfiguration `pulumi:"configuration"`
	Name          string                           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceYandexMetrica resource.
type SourceYandexMetricaArgs struct {
	Configuration SourceYandexMetricaConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceYandexMetricaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceYandexMetricaArgs)(nil)).Elem()
}

type SourceYandexMetricaInput interface {
	pulumi.Input

	ToSourceYandexMetricaOutput() SourceYandexMetricaOutput
	ToSourceYandexMetricaOutputWithContext(ctx context.Context) SourceYandexMetricaOutput
}

func (*SourceYandexMetrica) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceYandexMetrica)(nil)).Elem()
}

func (i *SourceYandexMetrica) ToSourceYandexMetricaOutput() SourceYandexMetricaOutput {
	return i.ToSourceYandexMetricaOutputWithContext(context.Background())
}

func (i *SourceYandexMetrica) ToSourceYandexMetricaOutputWithContext(ctx context.Context) SourceYandexMetricaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceYandexMetricaOutput)
}

// SourceYandexMetricaArrayInput is an input type that accepts SourceYandexMetricaArray and SourceYandexMetricaArrayOutput values.
// You can construct a concrete instance of `SourceYandexMetricaArrayInput` via:
//
//	SourceYandexMetricaArray{ SourceYandexMetricaArgs{...} }
type SourceYandexMetricaArrayInput interface {
	pulumi.Input

	ToSourceYandexMetricaArrayOutput() SourceYandexMetricaArrayOutput
	ToSourceYandexMetricaArrayOutputWithContext(context.Context) SourceYandexMetricaArrayOutput
}

type SourceYandexMetricaArray []SourceYandexMetricaInput

func (SourceYandexMetricaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceYandexMetrica)(nil)).Elem()
}

func (i SourceYandexMetricaArray) ToSourceYandexMetricaArrayOutput() SourceYandexMetricaArrayOutput {
	return i.ToSourceYandexMetricaArrayOutputWithContext(context.Background())
}

func (i SourceYandexMetricaArray) ToSourceYandexMetricaArrayOutputWithContext(ctx context.Context) SourceYandexMetricaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceYandexMetricaArrayOutput)
}

// SourceYandexMetricaMapInput is an input type that accepts SourceYandexMetricaMap and SourceYandexMetricaMapOutput values.
// You can construct a concrete instance of `SourceYandexMetricaMapInput` via:
//
//	SourceYandexMetricaMap{ "key": SourceYandexMetricaArgs{...} }
type SourceYandexMetricaMapInput interface {
	pulumi.Input

	ToSourceYandexMetricaMapOutput() SourceYandexMetricaMapOutput
	ToSourceYandexMetricaMapOutputWithContext(context.Context) SourceYandexMetricaMapOutput
}

type SourceYandexMetricaMap map[string]SourceYandexMetricaInput

func (SourceYandexMetricaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceYandexMetrica)(nil)).Elem()
}

func (i SourceYandexMetricaMap) ToSourceYandexMetricaMapOutput() SourceYandexMetricaMapOutput {
	return i.ToSourceYandexMetricaMapOutputWithContext(context.Background())
}

func (i SourceYandexMetricaMap) ToSourceYandexMetricaMapOutputWithContext(ctx context.Context) SourceYandexMetricaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceYandexMetricaMapOutput)
}

type SourceYandexMetricaOutput struct{ *pulumi.OutputState }

func (SourceYandexMetricaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceYandexMetrica)(nil)).Elem()
}

func (o SourceYandexMetricaOutput) ToSourceYandexMetricaOutput() SourceYandexMetricaOutput {
	return o
}

func (o SourceYandexMetricaOutput) ToSourceYandexMetricaOutputWithContext(ctx context.Context) SourceYandexMetricaOutput {
	return o
}

func (o SourceYandexMetricaOutput) Configuration() SourceYandexMetricaConfigurationOutput {
	return o.ApplyT(func(v *SourceYandexMetrica) SourceYandexMetricaConfigurationOutput { return v.Configuration }).(SourceYandexMetricaConfigurationOutput)
}

func (o SourceYandexMetricaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYandexMetrica) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceYandexMetricaOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceYandexMetrica) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceYandexMetricaOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYandexMetrica) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceYandexMetricaOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYandexMetrica) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceYandexMetricaOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYandexMetrica) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceYandexMetricaArrayOutput struct{ *pulumi.OutputState }

func (SourceYandexMetricaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceYandexMetrica)(nil)).Elem()
}

func (o SourceYandexMetricaArrayOutput) ToSourceYandexMetricaArrayOutput() SourceYandexMetricaArrayOutput {
	return o
}

func (o SourceYandexMetricaArrayOutput) ToSourceYandexMetricaArrayOutputWithContext(ctx context.Context) SourceYandexMetricaArrayOutput {
	return o
}

func (o SourceYandexMetricaArrayOutput) Index(i pulumi.IntInput) SourceYandexMetricaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceYandexMetrica {
		return vs[0].([]*SourceYandexMetrica)[vs[1].(int)]
	}).(SourceYandexMetricaOutput)
}

type SourceYandexMetricaMapOutput struct{ *pulumi.OutputState }

func (SourceYandexMetricaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceYandexMetrica)(nil)).Elem()
}

func (o SourceYandexMetricaMapOutput) ToSourceYandexMetricaMapOutput() SourceYandexMetricaMapOutput {
	return o
}

func (o SourceYandexMetricaMapOutput) ToSourceYandexMetricaMapOutputWithContext(ctx context.Context) SourceYandexMetricaMapOutput {
	return o
}

func (o SourceYandexMetricaMapOutput) MapIndex(k pulumi.StringInput) SourceYandexMetricaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceYandexMetrica {
		return vs[0].(map[string]*SourceYandexMetrica)[vs[1].(string)]
	}).(SourceYandexMetricaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceYandexMetricaInput)(nil)).Elem(), &SourceYandexMetrica{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceYandexMetricaArrayInput)(nil)).Elem(), SourceYandexMetricaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceYandexMetricaMapInput)(nil)).Elem(), SourceYandexMetricaMap{})
	pulumi.RegisterOutputType(SourceYandexMetricaOutput{})
	pulumi.RegisterOutputType(SourceYandexMetricaArrayOutput{})
	pulumi.RegisterOutputType(SourceYandexMetricaMapOutput{})
}
