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

// SourceSftpBulk Resource
type SourceSftpBulk struct {
	pulumi.CustomResourceState

	Configuration SourceSftpBulkConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceSftpBulk registers a new resource with the given unique name, arguments, and options.
func NewSourceSftpBulk(ctx *pulumi.Context,
	name string, args *SourceSftpBulkArgs, opts ...pulumi.ResourceOption) (*SourceSftpBulk, error) {
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
	var resource SourceSftpBulk
	err := ctx.RegisterResource("airbyte:index/sourceSftpBulk:SourceSftpBulk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSftpBulk gets an existing SourceSftpBulk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSftpBulk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSftpBulkState, opts ...pulumi.ResourceOption) (*SourceSftpBulk, error) {
	var resource SourceSftpBulk
	err := ctx.ReadResource("airbyte:index/sourceSftpBulk:SourceSftpBulk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSftpBulk resources.
type sourceSftpBulkState struct {
	Configuration *SourceSftpBulkConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceSftpBulkState struct {
	Configuration SourceSftpBulkConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceSftpBulkState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSftpBulkState)(nil)).Elem()
}

type sourceSftpBulkArgs struct {
	Configuration SourceSftpBulkConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSftpBulk resource.
type SourceSftpBulkArgs struct {
	Configuration SourceSftpBulkConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSftpBulkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSftpBulkArgs)(nil)).Elem()
}

type SourceSftpBulkInput interface {
	pulumi.Input

	ToSourceSftpBulkOutput() SourceSftpBulkOutput
	ToSourceSftpBulkOutputWithContext(ctx context.Context) SourceSftpBulkOutput
}

func (*SourceSftpBulk) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSftpBulk)(nil)).Elem()
}

func (i *SourceSftpBulk) ToSourceSftpBulkOutput() SourceSftpBulkOutput {
	return i.ToSourceSftpBulkOutputWithContext(context.Background())
}

func (i *SourceSftpBulk) ToSourceSftpBulkOutputWithContext(ctx context.Context) SourceSftpBulkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSftpBulkOutput)
}

// SourceSftpBulkArrayInput is an input type that accepts SourceSftpBulkArray and SourceSftpBulkArrayOutput values.
// You can construct a concrete instance of `SourceSftpBulkArrayInput` via:
//
//	SourceSftpBulkArray{ SourceSftpBulkArgs{...} }
type SourceSftpBulkArrayInput interface {
	pulumi.Input

	ToSourceSftpBulkArrayOutput() SourceSftpBulkArrayOutput
	ToSourceSftpBulkArrayOutputWithContext(context.Context) SourceSftpBulkArrayOutput
}

type SourceSftpBulkArray []SourceSftpBulkInput

func (SourceSftpBulkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSftpBulk)(nil)).Elem()
}

func (i SourceSftpBulkArray) ToSourceSftpBulkArrayOutput() SourceSftpBulkArrayOutput {
	return i.ToSourceSftpBulkArrayOutputWithContext(context.Background())
}

func (i SourceSftpBulkArray) ToSourceSftpBulkArrayOutputWithContext(ctx context.Context) SourceSftpBulkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSftpBulkArrayOutput)
}

// SourceSftpBulkMapInput is an input type that accepts SourceSftpBulkMap and SourceSftpBulkMapOutput values.
// You can construct a concrete instance of `SourceSftpBulkMapInput` via:
//
//	SourceSftpBulkMap{ "key": SourceSftpBulkArgs{...} }
type SourceSftpBulkMapInput interface {
	pulumi.Input

	ToSourceSftpBulkMapOutput() SourceSftpBulkMapOutput
	ToSourceSftpBulkMapOutputWithContext(context.Context) SourceSftpBulkMapOutput
}

type SourceSftpBulkMap map[string]SourceSftpBulkInput

func (SourceSftpBulkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSftpBulk)(nil)).Elem()
}

func (i SourceSftpBulkMap) ToSourceSftpBulkMapOutput() SourceSftpBulkMapOutput {
	return i.ToSourceSftpBulkMapOutputWithContext(context.Background())
}

func (i SourceSftpBulkMap) ToSourceSftpBulkMapOutputWithContext(ctx context.Context) SourceSftpBulkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSftpBulkMapOutput)
}

type SourceSftpBulkOutput struct{ *pulumi.OutputState }

func (SourceSftpBulkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSftpBulk)(nil)).Elem()
}

func (o SourceSftpBulkOutput) ToSourceSftpBulkOutput() SourceSftpBulkOutput {
	return o
}

func (o SourceSftpBulkOutput) ToSourceSftpBulkOutputWithContext(ctx context.Context) SourceSftpBulkOutput {
	return o
}

func (o SourceSftpBulkOutput) Configuration() SourceSftpBulkConfigurationOutput {
	return o.ApplyT(func(v *SourceSftpBulk) SourceSftpBulkConfigurationOutput { return v.Configuration }).(SourceSftpBulkConfigurationOutput)
}

func (o SourceSftpBulkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSftpBulk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceSftpBulkOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSftpBulk) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSftpBulkOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSftpBulk) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSftpBulkOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSftpBulk) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSftpBulkOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSftpBulk) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceSftpBulkArrayOutput struct{ *pulumi.OutputState }

func (SourceSftpBulkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSftpBulk)(nil)).Elem()
}

func (o SourceSftpBulkArrayOutput) ToSourceSftpBulkArrayOutput() SourceSftpBulkArrayOutput {
	return o
}

func (o SourceSftpBulkArrayOutput) ToSourceSftpBulkArrayOutputWithContext(ctx context.Context) SourceSftpBulkArrayOutput {
	return o
}

func (o SourceSftpBulkArrayOutput) Index(i pulumi.IntInput) SourceSftpBulkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceSftpBulk {
		return vs[0].([]*SourceSftpBulk)[vs[1].(int)]
	}).(SourceSftpBulkOutput)
}

type SourceSftpBulkMapOutput struct{ *pulumi.OutputState }

func (SourceSftpBulkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSftpBulk)(nil)).Elem()
}

func (o SourceSftpBulkMapOutput) ToSourceSftpBulkMapOutput() SourceSftpBulkMapOutput {
	return o
}

func (o SourceSftpBulkMapOutput) ToSourceSftpBulkMapOutputWithContext(ctx context.Context) SourceSftpBulkMapOutput {
	return o
}

func (o SourceSftpBulkMapOutput) MapIndex(k pulumi.StringInput) SourceSftpBulkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceSftpBulk {
		return vs[0].(map[string]*SourceSftpBulk)[vs[1].(string)]
	}).(SourceSftpBulkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSftpBulkInput)(nil)).Elem(), &SourceSftpBulk{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSftpBulkArrayInput)(nil)).Elem(), SourceSftpBulkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSftpBulkMapInput)(nil)).Elem(), SourceSftpBulkMap{})
	pulumi.RegisterOutputType(SourceSftpBulkOutput{})
	pulumi.RegisterOutputType(SourceSftpBulkArrayOutput{})
	pulumi.RegisterOutputType(SourceSftpBulkMapOutput{})
}
