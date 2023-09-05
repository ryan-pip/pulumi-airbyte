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

// SourceSurveySparrow Resource
type SourceSurveySparrow struct {
	pulumi.CustomResourceState

	Configuration SourceSurveySparrowConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceSurveySparrow registers a new resource with the given unique name, arguments, and options.
func NewSourceSurveySparrow(ctx *pulumi.Context,
	name string, args *SourceSurveySparrowArgs, opts ...pulumi.ResourceOption) (*SourceSurveySparrow, error) {
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
	var resource SourceSurveySparrow
	err := ctx.RegisterResource("airbyte:index/sourceSurveySparrow:SourceSurveySparrow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSurveySparrow gets an existing SourceSurveySparrow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSurveySparrow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSurveySparrowState, opts ...pulumi.ResourceOption) (*SourceSurveySparrow, error) {
	var resource SourceSurveySparrow
	err := ctx.ReadResource("airbyte:index/sourceSurveySparrow:SourceSurveySparrow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSurveySparrow resources.
type sourceSurveySparrowState struct {
	Configuration *SourceSurveySparrowConfiguration `pulumi:"configuration"`
	Name          *string                           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceSurveySparrowState struct {
	Configuration SourceSurveySparrowConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceSurveySparrowState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSurveySparrowState)(nil)).Elem()
}

type sourceSurveySparrowArgs struct {
	Configuration SourceSurveySparrowConfiguration `pulumi:"configuration"`
	Name          string                           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSurveySparrow resource.
type SourceSurveySparrowArgs struct {
	Configuration SourceSurveySparrowConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSurveySparrowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSurveySparrowArgs)(nil)).Elem()
}

type SourceSurveySparrowInput interface {
	pulumi.Input

	ToSourceSurveySparrowOutput() SourceSurveySparrowOutput
	ToSourceSurveySparrowOutputWithContext(ctx context.Context) SourceSurveySparrowOutput
}

func (*SourceSurveySparrow) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSurveySparrow)(nil)).Elem()
}

func (i *SourceSurveySparrow) ToSourceSurveySparrowOutput() SourceSurveySparrowOutput {
	return i.ToSourceSurveySparrowOutputWithContext(context.Background())
}

func (i *SourceSurveySparrow) ToSourceSurveySparrowOutputWithContext(ctx context.Context) SourceSurveySparrowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSurveySparrowOutput)
}

// SourceSurveySparrowArrayInput is an input type that accepts SourceSurveySparrowArray and SourceSurveySparrowArrayOutput values.
// You can construct a concrete instance of `SourceSurveySparrowArrayInput` via:
//
//	SourceSurveySparrowArray{ SourceSurveySparrowArgs{...} }
type SourceSurveySparrowArrayInput interface {
	pulumi.Input

	ToSourceSurveySparrowArrayOutput() SourceSurveySparrowArrayOutput
	ToSourceSurveySparrowArrayOutputWithContext(context.Context) SourceSurveySparrowArrayOutput
}

type SourceSurveySparrowArray []SourceSurveySparrowInput

func (SourceSurveySparrowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSurveySparrow)(nil)).Elem()
}

func (i SourceSurveySparrowArray) ToSourceSurveySparrowArrayOutput() SourceSurveySparrowArrayOutput {
	return i.ToSourceSurveySparrowArrayOutputWithContext(context.Background())
}

func (i SourceSurveySparrowArray) ToSourceSurveySparrowArrayOutputWithContext(ctx context.Context) SourceSurveySparrowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSurveySparrowArrayOutput)
}

// SourceSurveySparrowMapInput is an input type that accepts SourceSurveySparrowMap and SourceSurveySparrowMapOutput values.
// You can construct a concrete instance of `SourceSurveySparrowMapInput` via:
//
//	SourceSurveySparrowMap{ "key": SourceSurveySparrowArgs{...} }
type SourceSurveySparrowMapInput interface {
	pulumi.Input

	ToSourceSurveySparrowMapOutput() SourceSurveySparrowMapOutput
	ToSourceSurveySparrowMapOutputWithContext(context.Context) SourceSurveySparrowMapOutput
}

type SourceSurveySparrowMap map[string]SourceSurveySparrowInput

func (SourceSurveySparrowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSurveySparrow)(nil)).Elem()
}

func (i SourceSurveySparrowMap) ToSourceSurveySparrowMapOutput() SourceSurveySparrowMapOutput {
	return i.ToSourceSurveySparrowMapOutputWithContext(context.Background())
}

func (i SourceSurveySparrowMap) ToSourceSurveySparrowMapOutputWithContext(ctx context.Context) SourceSurveySparrowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSurveySparrowMapOutput)
}

type SourceSurveySparrowOutput struct{ *pulumi.OutputState }

func (SourceSurveySparrowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSurveySparrow)(nil)).Elem()
}

func (o SourceSurveySparrowOutput) ToSourceSurveySparrowOutput() SourceSurveySparrowOutput {
	return o
}

func (o SourceSurveySparrowOutput) ToSourceSurveySparrowOutputWithContext(ctx context.Context) SourceSurveySparrowOutput {
	return o
}

func (o SourceSurveySparrowOutput) Configuration() SourceSurveySparrowConfigurationOutput {
	return o.ApplyT(func(v *SourceSurveySparrow) SourceSurveySparrowConfigurationOutput { return v.Configuration }).(SourceSurveySparrowConfigurationOutput)
}

func (o SourceSurveySparrowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSurveySparrow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceSurveySparrowOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSurveySparrow) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSurveySparrowOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSurveySparrow) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSurveySparrowOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSurveySparrow) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSurveySparrowOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSurveySparrow) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceSurveySparrowArrayOutput struct{ *pulumi.OutputState }

func (SourceSurveySparrowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSurveySparrow)(nil)).Elem()
}

func (o SourceSurveySparrowArrayOutput) ToSourceSurveySparrowArrayOutput() SourceSurveySparrowArrayOutput {
	return o
}

func (o SourceSurveySparrowArrayOutput) ToSourceSurveySparrowArrayOutputWithContext(ctx context.Context) SourceSurveySparrowArrayOutput {
	return o
}

func (o SourceSurveySparrowArrayOutput) Index(i pulumi.IntInput) SourceSurveySparrowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceSurveySparrow {
		return vs[0].([]*SourceSurveySparrow)[vs[1].(int)]
	}).(SourceSurveySparrowOutput)
}

type SourceSurveySparrowMapOutput struct{ *pulumi.OutputState }

func (SourceSurveySparrowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSurveySparrow)(nil)).Elem()
}

func (o SourceSurveySparrowMapOutput) ToSourceSurveySparrowMapOutput() SourceSurveySparrowMapOutput {
	return o
}

func (o SourceSurveySparrowMapOutput) ToSourceSurveySparrowMapOutputWithContext(ctx context.Context) SourceSurveySparrowMapOutput {
	return o
}

func (o SourceSurveySparrowMapOutput) MapIndex(k pulumi.StringInput) SourceSurveySparrowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceSurveySparrow {
		return vs[0].(map[string]*SourceSurveySparrow)[vs[1].(string)]
	}).(SourceSurveySparrowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSurveySparrowInput)(nil)).Elem(), &SourceSurveySparrow{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSurveySparrowArrayInput)(nil)).Elem(), SourceSurveySparrowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSurveySparrowMapInput)(nil)).Elem(), SourceSurveySparrowMap{})
	pulumi.RegisterOutputType(SourceSurveySparrowOutput{})
	pulumi.RegisterOutputType(SourceSurveySparrowArrayOutput{})
	pulumi.RegisterOutputType(SourceSurveySparrowMapOutput{})
}
