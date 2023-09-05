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

// SourceDynamodb Resource
type SourceDynamodb struct {
	pulumi.CustomResourceState

	Configuration SourceDynamodbConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceDynamodb registers a new resource with the given unique name, arguments, and options.
func NewSourceDynamodb(ctx *pulumi.Context,
	name string, args *SourceDynamodbArgs, opts ...pulumi.ResourceOption) (*SourceDynamodb, error) {
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
	var resource SourceDynamodb
	err := ctx.RegisterResource("airbyte:index/sourceDynamodb:SourceDynamodb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceDynamodb gets an existing SourceDynamodb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceDynamodb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceDynamodbState, opts ...pulumi.ResourceOption) (*SourceDynamodb, error) {
	var resource SourceDynamodb
	err := ctx.ReadResource("airbyte:index/sourceDynamodb:SourceDynamodb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceDynamodb resources.
type sourceDynamodbState struct {
	Configuration *SourceDynamodbConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceDynamodbState struct {
	Configuration SourceDynamodbConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceDynamodbState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceDynamodbState)(nil)).Elem()
}

type sourceDynamodbArgs struct {
	Configuration SourceDynamodbConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceDynamodb resource.
type SourceDynamodbArgs struct {
	Configuration SourceDynamodbConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceDynamodbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceDynamodbArgs)(nil)).Elem()
}

type SourceDynamodbInput interface {
	pulumi.Input

	ToSourceDynamodbOutput() SourceDynamodbOutput
	ToSourceDynamodbOutputWithContext(ctx context.Context) SourceDynamodbOutput
}

func (*SourceDynamodb) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceDynamodb)(nil)).Elem()
}

func (i *SourceDynamodb) ToSourceDynamodbOutput() SourceDynamodbOutput {
	return i.ToSourceDynamodbOutputWithContext(context.Background())
}

func (i *SourceDynamodb) ToSourceDynamodbOutputWithContext(ctx context.Context) SourceDynamodbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceDynamodbOutput)
}

type SourceDynamodbOutput struct{ *pulumi.OutputState }

func (SourceDynamodbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceDynamodb)(nil)).Elem()
}

func (o SourceDynamodbOutput) ToSourceDynamodbOutput() SourceDynamodbOutput {
	return o
}

func (o SourceDynamodbOutput) ToSourceDynamodbOutputWithContext(ctx context.Context) SourceDynamodbOutput {
	return o
}

func (o SourceDynamodbOutput) Configuration() SourceDynamodbConfigurationOutput {
	return o.ApplyT(func(v *SourceDynamodb) SourceDynamodbConfigurationOutput { return v.Configuration }).(SourceDynamodbConfigurationOutput)
}

func (o SourceDynamodbOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceDynamodb) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceDynamodbOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceDynamodb) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceDynamodbOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceDynamodb) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceDynamodbOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceDynamodb) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceDynamodbOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceDynamodb) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceDynamodbInput)(nil)).Elem(), &SourceDynamodb{})
	pulumi.RegisterOutputType(SourceDynamodbOutput{})
}
