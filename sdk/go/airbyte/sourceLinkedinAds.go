// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceLinkedinAds struct {
	pulumi.CustomResourceState

	Configuration SourceLinkedinAdsConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceLinkedinAds registers a new resource with the given unique name, arguments, and options.
func NewSourceLinkedinAds(ctx *pulumi.Context,
	name string, args *SourceLinkedinAdsArgs, opts ...pulumi.ResourceOption) (*SourceLinkedinAds, error) {
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
	var resource SourceLinkedinAds
	err := ctx.RegisterResource("airbyte:index/sourceLinkedinAds:SourceLinkedinAds", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceLinkedinAds gets an existing SourceLinkedinAds resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceLinkedinAds(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceLinkedinAdsState, opts ...pulumi.ResourceOption) (*SourceLinkedinAds, error) {
	var resource SourceLinkedinAds
	err := ctx.ReadResource("airbyte:index/sourceLinkedinAds:SourceLinkedinAds", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceLinkedinAds resources.
type sourceLinkedinAdsState struct {
	Configuration *SourceLinkedinAdsConfiguration `pulumi:"configuration"`
	Name          *string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceLinkedinAdsState struct {
	Configuration SourceLinkedinAdsConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceLinkedinAdsState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLinkedinAdsState)(nil)).Elem()
}

type sourceLinkedinAdsArgs struct {
	Configuration SourceLinkedinAdsConfiguration `pulumi:"configuration"`
	Name          string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceLinkedinAds resource.
type SourceLinkedinAdsArgs struct {
	Configuration SourceLinkedinAdsConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceLinkedinAdsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLinkedinAdsArgs)(nil)).Elem()
}

type SourceLinkedinAdsInput interface {
	pulumi.Input

	ToSourceLinkedinAdsOutput() SourceLinkedinAdsOutput
	ToSourceLinkedinAdsOutputWithContext(ctx context.Context) SourceLinkedinAdsOutput
}

func (*SourceLinkedinAds) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLinkedinAds)(nil)).Elem()
}

func (i *SourceLinkedinAds) ToSourceLinkedinAdsOutput() SourceLinkedinAdsOutput {
	return i.ToSourceLinkedinAdsOutputWithContext(context.Background())
}

func (i *SourceLinkedinAds) ToSourceLinkedinAdsOutputWithContext(ctx context.Context) SourceLinkedinAdsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLinkedinAdsOutput)
}

type SourceLinkedinAdsOutput struct{ *pulumi.OutputState }

func (SourceLinkedinAdsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLinkedinAds)(nil)).Elem()
}

func (o SourceLinkedinAdsOutput) ToSourceLinkedinAdsOutput() SourceLinkedinAdsOutput {
	return o
}

func (o SourceLinkedinAdsOutput) ToSourceLinkedinAdsOutputWithContext(ctx context.Context) SourceLinkedinAdsOutput {
	return o
}

func (o SourceLinkedinAdsOutput) Configuration() SourceLinkedinAdsConfigurationOutput {
	return o.ApplyT(func(v *SourceLinkedinAds) SourceLinkedinAdsConfigurationOutput { return v.Configuration }).(SourceLinkedinAdsConfigurationOutput)
}

func (o SourceLinkedinAdsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinkedinAds) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceLinkedinAdsOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLinkedinAds) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceLinkedinAdsOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinkedinAds) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceLinkedinAdsOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinkedinAds) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceLinkedinAdsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinkedinAds) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLinkedinAdsInput)(nil)).Elem(), &SourceLinkedinAds{})
	pulumi.RegisterOutputType(SourceLinkedinAdsOutput{})
}