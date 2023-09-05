// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceBigcommerce Resource
type SourceBigcommerce struct {
	pulumi.CustomResourceState

	Configuration SourceBigcommerceConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceBigcommerce registers a new resource with the given unique name, arguments, and options.
func NewSourceBigcommerce(ctx *pulumi.Context,
	name string, args *SourceBigcommerceArgs, opts ...pulumi.ResourceOption) (*SourceBigcommerce, error) {
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
	var resource SourceBigcommerce
	err := ctx.RegisterResource("airbyte:index/sourceBigcommerce:SourceBigcommerce", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceBigcommerce gets an existing SourceBigcommerce resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceBigcommerce(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceBigcommerceState, opts ...pulumi.ResourceOption) (*SourceBigcommerce, error) {
	var resource SourceBigcommerce
	err := ctx.ReadResource("airbyte:index/sourceBigcommerce:SourceBigcommerce", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceBigcommerce resources.
type sourceBigcommerceState struct {
	Configuration *SourceBigcommerceConfiguration `pulumi:"configuration"`
	Name          *string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceBigcommerceState struct {
	Configuration SourceBigcommerceConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceBigcommerceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceBigcommerceState)(nil)).Elem()
}

type sourceBigcommerceArgs struct {
	Configuration SourceBigcommerceConfiguration `pulumi:"configuration"`
	Name          string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceBigcommerce resource.
type SourceBigcommerceArgs struct {
	Configuration SourceBigcommerceConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceBigcommerceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceBigcommerceArgs)(nil)).Elem()
}

type SourceBigcommerceInput interface {
	pulumi.Input

	ToSourceBigcommerceOutput() SourceBigcommerceOutput
	ToSourceBigcommerceOutputWithContext(ctx context.Context) SourceBigcommerceOutput
}

func (*SourceBigcommerce) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceBigcommerce)(nil)).Elem()
}

func (i *SourceBigcommerce) ToSourceBigcommerceOutput() SourceBigcommerceOutput {
	return i.ToSourceBigcommerceOutputWithContext(context.Background())
}

func (i *SourceBigcommerce) ToSourceBigcommerceOutputWithContext(ctx context.Context) SourceBigcommerceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceBigcommerceOutput)
}

type SourceBigcommerceOutput struct{ *pulumi.OutputState }

func (SourceBigcommerceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceBigcommerce)(nil)).Elem()
}

func (o SourceBigcommerceOutput) ToSourceBigcommerceOutput() SourceBigcommerceOutput {
	return o
}

func (o SourceBigcommerceOutput) ToSourceBigcommerceOutputWithContext(ctx context.Context) SourceBigcommerceOutput {
	return o
}

func (o SourceBigcommerceOutput) Configuration() SourceBigcommerceConfigurationOutput {
	return o.ApplyT(func(v *SourceBigcommerce) SourceBigcommerceConfigurationOutput { return v.Configuration }).(SourceBigcommerceConfigurationOutput)
}

func (o SourceBigcommerceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBigcommerce) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceBigcommerceOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceBigcommerce) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceBigcommerceOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBigcommerce) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceBigcommerceOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBigcommerce) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceBigcommerceOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBigcommerce) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceBigcommerceInput)(nil)).Elem(), &SourceBigcommerce{})
	pulumi.RegisterOutputType(SourceBigcommerceOutput{})
}
