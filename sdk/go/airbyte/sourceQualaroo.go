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

// SourceQualaroo Resource
type SourceQualaroo struct {
	pulumi.CustomResourceState

	Configuration SourceQualarooConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceQualaroo registers a new resource with the given unique name, arguments, and options.
func NewSourceQualaroo(ctx *pulumi.Context,
	name string, args *SourceQualarooArgs, opts ...pulumi.ResourceOption) (*SourceQualaroo, error) {
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
	var resource SourceQualaroo
	err := ctx.RegisterResource("airbyte:index/sourceQualaroo:SourceQualaroo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceQualaroo gets an existing SourceQualaroo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceQualaroo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceQualarooState, opts ...pulumi.ResourceOption) (*SourceQualaroo, error) {
	var resource SourceQualaroo
	err := ctx.ReadResource("airbyte:index/sourceQualaroo:SourceQualaroo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceQualaroo resources.
type sourceQualarooState struct {
	Configuration *SourceQualarooConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceQualarooState struct {
	Configuration SourceQualarooConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceQualarooState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceQualarooState)(nil)).Elem()
}

type sourceQualarooArgs struct {
	Configuration SourceQualarooConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceQualaroo resource.
type SourceQualarooArgs struct {
	Configuration SourceQualarooConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceQualarooArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceQualarooArgs)(nil)).Elem()
}

type SourceQualarooInput interface {
	pulumi.Input

	ToSourceQualarooOutput() SourceQualarooOutput
	ToSourceQualarooOutputWithContext(ctx context.Context) SourceQualarooOutput
}

func (*SourceQualaroo) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceQualaroo)(nil)).Elem()
}

func (i *SourceQualaroo) ToSourceQualarooOutput() SourceQualarooOutput {
	return i.ToSourceQualarooOutputWithContext(context.Background())
}

func (i *SourceQualaroo) ToSourceQualarooOutputWithContext(ctx context.Context) SourceQualarooOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceQualarooOutput)
}

type SourceQualarooOutput struct{ *pulumi.OutputState }

func (SourceQualarooOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceQualaroo)(nil)).Elem()
}

func (o SourceQualarooOutput) ToSourceQualarooOutput() SourceQualarooOutput {
	return o
}

func (o SourceQualarooOutput) ToSourceQualarooOutputWithContext(ctx context.Context) SourceQualarooOutput {
	return o
}

func (o SourceQualarooOutput) Configuration() SourceQualarooConfigurationOutput {
	return o.ApplyT(func(v *SourceQualaroo) SourceQualarooConfigurationOutput { return v.Configuration }).(SourceQualarooConfigurationOutput)
}

func (o SourceQualarooOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceQualaroo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceQualarooOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceQualaroo) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceQualarooOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceQualaroo) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceQualarooOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceQualaroo) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceQualarooOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceQualaroo) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceQualarooInput)(nil)).Elem(), &SourceQualaroo{})
	pulumi.RegisterOutputType(SourceQualarooOutput{})
}
