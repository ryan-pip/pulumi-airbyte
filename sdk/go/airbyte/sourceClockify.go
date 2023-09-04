// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceClockify struct {
	pulumi.CustomResourceState

	Configuration SourceClockifyConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceClockify registers a new resource with the given unique name, arguments, and options.
func NewSourceClockify(ctx *pulumi.Context,
	name string, args *SourceClockifyArgs, opts ...pulumi.ResourceOption) (*SourceClockify, error) {
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
	var resource SourceClockify
	err := ctx.RegisterResource("airbyte:index/sourceClockify:SourceClockify", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceClockify gets an existing SourceClockify resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceClockify(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceClockifyState, opts ...pulumi.ResourceOption) (*SourceClockify, error) {
	var resource SourceClockify
	err := ctx.ReadResource("airbyte:index/sourceClockify:SourceClockify", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceClockify resources.
type sourceClockifyState struct {
	Configuration *SourceClockifyConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceClockifyState struct {
	Configuration SourceClockifyConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceClockifyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceClockifyState)(nil)).Elem()
}

type sourceClockifyArgs struct {
	Configuration SourceClockifyConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceClockify resource.
type SourceClockifyArgs struct {
	Configuration SourceClockifyConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceClockifyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceClockifyArgs)(nil)).Elem()
}

type SourceClockifyInput interface {
	pulumi.Input

	ToSourceClockifyOutput() SourceClockifyOutput
	ToSourceClockifyOutputWithContext(ctx context.Context) SourceClockifyOutput
}

func (*SourceClockify) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceClockify)(nil)).Elem()
}

func (i *SourceClockify) ToSourceClockifyOutput() SourceClockifyOutput {
	return i.ToSourceClockifyOutputWithContext(context.Background())
}

func (i *SourceClockify) ToSourceClockifyOutputWithContext(ctx context.Context) SourceClockifyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceClockifyOutput)
}

type SourceClockifyOutput struct{ *pulumi.OutputState }

func (SourceClockifyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceClockify)(nil)).Elem()
}

func (o SourceClockifyOutput) ToSourceClockifyOutput() SourceClockifyOutput {
	return o
}

func (o SourceClockifyOutput) ToSourceClockifyOutputWithContext(ctx context.Context) SourceClockifyOutput {
	return o
}

func (o SourceClockifyOutput) Configuration() SourceClockifyConfigurationOutput {
	return o.ApplyT(func(v *SourceClockify) SourceClockifyConfigurationOutput { return v.Configuration }).(SourceClockifyConfigurationOutput)
}

func (o SourceClockifyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClockify) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceClockifyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceClockify) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceClockifyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClockify) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceClockifyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClockify) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceClockifyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceClockify) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceClockifyInput)(nil)).Elem(), &SourceClockify{})
	pulumi.RegisterOutputType(SourceClockifyOutput{})
}
