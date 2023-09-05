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

// SourceGoogleAnalyticsDataAPI Resource
type SourceGoogleAnalyticsDataApi struct {
	pulumi.CustomResourceState

	Configuration SourceGoogleAnalyticsDataApiConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceGoogleAnalyticsDataApi registers a new resource with the given unique name, arguments, and options.
func NewSourceGoogleAnalyticsDataApi(ctx *pulumi.Context,
	name string, args *SourceGoogleAnalyticsDataApiArgs, opts ...pulumi.ResourceOption) (*SourceGoogleAnalyticsDataApi, error) {
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
	var resource SourceGoogleAnalyticsDataApi
	err := ctx.RegisterResource("airbyte:index/sourceGoogleAnalyticsDataApi:SourceGoogleAnalyticsDataApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceGoogleAnalyticsDataApi gets an existing SourceGoogleAnalyticsDataApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceGoogleAnalyticsDataApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceGoogleAnalyticsDataApiState, opts ...pulumi.ResourceOption) (*SourceGoogleAnalyticsDataApi, error) {
	var resource SourceGoogleAnalyticsDataApi
	err := ctx.ReadResource("airbyte:index/sourceGoogleAnalyticsDataApi:SourceGoogleAnalyticsDataApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceGoogleAnalyticsDataApi resources.
type sourceGoogleAnalyticsDataApiState struct {
	Configuration *SourceGoogleAnalyticsDataApiConfiguration `pulumi:"configuration"`
	Name          *string                                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceGoogleAnalyticsDataApiState struct {
	Configuration SourceGoogleAnalyticsDataApiConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceGoogleAnalyticsDataApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGoogleAnalyticsDataApiState)(nil)).Elem()
}

type sourceGoogleAnalyticsDataApiArgs struct {
	Configuration SourceGoogleAnalyticsDataApiConfiguration `pulumi:"configuration"`
	Name          string                                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceGoogleAnalyticsDataApi resource.
type SourceGoogleAnalyticsDataApiArgs struct {
	Configuration SourceGoogleAnalyticsDataApiConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceGoogleAnalyticsDataApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGoogleAnalyticsDataApiArgs)(nil)).Elem()
}

type SourceGoogleAnalyticsDataApiInput interface {
	pulumi.Input

	ToSourceGoogleAnalyticsDataApiOutput() SourceGoogleAnalyticsDataApiOutput
	ToSourceGoogleAnalyticsDataApiOutputWithContext(ctx context.Context) SourceGoogleAnalyticsDataApiOutput
}

func (*SourceGoogleAnalyticsDataApi) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGoogleAnalyticsDataApi)(nil)).Elem()
}

func (i *SourceGoogleAnalyticsDataApi) ToSourceGoogleAnalyticsDataApiOutput() SourceGoogleAnalyticsDataApiOutput {
	return i.ToSourceGoogleAnalyticsDataApiOutputWithContext(context.Background())
}

func (i *SourceGoogleAnalyticsDataApi) ToSourceGoogleAnalyticsDataApiOutputWithContext(ctx context.Context) SourceGoogleAnalyticsDataApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleAnalyticsDataApiOutput)
}

// SourceGoogleAnalyticsDataApiArrayInput is an input type that accepts SourceGoogleAnalyticsDataApiArray and SourceGoogleAnalyticsDataApiArrayOutput values.
// You can construct a concrete instance of `SourceGoogleAnalyticsDataApiArrayInput` via:
//
//	SourceGoogleAnalyticsDataApiArray{ SourceGoogleAnalyticsDataApiArgs{...} }
type SourceGoogleAnalyticsDataApiArrayInput interface {
	pulumi.Input

	ToSourceGoogleAnalyticsDataApiArrayOutput() SourceGoogleAnalyticsDataApiArrayOutput
	ToSourceGoogleAnalyticsDataApiArrayOutputWithContext(context.Context) SourceGoogleAnalyticsDataApiArrayOutput
}

type SourceGoogleAnalyticsDataApiArray []SourceGoogleAnalyticsDataApiInput

func (SourceGoogleAnalyticsDataApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGoogleAnalyticsDataApi)(nil)).Elem()
}

func (i SourceGoogleAnalyticsDataApiArray) ToSourceGoogleAnalyticsDataApiArrayOutput() SourceGoogleAnalyticsDataApiArrayOutput {
	return i.ToSourceGoogleAnalyticsDataApiArrayOutputWithContext(context.Background())
}

func (i SourceGoogleAnalyticsDataApiArray) ToSourceGoogleAnalyticsDataApiArrayOutputWithContext(ctx context.Context) SourceGoogleAnalyticsDataApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleAnalyticsDataApiArrayOutput)
}

// SourceGoogleAnalyticsDataApiMapInput is an input type that accepts SourceGoogleAnalyticsDataApiMap and SourceGoogleAnalyticsDataApiMapOutput values.
// You can construct a concrete instance of `SourceGoogleAnalyticsDataApiMapInput` via:
//
//	SourceGoogleAnalyticsDataApiMap{ "key": SourceGoogleAnalyticsDataApiArgs{...} }
type SourceGoogleAnalyticsDataApiMapInput interface {
	pulumi.Input

	ToSourceGoogleAnalyticsDataApiMapOutput() SourceGoogleAnalyticsDataApiMapOutput
	ToSourceGoogleAnalyticsDataApiMapOutputWithContext(context.Context) SourceGoogleAnalyticsDataApiMapOutput
}

