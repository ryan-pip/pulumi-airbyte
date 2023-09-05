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

// SourceGoogleAds Resource
type SourceGoogleAds struct {
	pulumi.CustomResourceState

	Configuration SourceGoogleAdsConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceGoogleAds registers a new resource with the given unique name, arguments, and options.
func NewSourceGoogleAds(ctx *pulumi.Context,
	name string, args *SourceGoogleAdsArgs, opts ...pulumi.ResourceOption) (*SourceGoogleAds, error) {
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
	var resource SourceGoogleAds
	err := ctx.RegisterResource("airbyte:index/sourceGoogleAds:SourceGoogleAds", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceGoogleAds gets an existing SourceGoogleAds resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceGoogleAds(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceGoogleAdsState, opts ...pulumi.ResourceOption) (*SourceGoogleAds, error) {
	var resource SourceGoogleAds
	err := ctx.ReadResource("airbyte:index/sourceGoogleAds:SourceGoogleAds", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceGoogleAds resources.
type sourceGoogleAdsState struct {
	Configuration *SourceGoogleAdsConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceGoogleAdsState struct {
	Configuration SourceGoogleAdsConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceGoogleAdsState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGoogleAdsState)(nil)).Elem()
}

type sourceGoogleAdsArgs struct {
	Configuration SourceGoogleAdsConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceGoogleAds resource.
type SourceGoogleAdsArgs struct {
	Configuration SourceGoogleAdsConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceGoogleAdsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGoogleAdsArgs)(nil)).Elem()
}

type SourceGoogleAdsInput interface {
	pulumi.Input

	ToSourceGoogleAdsOutput() SourceGoogleAdsOutput
	ToSourceGoogleAdsOutputWithContext(ctx context.Context) SourceGoogleAdsOutput
}

func (*SourceGoogleAds) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGoogleAds)(nil)).Elem()
}

func (i *SourceGoogleAds) ToSourceGoogleAdsOutput() SourceGoogleAdsOutput {
	return i.ToSourceGoogleAdsOutputWithContext(context.Background())
}

func (i *SourceGoogleAds) ToSourceGoogleAdsOutputWithContext(ctx context.Context) SourceGoogleAdsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleAdsOutput)
}

// SourceGoogleAdsArrayInput is an input type that accepts SourceGoogleAdsArray and SourceGoogleAdsArrayOutput values.
// You can construct a concrete instance of `SourceGoogleAdsArrayInput` via:
//
//	SourceGoogleAdsArray{ SourceGoogleAdsArgs{...} }
type SourceGoogleAdsArrayInput interface {
	pulumi.Input

	ToSourceGoogleAdsArrayOutput() SourceGoogleAdsArrayOutput
	ToSourceGoogleAdsArrayOutputWithContext(context.Context) SourceGoogleAdsArrayOutput
}

type SourceGoogleAdsArray []SourceGoogleAdsInput

func (SourceGoogleAdsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGoogleAds)(nil)).Elem()
}

func (i SourceGoogleAdsArray) ToSourceGoogleAdsArrayOutput() SourceGoogleAdsArrayOutput {
	return i.ToSourceGoogleAdsArrayOutputWithContext(context.Background())
}

func (i SourceGoogleAdsArray) ToSourceGoogleAdsArrayOutputWithContext(ctx context.Context) SourceGoogleAdsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleAdsArrayOutput)
}

// SourceGoogleAdsMapInput is an input type that accepts SourceGoogleAdsMap and SourceGoogleAdsMapOutput values.
// You can construct a concrete instance of `SourceGoogleAdsMapInput` via:
//
//	SourceGoogleAdsMap{ "key": SourceGoogleAdsArgs{...} }
type SourceGoogleAdsMapInput interface {
	pulumi.Input

	ToSourceGoogleAdsMapOutput() SourceGoogleAdsMapOutput
	ToSourceGoogleAdsMapOutputWithContext(context.Context) SourceGoogleAdsMapOutput
}

type SourceGoogleAdsMap map[string]SourceGoogleAdsInput

func (SourceGoogleAdsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGoogleAds)(nil)).Elem()
}

func (i SourceGoogleAdsMap) ToSourceGoogleAdsMapOutput() SourceGoogleAdsMapOutput {
	return i.ToSourceGoogleAdsMapOutputWithContext(context.Background())
}

func (i SourceGoogleAdsMap) ToSourceGoogleAdsMapOutputWithContext(ctx context.Context) SourceGoogleAdsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleAdsMapOutput)
}

type SourceGoogleAdsOutput struct{ *pulumi.OutputState }

func (SourceGoogleAdsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGoogleAds)(nil)).Elem()
}

func (o SourceGoogleAdsOutput) ToSourceGoogleAdsOutput() SourceGoogleAdsOutput {
	return o
}

func (o SourceGoogleAdsOutput) ToSourceGoogleAdsOutputWithContext(ctx context.Context) SourceGoogleAdsOutput {
	return o
}

func (o SourceGoogleAdsOutput) Configuration() SourceGoogleAdsConfigurationOutput {
	return o.ApplyT(func(v *SourceGoogleAds) SourceGoogleAdsConfigurationOutput { return v.Configuration }).(SourceGoogleAdsConfigurationOutput)
}

func (o SourceGoogleAdsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleAds) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceGoogleAdsOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGoogleAds) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceGoogleAdsOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleAds) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceGoogleAdsOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleAds) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceGoogleAdsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleAds) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceGoogleAdsArrayOutput struct{ *pulumi.OutputState }

func (SourceGoogleAdsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGoogleAds)(nil)).Elem()
}

func (o SourceGoogleAdsArrayOutput) ToSourceGoogleAdsArrayOutput() SourceGoogleAdsArrayOutput {
	return o
}

func (o SourceGoogleAdsArrayOutput) ToSourceGoogleAdsArrayOutputWithContext(ctx context.Context) SourceGoogleAdsArrayOutput {
	return o
}

func (o SourceGoogleAdsArrayOutput) Index(i pulumi.IntInput) SourceGoogleAdsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceGoogleAds {
		return vs[0].([]*SourceGoogleAds)[vs[1].(int)]
	}).(SourceGoogleAdsOutput)
}

type SourceGoogleAdsMapOutput struct{ *pulumi.OutputState }

func (SourceGoogleAdsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGoogleAds)(nil)).Elem()
}

func (o SourceGoogleAdsMapOutput) ToSourceGoogleAdsMapOutput() SourceGoogleAdsMapOutput {
	return o
}

func (o SourceGoogleAdsMapOutput) ToSourceGoogleAdsMapOutputWithContext(ctx context.Context) SourceGoogleAdsMapOutput {
	return o
}

func (o SourceGoogleAdsMapOutput) MapIndex(k pulumi.StringInput) SourceGoogleAdsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceGoogleAds {
		return vs[0].(map[string]*SourceGoogleAds)[vs[1].(string)]
	}).(SourceGoogleAdsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleAdsInput)(nil)).Elem(), &SourceGoogleAds{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleAdsArrayInput)(nil)).Elem(), SourceGoogleAdsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleAdsMapInput)(nil)).Elem(), SourceGoogleAdsMap{})
	pulumi.RegisterOutputType(SourceGoogleAdsOutput{})
	pulumi.RegisterOutputType(SourceGoogleAdsArrayOutput{})
	pulumi.RegisterOutputType(SourceGoogleAdsMapOutput{})
}
