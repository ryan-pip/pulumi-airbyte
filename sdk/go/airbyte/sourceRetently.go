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

// SourceRetently Resource
type SourceRetently struct {
	pulumi.CustomResourceState

	Configuration SourceRetentlyConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceRetently registers a new resource with the given unique name, arguments, and options.
func NewSourceRetently(ctx *pulumi.Context,
	name string, args *SourceRetentlyArgs, opts ...pulumi.ResourceOption) (*SourceRetently, error) {
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
	var resource SourceRetently
	err := ctx.RegisterResource("airbyte:index/sourceRetently:SourceRetently", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceRetently gets an existing SourceRetently resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceRetently(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceRetentlyState, opts ...pulumi.ResourceOption) (*SourceRetently, error) {
	var resource SourceRetently
	err := ctx.ReadResource("airbyte:index/sourceRetently:SourceRetently", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceRetently resources.
type sourceRetentlyState struct {
	Configuration *SourceRetentlyConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceRetentlyState struct {
	Configuration SourceRetentlyConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceRetentlyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRetentlyState)(nil)).Elem()
}

type sourceRetentlyArgs struct {
	Configuration SourceRetentlyConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceRetently resource.
type SourceRetentlyArgs struct {
	Configuration SourceRetentlyConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceRetentlyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRetentlyArgs)(nil)).Elem()
}

type SourceRetentlyInput interface {
	pulumi.Input

	ToSourceRetentlyOutput() SourceRetentlyOutput
	ToSourceRetentlyOutputWithContext(ctx context.Context) SourceRetentlyOutput
}

func (*SourceRetently) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRetently)(nil)).Elem()
}

func (i *SourceRetently) ToSourceRetentlyOutput() SourceRetentlyOutput {
	return i.ToSourceRetentlyOutputWithContext(context.Background())
}

func (i *SourceRetently) ToSourceRetentlyOutputWithContext(ctx context.Context) SourceRetentlyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRetentlyOutput)
}

// SourceRetentlyArrayInput is an input type that accepts SourceRetentlyArray and SourceRetentlyArrayOutput values.
// You can construct a concrete instance of `SourceRetentlyArrayInput` via:
//
//	SourceRetentlyArray{ SourceRetentlyArgs{...} }
type SourceRetentlyArrayInput interface {
	pulumi.Input

	ToSourceRetentlyArrayOutput() SourceRetentlyArrayOutput
	ToSourceRetentlyArrayOutputWithContext(context.Context) SourceRetentlyArrayOutput
}

type SourceRetentlyArray []SourceRetentlyInput

func (SourceRetentlyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceRetently)(nil)).Elem()
}

func (i SourceRetentlyArray) ToSourceRetentlyArrayOutput() SourceRetentlyArrayOutput {
	return i.ToSourceRetentlyArrayOutputWithContext(context.Background())
}

func (i SourceRetentlyArray) ToSourceRetentlyArrayOutputWithContext(ctx context.Context) SourceRetentlyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRetentlyArrayOutput)
}

// SourceRetentlyMapInput is an input type that accepts SourceRetentlyMap and SourceRetentlyMapOutput values.
// You can construct a concrete instance of `SourceRetentlyMapInput` via:
//
//	SourceRetentlyMap{ "key": SourceRetentlyArgs{...} }
type SourceRetentlyMapInput interface {
	pulumi.Input

	ToSourceRetentlyMapOutput() SourceRetentlyMapOutput
	ToSourceRetentlyMapOutputWithContext(context.Context) SourceRetentlyMapOutput
}

type SourceRetentlyMap map[string]SourceRetentlyInput

func (SourceRetentlyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceRetently)(nil)).Elem()
}

func (i SourceRetentlyMap) ToSourceRetentlyMapOutput() SourceRetentlyMapOutput {
	return i.ToSourceRetentlyMapOutputWithContext(context.Background())
}

func (i SourceRetentlyMap) ToSourceRetentlyMapOutputWithContext(ctx context.Context) SourceRetentlyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRetentlyMapOutput)
}

type SourceRetentlyOutput struct{ *pulumi.OutputState }

func (SourceRetentlyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRetently)(nil)).Elem()
}

func (o SourceRetentlyOutput) ToSourceRetentlyOutput() SourceRetentlyOutput {
	return o
}

func (o SourceRetentlyOutput) ToSourceRetentlyOutputWithContext(ctx context.Context) SourceRetentlyOutput {
	return o
}

func (o SourceRetentlyOutput) Configuration() SourceRetentlyConfigurationOutput {
	return o.ApplyT(func(v *SourceRetently) SourceRetentlyConfigurationOutput { return v.Configuration }).(SourceRetentlyConfigurationOutput)
}

func (o SourceRetentlyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRetently) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceRetentlyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceRetently) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceRetentlyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRetently) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceRetentlyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRetently) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceRetentlyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRetently) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceRetentlyArrayOutput struct{ *pulumi.OutputState }

func (SourceRetentlyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceRetently)(nil)).Elem()
}

func (o SourceRetentlyArrayOutput) ToSourceRetentlyArrayOutput() SourceRetentlyArrayOutput {
	return o
}

func (o SourceRetentlyArrayOutput) ToSourceRetentlyArrayOutputWithContext(ctx context.Context) SourceRetentlyArrayOutput {
	return o
}

func (o SourceRetentlyArrayOutput) Index(i pulumi.IntInput) SourceRetentlyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceRetently {
		return vs[0].([]*SourceRetently)[vs[1].(int)]
	}).(SourceRetentlyOutput)
}

type SourceRetentlyMapOutput struct{ *pulumi.OutputState }

func (SourceRetentlyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceRetently)(nil)).Elem()
}

func (o SourceRetentlyMapOutput) ToSourceRetentlyMapOutput() SourceRetentlyMapOutput {
	return o
}

func (o SourceRetentlyMapOutput) ToSourceRetentlyMapOutputWithContext(ctx context.Context) SourceRetentlyMapOutput {
	return o
}

func (o SourceRetentlyMapOutput) MapIndex(k pulumi.StringInput) SourceRetentlyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceRetently {
		return vs[0].(map[string]*SourceRetently)[vs[1].(string)]
	}).(SourceRetentlyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRetentlyInput)(nil)).Elem(), &SourceRetently{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRetentlyArrayInput)(nil)).Elem(), SourceRetentlyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRetentlyMapInput)(nil)).Elem(), SourceRetentlyMap{})
	pulumi.RegisterOutputType(SourceRetentlyOutput{})
	pulumi.RegisterOutputType(SourceRetentlyArrayOutput{})
	pulumi.RegisterOutputType(SourceRetentlyMapOutput{})
}
