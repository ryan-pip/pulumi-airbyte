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

// SourceGetlago Resource
type SourceGetlago struct {
	pulumi.CustomResourceState

	Configuration SourceGetlagoConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceGetlago registers a new resource with the given unique name, arguments, and options.
func NewSourceGetlago(ctx *pulumi.Context,
	name string, args *SourceGetlagoArgs, opts ...pulumi.ResourceOption) (*SourceGetlago, error) {
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
	var resource SourceGetlago
	err := ctx.RegisterResource("airbyte:index/sourceGetlago:SourceGetlago", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceGetlago gets an existing SourceGetlago resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceGetlago(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceGetlagoState, opts ...pulumi.ResourceOption) (*SourceGetlago, error) {
	var resource SourceGetlago
	err := ctx.ReadResource("airbyte:index/sourceGetlago:SourceGetlago", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceGetlago resources.
type sourceGetlagoState struct {
	Configuration *SourceGetlagoConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceGetlagoState struct {
	Configuration SourceGetlagoConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceGetlagoState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGetlagoState)(nil)).Elem()
}

type sourceGetlagoArgs struct {
	Configuration SourceGetlagoConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceGetlago resource.
type SourceGetlagoArgs struct {
	Configuration SourceGetlagoConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceGetlagoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGetlagoArgs)(nil)).Elem()
}

type SourceGetlagoInput interface {
	pulumi.Input

	ToSourceGetlagoOutput() SourceGetlagoOutput
	ToSourceGetlagoOutputWithContext(ctx context.Context) SourceGetlagoOutput
}

func (*SourceGetlago) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGetlago)(nil)).Elem()
}

func (i *SourceGetlago) ToSourceGetlagoOutput() SourceGetlagoOutput {
	return i.ToSourceGetlagoOutputWithContext(context.Background())
}

func (i *SourceGetlago) ToSourceGetlagoOutputWithContext(ctx context.Context) SourceGetlagoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGetlagoOutput)
}

type SourceGetlagoOutput struct{ *pulumi.OutputState }

func (SourceGetlagoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGetlago)(nil)).Elem()
}

func (o SourceGetlagoOutput) ToSourceGetlagoOutput() SourceGetlagoOutput {
	return o
}

func (o SourceGetlagoOutput) ToSourceGetlagoOutputWithContext(ctx context.Context) SourceGetlagoOutput {
	return o
}

func (o SourceGetlagoOutput) Configuration() SourceGetlagoConfigurationOutput {
	return o.ApplyT(func(v *SourceGetlago) SourceGetlagoConfigurationOutput { return v.Configuration }).(SourceGetlagoConfigurationOutput)
}

func (o SourceGetlagoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGetlago) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceGetlagoOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGetlago) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceGetlagoOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGetlago) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceGetlagoOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGetlago) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceGetlagoOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGetlago) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGetlagoInput)(nil)).Elem(), &SourceGetlago{})
	pulumi.RegisterOutputType(SourceGetlagoOutput{})
}
