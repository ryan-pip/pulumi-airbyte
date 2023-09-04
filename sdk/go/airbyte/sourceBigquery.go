// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceBigquery struct {
	pulumi.CustomResourceState

	Configuration SourceBigqueryConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceBigquery registers a new resource with the given unique name, arguments, and options.
func NewSourceBigquery(ctx *pulumi.Context,
	name string, args *SourceBigqueryArgs, opts ...pulumi.ResourceOption) (*SourceBigquery, error) {
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
	var resource SourceBigquery
	err := ctx.RegisterResource("airbyte:index/sourceBigquery:SourceBigquery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceBigquery gets an existing SourceBigquery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceBigquery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceBigqueryState, opts ...pulumi.ResourceOption) (*SourceBigquery, error) {
	var resource SourceBigquery
	err := ctx.ReadResource("airbyte:index/sourceBigquery:SourceBigquery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceBigquery resources.
type sourceBigqueryState struct {
	Configuration *SourceBigqueryConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceBigqueryState struct {
	Configuration SourceBigqueryConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceBigqueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceBigqueryState)(nil)).Elem()
}

type sourceBigqueryArgs struct {
	Configuration SourceBigqueryConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceBigquery resource.
type SourceBigqueryArgs struct {
	Configuration SourceBigqueryConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceBigqueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceBigqueryArgs)(nil)).Elem()
}

type SourceBigqueryInput interface {
	pulumi.Input

	ToSourceBigqueryOutput() SourceBigqueryOutput
	ToSourceBigqueryOutputWithContext(ctx context.Context) SourceBigqueryOutput
}

func (*SourceBigquery) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceBigquery)(nil)).Elem()
}

func (i *SourceBigquery) ToSourceBigqueryOutput() SourceBigqueryOutput {
	return i.ToSourceBigqueryOutputWithContext(context.Background())
}

func (i *SourceBigquery) ToSourceBigqueryOutputWithContext(ctx context.Context) SourceBigqueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceBigqueryOutput)
}

type SourceBigqueryOutput struct{ *pulumi.OutputState }

func (SourceBigqueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceBigquery)(nil)).Elem()
}

func (o SourceBigqueryOutput) ToSourceBigqueryOutput() SourceBigqueryOutput {
	return o
}

func (o SourceBigqueryOutput) ToSourceBigqueryOutputWithContext(ctx context.Context) SourceBigqueryOutput {
	return o
}

func (o SourceBigqueryOutput) Configuration() SourceBigqueryConfigurationOutput {
	return o.ApplyT(func(v *SourceBigquery) SourceBigqueryConfigurationOutput { return v.Configuration }).(SourceBigqueryConfigurationOutput)
}

func (o SourceBigqueryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBigquery) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceBigqueryOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceBigquery) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceBigqueryOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBigquery) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceBigqueryOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBigquery) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceBigqueryOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceBigquery) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceBigqueryInput)(nil)).Elem(), &SourceBigquery{})
	pulumi.RegisterOutputType(SourceBigqueryOutput{})
}