// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// DestinationMysql Resource
type DestinationMysql struct {
	pulumi.CustomResourceState

	Configuration   DestinationMysqlConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                 `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                 `pulumi:"destinationType"`
	Name            pulumi.StringOutput                 `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                 `pulumi:"workspaceId"`
}

// NewDestinationMysql registers a new resource with the given unique name, arguments, and options.
func NewDestinationMysql(ctx *pulumi.Context,
	name string, args *DestinationMysqlArgs, opts ...pulumi.ResourceOption) (*DestinationMysql, error) {
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
	var resource DestinationMysql
	err := ctx.RegisterResource("airbyte:index/destinationMysql:DestinationMysql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationMysql gets an existing DestinationMysql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationMysql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationMysqlState, opts ...pulumi.ResourceOption) (*DestinationMysql, error) {
	var resource DestinationMysql
	err := ctx.ReadResource("airbyte:index/destinationMysql:DestinationMysql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationMysql resources.
type destinationMysqlState struct {
	Configuration   *DestinationMysqlConfiguration `pulumi:"configuration"`
	DestinationId   *string                        `pulumi:"destinationId"`
	DestinationType *string                        `pulumi:"destinationType"`
	Name            *string                        `pulumi:"name"`
	WorkspaceId     *string                        `pulumi:"workspaceId"`
}

type DestinationMysqlState struct {
	Configuration   DestinationMysqlConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationMysqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationMysqlState)(nil)).Elem()
}

type destinationMysqlArgs struct {
	Configuration DestinationMysqlConfiguration `pulumi:"configuration"`
	Name          string                        `pulumi:"name"`
	WorkspaceId   string                        `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationMysql resource.
type DestinationMysqlArgs struct {
	Configuration DestinationMysqlConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationMysqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationMysqlArgs)(nil)).Elem()
}

type DestinationMysqlInput interface {
	pulumi.Input

	ToDestinationMysqlOutput() DestinationMysqlOutput
	ToDestinationMysqlOutputWithContext(ctx context.Context) DestinationMysqlOutput
}

func (*DestinationMysql) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationMysql)(nil)).Elem()
}

func (i *DestinationMysql) ToDestinationMysqlOutput() DestinationMysqlOutput {
	return i.ToDestinationMysqlOutputWithContext(context.Background())
}

func (i *DestinationMysql) ToDestinationMysqlOutputWithContext(ctx context.Context) DestinationMysqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationMysqlOutput)
}

// DestinationMysqlArrayInput is an input type that accepts DestinationMysqlArray and DestinationMysqlArrayOutput values.
// You can construct a concrete instance of `DestinationMysqlArrayInput` via:
//
//	DestinationMysqlArray{ DestinationMysqlArgs{...} }
type DestinationMysqlArrayInput interface {
	pulumi.Input

	ToDestinationMysqlArrayOutput() DestinationMysqlArrayOutput
	ToDestinationMysqlArrayOutputWithContext(context.Context) DestinationMysqlArrayOutput
}

type DestinationMysqlArray []DestinationMysqlInput

func (DestinationMysqlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationMysql)(nil)).Elem()
}

func (i DestinationMysqlArray) ToDestinationMysqlArrayOutput() DestinationMysqlArrayOutput {
	return i.ToDestinationMysqlArrayOutputWithContext(context.Background())
}

func (i DestinationMysqlArray) ToDestinationMysqlArrayOutputWithContext(ctx context.Context) DestinationMysqlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationMysqlArrayOutput)
}

// DestinationMysqlMapInput is an input type that accepts DestinationMysqlMap and DestinationMysqlMapOutput values.
// You can construct a concrete instance of `DestinationMysqlMapInput` via:
//
//	DestinationMysqlMap{ "key": DestinationMysqlArgs{...} }
type DestinationMysqlMapInput interface {
	pulumi.Input

	ToDestinationMysqlMapOutput() DestinationMysqlMapOutput
	ToDestinationMysqlMapOutputWithContext(context.Context) DestinationMysqlMapOutput
}

type DestinationMysqlMap map[string]DestinationMysqlInput

func (DestinationMysqlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationMysql)(nil)).Elem()
}

func (i DestinationMysqlMap) ToDestinationMysqlMapOutput() DestinationMysqlMapOutput {
	return i.ToDestinationMysqlMapOutputWithContext(context.Background())
}

func (i DestinationMysqlMap) ToDestinationMysqlMapOutputWithContext(ctx context.Context) DestinationMysqlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationMysqlMapOutput)
}

type DestinationMysqlOutput struct{ *pulumi.OutputState }

func (DestinationMysqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationMysql)(nil)).Elem()
}

func (o DestinationMysqlOutput) ToDestinationMysqlOutput() DestinationMysqlOutput {
	return o
}

func (o DestinationMysqlOutput) ToDestinationMysqlOutputWithContext(ctx context.Context) DestinationMysqlOutput {
	return o
}

func (o DestinationMysqlOutput) Configuration() DestinationMysqlConfigurationOutput {
	return o.ApplyT(func(v *DestinationMysql) DestinationMysqlConfigurationOutput { return v.Configuration }).(DestinationMysqlConfigurationOutput)
}

func (o DestinationMysqlOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationMysql) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationMysqlOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationMysql) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationMysqlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationMysql) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationMysqlOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationMysql) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type DestinationMysqlArrayOutput struct{ *pulumi.OutputState }

func (DestinationMysqlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationMysql)(nil)).Elem()
}

func (o DestinationMysqlArrayOutput) ToDestinationMysqlArrayOutput() DestinationMysqlArrayOutput {
	return o
}

func (o DestinationMysqlArrayOutput) ToDestinationMysqlArrayOutputWithContext(ctx context.Context) DestinationMysqlArrayOutput {
	return o
}

func (o DestinationMysqlArrayOutput) Index(i pulumi.IntInput) DestinationMysqlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationMysql {
		return vs[0].([]*DestinationMysql)[vs[1].(int)]
	}).(DestinationMysqlOutput)
}

type DestinationMysqlMapOutput struct{ *pulumi.OutputState }

func (DestinationMysqlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationMysql)(nil)).Elem()
}

func (o DestinationMysqlMapOutput) ToDestinationMysqlMapOutput() DestinationMysqlMapOutput {
	return o
}

func (o DestinationMysqlMapOutput) ToDestinationMysqlMapOutputWithContext(ctx context.Context) DestinationMysqlMapOutput {
	return o
}

func (o DestinationMysqlMapOutput) MapIndex(k pulumi.StringInput) DestinationMysqlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationMysql {
		return vs[0].(map[string]*DestinationMysql)[vs[1].(string)]
	}).(DestinationMysqlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationMysqlInput)(nil)).Elem(), &DestinationMysql{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationMysqlArrayInput)(nil)).Elem(), DestinationMysqlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationMysqlMapInput)(nil)).Elem(), DestinationMysqlMap{})
	pulumi.RegisterOutputType(DestinationMysqlOutput{})
	pulumi.RegisterOutputType(DestinationMysqlArrayOutput{})
	pulumi.RegisterOutputType(DestinationMysqlMapOutput{})
}
