// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourcePublicApis struct {
	pulumi.CustomResourceState

	Configuration SourcePublicApisConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                 `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourcePublicApis registers a new resource with the given unique name, arguments, and options.
func NewSourcePublicApis(ctx *pulumi.Context,
	name string, args *SourcePublicApisArgs, opts ...pulumi.ResourceOption) (*SourcePublicApis, error) {
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
	var resource SourcePublicApis
	err := ctx.RegisterResource("airbyte:index/sourcePublicApis:SourcePublicApis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourcePublicApis gets an existing SourcePublicApis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourcePublicApis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourcePublicApisState, opts ...pulumi.ResourceOption) (*SourcePublicApis, error) {
	var resource SourcePublicApis
	err := ctx.ReadResource("airbyte:index/sourcePublicApis:SourcePublicApis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourcePublicApis resources.
type sourcePublicApisState struct {
	Configuration *SourcePublicApisConfiguration `pulumi:"configuration"`
	Name          *string                        `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourcePublicApisState struct {
	Configuration SourcePublicApisConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourcePublicApisState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePublicApisState)(nil)).Elem()
}

type sourcePublicApisArgs struct {
	Configuration SourcePublicApisConfiguration `pulumi:"configuration"`
	Name          string                        `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourcePublicApis resource.
type SourcePublicApisArgs struct {
	Configuration SourcePublicApisConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourcePublicApisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePublicApisArgs)(nil)).Elem()
}

type SourcePublicApisInput interface {
	pulumi.Input

	ToSourcePublicApisOutput() SourcePublicApisOutput
	ToSourcePublicApisOutputWithContext(ctx context.Context) SourcePublicApisOutput
}

func (*SourcePublicApis) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePublicApis)(nil)).Elem()
}

func (i *SourcePublicApis) ToSourcePublicApisOutput() SourcePublicApisOutput {
	return i.ToSourcePublicApisOutputWithContext(context.Background())
}

func (i *SourcePublicApis) ToSourcePublicApisOutputWithContext(ctx context.Context) SourcePublicApisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePublicApisOutput)
}

type SourcePublicApisOutput struct{ *pulumi.OutputState }

func (SourcePublicApisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePublicApis)(nil)).Elem()
}

func (o SourcePublicApisOutput) ToSourcePublicApisOutput() SourcePublicApisOutput {
	return o
}

func (o SourcePublicApisOutput) ToSourcePublicApisOutputWithContext(ctx context.Context) SourcePublicApisOutput {
	return o
}

func (o SourcePublicApisOutput) Configuration() SourcePublicApisConfigurationOutput {
	return o.ApplyT(func(v *SourcePublicApis) SourcePublicApisConfigurationOutput { return v.Configuration }).(SourcePublicApisConfigurationOutput)
}

func (o SourcePublicApisOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePublicApis) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourcePublicApisOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePublicApis) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourcePublicApisOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePublicApis) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourcePublicApisOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePublicApis) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourcePublicApisOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePublicApis) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePublicApisInput)(nil)).Elem(), &SourcePublicApis{})
	pulumi.RegisterOutputType(SourcePublicApisOutput{})
}
