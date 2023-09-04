// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceAsana struct {
	pulumi.CustomResourceState

	Configuration SourceAsanaConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceAsana registers a new resource with the given unique name, arguments, and options.
func NewSourceAsana(ctx *pulumi.Context,
	name string, args *SourceAsanaArgs, opts ...pulumi.ResourceOption) (*SourceAsana, error) {
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
	var resource SourceAsana
	err := ctx.RegisterResource("airbyte:index/sourceAsana:SourceAsana", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceAsana gets an existing SourceAsana resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceAsana(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceAsanaState, opts ...pulumi.ResourceOption) (*SourceAsana, error) {
	var resource SourceAsana
	err := ctx.ReadResource("airbyte:index/sourceAsana:SourceAsana", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceAsana resources.
type sourceAsanaState struct {
	Configuration *SourceAsanaConfiguration `pulumi:"configuration"`
	Name          *string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceAsanaState struct {
	Configuration SourceAsanaConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceAsanaState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAsanaState)(nil)).Elem()
}

type sourceAsanaArgs struct {
	Configuration SourceAsanaConfiguration `pulumi:"configuration"`
	Name          string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceAsana resource.
type SourceAsanaArgs struct {
	Configuration SourceAsanaConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceAsanaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAsanaArgs)(nil)).Elem()
}

type SourceAsanaInput interface {
	pulumi.Input

	ToSourceAsanaOutput() SourceAsanaOutput
	ToSourceAsanaOutputWithContext(ctx context.Context) SourceAsanaOutput
}

func (*SourceAsana) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAsana)(nil)).Elem()
}

func (i *SourceAsana) ToSourceAsanaOutput() SourceAsanaOutput {
	return i.ToSourceAsanaOutputWithContext(context.Background())
}

func (i *SourceAsana) ToSourceAsanaOutputWithContext(ctx context.Context) SourceAsanaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAsanaOutput)
}

type SourceAsanaOutput struct{ *pulumi.OutputState }

func (SourceAsanaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAsana)(nil)).Elem()
}

func (o SourceAsanaOutput) ToSourceAsanaOutput() SourceAsanaOutput {
	return o
}

func (o SourceAsanaOutput) ToSourceAsanaOutputWithContext(ctx context.Context) SourceAsanaOutput {
	return o
}

func (o SourceAsanaOutput) Configuration() SourceAsanaConfigurationOutput {
	return o.ApplyT(func(v *SourceAsana) SourceAsanaConfigurationOutput { return v.Configuration }).(SourceAsanaConfigurationOutput)
}

func (o SourceAsanaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAsana) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceAsanaOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAsana) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceAsanaOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAsana) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceAsanaOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAsana) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceAsanaOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAsana) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAsanaInput)(nil)).Elem(), &SourceAsana{})
	pulumi.RegisterOutputType(SourceAsanaOutput{})
}