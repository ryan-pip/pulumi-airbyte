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

// DestinationLangchain Resource
type DestinationLangchain struct {
	pulumi.CustomResourceState

	Configuration   DestinationLangchainConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                     `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                     `pulumi:"destinationType"`
	Name            pulumi.StringOutput                     `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                     `pulumi:"workspaceId"`
}

// NewDestinationLangchain registers a new resource with the given unique name, arguments, and options.
func NewDestinationLangchain(ctx *pulumi.Context,
	name string, args *DestinationLangchainArgs, opts ...pulumi.ResourceOption) (*DestinationLangchain, error) {
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
	var resource DestinationLangchain
	err := ctx.RegisterResource("airbyte:index/destinationLangchain:DestinationLangchain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationLangchain gets an existing DestinationLangchain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationLangchain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationLangchainState, opts ...pulumi.ResourceOption) (*DestinationLangchain, error) {
	var resource DestinationLangchain
	err := ctx.ReadResource("airbyte:index/destinationLangchain:DestinationLangchain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationLangchain resources.
type destinationLangchainState struct {
	Configuration   *DestinationLangchainConfiguration `pulumi:"configuration"`
	DestinationId   *string                            `pulumi:"destinationId"`
	DestinationType *string                            `pulumi:"destinationType"`
	Name            *string                            `pulumi:"name"`
	WorkspaceId     *string                            `pulumi:"workspaceId"`
}

type DestinationLangchainState struct {
	Configuration   DestinationLangchainConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationLangchainState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationLangchainState)(nil)).Elem()
}

type destinationLangchainArgs struct {
	Configuration DestinationLangchainConfiguration `pulumi:"configuration"`
	Name          string                            `pulumi:"name"`
	WorkspaceId   string                            `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationLangchain resource.
type DestinationLangchainArgs struct {
	Configuration DestinationLangchainConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationLangchainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationLangchainArgs)(nil)).Elem()
}

type DestinationLangchainInput interface {
	pulumi.Input

	ToDestinationLangchainOutput() DestinationLangchainOutput
	ToDestinationLangchainOutputWithContext(ctx context.Context) DestinationLangchainOutput
}

func (*DestinationLangchain) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationLangchain)(nil)).Elem()
}

func (i *DestinationLangchain) ToDestinationLangchainOutput() DestinationLangchainOutput {
	return i.ToDestinationLangchainOutputWithContext(context.Background())
}

func (i *DestinationLangchain) ToDestinationLangchainOutputWithContext(ctx context.Context) DestinationLangchainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationLangchainOutput)
}

// DestinationLangchainArrayInput is an input type that accepts DestinationLangchainArray and DestinationLangchainArrayOutput values.
// You can construct a concrete instance of `DestinationLangchainArrayInput` via:
//
//	DestinationLangchainArray{ DestinationLangchainArgs{...} }
type DestinationLangchainArrayInput interface {
	pulumi.Input

	ToDestinationLangchainArrayOutput() DestinationLangchainArrayOutput
	ToDestinationLangchainArrayOutputWithContext(context.Context) DestinationLangchainArrayOutput
}

type DestinationLangchainArray []DestinationLangchainInput

func (DestinationLangchainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationLangchain)(nil)).Elem()
}

func (i DestinationLangchainArray) ToDestinationLangchainArrayOutput() DestinationLangchainArrayOutput {
	return i.ToDestinationLangchainArrayOutputWithContext(context.Background())
}

func (i DestinationLangchainArray) ToDestinationLangchainArrayOutputWithContext(ctx context.Context) DestinationLangchainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationLangchainArrayOutput)
}

// DestinationLangchainMapInput is an input type that accepts DestinationLangchainMap and DestinationLangchainMapOutput values.
// You can construct a concrete instance of `DestinationLangchainMapInput` via:
//
//	DestinationLangchainMap{ "key": DestinationLangchainArgs{...} }
type DestinationLangchainMapInput interface {
	pulumi.Input

	ToDestinationLangchainMapOutput() DestinationLangchainMapOutput
	ToDestinationLangchainMapOutputWithContext(context.Context) DestinationLangchainMapOutput
}

type DestinationLangchainMap map[string]DestinationLangchainInput

func (DestinationLangchainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationLangchain)(nil)).Elem()
}

func (i DestinationLangchainMap) ToDestinationLangchainMapOutput() DestinationLangchainMapOutput {
	return i.ToDestinationLangchainMapOutputWithContext(context.Background())
}

func (i DestinationLangchainMap) ToDestinationLangchainMapOutputWithContext(ctx context.Context) DestinationLangchainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationLangchainMapOutput)
}

type DestinationLangchainOutput struct{ *pulumi.OutputState }

func (DestinationLangchainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationLangchain)(nil)).Elem()
}

func (o DestinationLangchainOutput) ToDestinationLangchainOutput() DestinationLangchainOutput {
	return o
}

func (o DestinationLangchainOutput) ToDestinationLangchainOutputWithContext(ctx context.Context) DestinationLangchainOutput {
	return o
}

func (o DestinationLangchainOutput) Configuration() DestinationLangchainConfigurationOutput {
	return o.ApplyT(func(v *DestinationLangchain) DestinationLangchainConfigurationOutput { return v.Configuration }).(DestinationLangchainConfigurationOutput)
}

func (o DestinationLangchainOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationLangchain) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationLangchainOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationLangchain) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationLangchainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationLangchain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationLangchainOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationLangchain) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type DestinationLangchainArrayOutput struct{ *pulumi.OutputState }

func (DestinationLangchainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationLangchain)(nil)).Elem()
}

func (o DestinationLangchainArrayOutput) ToDestinationLangchainArrayOutput() DestinationLangchainArrayOutput {
	return o
}

func (o DestinationLangchainArrayOutput) ToDestinationLangchainArrayOutputWithContext(ctx context.Context) DestinationLangchainArrayOutput {
	return o
}

func (o DestinationLangchainArrayOutput) Index(i pulumi.IntInput) DestinationLangchainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationLangchain {
		return vs[0].([]*DestinationLangchain)[vs[1].(int)]
	}).(DestinationLangchainOutput)
}

type DestinationLangchainMapOutput struct{ *pulumi.OutputState }

func (DestinationLangchainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationLangchain)(nil)).Elem()
}

func (o DestinationLangchainMapOutput) ToDestinationLangchainMapOutput() DestinationLangchainMapOutput {
	return o
}

func (o DestinationLangchainMapOutput) ToDestinationLangchainMapOutputWithContext(ctx context.Context) DestinationLangchainMapOutput {
	return o
}

func (o DestinationLangchainMapOutput) MapIndex(k pulumi.StringInput) DestinationLangchainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationLangchain {
		return vs[0].(map[string]*DestinationLangchain)[vs[1].(string)]
	}).(DestinationLangchainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationLangchainInput)(nil)).Elem(), &DestinationLangchain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationLangchainArrayInput)(nil)).Elem(), DestinationLangchainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationLangchainMapInput)(nil)).Elem(), DestinationLangchainMap{})
	pulumi.RegisterOutputType(DestinationLangchainOutput{})
	pulumi.RegisterOutputType(DestinationLangchainArrayOutput{})
	pulumi.RegisterOutputType(DestinationLangchainMapOutput{})
}
