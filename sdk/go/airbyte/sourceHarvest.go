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

// SourceHarvest Resource
type SourceHarvest struct {
	pulumi.CustomResourceState

	Configuration SourceHarvestConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceHarvest registers a new resource with the given unique name, arguments, and options.
func NewSourceHarvest(ctx *pulumi.Context,
	name string, args *SourceHarvestArgs, opts ...pulumi.ResourceOption) (*SourceHarvest, error) {
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
	var resource SourceHarvest
	err := ctx.RegisterResource("airbyte:index/sourceHarvest:SourceHarvest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceHarvest gets an existing SourceHarvest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceHarvest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceHarvestState, opts ...pulumi.ResourceOption) (*SourceHarvest, error) {
	var resource SourceHarvest
	err := ctx.ReadResource("airbyte:index/sourceHarvest:SourceHarvest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceHarvest resources.
type sourceHarvestState struct {
	Configuration *SourceHarvestConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceHarvestState struct {
	Configuration SourceHarvestConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceHarvestState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceHarvestState)(nil)).Elem()
}

type sourceHarvestArgs struct {
	Configuration SourceHarvestConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceHarvest resource.
type SourceHarvestArgs struct {
	Configuration SourceHarvestConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceHarvestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceHarvestArgs)(nil)).Elem()
}

type SourceHarvestInput interface {
	pulumi.Input

	ToSourceHarvestOutput() SourceHarvestOutput
	ToSourceHarvestOutputWithContext(ctx context.Context) SourceHarvestOutput
}

func (*SourceHarvest) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceHarvest)(nil)).Elem()
}

func (i *SourceHarvest) ToSourceHarvestOutput() SourceHarvestOutput {
	return i.ToSourceHarvestOutputWithContext(context.Background())
}

func (i *SourceHarvest) ToSourceHarvestOutputWithContext(ctx context.Context) SourceHarvestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceHarvestOutput)
}

// SourceHarvestArrayInput is an input type that accepts SourceHarvestArray and SourceHarvestArrayOutput values.
// You can construct a concrete instance of `SourceHarvestArrayInput` via:
//
//	SourceHarvestArray{ SourceHarvestArgs{...} }
type SourceHarvestArrayInput interface {
	pulumi.Input

	ToSourceHarvestArrayOutput() SourceHarvestArrayOutput
	ToSourceHarvestArrayOutputWithContext(context.Context) SourceHarvestArrayOutput
}

type SourceHarvestArray []SourceHarvestInput

func (SourceHarvestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceHarvest)(nil)).Elem()
}

func (i SourceHarvestArray) ToSourceHarvestArrayOutput() SourceHarvestArrayOutput {
	return i.ToSourceHarvestArrayOutputWithContext(context.Background())
}

func (i SourceHarvestArray) ToSourceHarvestArrayOutputWithContext(ctx context.Context) SourceHarvestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceHarvestArrayOutput)
}

// SourceHarvestMapInput is an input type that accepts SourceHarvestMap and SourceHarvestMapOutput values.
// You can construct a concrete instance of `SourceHarvestMapInput` via:
//
//	SourceHarvestMap{ "key": SourceHarvestArgs{...} }
type SourceHarvestMapInput interface {
	pulumi.Input

	ToSourceHarvestMapOutput() SourceHarvestMapOutput
	ToSourceHarvestMapOutputWithContext(context.Context) SourceHarvestMapOutput
}

type SourceHarvestMap map[string]SourceHarvestInput

func (SourceHarvestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceHarvest)(nil)).Elem()
}

func (i SourceHarvestMap) ToSourceHarvestMapOutput() SourceHarvestMapOutput {
	return i.ToSourceHarvestMapOutputWithContext(context.Background())
}

func (i SourceHarvestMap) ToSourceHarvestMapOutputWithContext(ctx context.Context) SourceHarvestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceHarvestMapOutput)
}

type SourceHarvestOutput struct{ *pulumi.OutputState }

func (SourceHarvestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceHarvest)(nil)).Elem()
}

func (o SourceHarvestOutput) ToSourceHarvestOutput() SourceHarvestOutput {
	return o
}

func (o SourceHarvestOutput) ToSourceHarvestOutputWithContext(ctx context.Context) SourceHarvestOutput {
	return o
}

func (o SourceHarvestOutput) Configuration() SourceHarvestConfigurationOutput {
	return o.ApplyT(func(v *SourceHarvest) SourceHarvestConfigurationOutput { return v.Configuration }).(SourceHarvestConfigurationOutput)
}

func (o SourceHarvestOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceHarvest) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceHarvestOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceHarvest) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceHarvestOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceHarvest) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceHarvestOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceHarvest) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceHarvestOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceHarvest) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceHarvestArrayOutput struct{ *pulumi.OutputState }

func (SourceHarvestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceHarvest)(nil)).Elem()
}

func (o SourceHarvestArrayOutput) ToSourceHarvestArrayOutput() SourceHarvestArrayOutput {
	return o
}

func (o SourceHarvestArrayOutput) ToSourceHarvestArrayOutputWithContext(ctx context.Context) SourceHarvestArrayOutput {
	return o
}

func (o SourceHarvestArrayOutput) Index(i pulumi.IntInput) SourceHarvestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceHarvest {
		return vs[0].([]*SourceHarvest)[vs[1].(int)]
	}).(SourceHarvestOutput)
}

type SourceHarvestMapOutput struct{ *pulumi.OutputState }

func (SourceHarvestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceHarvest)(nil)).Elem()
}

func (o SourceHarvestMapOutput) ToSourceHarvestMapOutput() SourceHarvestMapOutput {
	return o
}

func (o SourceHarvestMapOutput) ToSourceHarvestMapOutputWithContext(ctx context.Context) SourceHarvestMapOutput {
	return o
}

func (o SourceHarvestMapOutput) MapIndex(k pulumi.StringInput) SourceHarvestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceHarvest {
		return vs[0].(map[string]*SourceHarvest)[vs[1].(string)]
	}).(SourceHarvestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceHarvestInput)(nil)).Elem(), &SourceHarvest{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceHarvestArrayInput)(nil)).Elem(), SourceHarvestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceHarvestMapInput)(nil)).Elem(), SourceHarvestMap{})
	pulumi.RegisterOutputType(SourceHarvestOutput{})
	pulumi.RegisterOutputType(SourceHarvestArrayOutput{})
	pulumi.RegisterOutputType(SourceHarvestMapOutput{})
}
