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

// DestinationPubsub Resource
type DestinationPubsub struct {
	pulumi.CustomResourceState

	Configuration   DestinationPubsubConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                  `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                  `pulumi:"destinationType"`
	Name            pulumi.StringOutput                  `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                  `pulumi:"workspaceId"`
}

// NewDestinationPubsub registers a new resource with the given unique name, arguments, and options.
func NewDestinationPubsub(ctx *pulumi.Context,
	name string, args *DestinationPubsubArgs, opts ...pulumi.ResourceOption) (*DestinationPubsub, error) {
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
	var resource DestinationPubsub
	err := ctx.RegisterResource("airbyte:index/destinationPubsub:DestinationPubsub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationPubsub gets an existing DestinationPubsub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationPubsub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationPubsubState, opts ...pulumi.ResourceOption) (*DestinationPubsub, error) {
	var resource DestinationPubsub
	err := ctx.ReadResource("airbyte:index/destinationPubsub:DestinationPubsub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationPubsub resources.
type destinationPubsubState struct {
	Configuration   *DestinationPubsubConfiguration `pulumi:"configuration"`
	DestinationId   *string                         `pulumi:"destinationId"`
	DestinationType *string                         `pulumi:"destinationType"`
	Name            *string                         `pulumi:"name"`
	WorkspaceId     *string                         `pulumi:"workspaceId"`
}

type DestinationPubsubState struct {
	Configuration   DestinationPubsubConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationPubsubState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationPubsubState)(nil)).Elem()
}

type destinationPubsubArgs struct {
	Configuration DestinationPubsubConfiguration `pulumi:"configuration"`
	Name          string                         `pulumi:"name"`
	WorkspaceId   string                         `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationPubsub resource.
type DestinationPubsubArgs struct {
	Configuration DestinationPubsubConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationPubsubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationPubsubArgs)(nil)).Elem()
}

type DestinationPubsubInput interface {
	pulumi.Input

	ToDestinationPubsubOutput() DestinationPubsubOutput
	ToDestinationPubsubOutputWithContext(ctx context.Context) DestinationPubsubOutput
}

func (*DestinationPubsub) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationPubsub)(nil)).Elem()
}

func (i *DestinationPubsub) ToDestinationPubsubOutput() DestinationPubsubOutput {
	return i.ToDestinationPubsubOutputWithContext(context.Background())
}

func (i *DestinationPubsub) ToDestinationPubsubOutputWithContext(ctx context.Context) DestinationPubsubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationPubsubOutput)
}

// DestinationPubsubArrayInput is an input type that accepts DestinationPubsubArray and DestinationPubsubArrayOutput values.
// You can construct a concrete instance of `DestinationPubsubArrayInput` via:
//
//	DestinationPubsubArray{ DestinationPubsubArgs{...} }
type DestinationPubsubArrayInput interface {
	pulumi.Input

	ToDestinationPubsubArrayOutput() DestinationPubsubArrayOutput
	ToDestinationPubsubArrayOutputWithContext(context.Context) DestinationPubsubArrayOutput
}

type DestinationPubsubArray []DestinationPubsubInput

func (DestinationPubsubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationPubsub)(nil)).Elem()
}

func (i DestinationPubsubArray) ToDestinationPubsubArrayOutput() DestinationPubsubArrayOutput {
	return i.ToDestinationPubsubArrayOutputWithContext(context.Background())
}

func (i DestinationPubsubArray) ToDestinationPubsubArrayOutputWithContext(ctx context.Context) DestinationPubsubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationPubsubArrayOutput)
}

// DestinationPubsubMapInput is an input type that accepts DestinationPubsubMap and DestinationPubsubMapOutput values.
// You can construct a concrete instance of `DestinationPubsubMapInput` via:
//
//	DestinationPubsubMap{ "key": DestinationPubsubArgs{...} }
type DestinationPubsubMapInput interface {
	pulumi.Input

	ToDestinationPubsubMapOutput() DestinationPubsubMapOutput
	ToDestinationPubsubMapOutputWithContext(context.Context) DestinationPubsubMapOutput
}

type DestinationPubsubMap map[string]DestinationPubsubInput

func (DestinationPubsubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationPubsub)(nil)).Elem()
}

func (i DestinationPubsubMap) ToDestinationPubsubMapOutput() DestinationPubsubMapOutput {
	return i.ToDestinationPubsubMapOutputWithContext(context.Background())
}

func (i DestinationPubsubMap) ToDestinationPubsubMapOutputWithContext(ctx context.Context) DestinationPubsubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationPubsubMapOutput)
}

type DestinationPubsubOutput struct{ *pulumi.OutputState }

func (DestinationPubsubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationPubsub)(nil)).Elem()
}

func (o DestinationPubsubOutput) ToDestinationPubsubOutput() DestinationPubsubOutput {
	return o
}

func (o DestinationPubsubOutput) ToDestinationPubsubOutputWithContext(ctx context.Context) DestinationPubsubOutput {
	return o
}

func (o DestinationPubsubOutput) Configuration() DestinationPubsubConfigurationOutput {
	return o.ApplyT(func(v *DestinationPubsub) DestinationPubsubConfigurationOutput { return v.Configuration }).(DestinationPubsubConfigurationOutput)
}

func (o DestinationPubsubOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationPubsub) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationPubsubOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationPubsub) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationPubsubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationPubsub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationPubsubOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationPubsub) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type DestinationPubsubArrayOutput struct{ *pulumi.OutputState }

func (DestinationPubsubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationPubsub)(nil)).Elem()
}

func (o DestinationPubsubArrayOutput) ToDestinationPubsubArrayOutput() DestinationPubsubArrayOutput {
	return o
}

func (o DestinationPubsubArrayOutput) ToDestinationPubsubArrayOutputWithContext(ctx context.Context) DestinationPubsubArrayOutput {
	return o
}

func (o DestinationPubsubArrayOutput) Index(i pulumi.IntInput) DestinationPubsubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationPubsub {
		return vs[0].([]*DestinationPubsub)[vs[1].(int)]
	}).(DestinationPubsubOutput)
}

type DestinationPubsubMapOutput struct{ *pulumi.OutputState }

func (DestinationPubsubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationPubsub)(nil)).Elem()
}

func (o DestinationPubsubMapOutput) ToDestinationPubsubMapOutput() DestinationPubsubMapOutput {
	return o
}

func (o DestinationPubsubMapOutput) ToDestinationPubsubMapOutputWithContext(ctx context.Context) DestinationPubsubMapOutput {
	return o
}

func (o DestinationPubsubMapOutput) MapIndex(k pulumi.StringInput) DestinationPubsubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationPubsub {
		return vs[0].(map[string]*DestinationPubsub)[vs[1].(string)]
	}).(DestinationPubsubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationPubsubInput)(nil)).Elem(), &DestinationPubsub{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationPubsubArrayInput)(nil)).Elem(), DestinationPubsubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationPubsubMapInput)(nil)).Elem(), DestinationPubsubMap{})
	pulumi.RegisterOutputType(DestinationPubsubOutput{})
	pulumi.RegisterOutputType(DestinationPubsubArrayOutput{})
	pulumi.RegisterOutputType(DestinationPubsubMapOutput{})
}