type SourceGoogleAnalyticsDataApiMap map[string]SourceGoogleAnalyticsDataApiInput

func (SourceGoogleAnalyticsDataApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGoogleAnalyticsDataApi)(nil)).Elem()
}

func (i SourceGoogleAnalyticsDataApiMap) ToSourceGoogleAnalyticsDataApiMapOutput() SourceGoogleAnalyticsDataApiMapOutput {
	return i.ToSourceGoogleAnalyticsDataApiMapOutputWithContext(context.Background())
}

func (i SourceGoogleAnalyticsDataApiMap) ToSourceGoogleAnalyticsDataApiMapOutputWithContext(ctx context.Context) SourceGoogleAnalyticsDataApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleAnalyticsDataApiMapOutput)
}

type SourceGoogleAnalyticsDataApiOutput struct{ *pulumi.OutputState }

func (SourceGoogleAnalyticsDataApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGoogleAnalyticsDataApi)(nil)).Elem()
}

func (o SourceGoogleAnalyticsDataApiOutput) ToSourceGoogleAnalyticsDataApiOutput() SourceGoogleAnalyticsDataApiOutput {
	return o
}

func (o SourceGoogleAnalyticsDataApiOutput) ToSourceGoogleAnalyticsDataApiOutputWithContext(ctx context.Context) SourceGoogleAnalyticsDataApiOutput {
	return o
}

func (o SourceGoogleAnalyticsDataApiOutput) Configuration() SourceGoogleAnalyticsDataApiConfigurationOutput {
	return o.ApplyT(func(v *SourceGoogleAnalyticsDataApi) SourceGoogleAnalyticsDataApiConfigurationOutput {
		return v.Configuration
	}).(SourceGoogleAnalyticsDataApiConfigurationOutput)
}

func (o SourceGoogleAnalyticsDataApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleAnalyticsDataApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceGoogleAnalyticsDataApiOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGoogleAnalyticsDataApi) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceGoogleAnalyticsDataApiOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleAnalyticsDataApi) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceGoogleAnalyticsDataApiOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleAnalyticsDataApi) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceGoogleAnalyticsDataApiOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleAnalyticsDataApi) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceGoogleAnalyticsDataApiArrayOutput struct{ *pulumi.OutputState }

func (SourceGoogleAnalyticsDataApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGoogleAnalyticsDataApi)(nil)).Elem()
}

func (o SourceGoogleAnalyticsDataApiArrayOutput) ToSourceGoogleAnalyticsDataApiArrayOutput() SourceGoogleAnalyticsDataApiArrayOutput {
	return o
}

func (o SourceGoogleAnalyticsDataApiArrayOutput) ToSourceGoogleAnalyticsDataApiArrayOutputWithContext(ctx context.Context) SourceGoogleAnalyticsDataApiArrayOutput {
	return o
}

func (o SourceGoogleAnalyticsDataApiArrayOutput) Index(i pulumi.IntInput) SourceGoogleAnalyticsDataApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceGoogleAnalyticsDataApi {
		return vs[0].([]*SourceGoogleAnalyticsDataApi)[vs[1].(int)]
	}).(SourceGoogleAnalyticsDataApiOutput)
}

type SourceGoogleAnalyticsDataApiMapOutput struct{ *pulumi.OutputState }

func (SourceGoogleAnalyticsDataApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGoogleAnalyticsDataApi)(nil)).Elem()
}

func (o SourceGoogleAnalyticsDataApiMapOutput) ToSourceGoogleAnalyticsDataApiMapOutput() SourceGoogleAnalyticsDataApiMapOutput {
	return o
}

func (o SourceGoogleAnalyticsDataApiMapOutput) ToSourceGoogleAnalyticsDataApiMapOutputWithContext(ctx context.Context) SourceGoogleAnalyticsDataApiMapOutput {
	return o
}

func (o SourceGoogleAnalyticsDataApiMapOutput) MapIndex(k pulumi.StringInput) SourceGoogleAnalyticsDataApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceGoogleAnalyticsDataApi {
		return vs[0].(map[string]*SourceGoogleAnalyticsDataApi)[vs[1].(string)]
	}).(SourceGoogleAnalyticsDataApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleAnalyticsDataApiInput)(nil)).Elem(), &SourceGoogleAnalyticsDataApi{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleAnalyticsDataApiArrayInput)(nil)).Elem(), SourceGoogleAnalyticsDataApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleAnalyticsDataApiMapInput)(nil)).Elem(), SourceGoogleAnalyticsDataApiMap{})
	pulumi.RegisterOutputType(SourceGoogleAnalyticsDataApiOutput{})
	pulumi.RegisterOutputType(SourceGoogleAnalyticsDataApiArrayOutput{})
	pulumi.RegisterOutputType(SourceGoogleAnalyticsDataApiMapOutput{})
}
