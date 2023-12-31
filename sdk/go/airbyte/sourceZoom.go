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

// SourceZoom Resource
type SourceZoom struct {
	pulumi.CustomResourceState

	Configuration SourceZoomConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceZoom registers a new resource with the given unique name, arguments, and options.
func NewSourceZoom(ctx *pulumi.Context,
	name string, args *SourceZoomArgs, opts ...pulumi.ResourceOption) (*SourceZoom, error) {
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
	var resource SourceZoom
	err := ctx.RegisterResource("airbyte:index/sourceZoom:SourceZoom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceZoom gets an existing SourceZoom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceZoom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceZoomState, opts ...pulumi.ResourceOption) (*SourceZoom, error) {
	var resource SourceZoom
	err := ctx.ReadResource("airbyte:index/sourceZoom:SourceZoom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceZoom resources.
type sourceZoomState struct {
	Configuration *SourceZoomConfiguration `pulumi:"configuration"`
	Name          *string                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceZoomState struct {
	Configuration SourceZoomConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceZoomState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceZoomState)(nil)).Elem()
}

type sourceZoomArgs struct {
	Configuration SourceZoomConfiguration `pulumi:"configuration"`
	Name          string                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceZoom resource.
type SourceZoomArgs struct {
	Configuration SourceZoomConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceZoomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceZoomArgs)(nil)).Elem()
}

type SourceZoomInput interface {
	pulumi.Input

	ToSourceZoomOutput() SourceZoomOutput
	ToSourceZoomOutputWithContext(ctx context.Context) SourceZoomOutput
}

func (*SourceZoom) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceZoom)(nil)).Elem()
}

func (i *SourceZoom) ToSourceZoomOutput() SourceZoomOutput {
	return i.ToSourceZoomOutputWithContext(context.Background())
}

func (i *SourceZoom) ToSourceZoomOutputWithContext(ctx context.Context) SourceZoomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceZoomOutput)
}

// SourceZoomArrayInput is an input type that accepts SourceZoomArray and SourceZoomArrayOutput values.
// You can construct a concrete instance of `SourceZoomArrayInput` via:
//
//	SourceZoomArray{ SourceZoomArgs{...} }
type SourceZoomArrayInput interface {
	pulumi.Input

	ToSourceZoomArrayOutput() SourceZoomArrayOutput
	ToSourceZoomArrayOutputWithContext(context.Context) SourceZoomArrayOutput
}

type SourceZoomArray []SourceZoomInput

func (SourceZoomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceZoom)(nil)).Elem()
}

func (i SourceZoomArray) ToSourceZoomArrayOutput() SourceZoomArrayOutput {
	return i.ToSourceZoomArrayOutputWithContext(context.Background())
}

func (i SourceZoomArray) ToSourceZoomArrayOutputWithContext(ctx context.Context) SourceZoomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceZoomArrayOutput)
}

// SourceZoomMapInput is an input type that accepts SourceZoomMap and SourceZoomMapOutput values.
// You can construct a concrete instance of `SourceZoomMapInput` via:
//
//	SourceZoomMap{ "key": SourceZoomArgs{...} }
type SourceZoomMapInput interface {
	pulumi.Input

	ToSourceZoomMapOutput() SourceZoomMapOutput
	ToSourceZoomMapOutputWithContext(context.Context) SourceZoomMapOutput
}

type SourceZoomMap map[string]SourceZoomInput

func (SourceZoomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceZoom)(nil)).Elem()
}

func (i SourceZoomMap) ToSourceZoomMapOutput() SourceZoomMapOutput {
	return i.ToSourceZoomMapOutputWithContext(context.Background())
}

func (i SourceZoomMap) ToSourceZoomMapOutputWithContext(ctx context.Context) SourceZoomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceZoomMapOutput)
}

type SourceZoomOutput struct{ *pulumi.OutputState }

func (SourceZoomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceZoom)(nil)).Elem()
}

func (o SourceZoomOutput) ToSourceZoomOutput() SourceZoomOutput {
	return o
}

func (o SourceZoomOutput) ToSourceZoomOutputWithContext(ctx context.Context) SourceZoomOutput {
	return o
}

func (o SourceZoomOutput) Configuration() SourceZoomConfigurationOutput {
	return o.ApplyT(func(v *SourceZoom) SourceZoomConfigurationOutput { return v.Configuration }).(SourceZoomConfigurationOutput)
}

func (o SourceZoomOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZoom) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceZoomOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceZoom) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceZoomOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZoom) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceZoomOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZoom) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceZoomOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZoom) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceZoomArrayOutput struct{ *pulumi.OutputState }

func (SourceZoomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceZoom)(nil)).Elem()
}

func (o SourceZoomArrayOutput) ToSourceZoomArrayOutput() SourceZoomArrayOutput {
	return o
}

func (o SourceZoomArrayOutput) ToSourceZoomArrayOutputWithContext(ctx context.Context) SourceZoomArrayOutput {
	return o
}

func (o SourceZoomArrayOutput) Index(i pulumi.IntInput) SourceZoomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceZoom {
		return vs[0].([]*SourceZoom)[vs[1].(int)]
	}).(SourceZoomOutput)
}

type SourceZoomMapOutput struct{ *pulumi.OutputState }

func (SourceZoomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceZoom)(nil)).Elem()
}

func (o SourceZoomMapOutput) ToSourceZoomMapOutput() SourceZoomMapOutput {
	return o
}

func (o SourceZoomMapOutput) ToSourceZoomMapOutputWithContext(ctx context.Context) SourceZoomMapOutput {
	return o
}

func (o SourceZoomMapOutput) MapIndex(k pulumi.StringInput) SourceZoomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceZoom {
		return vs[0].(map[string]*SourceZoom)[vs[1].(string)]
	}).(SourceZoomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceZoomInput)(nil)).Elem(), &SourceZoom{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceZoomArrayInput)(nil)).Elem(), SourceZoomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceZoomMapInput)(nil)).Elem(), SourceZoomMap{})
	pulumi.RegisterOutputType(SourceZoomOutput{})
	pulumi.RegisterOutputType(SourceZoomArrayOutput{})
	pulumi.RegisterOutputType(SourceZoomMapOutput{})
}
