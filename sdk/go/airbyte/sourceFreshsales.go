// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceFreshsales struct {
	pulumi.CustomResourceState

	Configuration SourceFreshsalesConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                 `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceFreshsales registers a new resource with the given unique name, arguments, and options.
func NewSourceFreshsales(ctx *pulumi.Context,
	name string, args *SourceFreshsalesArgs, opts ...pulumi.ResourceOption) (*SourceFreshsales, error) {
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
	var resource SourceFreshsales
	err := ctx.RegisterResource("airbyte:index/sourceFreshsales:SourceFreshsales", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceFreshsales gets an existing SourceFreshsales resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceFreshsales(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceFreshsalesState, opts ...pulumi.ResourceOption) (*SourceFreshsales, error) {
	var resource SourceFreshsales
	err := ctx.ReadResource("airbyte:index/sourceFreshsales:SourceFreshsales", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceFreshsales resources.
type sourceFreshsalesState struct {
	Configuration *SourceFreshsalesConfiguration `pulumi:"configuration"`
	Name          *string                        `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceFreshsalesState struct {
	Configuration SourceFreshsalesConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceFreshsalesState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFreshsalesState)(nil)).Elem()
}

type sourceFreshsalesArgs struct {
	Configuration SourceFreshsalesConfiguration `pulumi:"configuration"`
	Name          string                        `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceFreshsales resource.
type SourceFreshsalesArgs struct {
	Configuration SourceFreshsalesConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceFreshsalesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFreshsalesArgs)(nil)).Elem()
}

type SourceFreshsalesInput interface {
	pulumi.Input

	ToSourceFreshsalesOutput() SourceFreshsalesOutput
	ToSourceFreshsalesOutputWithContext(ctx context.Context) SourceFreshsalesOutput
}

func (*SourceFreshsales) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFreshsales)(nil)).Elem()
}

func (i *SourceFreshsales) ToSourceFreshsalesOutput() SourceFreshsalesOutput {
	return i.ToSourceFreshsalesOutputWithContext(context.Background())
}

func (i *SourceFreshsales) ToSourceFreshsalesOutputWithContext(ctx context.Context) SourceFreshsalesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFreshsalesOutput)
}

type SourceFreshsalesOutput struct{ *pulumi.OutputState }

func (SourceFreshsalesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFreshsales)(nil)).Elem()
}

func (o SourceFreshsalesOutput) ToSourceFreshsalesOutput() SourceFreshsalesOutput {
	return o
}

func (o SourceFreshsalesOutput) ToSourceFreshsalesOutputWithContext(ctx context.Context) SourceFreshsalesOutput {
	return o
}

func (o SourceFreshsalesOutput) Configuration() SourceFreshsalesConfigurationOutput {
	return o.ApplyT(func(v *SourceFreshsales) SourceFreshsalesConfigurationOutput { return v.Configuration }).(SourceFreshsalesConfigurationOutput)
}

func (o SourceFreshsalesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFreshsales) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceFreshsalesOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceFreshsales) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceFreshsalesOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFreshsales) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceFreshsalesOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFreshsales) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceFreshsalesOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFreshsales) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFreshsalesInput)(nil)).Elem(), &SourceFreshsales{})
	pulumi.RegisterOutputType(SourceFreshsalesOutput{})
}
