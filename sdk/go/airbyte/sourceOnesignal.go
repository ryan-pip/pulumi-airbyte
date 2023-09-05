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

// SourceOnesignal Resource
type SourceOnesignal struct {
	pulumi.CustomResourceState

	Configuration SourceOnesignalConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceOnesignal registers a new resource with the given unique name, arguments, and options.
func NewSourceOnesignal(ctx *pulumi.Context,
	name string, args *SourceOnesignalArgs, opts ...pulumi.ResourceOption) (*SourceOnesignal, error) {
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
	var resource SourceOnesignal
	err := ctx.RegisterResource("airbyte:index/sourceOnesignal:SourceOnesignal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceOnesignal gets an existing SourceOnesignal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceOnesignal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceOnesignalState, opts ...pulumi.ResourceOption) (*SourceOnesignal, error) {
	var resource SourceOnesignal
	err := ctx.ReadResource("airbyte:index/sourceOnesignal:SourceOnesignal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceOnesignal resources.
type sourceOnesignalState struct {
	Configuration *SourceOnesignalConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceOnesignalState struct {
	Configuration SourceOnesignalConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceOnesignalState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOnesignalState)(nil)).Elem()
}

type sourceOnesignalArgs struct {
	Configuration SourceOnesignalConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceOnesignal resource.
type SourceOnesignalArgs struct {
	Configuration SourceOnesignalConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceOnesignalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOnesignalArgs)(nil)).Elem()
}

type SourceOnesignalInput interface {
	pulumi.Input

	ToSourceOnesignalOutput() SourceOnesignalOutput
	ToSourceOnesignalOutputWithContext(ctx context.Context) SourceOnesignalOutput
}

func (*SourceOnesignal) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOnesignal)(nil)).Elem()
}

func (i *SourceOnesignal) ToSourceOnesignalOutput() SourceOnesignalOutput {
	return i.ToSourceOnesignalOutputWithContext(context.Background())
}

func (i *SourceOnesignal) ToSourceOnesignalOutputWithContext(ctx context.Context) SourceOnesignalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOnesignalOutput)
}

// SourceOnesignalArrayInput is an input type that accepts SourceOnesignalArray and SourceOnesignalArrayOutput values.
// You can construct a concrete instance of `SourceOnesignalArrayInput` via:
//
//	SourceOnesignalArray{ SourceOnesignalArgs{...} }
type SourceOnesignalArrayInput interface {
	pulumi.Input

	ToSourceOnesignalArrayOutput() SourceOnesignalArrayOutput
	ToSourceOnesignalArrayOutputWithContext(context.Context) SourceOnesignalArrayOutput
}

type SourceOnesignalArray []SourceOnesignalInput

func (SourceOnesignalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceOnesignal)(nil)).Elem()
}

func (i SourceOnesignalArray) ToSourceOnesignalArrayOutput() SourceOnesignalArrayOutput {
	return i.ToSourceOnesignalArrayOutputWithContext(context.Background())
}

func (i SourceOnesignalArray) ToSourceOnesignalArrayOutputWithContext(ctx context.Context) SourceOnesignalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOnesignalArrayOutput)
}

// SourceOnesignalMapInput is an input type that accepts SourceOnesignalMap and SourceOnesignalMapOutput values.
// You can construct a concrete instance of `SourceOnesignalMapInput` via:
//
//	SourceOnesignalMap{ "key": SourceOnesignalArgs{...} }
type SourceOnesignalMapInput interface {
	pulumi.Input

	ToSourceOnesignalMapOutput() SourceOnesignalMapOutput
	ToSourceOnesignalMapOutputWithContext(context.Context) SourceOnesignalMapOutput
}

type SourceOnesignalMap map[string]SourceOnesignalInput

func (SourceOnesignalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceOnesignal)(nil)).Elem()
}

func (i SourceOnesignalMap) ToSourceOnesignalMapOutput() SourceOnesignalMapOutput {
	return i.ToSourceOnesignalMapOutputWithContext(context.Background())
}

func (i SourceOnesignalMap) ToSourceOnesignalMapOutputWithContext(ctx context.Context) SourceOnesignalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOnesignalMapOutput)
}

type SourceOnesignalOutput struct{ *pulumi.OutputState }

func (SourceOnesignalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOnesignal)(nil)).Elem()
}

func (o SourceOnesignalOutput) ToSourceOnesignalOutput() SourceOnesignalOutput {
	return o
}

func (o SourceOnesignalOutput) ToSourceOnesignalOutputWithContext(ctx context.Context) SourceOnesignalOutput {
	return o
}

func (o SourceOnesignalOutput) Configuration() SourceOnesignalConfigurationOutput {
	return o.ApplyT(func(v *SourceOnesignal) SourceOnesignalConfigurationOutput { return v.Configuration }).(SourceOnesignalConfigurationOutput)
}

func (o SourceOnesignalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOnesignal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceOnesignalOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOnesignal) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceOnesignalOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOnesignal) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceOnesignalOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOnesignal) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceOnesignalOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOnesignal) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceOnesignalArrayOutput struct{ *pulumi.OutputState }

func (SourceOnesignalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceOnesignal)(nil)).Elem()
}

func (o SourceOnesignalArrayOutput) ToSourceOnesignalArrayOutput() SourceOnesignalArrayOutput {
	return o
}

func (o SourceOnesignalArrayOutput) ToSourceOnesignalArrayOutputWithContext(ctx context.Context) SourceOnesignalArrayOutput {
	return o
}

func (o SourceOnesignalArrayOutput) Index(i pulumi.IntInput) SourceOnesignalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceOnesignal {
		return vs[0].([]*SourceOnesignal)[vs[1].(int)]
	}).(SourceOnesignalOutput)
}

type SourceOnesignalMapOutput struct{ *pulumi.OutputState }

func (SourceOnesignalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceOnesignal)(nil)).Elem()
}

func (o SourceOnesignalMapOutput) ToSourceOnesignalMapOutput() SourceOnesignalMapOutput {
	return o
}

func (o SourceOnesignalMapOutput) ToSourceOnesignalMapOutputWithContext(ctx context.Context) SourceOnesignalMapOutput {
	return o
}

func (o SourceOnesignalMapOutput) MapIndex(k pulumi.StringInput) SourceOnesignalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceOnesignal {
		return vs[0].(map[string]*SourceOnesignal)[vs[1].(string)]
	}).(SourceOnesignalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOnesignalInput)(nil)).Elem(), &SourceOnesignal{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOnesignalArrayInput)(nil)).Elem(), SourceOnesignalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOnesignalMapInput)(nil)).Elem(), SourceOnesignalMap{})
	pulumi.RegisterOutputType(SourceOnesignalOutput{})
	pulumi.RegisterOutputType(SourceOnesignalArrayOutput{})
	pulumi.RegisterOutputType(SourceOnesignalMapOutput{})
}
