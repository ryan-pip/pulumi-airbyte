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

// SourceStripe Resource
type SourceStripe struct {
	pulumi.CustomResourceState

	Configuration SourceStripeConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceStripe registers a new resource with the given unique name, arguments, and options.
func NewSourceStripe(ctx *pulumi.Context,
	name string, args *SourceStripeArgs, opts ...pulumi.ResourceOption) (*SourceStripe, error) {
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
	var resource SourceStripe
	err := ctx.RegisterResource("airbyte:index/sourceStripe:SourceStripe", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceStripe gets an existing SourceStripe resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceStripe(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceStripeState, opts ...pulumi.ResourceOption) (*SourceStripe, error) {
	var resource SourceStripe
	err := ctx.ReadResource("airbyte:index/sourceStripe:SourceStripe", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceStripe resources.
type sourceStripeState struct {
	Configuration *SourceStripeConfiguration `pulumi:"configuration"`
	Name          *string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceStripeState struct {
	Configuration SourceStripeConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceStripeState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceStripeState)(nil)).Elem()
}

type sourceStripeArgs struct {
	Configuration SourceStripeConfiguration `pulumi:"configuration"`
	Name          string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceStripe resource.
type SourceStripeArgs struct {
	Configuration SourceStripeConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceStripeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceStripeArgs)(nil)).Elem()
}

type SourceStripeInput interface {
	pulumi.Input

	ToSourceStripeOutput() SourceStripeOutput
	ToSourceStripeOutputWithContext(ctx context.Context) SourceStripeOutput
}

func (*SourceStripe) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceStripe)(nil)).Elem()
}

func (i *SourceStripe) ToSourceStripeOutput() SourceStripeOutput {
	return i.ToSourceStripeOutputWithContext(context.Background())
}

func (i *SourceStripe) ToSourceStripeOutputWithContext(ctx context.Context) SourceStripeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceStripeOutput)
}

// SourceStripeArrayInput is an input type that accepts SourceStripeArray and SourceStripeArrayOutput values.
// You can construct a concrete instance of `SourceStripeArrayInput` via:
//
//	SourceStripeArray{ SourceStripeArgs{...} }
type SourceStripeArrayInput interface {
	pulumi.Input

	ToSourceStripeArrayOutput() SourceStripeArrayOutput
	ToSourceStripeArrayOutputWithContext(context.Context) SourceStripeArrayOutput
}

type SourceStripeArray []SourceStripeInput

func (SourceStripeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceStripe)(nil)).Elem()
}

func (i SourceStripeArray) ToSourceStripeArrayOutput() SourceStripeArrayOutput {
	return i.ToSourceStripeArrayOutputWithContext(context.Background())
}

func (i SourceStripeArray) ToSourceStripeArrayOutputWithContext(ctx context.Context) SourceStripeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceStripeArrayOutput)
}

// SourceStripeMapInput is an input type that accepts SourceStripeMap and SourceStripeMapOutput values.
// You can construct a concrete instance of `SourceStripeMapInput` via:
//
//	SourceStripeMap{ "key": SourceStripeArgs{...} }
type SourceStripeMapInput interface {
	pulumi.Input

	ToSourceStripeMapOutput() SourceStripeMapOutput
	ToSourceStripeMapOutputWithContext(context.Context) SourceStripeMapOutput
}

type SourceStripeMap map[string]SourceStripeInput

func (SourceStripeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceStripe)(nil)).Elem()
}

func (i SourceStripeMap) ToSourceStripeMapOutput() SourceStripeMapOutput {
	return i.ToSourceStripeMapOutputWithContext(context.Background())
}

func (i SourceStripeMap) ToSourceStripeMapOutputWithContext(ctx context.Context) SourceStripeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceStripeMapOutput)
}

type SourceStripeOutput struct{ *pulumi.OutputState }

func (SourceStripeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceStripe)(nil)).Elem()
}

func (o SourceStripeOutput) ToSourceStripeOutput() SourceStripeOutput {
	return o
}

func (o SourceStripeOutput) ToSourceStripeOutputWithContext(ctx context.Context) SourceStripeOutput {
	return o
}

func (o SourceStripeOutput) Configuration() SourceStripeConfigurationOutput {
	return o.ApplyT(func(v *SourceStripe) SourceStripeConfigurationOutput { return v.Configuration }).(SourceStripeConfigurationOutput)
}

func (o SourceStripeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceStripe) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceStripeOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceStripe) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceStripeOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceStripe) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceStripeOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceStripe) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceStripeOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceStripe) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceStripeArrayOutput struct{ *pulumi.OutputState }

func (SourceStripeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceStripe)(nil)).Elem()
}

func (o SourceStripeArrayOutput) ToSourceStripeArrayOutput() SourceStripeArrayOutput {
	return o
}

func (o SourceStripeArrayOutput) ToSourceStripeArrayOutputWithContext(ctx context.Context) SourceStripeArrayOutput {
	return o
}

func (o SourceStripeArrayOutput) Index(i pulumi.IntInput) SourceStripeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceStripe {
		return vs[0].([]*SourceStripe)[vs[1].(int)]
	}).(SourceStripeOutput)
}

type SourceStripeMapOutput struct{ *pulumi.OutputState }

func (SourceStripeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceStripe)(nil)).Elem()
}

func (o SourceStripeMapOutput) ToSourceStripeMapOutput() SourceStripeMapOutput {
	return o
}

func (o SourceStripeMapOutput) ToSourceStripeMapOutputWithContext(ctx context.Context) SourceStripeMapOutput {
	return o
}

func (o SourceStripeMapOutput) MapIndex(k pulumi.StringInput) SourceStripeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceStripe {
		return vs[0].(map[string]*SourceStripe)[vs[1].(string)]
	}).(SourceStripeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceStripeInput)(nil)).Elem(), &SourceStripe{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceStripeArrayInput)(nil)).Elem(), SourceStripeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceStripeMapInput)(nil)).Elem(), SourceStripeMap{})
	pulumi.RegisterOutputType(SourceStripeOutput{})
	pulumi.RegisterOutputType(SourceStripeArrayOutput{})
	pulumi.RegisterOutputType(SourceStripeMapOutput{})
}
