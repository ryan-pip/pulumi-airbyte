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

// SourceZendeskSupport Resource
type SourceZendeskSupport struct {
	pulumi.CustomResourceState

	Configuration SourceZendeskSupportConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceZendeskSupport registers a new resource with the given unique name, arguments, and options.
func NewSourceZendeskSupport(ctx *pulumi.Context,
	name string, args *SourceZendeskSupportArgs, opts ...pulumi.ResourceOption) (*SourceZendeskSupport, error) {
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
	var resource SourceZendeskSupport
	err := ctx.RegisterResource("airbyte:index/sourceZendeskSupport:SourceZendeskSupport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceZendeskSupport gets an existing SourceZendeskSupport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceZendeskSupport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceZendeskSupportState, opts ...pulumi.ResourceOption) (*SourceZendeskSupport, error) {
	var resource SourceZendeskSupport
	err := ctx.ReadResource("airbyte:index/sourceZendeskSupport:SourceZendeskSupport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceZendeskSupport resources.
type sourceZendeskSupportState struct {
	Configuration *SourceZendeskSupportConfiguration `pulumi:"configuration"`
	Name          *string                            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceZendeskSupportState struct {
	Configuration SourceZendeskSupportConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceZendeskSupportState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceZendeskSupportState)(nil)).Elem()
}

type sourceZendeskSupportArgs struct {
	Configuration SourceZendeskSupportConfiguration `pulumi:"configuration"`
	Name          string                            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceZendeskSupport resource.
type SourceZendeskSupportArgs struct {
	Configuration SourceZendeskSupportConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceZendeskSupportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceZendeskSupportArgs)(nil)).Elem()
}

type SourceZendeskSupportInput interface {
	pulumi.Input

	ToSourceZendeskSupportOutput() SourceZendeskSupportOutput
	ToSourceZendeskSupportOutputWithContext(ctx context.Context) SourceZendeskSupportOutput
}

func (*SourceZendeskSupport) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceZendeskSupport)(nil)).Elem()
}

func (i *SourceZendeskSupport) ToSourceZendeskSupportOutput() SourceZendeskSupportOutput {
	return i.ToSourceZendeskSupportOutputWithContext(context.Background())
}

func (i *SourceZendeskSupport) ToSourceZendeskSupportOutputWithContext(ctx context.Context) SourceZendeskSupportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceZendeskSupportOutput)
}

// SourceZendeskSupportArrayInput is an input type that accepts SourceZendeskSupportArray and SourceZendeskSupportArrayOutput values.
// You can construct a concrete instance of `SourceZendeskSupportArrayInput` via:
//
//	SourceZendeskSupportArray{ SourceZendeskSupportArgs{...} }
type SourceZendeskSupportArrayInput interface {
	pulumi.Input

	ToSourceZendeskSupportArrayOutput() SourceZendeskSupportArrayOutput
	ToSourceZendeskSupportArrayOutputWithContext(context.Context) SourceZendeskSupportArrayOutput
}

type SourceZendeskSupportArray []SourceZendeskSupportInput

func (SourceZendeskSupportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceZendeskSupport)(nil)).Elem()
}

func (i SourceZendeskSupportArray) ToSourceZendeskSupportArrayOutput() SourceZendeskSupportArrayOutput {
	return i.ToSourceZendeskSupportArrayOutputWithContext(context.Background())
}

func (i SourceZendeskSupportArray) ToSourceZendeskSupportArrayOutputWithContext(ctx context.Context) SourceZendeskSupportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceZendeskSupportArrayOutput)
}

// SourceZendeskSupportMapInput is an input type that accepts SourceZendeskSupportMap and SourceZendeskSupportMapOutput values.
// You can construct a concrete instance of `SourceZendeskSupportMapInput` via:
//
//	SourceZendeskSupportMap{ "key": SourceZendeskSupportArgs{...} }
type SourceZendeskSupportMapInput interface {
	pulumi.Input

	ToSourceZendeskSupportMapOutput() SourceZendeskSupportMapOutput
	ToSourceZendeskSupportMapOutputWithContext(context.Context) SourceZendeskSupportMapOutput
}

