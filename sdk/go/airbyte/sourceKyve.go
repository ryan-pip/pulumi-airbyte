// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceKyve struct {
	pulumi.CustomResourceState

	Configuration SourceKyveConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceKyve registers a new resource with the given unique name, arguments, and options.
func NewSourceKyve(ctx *pulumi.Context,
	name string, args *SourceKyveArgs, opts ...pulumi.ResourceOption) (*SourceKyve, error) {
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
	var resource SourceKyve
	err := ctx.RegisterResource("airbyte:index/sourceKyve:SourceKyve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceKyve gets an existing SourceKyve resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceKyve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceKyveState, opts ...pulumi.ResourceOption) (*SourceKyve, error) {
	var resource SourceKyve
	err := ctx.ReadResource("airbyte:index/sourceKyve:SourceKyve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceKyve resources.
type sourceKyveState struct {
	Configuration *SourceKyveConfiguration `pulumi:"configuration"`
	Name          *string                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceKyveState struct {
	Configuration SourceKyveConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceKyveState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceKyveState)(nil)).Elem()
}

type sourceKyveArgs struct {
	Configuration SourceKyveConfiguration `pulumi:"configuration"`
	Name          string                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceKyve resource.
type SourceKyveArgs struct {
	Configuration SourceKyveConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceKyveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceKyveArgs)(nil)).Elem()
}

type SourceKyveInput interface {
	pulumi.Input

	ToSourceKyveOutput() SourceKyveOutput
	ToSourceKyveOutputWithContext(ctx context.Context) SourceKyveOutput
}

func (*SourceKyve) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceKyve)(nil)).Elem()
}

func (i *SourceKyve) ToSourceKyveOutput() SourceKyveOutput {
	return i.ToSourceKyveOutputWithContext(context.Background())
}

func (i *SourceKyve) ToSourceKyveOutputWithContext(ctx context.Context) SourceKyveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceKyveOutput)
}

type SourceKyveOutput struct{ *pulumi.OutputState }

func (SourceKyveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceKyve)(nil)).Elem()
}

func (o SourceKyveOutput) ToSourceKyveOutput() SourceKyveOutput {
	return o
}

func (o SourceKyveOutput) ToSourceKyveOutputWithContext(ctx context.Context) SourceKyveOutput {
	return o
}

func (o SourceKyveOutput) Configuration() SourceKyveConfigurationOutput {
	return o.ApplyT(func(v *SourceKyve) SourceKyveConfigurationOutput { return v.Configuration }).(SourceKyveConfigurationOutput)
}

func (o SourceKyveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceKyve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceKyveOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceKyve) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceKyveOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceKyve) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceKyveOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceKyve) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceKyveOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceKyve) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceKyveInput)(nil)).Elem(), &SourceKyve{})
	pulumi.RegisterOutputType(SourceKyveOutput{})
}
