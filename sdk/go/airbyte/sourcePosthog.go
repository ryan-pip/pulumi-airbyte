// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourcePosthog struct {
	pulumi.CustomResourceState

	Configuration SourcePosthogConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourcePosthog registers a new resource with the given unique name, arguments, and options.
func NewSourcePosthog(ctx *pulumi.Context,
	name string, args *SourcePosthogArgs, opts ...pulumi.ResourceOption) (*SourcePosthog, error) {
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
	var resource SourcePosthog
	err := ctx.RegisterResource("airbyte:index/sourcePosthog:SourcePosthog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourcePosthog gets an existing SourcePosthog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourcePosthog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourcePosthogState, opts ...pulumi.ResourceOption) (*SourcePosthog, error) {
	var resource SourcePosthog
	err := ctx.ReadResource("airbyte:index/sourcePosthog:SourcePosthog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourcePosthog resources.
type sourcePosthogState struct {
	Configuration *SourcePosthogConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourcePosthogState struct {
	Configuration SourcePosthogConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourcePosthogState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePosthogState)(nil)).Elem()
}

type sourcePosthogArgs struct {
	Configuration SourcePosthogConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourcePosthog resource.
type SourcePosthogArgs struct {
	Configuration SourcePosthogConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourcePosthogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePosthogArgs)(nil)).Elem()
}

type SourcePosthogInput interface {
	pulumi.Input

	ToSourcePosthogOutput() SourcePosthogOutput
	ToSourcePosthogOutputWithContext(ctx context.Context) SourcePosthogOutput
}

func (*SourcePosthog) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePosthog)(nil)).Elem()
}

func (i *SourcePosthog) ToSourcePosthogOutput() SourcePosthogOutput {
	return i.ToSourcePosthogOutputWithContext(context.Background())
}

func (i *SourcePosthog) ToSourcePosthogOutputWithContext(ctx context.Context) SourcePosthogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePosthogOutput)
}

type SourcePosthogOutput struct{ *pulumi.OutputState }

func (SourcePosthogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePosthog)(nil)).Elem()
}

func (o SourcePosthogOutput) ToSourcePosthogOutput() SourcePosthogOutput {
	return o
}

func (o SourcePosthogOutput) ToSourcePosthogOutputWithContext(ctx context.Context) SourcePosthogOutput {
	return o
}

func (o SourcePosthogOutput) Configuration() SourcePosthogConfigurationOutput {
	return o.ApplyT(func(v *SourcePosthog) SourcePosthogConfigurationOutput { return v.Configuration }).(SourcePosthogConfigurationOutput)
}

func (o SourcePosthogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePosthog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourcePosthogOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePosthog) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourcePosthogOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePosthog) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourcePosthogOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePosthog) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourcePosthogOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePosthog) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePosthogInput)(nil)).Elem(), &SourcePosthog{})
	pulumi.RegisterOutputType(SourcePosthogOutput{})
}