type SourceZendeskSupportMap map[string]SourceZendeskSupportInput

func (SourceZendeskSupportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceZendeskSupport)(nil)).Elem()
}

func (i SourceZendeskSupportMap) ToSourceZendeskSupportMapOutput() SourceZendeskSupportMapOutput {
	return i.ToSourceZendeskSupportMapOutputWithContext(context.Background())
}

func (i SourceZendeskSupportMap) ToSourceZendeskSupportMapOutputWithContext(ctx context.Context) SourceZendeskSupportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceZendeskSupportMapOutput)
}

type SourceZendeskSupportOutput struct{ *pulumi.OutputState }

func (SourceZendeskSupportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceZendeskSupport)(nil)).Elem()
}

func (o SourceZendeskSupportOutput) ToSourceZendeskSupportOutput() SourceZendeskSupportOutput {
	return o
}

func (o SourceZendeskSupportOutput) ToSourceZendeskSupportOutputWithContext(ctx context.Context) SourceZendeskSupportOutput {
	return o
}

func (o SourceZendeskSupportOutput) Configuration() SourceZendeskSupportConfigurationOutput {
	return o.ApplyT(func(v *SourceZendeskSupport) SourceZendeskSupportConfigurationOutput { return v.Configuration }).(SourceZendeskSupportConfigurationOutput)
}

func (o SourceZendeskSupportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZendeskSupport) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceZendeskSupportOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceZendeskSupport) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceZendeskSupportOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZendeskSupport) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceZendeskSupportOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZendeskSupport) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceZendeskSupportOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceZendeskSupport) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceZendeskSupportArrayOutput struct{ *pulumi.OutputState }

func (SourceZendeskSupportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceZendeskSupport)(nil)).Elem()
}

func (o SourceZendeskSupportArrayOutput) ToSourceZendeskSupportArrayOutput() SourceZendeskSupportArrayOutput {
	return o
}

func (o SourceZendeskSupportArrayOutput) ToSourceZendeskSupportArrayOutputWithContext(ctx context.Context) SourceZendeskSupportArrayOutput {
	return o
}

func (o SourceZendeskSupportArrayOutput) Index(i pulumi.IntInput) SourceZendeskSupportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceZendeskSupport {
		return vs[0].([]*SourceZendeskSupport)[vs[1].(int)]
	}).(SourceZendeskSupportOutput)
}

type SourceZendeskSupportMapOutput struct{ *pulumi.OutputState }

func (SourceZendeskSupportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceZendeskSupport)(nil)).Elem()
}

func (o SourceZendeskSupportMapOutput) ToSourceZendeskSupportMapOutput() SourceZendeskSupportMapOutput {
	return o
}

func (o SourceZendeskSupportMapOutput) ToSourceZendeskSupportMapOutputWithContext(ctx context.Context) SourceZendeskSupportMapOutput {
	return o
}

func (o SourceZendeskSupportMapOutput) MapIndex(k pulumi.StringInput) SourceZendeskSupportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceZendeskSupport {
		return vs[0].(map[string]*SourceZendeskSupport)[vs[1].(string)]
	}).(SourceZendeskSupportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceZendeskSupportInput)(nil)).Elem(), &SourceZendeskSupport{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceZendeskSupportArrayInput)(nil)).Elem(), SourceZendeskSupportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceZendeskSupportMapInput)(nil)).Elem(), SourceZendeskSupportMap{})
	pulumi.RegisterOutputType(SourceZendeskSupportOutput{})
	pulumi.RegisterOutputType(SourceZendeskSupportArrayOutput{})
	pulumi.RegisterOutputType(SourceZendeskSupportMapOutput{})
}
