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

// SourceLinnworks Resource
type SourceLinnworks struct {
	pulumi.CustomResourceState

	Configuration SourceLinnworksConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceLinnworks registers a new resource with the given unique name, arguments, and options.
func NewSourceLinnworks(ctx *pulumi.Context,
	name string, args *SourceLinnworksArgs, opts ...pulumi.ResourceOption) (*SourceLinnworks, error) {
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
	var resource SourceLinnworks
	err := ctx.RegisterResource("airbyte:index/sourceLinnworks:SourceLinnworks", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceLinnworks gets an existing SourceLinnworks resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceLinnworks(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceLinnworksState, opts ...pulumi.ResourceOption) (*SourceLinnworks, error) {
	var resource SourceLinnworks
	err := ctx.ReadResource("airbyte:index/sourceLinnworks:SourceLinnworks", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceLinnworks resources.
type sourceLinnworksState struct {
	Configuration *SourceLinnworksConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceLinnworksState struct {
	Configuration SourceLinnworksConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceLinnworksState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLinnworksState)(nil)).Elem()
}

type sourceLinnworksArgs struct {
	Configuration SourceLinnworksConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceLinnworks resource.
type SourceLinnworksArgs struct {
	Configuration SourceLinnworksConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceLinnworksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLinnworksArgs)(nil)).Elem()
}

type SourceLinnworksInput interface {
	pulumi.Input

	ToSourceLinnworksOutput() SourceLinnworksOutput
	ToSourceLinnworksOutputWithContext(ctx context.Context) SourceLinnworksOutput
}

func (*SourceLinnworks) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLinnworks)(nil)).Elem()
}

func (i *SourceLinnworks) ToSourceLinnworksOutput() SourceLinnworksOutput {
	return i.ToSourceLinnworksOutputWithContext(context.Background())
}

func (i *SourceLinnworks) ToSourceLinnworksOutputWithContext(ctx context.Context) SourceLinnworksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLinnworksOutput)
}

type SourceLinnworksOutput struct{ *pulumi.OutputState }

func (SourceLinnworksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLinnworks)(nil)).Elem()
}

func (o SourceLinnworksOutput) ToSourceLinnworksOutput() SourceLinnworksOutput {
	return o
}

func (o SourceLinnworksOutput) ToSourceLinnworksOutputWithContext(ctx context.Context) SourceLinnworksOutput {
	return o
}

func (o SourceLinnworksOutput) Configuration() SourceLinnworksConfigurationOutput {
	return o.ApplyT(func(v *SourceLinnworks) SourceLinnworksConfigurationOutput { return v.Configuration }).(SourceLinnworksConfigurationOutput)
}

func (o SourceLinnworksOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinnworks) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceLinnworksOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLinnworks) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceLinnworksOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinnworks) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceLinnworksOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinnworks) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceLinnworksOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinnworks) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLinnworksInput)(nil)).Elem(), &SourceLinnworks{})
	pulumi.RegisterOutputType(SourceLinnworksOutput{})
}
