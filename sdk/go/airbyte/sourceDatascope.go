// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceDatascope Resource
type SourceDatascope struct {
	pulumi.CustomResourceState

	Configuration SourceDatascopeConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceDatascope registers a new resource with the given unique name, arguments, and options.
func NewSourceDatascope(ctx *pulumi.Context,
	name string, args *SourceDatascopeArgs, opts ...pulumi.ResourceOption) (*SourceDatascope, error) {
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
	var resource SourceDatascope
	err := ctx.RegisterResource("airbyte:index/sourceDatascope:SourceDatascope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceDatascope gets an existing SourceDatascope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceDatascope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceDatascopeState, opts ...pulumi.ResourceOption) (*SourceDatascope, error) {
	var resource SourceDatascope
	err := ctx.ReadResource("airbyte:index/sourceDatascope:SourceDatascope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceDatascope resources.
type sourceDatascopeState struct {
	Configuration *SourceDatascopeConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceDatascopeState struct {
	Configuration SourceDatascopeConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceDatascopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceDatascopeState)(nil)).Elem()
}

type sourceDatascopeArgs struct {
	Configuration SourceDatascopeConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceDatascope resource.
type SourceDatascopeArgs struct {
	Configuration SourceDatascopeConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceDatascopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceDatascopeArgs)(nil)).Elem()
}

type SourceDatascopeInput interface {
	pulumi.Input

	ToSourceDatascopeOutput() SourceDatascopeOutput
	ToSourceDatascopeOutputWithContext(ctx context.Context) SourceDatascopeOutput
}

func (*SourceDatascope) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceDatascope)(nil)).Elem()
}

func (i *SourceDatascope) ToSourceDatascopeOutput() SourceDatascopeOutput {
	return i.ToSourceDatascopeOutputWithContext(context.Background())
}

func (i *SourceDatascope) ToSourceDatascopeOutputWithContext(ctx context.Context) SourceDatascopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceDatascopeOutput)
}

// SourceDatascopeArrayInput is an input type that accepts SourceDatascopeArray and SourceDatascopeArrayOutput values.
// You can construct a concrete instance of `SourceDatascopeArrayInput` via:
//
//	SourceDatascopeArray{ SourceDatascopeArgs{...} }
type SourceDatascopeArrayInput interface {
	pulumi.Input

	ToSourceDatascopeArrayOutput() SourceDatascopeArrayOutput
	ToSourceDatascopeArrayOutputWithContext(context.Context) SourceDatascopeArrayOutput
}

type SourceDatascopeArray []SourceDatascopeInput

func (SourceDatascopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceDatascope)(nil)).Elem()
}

func (i SourceDatascopeArray) ToSourceDatascopeArrayOutput() SourceDatascopeArrayOutput {
	return i.ToSourceDatascopeArrayOutputWithContext(context.Background())
}

func (i SourceDatascopeArray) ToSourceDatascopeArrayOutputWithContext(ctx context.Context) SourceDatascopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceDatascopeArrayOutput)
}

// SourceDatascopeMapInput is an input type that accepts SourceDatascopeMap and SourceDatascopeMapOutput values.
// You can construct a concrete instance of `SourceDatascopeMapInput` via:
//
//	SourceDatascopeMap{ "key": SourceDatascopeArgs{...} }
type SourceDatascopeMapInput interface {
	pulumi.Input

	ToSourceDatascopeMapOutput() SourceDatascopeMapOutput
	ToSourceDatascopeMapOutputWithContext(context.Context) SourceDatascopeMapOutput
}

type SourceDatascopeMap map[string]SourceDatascopeInput

func (SourceDatascopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceDatascope)(nil)).Elem()
}

func (i SourceDatascopeMap) ToSourceDatascopeMapOutput() SourceDatascopeMapOutput {
	return i.ToSourceDatascopeMapOutputWithContext(context.Background())
}

func (i SourceDatascopeMap) ToSourceDatascopeMapOutputWithContext(ctx context.Context) SourceDatascopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceDatascopeMapOutput)
}

type SourceDatascopeOutput struct{ *pulumi.OutputState }

func (SourceDatascopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceDatascope)(nil)).Elem()
}

func (o SourceDatascopeOutput) ToSourceDatascopeOutput() SourceDatascopeOutput {
	return o
}

func (o SourceDatascopeOutput) ToSourceDatascopeOutputWithContext(ctx context.Context) SourceDatascopeOutput {
	return o
}

func (o SourceDatascopeOutput) Configuration() SourceDatascopeConfigurationOutput {
	return o.ApplyT(func(v *SourceDatascope) SourceDatascopeConfigurationOutput { return v.Configuration }).(SourceDatascopeConfigurationOutput)
}

func (o SourceDatascopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceDatascope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceDatascopeOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceDatascope) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceDatascopeOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceDatascope) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceDatascopeOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceDatascope) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceDatascopeOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceDatascope) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceDatascopeArrayOutput struct{ *pulumi.OutputState }

func (SourceDatascopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceDatascope)(nil)).Elem()
}

func (o SourceDatascopeArrayOutput) ToSourceDatascopeArrayOutput() SourceDatascopeArrayOutput {
	return o
}

func (o SourceDatascopeArrayOutput) ToSourceDatascopeArrayOutputWithContext(ctx context.Context) SourceDatascopeArrayOutput {
	return o
}

func (o SourceDatascopeArrayOutput) Index(i pulumi.IntInput) SourceDatascopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceDatascope {
		return vs[0].([]*SourceDatascope)[vs[1].(int)]
	}).(SourceDatascopeOutput)
}

type SourceDatascopeMapOutput struct{ *pulumi.OutputState }

func (SourceDatascopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceDatascope)(nil)).Elem()
}

func (o SourceDatascopeMapOutput) ToSourceDatascopeMapOutput() SourceDatascopeMapOutput {
	return o
}

func (o SourceDatascopeMapOutput) ToSourceDatascopeMapOutputWithContext(ctx context.Context) SourceDatascopeMapOutput {
	return o
}

func (o SourceDatascopeMapOutput) MapIndex(k pulumi.StringInput) SourceDatascopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceDatascope {
		return vs[0].(map[string]*SourceDatascope)[vs[1].(string)]
	}).(SourceDatascopeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceDatascopeInput)(nil)).Elem(), &SourceDatascope{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceDatascopeArrayInput)(nil)).Elem(), SourceDatascopeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceDatascopeMapInput)(nil)).Elem(), SourceDatascopeMap{})
	pulumi.RegisterOutputType(SourceDatascopeOutput{})
	pulumi.RegisterOutputType(SourceDatascopeArrayOutput{})
	pulumi.RegisterOutputType(SourceDatascopeMapOutput{})
}
