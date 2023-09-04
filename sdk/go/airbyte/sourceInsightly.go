// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceInsightly struct {
	pulumi.CustomResourceState

	Configuration SourceInsightlyConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceInsightly registers a new resource with the given unique name, arguments, and options.
func NewSourceInsightly(ctx *pulumi.Context,
	name string, args *SourceInsightlyArgs, opts ...pulumi.ResourceOption) (*SourceInsightly, error) {
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
	var resource SourceInsightly
	err := ctx.RegisterResource("airbyte:index/sourceInsightly:SourceInsightly", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceInsightly gets an existing SourceInsightly resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceInsightly(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceInsightlyState, opts ...pulumi.ResourceOption) (*SourceInsightly, error) {
	var resource SourceInsightly
	err := ctx.ReadResource("airbyte:index/sourceInsightly:SourceInsightly", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceInsightly resources.
type sourceInsightlyState struct {
	Configuration *SourceInsightlyConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceInsightlyState struct {
	Configuration SourceInsightlyConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceInsightlyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceInsightlyState)(nil)).Elem()
}

type sourceInsightlyArgs struct {
	Configuration SourceInsightlyConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceInsightly resource.
type SourceInsightlyArgs struct {
	Configuration SourceInsightlyConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceInsightlyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceInsightlyArgs)(nil)).Elem()
}

type SourceInsightlyInput interface {
	pulumi.Input

	ToSourceInsightlyOutput() SourceInsightlyOutput
	ToSourceInsightlyOutputWithContext(ctx context.Context) SourceInsightlyOutput
}

func (*SourceInsightly) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceInsightly)(nil)).Elem()
}

func (i *SourceInsightly) ToSourceInsightlyOutput() SourceInsightlyOutput {
	return i.ToSourceInsightlyOutputWithContext(context.Background())
}

func (i *SourceInsightly) ToSourceInsightlyOutputWithContext(ctx context.Context) SourceInsightlyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceInsightlyOutput)
}

type SourceInsightlyOutput struct{ *pulumi.OutputState }

func (SourceInsightlyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceInsightly)(nil)).Elem()
}

func (o SourceInsightlyOutput) ToSourceInsightlyOutput() SourceInsightlyOutput {
	return o
}

func (o SourceInsightlyOutput) ToSourceInsightlyOutputWithContext(ctx context.Context) SourceInsightlyOutput {
	return o
}

func (o SourceInsightlyOutput) Configuration() SourceInsightlyConfigurationOutput {
	return o.ApplyT(func(v *SourceInsightly) SourceInsightlyConfigurationOutput { return v.Configuration }).(SourceInsightlyConfigurationOutput)
}

func (o SourceInsightlyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceInsightly) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceInsightlyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceInsightly) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceInsightlyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceInsightly) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceInsightlyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceInsightly) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceInsightlyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceInsightly) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceInsightlyInput)(nil)).Elem(), &SourceInsightly{})
	pulumi.RegisterOutputType(SourceInsightlyOutput{})
}