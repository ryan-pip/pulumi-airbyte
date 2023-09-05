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

// SourceXero Resource
type SourceXero struct {
	pulumi.CustomResourceState

	Configuration SourceXeroConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceXero registers a new resource with the given unique name, arguments, and options.
func NewSourceXero(ctx *pulumi.Context,
	name string, args *SourceXeroArgs, opts ...pulumi.ResourceOption) (*SourceXero, error) {
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
	var resource SourceXero
	err := ctx.RegisterResource("airbyte:index/sourceXero:SourceXero", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceXero gets an existing SourceXero resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceXero(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceXeroState, opts ...pulumi.ResourceOption) (*SourceXero, error) {
	var resource SourceXero
	err := ctx.ReadResource("airbyte:index/sourceXero:SourceXero", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceXero resources.
type sourceXeroState struct {
	Configuration *SourceXeroConfiguration `pulumi:"configuration"`
	Name          *string                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceXeroState struct {
	Configuration SourceXeroConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceXeroState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceXeroState)(nil)).Elem()
}

type sourceXeroArgs struct {
	Configuration SourceXeroConfiguration `pulumi:"configuration"`
	Name          string                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceXero resource.
type SourceXeroArgs struct {
	Configuration SourceXeroConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceXeroArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceXeroArgs)(nil)).Elem()
}

type SourceXeroInput interface {
	pulumi.Input

	ToSourceXeroOutput() SourceXeroOutput
	ToSourceXeroOutputWithContext(ctx context.Context) SourceXeroOutput
}

func (*SourceXero) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceXero)(nil)).Elem()
}

func (i *SourceXero) ToSourceXeroOutput() SourceXeroOutput {
	return i.ToSourceXeroOutputWithContext(context.Background())
}

func (i *SourceXero) ToSourceXeroOutputWithContext(ctx context.Context) SourceXeroOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceXeroOutput)
}

type SourceXeroOutput struct{ *pulumi.OutputState }

func (SourceXeroOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceXero)(nil)).Elem()
}

func (o SourceXeroOutput) ToSourceXeroOutput() SourceXeroOutput {
	return o
}

func (o SourceXeroOutput) ToSourceXeroOutputWithContext(ctx context.Context) SourceXeroOutput {
	return o
}

func (o SourceXeroOutput) Configuration() SourceXeroConfigurationOutput {
	return o.ApplyT(func(v *SourceXero) SourceXeroConfigurationOutput { return v.Configuration }).(SourceXeroConfigurationOutput)
}

func (o SourceXeroOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceXero) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceXeroOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceXero) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceXeroOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceXero) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceXeroOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceXero) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceXeroOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceXero) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceXeroInput)(nil)).Elem(), &SourceXero{})
	pulumi.RegisterOutputType(SourceXeroOutput{})
}
