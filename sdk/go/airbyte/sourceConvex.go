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

// SourceConvex Resource
type SourceConvex struct {
	pulumi.CustomResourceState

	Configuration SourceConvexConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceConvex registers a new resource with the given unique name, arguments, and options.
func NewSourceConvex(ctx *pulumi.Context,
	name string, args *SourceConvexArgs, opts ...pulumi.ResourceOption) (*SourceConvex, error) {
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
	var resource SourceConvex
	err := ctx.RegisterResource("airbyte:index/sourceConvex:SourceConvex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceConvex gets an existing SourceConvex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceConvex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceConvexState, opts ...pulumi.ResourceOption) (*SourceConvex, error) {
	var resource SourceConvex
	err := ctx.ReadResource("airbyte:index/sourceConvex:SourceConvex", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceConvex resources.
type sourceConvexState struct {
	Configuration *SourceConvexConfiguration `pulumi:"configuration"`
	Name          *string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceConvexState struct {
	Configuration SourceConvexConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceConvexState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceConvexState)(nil)).Elem()
}

type sourceConvexArgs struct {
	Configuration SourceConvexConfiguration `pulumi:"configuration"`
	Name          string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceConvex resource.
type SourceConvexArgs struct {
	Configuration SourceConvexConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceConvexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceConvexArgs)(nil)).Elem()
}

type SourceConvexInput interface {
	pulumi.Input

	ToSourceConvexOutput() SourceConvexOutput
	ToSourceConvexOutputWithContext(ctx context.Context) SourceConvexOutput
}

func (*SourceConvex) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceConvex)(nil)).Elem()
}

func (i *SourceConvex) ToSourceConvexOutput() SourceConvexOutput {
	return i.ToSourceConvexOutputWithContext(context.Background())
}

func (i *SourceConvex) ToSourceConvexOutputWithContext(ctx context.Context) SourceConvexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceConvexOutput)
}

type SourceConvexOutput struct{ *pulumi.OutputState }

func (SourceConvexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceConvex)(nil)).Elem()
}

func (o SourceConvexOutput) ToSourceConvexOutput() SourceConvexOutput {
	return o
}

func (o SourceConvexOutput) ToSourceConvexOutputWithContext(ctx context.Context) SourceConvexOutput {
	return o
}

func (o SourceConvexOutput) Configuration() SourceConvexConfigurationOutput {
	return o.ApplyT(func(v *SourceConvex) SourceConvexConfigurationOutput { return v.Configuration }).(SourceConvexConfigurationOutput)
}

func (o SourceConvexOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceConvex) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceConvexOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceConvex) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceConvexOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceConvex) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceConvexOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceConvex) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceConvexOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceConvex) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceConvexInput)(nil)).Elem(), &SourceConvex{})
	pulumi.RegisterOutputType(SourceConvexOutput{})
}
