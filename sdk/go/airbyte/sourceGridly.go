// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceGridly struct {
	pulumi.CustomResourceState

	Configuration SourceGridlyConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceGridly registers a new resource with the given unique name, arguments, and options.
func NewSourceGridly(ctx *pulumi.Context,
	name string, args *SourceGridlyArgs, opts ...pulumi.ResourceOption) (*SourceGridly, error) {
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
	var resource SourceGridly
	err := ctx.RegisterResource("airbyte:index/sourceGridly:SourceGridly", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceGridly gets an existing SourceGridly resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceGridly(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceGridlyState, opts ...pulumi.ResourceOption) (*SourceGridly, error) {
	var resource SourceGridly
	err := ctx.ReadResource("airbyte:index/sourceGridly:SourceGridly", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceGridly resources.
type sourceGridlyState struct {
	Configuration *SourceGridlyConfiguration `pulumi:"configuration"`
	Name          *string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceGridlyState struct {
	Configuration SourceGridlyConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceGridlyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGridlyState)(nil)).Elem()
}

type sourceGridlyArgs struct {
	Configuration SourceGridlyConfiguration `pulumi:"configuration"`
	Name          string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceGridly resource.
type SourceGridlyArgs struct {
	Configuration SourceGridlyConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceGridlyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGridlyArgs)(nil)).Elem()
}

type SourceGridlyInput interface {
	pulumi.Input

	ToSourceGridlyOutput() SourceGridlyOutput
	ToSourceGridlyOutputWithContext(ctx context.Context) SourceGridlyOutput
}

func (*SourceGridly) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGridly)(nil)).Elem()
}

func (i *SourceGridly) ToSourceGridlyOutput() SourceGridlyOutput {
	return i.ToSourceGridlyOutputWithContext(context.Background())
}

func (i *SourceGridly) ToSourceGridlyOutputWithContext(ctx context.Context) SourceGridlyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGridlyOutput)
}

type SourceGridlyOutput struct{ *pulumi.OutputState }

func (SourceGridlyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGridly)(nil)).Elem()
}

func (o SourceGridlyOutput) ToSourceGridlyOutput() SourceGridlyOutput {
	return o
}

func (o SourceGridlyOutput) ToSourceGridlyOutputWithContext(ctx context.Context) SourceGridlyOutput {
	return o
}

func (o SourceGridlyOutput) Configuration() SourceGridlyConfigurationOutput {
	return o.ApplyT(func(v *SourceGridly) SourceGridlyConfigurationOutput { return v.Configuration }).(SourceGridlyConfigurationOutput)
}

func (o SourceGridlyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGridly) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceGridlyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGridly) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceGridlyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGridly) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceGridlyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGridly) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceGridlyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGridly) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGridlyInput)(nil)).Elem(), &SourceGridly{})
	pulumi.RegisterOutputType(SourceGridlyOutput{})
}