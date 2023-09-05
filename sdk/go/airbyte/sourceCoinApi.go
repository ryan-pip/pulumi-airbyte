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

// SourceCoinAPI Resource
type SourceCoinApi struct {
	pulumi.CustomResourceState

	Configuration SourceCoinApiConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceCoinApi registers a new resource with the given unique name, arguments, and options.
func NewSourceCoinApi(ctx *pulumi.Context,
	name string, args *SourceCoinApiArgs, opts ...pulumi.ResourceOption) (*SourceCoinApi, error) {
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
	var resource SourceCoinApi
	err := ctx.RegisterResource("airbyte:index/sourceCoinApi:SourceCoinApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceCoinApi gets an existing SourceCoinApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceCoinApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceCoinApiState, opts ...pulumi.ResourceOption) (*SourceCoinApi, error) {
	var resource SourceCoinApi
	err := ctx.ReadResource("airbyte:index/sourceCoinApi:SourceCoinApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceCoinApi resources.
type sourceCoinApiState struct {
	Configuration *SourceCoinApiConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceCoinApiState struct {
	Configuration SourceCoinApiConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceCoinApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCoinApiState)(nil)).Elem()
}

type sourceCoinApiArgs struct {
	Configuration SourceCoinApiConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceCoinApi resource.
type SourceCoinApiArgs struct {
	Configuration SourceCoinApiConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceCoinApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCoinApiArgs)(nil)).Elem()
}

type SourceCoinApiInput interface {
	pulumi.Input

	ToSourceCoinApiOutput() SourceCoinApiOutput
	ToSourceCoinApiOutputWithContext(ctx context.Context) SourceCoinApiOutput
}

func (*SourceCoinApi) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCoinApi)(nil)).Elem()
}

func (i *SourceCoinApi) ToSourceCoinApiOutput() SourceCoinApiOutput {
	return i.ToSourceCoinApiOutputWithContext(context.Background())
}

func (i *SourceCoinApi) ToSourceCoinApiOutputWithContext(ctx context.Context) SourceCoinApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCoinApiOutput)
}

// SourceCoinApiArrayInput is an input type that accepts SourceCoinApiArray and SourceCoinApiArrayOutput values.
// You can construct a concrete instance of `SourceCoinApiArrayInput` via:
//
//	SourceCoinApiArray{ SourceCoinApiArgs{...} }
type SourceCoinApiArrayInput interface {
	pulumi.Input

	ToSourceCoinApiArrayOutput() SourceCoinApiArrayOutput
	ToSourceCoinApiArrayOutputWithContext(context.Context) SourceCoinApiArrayOutput
}

type SourceCoinApiArray []SourceCoinApiInput

func (SourceCoinApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceCoinApi)(nil)).Elem()
}

func (i SourceCoinApiArray) ToSourceCoinApiArrayOutput() SourceCoinApiArrayOutput {
	return i.ToSourceCoinApiArrayOutputWithContext(context.Background())
}

func (i SourceCoinApiArray) ToSourceCoinApiArrayOutputWithContext(ctx context.Context) SourceCoinApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCoinApiArrayOutput)
}

// SourceCoinApiMapInput is an input type that accepts SourceCoinApiMap and SourceCoinApiMapOutput values.
// You can construct a concrete instance of `SourceCoinApiMapInput` via:
//
//	SourceCoinApiMap{ "key": SourceCoinApiArgs{...} }
type SourceCoinApiMapInput interface {
	pulumi.Input

	ToSourceCoinApiMapOutput() SourceCoinApiMapOutput
	ToSourceCoinApiMapOutputWithContext(context.Context) SourceCoinApiMapOutput
}

type SourceCoinApiMap map[string]SourceCoinApiInput

func (SourceCoinApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceCoinApi)(nil)).Elem()
}

func (i SourceCoinApiMap) ToSourceCoinApiMapOutput() SourceCoinApiMapOutput {
	return i.ToSourceCoinApiMapOutputWithContext(context.Background())
}

func (i SourceCoinApiMap) ToSourceCoinApiMapOutputWithContext(ctx context.Context) SourceCoinApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCoinApiMapOutput)
}

type SourceCoinApiOutput struct{ *pulumi.OutputState }

func (SourceCoinApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCoinApi)(nil)).Elem()
}

func (o SourceCoinApiOutput) ToSourceCoinApiOutput() SourceCoinApiOutput {
	return o
}

func (o SourceCoinApiOutput) ToSourceCoinApiOutputWithContext(ctx context.Context) SourceCoinApiOutput {
	return o
}

func (o SourceCoinApiOutput) Configuration() SourceCoinApiConfigurationOutput {
	return o.ApplyT(func(v *SourceCoinApi) SourceCoinApiConfigurationOutput { return v.Configuration }).(SourceCoinApiConfigurationOutput)
}

func (o SourceCoinApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoinApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceCoinApiOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCoinApi) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceCoinApiOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoinApi) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceCoinApiOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoinApi) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceCoinApiOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCoinApi) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceCoinApiArrayOutput struct{ *pulumi.OutputState }

func (SourceCoinApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceCoinApi)(nil)).Elem()
}

func (o SourceCoinApiArrayOutput) ToSourceCoinApiArrayOutput() SourceCoinApiArrayOutput {
	return o
}

func (o SourceCoinApiArrayOutput) ToSourceCoinApiArrayOutputWithContext(ctx context.Context) SourceCoinApiArrayOutput {
	return o
}

func (o SourceCoinApiArrayOutput) Index(i pulumi.IntInput) SourceCoinApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceCoinApi {
		return vs[0].([]*SourceCoinApi)[vs[1].(int)]
	}).(SourceCoinApiOutput)
}

type SourceCoinApiMapOutput struct{ *pulumi.OutputState }

func (SourceCoinApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceCoinApi)(nil)).Elem()
}

func (o SourceCoinApiMapOutput) ToSourceCoinApiMapOutput() SourceCoinApiMapOutput {
	return o
}

func (o SourceCoinApiMapOutput) ToSourceCoinApiMapOutputWithContext(ctx context.Context) SourceCoinApiMapOutput {
	return o
}

func (o SourceCoinApiMapOutput) MapIndex(k pulumi.StringInput) SourceCoinApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceCoinApi {
		return vs[0].(map[string]*SourceCoinApi)[vs[1].(string)]
	}).(SourceCoinApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCoinApiInput)(nil)).Elem(), &SourceCoinApi{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCoinApiArrayInput)(nil)).Elem(), SourceCoinApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCoinApiMapInput)(nil)).Elem(), SourceCoinApiMap{})
	pulumi.RegisterOutputType(SourceCoinApiOutput{})
	pulumi.RegisterOutputType(SourceCoinApiArrayOutput{})
	pulumi.RegisterOutputType(SourceCoinApiMapOutput{})
}
