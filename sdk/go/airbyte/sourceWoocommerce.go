// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceWoocommerce struct {
	pulumi.CustomResourceState

	Configuration SourceWoocommerceConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceWoocommerce registers a new resource with the given unique name, arguments, and options.
func NewSourceWoocommerce(ctx *pulumi.Context,
	name string, args *SourceWoocommerceArgs, opts ...pulumi.ResourceOption) (*SourceWoocommerce, error) {
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
	opts = pkgResourceDefaultOpts(opts)
	var resource SourceWoocommerce
	err := ctx.RegisterResource("airbyte:index/sourceWoocommerce:SourceWoocommerce", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceWoocommerce gets an existing SourceWoocommerce resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceWoocommerce(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceWoocommerceState, opts ...pulumi.ResourceOption) (*SourceWoocommerce, error) {
	var resource SourceWoocommerce
	err := ctx.ReadResource("airbyte:index/sourceWoocommerce:SourceWoocommerce", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceWoocommerce resources.
type sourceWoocommerceState struct {
	Configuration *SourceWoocommerceConfiguration `pulumi:"configuration"`
	Name          *string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceWoocommerceState struct {
	Configuration SourceWoocommerceConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceWoocommerceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceWoocommerceState)(nil)).Elem()
}

type sourceWoocommerceArgs struct {
	Configuration SourceWoocommerceConfiguration `pulumi:"configuration"`
	Name          string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceWoocommerce resource.
type SourceWoocommerceArgs struct {
	Configuration SourceWoocommerceConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceWoocommerceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceWoocommerceArgs)(nil)).Elem()
}

type SourceWoocommerceInput interface {
	pulumi.Input

	ToSourceWoocommerceOutput() SourceWoocommerceOutput
	ToSourceWoocommerceOutputWithContext(ctx context.Context) SourceWoocommerceOutput
}

func (*SourceWoocommerce) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceWoocommerce)(nil)).Elem()
}

func (i *SourceWoocommerce) ToSourceWoocommerceOutput() SourceWoocommerceOutput {
	return i.ToSourceWoocommerceOutputWithContext(context.Background())
}

func (i *SourceWoocommerce) ToSourceWoocommerceOutputWithContext(ctx context.Context) SourceWoocommerceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceWoocommerceOutput)
}

type SourceWoocommerceOutput struct{ *pulumi.OutputState }

func (SourceWoocommerceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceWoocommerce)(nil)).Elem()
}

func (o SourceWoocommerceOutput) ToSourceWoocommerceOutput() SourceWoocommerceOutput {
	return o
}

func (o SourceWoocommerceOutput) ToSourceWoocommerceOutputWithContext(ctx context.Context) SourceWoocommerceOutput {
	return o
}

func (o SourceWoocommerceOutput) Configuration() SourceWoocommerceConfigurationOutput {
	return o.ApplyT(func(v *SourceWoocommerce) SourceWoocommerceConfigurationOutput { return v.Configuration }).(SourceWoocommerceConfigurationOutput)
}

func (o SourceWoocommerceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceWoocommerce) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceWoocommerceOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceWoocommerce) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceWoocommerceOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceWoocommerce) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceWoocommerceOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceWoocommerce) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceWoocommerceOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceWoocommerce) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceWoocommerceInput)(nil)).Elem(), &SourceWoocommerce{})
	pulumi.RegisterOutputType(SourceWoocommerceOutput{})
}
