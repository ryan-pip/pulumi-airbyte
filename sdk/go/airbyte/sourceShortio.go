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

// SourceShortio Resource
type SourceShortio struct {
	pulumi.CustomResourceState

	Configuration SourceShortioConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceShortio registers a new resource with the given unique name, arguments, and options.
func NewSourceShortio(ctx *pulumi.Context,
	name string, args *SourceShortioArgs, opts ...pulumi.ResourceOption) (*SourceShortio, error) {
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
	var resource SourceShortio
	err := ctx.RegisterResource("airbyte:index/sourceShortio:SourceShortio", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceShortio gets an existing SourceShortio resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceShortio(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceShortioState, opts ...pulumi.ResourceOption) (*SourceShortio, error) {
	var resource SourceShortio
	err := ctx.ReadResource("airbyte:index/sourceShortio:SourceShortio", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceShortio resources.
type sourceShortioState struct {
	Configuration *SourceShortioConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceShortioState struct {
	Configuration SourceShortioConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceShortioState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceShortioState)(nil)).Elem()
}

type sourceShortioArgs struct {
	Configuration SourceShortioConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceShortio resource.
type SourceShortioArgs struct {
	Configuration SourceShortioConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceShortioArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceShortioArgs)(nil)).Elem()
}

type SourceShortioInput interface {
	pulumi.Input

	ToSourceShortioOutput() SourceShortioOutput
	ToSourceShortioOutputWithContext(ctx context.Context) SourceShortioOutput
}

func (*SourceShortio) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceShortio)(nil)).Elem()
}

func (i *SourceShortio) ToSourceShortioOutput() SourceShortioOutput {
	return i.ToSourceShortioOutputWithContext(context.Background())
}

func (i *SourceShortio) ToSourceShortioOutputWithContext(ctx context.Context) SourceShortioOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceShortioOutput)
}

type SourceShortioOutput struct{ *pulumi.OutputState }

func (SourceShortioOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceShortio)(nil)).Elem()
}

func (o SourceShortioOutput) ToSourceShortioOutput() SourceShortioOutput {
	return o
}

func (o SourceShortioOutput) ToSourceShortioOutputWithContext(ctx context.Context) SourceShortioOutput {
	return o
}

func (o SourceShortioOutput) Configuration() SourceShortioConfigurationOutput {
	return o.ApplyT(func(v *SourceShortio) SourceShortioConfigurationOutput { return v.Configuration }).(SourceShortioConfigurationOutput)
}

func (o SourceShortioOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceShortio) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceShortioOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceShortio) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceShortioOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceShortio) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceShortioOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceShortio) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceShortioOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceShortio) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceShortioInput)(nil)).Elem(), &SourceShortio{})
	pulumi.RegisterOutputType(SourceShortioOutput{})
}
