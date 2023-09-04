// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceIterable struct {
	pulumi.CustomResourceState

	Configuration SourceIterableConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceIterable registers a new resource with the given unique name, arguments, and options.
func NewSourceIterable(ctx *pulumi.Context,
	name string, args *SourceIterableArgs, opts ...pulumi.ResourceOption) (*SourceIterable, error) {
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
	var resource SourceIterable
	err := ctx.RegisterResource("airbyte:index/sourceIterable:SourceIterable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceIterable gets an existing SourceIterable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceIterable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceIterableState, opts ...pulumi.ResourceOption) (*SourceIterable, error) {
	var resource SourceIterable
	err := ctx.ReadResource("airbyte:index/sourceIterable:SourceIterable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceIterable resources.
type sourceIterableState struct {
	Configuration *SourceIterableConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceIterableState struct {
	Configuration SourceIterableConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceIterableState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceIterableState)(nil)).Elem()
}

type sourceIterableArgs struct {
	Configuration SourceIterableConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceIterable resource.
type SourceIterableArgs struct {
	Configuration SourceIterableConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceIterableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceIterableArgs)(nil)).Elem()
}

type SourceIterableInput interface {
	pulumi.Input

	ToSourceIterableOutput() SourceIterableOutput
	ToSourceIterableOutputWithContext(ctx context.Context) SourceIterableOutput
}

func (*SourceIterable) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceIterable)(nil)).Elem()
}

func (i *SourceIterable) ToSourceIterableOutput() SourceIterableOutput {
	return i.ToSourceIterableOutputWithContext(context.Background())
}

func (i *SourceIterable) ToSourceIterableOutputWithContext(ctx context.Context) SourceIterableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceIterableOutput)
}

type SourceIterableOutput struct{ *pulumi.OutputState }

func (SourceIterableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceIterable)(nil)).Elem()
}

func (o SourceIterableOutput) ToSourceIterableOutput() SourceIterableOutput {
	return o
}

func (o SourceIterableOutput) ToSourceIterableOutputWithContext(ctx context.Context) SourceIterableOutput {
	return o
}

func (o SourceIterableOutput) Configuration() SourceIterableConfigurationOutput {
	return o.ApplyT(func(v *SourceIterable) SourceIterableConfigurationOutput { return v.Configuration }).(SourceIterableConfigurationOutput)
}

func (o SourceIterableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceIterable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceIterableOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceIterable) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceIterableOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceIterable) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceIterableOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceIterable) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceIterableOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceIterable) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceIterableInput)(nil)).Elem(), &SourceIterable{})
	pulumi.RegisterOutputType(SourceIterableOutput{})
}