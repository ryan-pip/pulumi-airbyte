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

// SourceOutbrainAmplify Resource
type SourceOutbrainAmplify struct {
	pulumi.CustomResourceState

	Configuration SourceOutbrainAmplifyConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceOutbrainAmplify registers a new resource with the given unique name, arguments, and options.
func NewSourceOutbrainAmplify(ctx *pulumi.Context,
	name string, args *SourceOutbrainAmplifyArgs, opts ...pulumi.ResourceOption) (*SourceOutbrainAmplify, error) {
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
	var resource SourceOutbrainAmplify
	err := ctx.RegisterResource("airbyte:index/sourceOutbrainAmplify:SourceOutbrainAmplify", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceOutbrainAmplify gets an existing SourceOutbrainAmplify resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceOutbrainAmplify(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceOutbrainAmplifyState, opts ...pulumi.ResourceOption) (*SourceOutbrainAmplify, error) {
	var resource SourceOutbrainAmplify
	err := ctx.ReadResource("airbyte:index/sourceOutbrainAmplify:SourceOutbrainAmplify", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceOutbrainAmplify resources.
type sourceOutbrainAmplifyState struct {
	Configuration *SourceOutbrainAmplifyConfiguration `pulumi:"configuration"`
	Name          *string                             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceOutbrainAmplifyState struct {
	Configuration SourceOutbrainAmplifyConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceOutbrainAmplifyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOutbrainAmplifyState)(nil)).Elem()
}

type sourceOutbrainAmplifyArgs struct {
	Configuration SourceOutbrainAmplifyConfiguration `pulumi:"configuration"`
	Name          string                             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceOutbrainAmplify resource.
type SourceOutbrainAmplifyArgs struct {
	Configuration SourceOutbrainAmplifyConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceOutbrainAmplifyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOutbrainAmplifyArgs)(nil)).Elem()
}

type SourceOutbrainAmplifyInput interface {
	pulumi.Input

	ToSourceOutbrainAmplifyOutput() SourceOutbrainAmplifyOutput
	ToSourceOutbrainAmplifyOutputWithContext(ctx context.Context) SourceOutbrainAmplifyOutput
}

func (*SourceOutbrainAmplify) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOutbrainAmplify)(nil)).Elem()
}

func (i *SourceOutbrainAmplify) ToSourceOutbrainAmplifyOutput() SourceOutbrainAmplifyOutput {
	return i.ToSourceOutbrainAmplifyOutputWithContext(context.Background())
}

func (i *SourceOutbrainAmplify) ToSourceOutbrainAmplifyOutputWithContext(ctx context.Context) SourceOutbrainAmplifyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutbrainAmplifyOutput)
}

type SourceOutbrainAmplifyOutput struct{ *pulumi.OutputState }

func (SourceOutbrainAmplifyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOutbrainAmplify)(nil)).Elem()
}

func (o SourceOutbrainAmplifyOutput) ToSourceOutbrainAmplifyOutput() SourceOutbrainAmplifyOutput {
	return o
}

func (o SourceOutbrainAmplifyOutput) ToSourceOutbrainAmplifyOutputWithContext(ctx context.Context) SourceOutbrainAmplifyOutput {
	return o
}

func (o SourceOutbrainAmplifyOutput) Configuration() SourceOutbrainAmplifyConfigurationOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) SourceOutbrainAmplifyConfigurationOutput { return v.Configuration }).(SourceOutbrainAmplifyConfigurationOutput)
}

func (o SourceOutbrainAmplifyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceOutbrainAmplifyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceOutbrainAmplifyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceOutbrainAmplifyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceOutbrainAmplifyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOutbrainAmplifyInput)(nil)).Elem(), &SourceOutbrainAmplify{})
	pulumi.RegisterOutputType(SourceOutbrainAmplifyOutput{})
}
