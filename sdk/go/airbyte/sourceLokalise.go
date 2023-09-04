// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceLokalise struct {
	pulumi.CustomResourceState

	Configuration SourceLokaliseConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceLokalise registers a new resource with the given unique name, arguments, and options.
func NewSourceLokalise(ctx *pulumi.Context,
	name string, args *SourceLokaliseArgs, opts ...pulumi.ResourceOption) (*SourceLokalise, error) {
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
	var resource SourceLokalise
	err := ctx.RegisterResource("airbyte:index/sourceLokalise:SourceLokalise", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceLokalise gets an existing SourceLokalise resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceLokalise(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceLokaliseState, opts ...pulumi.ResourceOption) (*SourceLokalise, error) {
	var resource SourceLokalise
	err := ctx.ReadResource("airbyte:index/sourceLokalise:SourceLokalise", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceLokalise resources.
type sourceLokaliseState struct {
	Configuration *SourceLokaliseConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceLokaliseState struct {
	Configuration SourceLokaliseConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceLokaliseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLokaliseState)(nil)).Elem()
}

type sourceLokaliseArgs struct {
	Configuration SourceLokaliseConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceLokalise resource.
type SourceLokaliseArgs struct {
	Configuration SourceLokaliseConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceLokaliseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLokaliseArgs)(nil)).Elem()
}

type SourceLokaliseInput interface {
	pulumi.Input

	ToSourceLokaliseOutput() SourceLokaliseOutput
	ToSourceLokaliseOutputWithContext(ctx context.Context) SourceLokaliseOutput
}

func (*SourceLokalise) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLokalise)(nil)).Elem()
}

func (i *SourceLokalise) ToSourceLokaliseOutput() SourceLokaliseOutput {
	return i.ToSourceLokaliseOutputWithContext(context.Background())
}

func (i *SourceLokalise) ToSourceLokaliseOutputWithContext(ctx context.Context) SourceLokaliseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLokaliseOutput)
}

type SourceLokaliseOutput struct{ *pulumi.OutputState }

func (SourceLokaliseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLokalise)(nil)).Elem()
}

func (o SourceLokaliseOutput) ToSourceLokaliseOutput() SourceLokaliseOutput {
	return o
}

func (o SourceLokaliseOutput) ToSourceLokaliseOutputWithContext(ctx context.Context) SourceLokaliseOutput {
	return o
}

func (o SourceLokaliseOutput) Configuration() SourceLokaliseConfigurationOutput {
	return o.ApplyT(func(v *SourceLokalise) SourceLokaliseConfigurationOutput { return v.Configuration }).(SourceLokaliseConfigurationOutput)
}

func (o SourceLokaliseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLokalise) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceLokaliseOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLokalise) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceLokaliseOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLokalise) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceLokaliseOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLokalise) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceLokaliseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLokalise) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLokaliseInput)(nil)).Elem(), &SourceLokalise{})
	pulumi.RegisterOutputType(SourceLokaliseOutput{})
}
