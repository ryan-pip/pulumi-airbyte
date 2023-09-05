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

// SourceSmartengage Resource
type SourceSmartengage struct {
	pulumi.CustomResourceState

	Configuration SourceSmartengageConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceSmartengage registers a new resource with the given unique name, arguments, and options.
func NewSourceSmartengage(ctx *pulumi.Context,
	name string, args *SourceSmartengageArgs, opts ...pulumi.ResourceOption) (*SourceSmartengage, error) {
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
	var resource SourceSmartengage
	err := ctx.RegisterResource("airbyte:index/sourceSmartengage:SourceSmartengage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSmartengage gets an existing SourceSmartengage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSmartengage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSmartengageState, opts ...pulumi.ResourceOption) (*SourceSmartengage, error) {
	var resource SourceSmartengage
	err := ctx.ReadResource("airbyte:index/sourceSmartengage:SourceSmartengage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSmartengage resources.
type sourceSmartengageState struct {
	Configuration *SourceSmartengageConfiguration `pulumi:"configuration"`
	Name          *string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceSmartengageState struct {
	Configuration SourceSmartengageConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceSmartengageState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSmartengageState)(nil)).Elem()
}

type sourceSmartengageArgs struct {
	Configuration SourceSmartengageConfiguration `pulumi:"configuration"`
	Name          string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSmartengage resource.
type SourceSmartengageArgs struct {
	Configuration SourceSmartengageConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSmartengageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSmartengageArgs)(nil)).Elem()
}

type SourceSmartengageInput interface {
	pulumi.Input

	ToSourceSmartengageOutput() SourceSmartengageOutput
	ToSourceSmartengageOutputWithContext(ctx context.Context) SourceSmartengageOutput
}

func (*SourceSmartengage) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSmartengage)(nil)).Elem()
}

func (i *SourceSmartengage) ToSourceSmartengageOutput() SourceSmartengageOutput {
	return i.ToSourceSmartengageOutputWithContext(context.Background())
}

func (i *SourceSmartengage) ToSourceSmartengageOutputWithContext(ctx context.Context) SourceSmartengageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSmartengageOutput)
}

type SourceSmartengageOutput struct{ *pulumi.OutputState }

func (SourceSmartengageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSmartengage)(nil)).Elem()
}

func (o SourceSmartengageOutput) ToSourceSmartengageOutput() SourceSmartengageOutput {
	return o
}

func (o SourceSmartengageOutput) ToSourceSmartengageOutputWithContext(ctx context.Context) SourceSmartengageOutput {
	return o
}

func (o SourceSmartengageOutput) Configuration() SourceSmartengageConfigurationOutput {
	return o.ApplyT(func(v *SourceSmartengage) SourceSmartengageConfigurationOutput { return v.Configuration }).(SourceSmartengageConfigurationOutput)
}

func (o SourceSmartengageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmartengage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceSmartengageOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSmartengage) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSmartengageOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmartengage) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSmartengageOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmartengage) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSmartengageOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmartengage) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSmartengageInput)(nil)).Elem(), &SourceSmartengage{})
	pulumi.RegisterOutputType(SourceSmartengageOutput{})
}
