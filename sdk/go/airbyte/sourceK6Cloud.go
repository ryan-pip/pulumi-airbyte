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

// SourceK6Cloud Resource
type SourceK6Cloud struct {
	pulumi.CustomResourceState

	Configuration SourceK6CloudConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceK6Cloud registers a new resource with the given unique name, arguments, and options.
func NewSourceK6Cloud(ctx *pulumi.Context,
	name string, args *SourceK6CloudArgs, opts ...pulumi.ResourceOption) (*SourceK6Cloud, error) {
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
	var resource SourceK6Cloud
	err := ctx.RegisterResource("airbyte:index/sourceK6Cloud:SourceK6Cloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceK6Cloud gets an existing SourceK6Cloud resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceK6Cloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceK6CloudState, opts ...pulumi.ResourceOption) (*SourceK6Cloud, error) {
	var resource SourceK6Cloud
	err := ctx.ReadResource("airbyte:index/sourceK6Cloud:SourceK6Cloud", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceK6Cloud resources.
type sourceK6CloudState struct {
	Configuration *SourceK6CloudConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceK6CloudState struct {
	Configuration SourceK6CloudConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceK6CloudState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceK6CloudState)(nil)).Elem()
}

type sourceK6CloudArgs struct {
	Configuration SourceK6CloudConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceK6Cloud resource.
type SourceK6CloudArgs struct {
	Configuration SourceK6CloudConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceK6CloudArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceK6CloudArgs)(nil)).Elem()
}

type SourceK6CloudInput interface {
	pulumi.Input

	ToSourceK6CloudOutput() SourceK6CloudOutput
	ToSourceK6CloudOutputWithContext(ctx context.Context) SourceK6CloudOutput
}

func (*SourceK6Cloud) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceK6Cloud)(nil)).Elem()
}

func (i *SourceK6Cloud) ToSourceK6CloudOutput() SourceK6CloudOutput {
	return i.ToSourceK6CloudOutputWithContext(context.Background())
}

func (i *SourceK6Cloud) ToSourceK6CloudOutputWithContext(ctx context.Context) SourceK6CloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceK6CloudOutput)
}

// SourceK6CloudArrayInput is an input type that accepts SourceK6CloudArray and SourceK6CloudArrayOutput values.
// You can construct a concrete instance of `SourceK6CloudArrayInput` via:
//
//	SourceK6CloudArray{ SourceK6CloudArgs{...} }
type SourceK6CloudArrayInput interface {
	pulumi.Input

	ToSourceK6CloudArrayOutput() SourceK6CloudArrayOutput
	ToSourceK6CloudArrayOutputWithContext(context.Context) SourceK6CloudArrayOutput
}

type SourceK6CloudArray []SourceK6CloudInput

func (SourceK6CloudArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceK6Cloud)(nil)).Elem()
}

func (i SourceK6CloudArray) ToSourceK6CloudArrayOutput() SourceK6CloudArrayOutput {
	return i.ToSourceK6CloudArrayOutputWithContext(context.Background())
}

func (i SourceK6CloudArray) ToSourceK6CloudArrayOutputWithContext(ctx context.Context) SourceK6CloudArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceK6CloudArrayOutput)
}

// SourceK6CloudMapInput is an input type that accepts SourceK6CloudMap and SourceK6CloudMapOutput values.
// You can construct a concrete instance of `SourceK6CloudMapInput` via:
//
//	SourceK6CloudMap{ "key": SourceK6CloudArgs{...} }
type SourceK6CloudMapInput interface {
	pulumi.Input

	ToSourceK6CloudMapOutput() SourceK6CloudMapOutput
	ToSourceK6CloudMapOutputWithContext(context.Context) SourceK6CloudMapOutput
}

type SourceK6CloudMap map[string]SourceK6CloudInput

func (SourceK6CloudMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceK6Cloud)(nil)).Elem()
}

func (i SourceK6CloudMap) ToSourceK6CloudMapOutput() SourceK6CloudMapOutput {
	return i.ToSourceK6CloudMapOutputWithContext(context.Background())
}

func (i SourceK6CloudMap) ToSourceK6CloudMapOutputWithContext(ctx context.Context) SourceK6CloudMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceK6CloudMapOutput)
}

type SourceK6CloudOutput struct{ *pulumi.OutputState }

func (SourceK6CloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceK6Cloud)(nil)).Elem()
}

func (o SourceK6CloudOutput) ToSourceK6CloudOutput() SourceK6CloudOutput {
	return o
}

func (o SourceK6CloudOutput) ToSourceK6CloudOutputWithContext(ctx context.Context) SourceK6CloudOutput {
	return o
}

func (o SourceK6CloudOutput) Configuration() SourceK6CloudConfigurationOutput {
	return o.ApplyT(func(v *SourceK6Cloud) SourceK6CloudConfigurationOutput { return v.Configuration }).(SourceK6CloudConfigurationOutput)
}

func (o SourceK6CloudOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceK6Cloud) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceK6CloudOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceK6Cloud) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceK6CloudOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceK6Cloud) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceK6CloudOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceK6Cloud) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceK6CloudOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceK6Cloud) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceK6CloudArrayOutput struct{ *pulumi.OutputState }

func (SourceK6CloudArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceK6Cloud)(nil)).Elem()
}

func (o SourceK6CloudArrayOutput) ToSourceK6CloudArrayOutput() SourceK6CloudArrayOutput {
	return o
}

func (o SourceK6CloudArrayOutput) ToSourceK6CloudArrayOutputWithContext(ctx context.Context) SourceK6CloudArrayOutput {
	return o
}

func (o SourceK6CloudArrayOutput) Index(i pulumi.IntInput) SourceK6CloudOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceK6Cloud {
		return vs[0].([]*SourceK6Cloud)[vs[1].(int)]
	}).(SourceK6CloudOutput)
}

type SourceK6CloudMapOutput struct{ *pulumi.OutputState }

func (SourceK6CloudMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceK6Cloud)(nil)).Elem()
}

func (o SourceK6CloudMapOutput) ToSourceK6CloudMapOutput() SourceK6CloudMapOutput {
	return o
}

func (o SourceK6CloudMapOutput) ToSourceK6CloudMapOutputWithContext(ctx context.Context) SourceK6CloudMapOutput {
	return o
}

func (o SourceK6CloudMapOutput) MapIndex(k pulumi.StringInput) SourceK6CloudOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceK6Cloud {
		return vs[0].(map[string]*SourceK6Cloud)[vs[1].(string)]
	}).(SourceK6CloudOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceK6CloudInput)(nil)).Elem(), &SourceK6Cloud{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceK6CloudArrayInput)(nil)).Elem(), SourceK6CloudArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceK6CloudMapInput)(nil)).Elem(), SourceK6CloudMap{})
	pulumi.RegisterOutputType(SourceK6CloudOutput{})
	pulumi.RegisterOutputType(SourceK6CloudArrayOutput{})
	pulumi.RegisterOutputType(SourceK6CloudMapOutput{})
}
