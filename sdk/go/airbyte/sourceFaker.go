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

// SourceFaker Resource
type SourceFaker struct {
	pulumi.CustomResourceState

	Configuration SourceFakerConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceFaker registers a new resource with the given unique name, arguments, and options.
func NewSourceFaker(ctx *pulumi.Context,
	name string, args *SourceFakerArgs, opts ...pulumi.ResourceOption) (*SourceFaker, error) {
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
	var resource SourceFaker
	err := ctx.RegisterResource("airbyte:index/sourceFaker:SourceFaker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceFaker gets an existing SourceFaker resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceFaker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceFakerState, opts ...pulumi.ResourceOption) (*SourceFaker, error) {
	var resource SourceFaker
	err := ctx.ReadResource("airbyte:index/sourceFaker:SourceFaker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceFaker resources.
type sourceFakerState struct {
	Configuration *SourceFakerConfiguration `pulumi:"configuration"`
	Name          *string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceFakerState struct {
	Configuration SourceFakerConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceFakerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFakerState)(nil)).Elem()
}

type sourceFakerArgs struct {
	Configuration SourceFakerConfiguration `pulumi:"configuration"`
	Name          string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceFaker resource.
type SourceFakerArgs struct {
	Configuration SourceFakerConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceFakerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFakerArgs)(nil)).Elem()
}

type SourceFakerInput interface {
	pulumi.Input

	ToSourceFakerOutput() SourceFakerOutput
	ToSourceFakerOutputWithContext(ctx context.Context) SourceFakerOutput
}

func (*SourceFaker) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFaker)(nil)).Elem()
}

func (i *SourceFaker) ToSourceFakerOutput() SourceFakerOutput {
	return i.ToSourceFakerOutputWithContext(context.Background())
}

func (i *SourceFaker) ToSourceFakerOutputWithContext(ctx context.Context) SourceFakerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFakerOutput)
}

type SourceFakerOutput struct{ *pulumi.OutputState }

func (SourceFakerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFaker)(nil)).Elem()
}

func (o SourceFakerOutput) ToSourceFakerOutput() SourceFakerOutput {
	return o
}

func (o SourceFakerOutput) ToSourceFakerOutputWithContext(ctx context.Context) SourceFakerOutput {
	return o
}

func (o SourceFakerOutput) Configuration() SourceFakerConfigurationOutput {
	return o.ApplyT(func(v *SourceFaker) SourceFakerConfigurationOutput { return v.Configuration }).(SourceFakerConfigurationOutput)
}

func (o SourceFakerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFaker) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceFakerOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceFaker) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceFakerOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFaker) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceFakerOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFaker) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceFakerOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFaker) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFakerInput)(nil)).Elem(), &SourceFaker{})
	pulumi.RegisterOutputType(SourceFakerOutput{})
}
