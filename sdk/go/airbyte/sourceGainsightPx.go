// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceGainsightPx struct {
	pulumi.CustomResourceState

	Configuration SourceGainsightPxConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceGainsightPx registers a new resource with the given unique name, arguments, and options.
func NewSourceGainsightPx(ctx *pulumi.Context,
	name string, args *SourceGainsightPxArgs, opts ...pulumi.ResourceOption) (*SourceGainsightPx, error) {
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
	var resource SourceGainsightPx
	err := ctx.RegisterResource("airbyte:index/sourceGainsightPx:SourceGainsightPx", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceGainsightPx gets an existing SourceGainsightPx resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceGainsightPx(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceGainsightPxState, opts ...pulumi.ResourceOption) (*SourceGainsightPx, error) {
	var resource SourceGainsightPx
	err := ctx.ReadResource("airbyte:index/sourceGainsightPx:SourceGainsightPx", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceGainsightPx resources.
type sourceGainsightPxState struct {
	Configuration *SourceGainsightPxConfiguration `pulumi:"configuration"`
	Name          *string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceGainsightPxState struct {
	Configuration SourceGainsightPxConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceGainsightPxState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGainsightPxState)(nil)).Elem()
}

type sourceGainsightPxArgs struct {
	Configuration SourceGainsightPxConfiguration `pulumi:"configuration"`
	Name          string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceGainsightPx resource.
type SourceGainsightPxArgs struct {
	Configuration SourceGainsightPxConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceGainsightPxArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGainsightPxArgs)(nil)).Elem()
}

type SourceGainsightPxInput interface {
	pulumi.Input

	ToSourceGainsightPxOutput() SourceGainsightPxOutput
	ToSourceGainsightPxOutputWithContext(ctx context.Context) SourceGainsightPxOutput
}

func (*SourceGainsightPx) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGainsightPx)(nil)).Elem()
}

func (i *SourceGainsightPx) ToSourceGainsightPxOutput() SourceGainsightPxOutput {
	return i.ToSourceGainsightPxOutputWithContext(context.Background())
}

func (i *SourceGainsightPx) ToSourceGainsightPxOutputWithContext(ctx context.Context) SourceGainsightPxOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGainsightPxOutput)
}

type SourceGainsightPxOutput struct{ *pulumi.OutputState }

func (SourceGainsightPxOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGainsightPx)(nil)).Elem()
}

func (o SourceGainsightPxOutput) ToSourceGainsightPxOutput() SourceGainsightPxOutput {
	return o
}

func (o SourceGainsightPxOutput) ToSourceGainsightPxOutputWithContext(ctx context.Context) SourceGainsightPxOutput {
	return o
}

func (o SourceGainsightPxOutput) Configuration() SourceGainsightPxConfigurationOutput {
	return o.ApplyT(func(v *SourceGainsightPx) SourceGainsightPxConfigurationOutput { return v.Configuration }).(SourceGainsightPxConfigurationOutput)
}

func (o SourceGainsightPxOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGainsightPx) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceGainsightPxOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGainsightPx) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceGainsightPxOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGainsightPx) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceGainsightPxOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGainsightPx) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceGainsightPxOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGainsightPx) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGainsightPxInput)(nil)).Elem(), &SourceGainsightPx{})
	pulumi.RegisterOutputType(SourceGainsightPxOutput{})
}