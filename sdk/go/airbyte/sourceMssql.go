// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceMssql struct {
	pulumi.CustomResourceState

	Configuration SourceMssqlConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceMssql registers a new resource with the given unique name, arguments, and options.
func NewSourceMssql(ctx *pulumi.Context,
	name string, args *SourceMssqlArgs, opts ...pulumi.ResourceOption) (*SourceMssql, error) {
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
	var resource SourceMssql
	err := ctx.RegisterResource("airbyte:index/sourceMssql:SourceMssql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceMssql gets an existing SourceMssql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceMssql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceMssqlState, opts ...pulumi.ResourceOption) (*SourceMssql, error) {
	var resource SourceMssql
	err := ctx.ReadResource("airbyte:index/sourceMssql:SourceMssql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceMssql resources.
type sourceMssqlState struct {
	Configuration *SourceMssqlConfiguration `pulumi:"configuration"`
	Name          *string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceMssqlState struct {
	Configuration SourceMssqlConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceMssqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMssqlState)(nil)).Elem()
}

type sourceMssqlArgs struct {
	Configuration SourceMssqlConfiguration `pulumi:"configuration"`
	Name          string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceMssql resource.
type SourceMssqlArgs struct {
	Configuration SourceMssqlConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceMssqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMssqlArgs)(nil)).Elem()
}

type SourceMssqlInput interface {
	pulumi.Input

	ToSourceMssqlOutput() SourceMssqlOutput
	ToSourceMssqlOutputWithContext(ctx context.Context) SourceMssqlOutput
}

func (*SourceMssql) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMssql)(nil)).Elem()
}

func (i *SourceMssql) ToSourceMssqlOutput() SourceMssqlOutput {
	return i.ToSourceMssqlOutputWithContext(context.Background())
}

func (i *SourceMssql) ToSourceMssqlOutputWithContext(ctx context.Context) SourceMssqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMssqlOutput)
}

type SourceMssqlOutput struct{ *pulumi.OutputState }

func (SourceMssqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMssql)(nil)).Elem()
}

func (o SourceMssqlOutput) ToSourceMssqlOutput() SourceMssqlOutput {
	return o
}

func (o SourceMssqlOutput) ToSourceMssqlOutputWithContext(ctx context.Context) SourceMssqlOutput {
	return o
}

func (o SourceMssqlOutput) Configuration() SourceMssqlConfigurationOutput {
	return o.ApplyT(func(v *SourceMssql) SourceMssqlConfigurationOutput { return v.Configuration }).(SourceMssqlConfigurationOutput)
}

func (o SourceMssqlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMssql) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceMssqlOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceMssql) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceMssqlOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMssql) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceMssqlOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMssql) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceMssqlOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMssql) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMssqlInput)(nil)).Elem(), &SourceMssql{})
	pulumi.RegisterOutputType(SourceMssqlOutput{})
}