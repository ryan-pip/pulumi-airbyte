// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceAmplitude struct {
	pulumi.CustomResourceState

	Configuration SourceAmplitudeConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceAmplitude registers a new resource with the given unique name, arguments, and options.
func NewSourceAmplitude(ctx *pulumi.Context,
	name string, args *SourceAmplitudeArgs, opts ...pulumi.ResourceOption) (*SourceAmplitude, error) {
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
	var resource SourceAmplitude
	err := ctx.RegisterResource("airbyte:index/sourceAmplitude:SourceAmplitude", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceAmplitude gets an existing SourceAmplitude resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceAmplitude(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceAmplitudeState, opts ...pulumi.ResourceOption) (*SourceAmplitude, error) {
	var resource SourceAmplitude
	err := ctx.ReadResource("airbyte:index/sourceAmplitude:SourceAmplitude", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceAmplitude resources.
type sourceAmplitudeState struct {
	Configuration *SourceAmplitudeConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceAmplitudeState struct {
	Configuration SourceAmplitudeConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceAmplitudeState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAmplitudeState)(nil)).Elem()
}

type sourceAmplitudeArgs struct {
	Configuration SourceAmplitudeConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceAmplitude resource.
type SourceAmplitudeArgs struct {
	Configuration SourceAmplitudeConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceAmplitudeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAmplitudeArgs)(nil)).Elem()
}

type SourceAmplitudeInput interface {
	pulumi.Input

	ToSourceAmplitudeOutput() SourceAmplitudeOutput
	ToSourceAmplitudeOutputWithContext(ctx context.Context) SourceAmplitudeOutput
}

func (*SourceAmplitude) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAmplitude)(nil)).Elem()
}

func (i *SourceAmplitude) ToSourceAmplitudeOutput() SourceAmplitudeOutput {
	return i.ToSourceAmplitudeOutputWithContext(context.Background())
}

func (i *SourceAmplitude) ToSourceAmplitudeOutputWithContext(ctx context.Context) SourceAmplitudeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAmplitudeOutput)
}

type SourceAmplitudeOutput struct{ *pulumi.OutputState }

func (SourceAmplitudeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAmplitude)(nil)).Elem()
}

func (o SourceAmplitudeOutput) ToSourceAmplitudeOutput() SourceAmplitudeOutput {
	return o
}

func (o SourceAmplitudeOutput) ToSourceAmplitudeOutputWithContext(ctx context.Context) SourceAmplitudeOutput {
	return o
}

func (o SourceAmplitudeOutput) Configuration() SourceAmplitudeConfigurationOutput {
	return o.ApplyT(func(v *SourceAmplitude) SourceAmplitudeConfigurationOutput { return v.Configuration }).(SourceAmplitudeConfigurationOutput)
}

func (o SourceAmplitudeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmplitude) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceAmplitudeOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAmplitude) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceAmplitudeOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmplitude) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceAmplitudeOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmplitude) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceAmplitudeOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmplitude) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAmplitudeInput)(nil)).Elem(), &SourceAmplitude{})
	pulumi.RegisterOutputType(SourceAmplitudeOutput{})
}