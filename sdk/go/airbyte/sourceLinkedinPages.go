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

// SourceLinkedinPages Resource
type SourceLinkedinPages struct {
	pulumi.CustomResourceState

	Configuration SourceLinkedinPagesConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceLinkedinPages registers a new resource with the given unique name, arguments, and options.
func NewSourceLinkedinPages(ctx *pulumi.Context,
	name string, args *SourceLinkedinPagesArgs, opts ...pulumi.ResourceOption) (*SourceLinkedinPages, error) {
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
	var resource SourceLinkedinPages
	err := ctx.RegisterResource("airbyte:index/sourceLinkedinPages:SourceLinkedinPages", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceLinkedinPages gets an existing SourceLinkedinPages resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceLinkedinPages(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceLinkedinPagesState, opts ...pulumi.ResourceOption) (*SourceLinkedinPages, error) {
	var resource SourceLinkedinPages
	err := ctx.ReadResource("airbyte:index/sourceLinkedinPages:SourceLinkedinPages", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceLinkedinPages resources.
type sourceLinkedinPagesState struct {
	Configuration *SourceLinkedinPagesConfiguration `pulumi:"configuration"`
	Name          *string                           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceLinkedinPagesState struct {
	Configuration SourceLinkedinPagesConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceLinkedinPagesState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLinkedinPagesState)(nil)).Elem()
}

type sourceLinkedinPagesArgs struct {
	Configuration SourceLinkedinPagesConfiguration `pulumi:"configuration"`
	Name          string                           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceLinkedinPages resource.
type SourceLinkedinPagesArgs struct {
	Configuration SourceLinkedinPagesConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceLinkedinPagesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceLinkedinPagesArgs)(nil)).Elem()
}

type SourceLinkedinPagesInput interface {
	pulumi.Input

	ToSourceLinkedinPagesOutput() SourceLinkedinPagesOutput
	ToSourceLinkedinPagesOutputWithContext(ctx context.Context) SourceLinkedinPagesOutput
}

func (*SourceLinkedinPages) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLinkedinPages)(nil)).Elem()
}

func (i *SourceLinkedinPages) ToSourceLinkedinPagesOutput() SourceLinkedinPagesOutput {
	return i.ToSourceLinkedinPagesOutputWithContext(context.Background())
}

func (i *SourceLinkedinPages) ToSourceLinkedinPagesOutputWithContext(ctx context.Context) SourceLinkedinPagesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLinkedinPagesOutput)
}

// SourceLinkedinPagesArrayInput is an input type that accepts SourceLinkedinPagesArray and SourceLinkedinPagesArrayOutput values.
// You can construct a concrete instance of `SourceLinkedinPagesArrayInput` via:
//
//	SourceLinkedinPagesArray{ SourceLinkedinPagesArgs{...} }
type SourceLinkedinPagesArrayInput interface {
	pulumi.Input

	ToSourceLinkedinPagesArrayOutput() SourceLinkedinPagesArrayOutput
	ToSourceLinkedinPagesArrayOutputWithContext(context.Context) SourceLinkedinPagesArrayOutput
}

type SourceLinkedinPagesArray []SourceLinkedinPagesInput

func (SourceLinkedinPagesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceLinkedinPages)(nil)).Elem()
}

func (i SourceLinkedinPagesArray) ToSourceLinkedinPagesArrayOutput() SourceLinkedinPagesArrayOutput {
	return i.ToSourceLinkedinPagesArrayOutputWithContext(context.Background())
}

func (i SourceLinkedinPagesArray) ToSourceLinkedinPagesArrayOutputWithContext(ctx context.Context) SourceLinkedinPagesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLinkedinPagesArrayOutput)
}

// SourceLinkedinPagesMapInput is an input type that accepts SourceLinkedinPagesMap and SourceLinkedinPagesMapOutput values.
// You can construct a concrete instance of `SourceLinkedinPagesMapInput` via:
//
//	SourceLinkedinPagesMap{ "key": SourceLinkedinPagesArgs{...} }
type SourceLinkedinPagesMapInput interface {
	pulumi.Input

	ToSourceLinkedinPagesMapOutput() SourceLinkedinPagesMapOutput
	ToSourceLinkedinPagesMapOutputWithContext(context.Context) SourceLinkedinPagesMapOutput
}

type SourceLinkedinPagesMap map[string]SourceLinkedinPagesInput

func (SourceLinkedinPagesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceLinkedinPages)(nil)).Elem()
}

func (i SourceLinkedinPagesMap) ToSourceLinkedinPagesMapOutput() SourceLinkedinPagesMapOutput {
	return i.ToSourceLinkedinPagesMapOutputWithContext(context.Background())
}

func (i SourceLinkedinPagesMap) ToSourceLinkedinPagesMapOutputWithContext(ctx context.Context) SourceLinkedinPagesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceLinkedinPagesMapOutput)
}

type SourceLinkedinPagesOutput struct{ *pulumi.OutputState }

func (SourceLinkedinPagesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceLinkedinPages)(nil)).Elem()
}

func (o SourceLinkedinPagesOutput) ToSourceLinkedinPagesOutput() SourceLinkedinPagesOutput {
	return o
}

func (o SourceLinkedinPagesOutput) ToSourceLinkedinPagesOutputWithContext(ctx context.Context) SourceLinkedinPagesOutput {
	return o
}

func (o SourceLinkedinPagesOutput) Configuration() SourceLinkedinPagesConfigurationOutput {
	return o.ApplyT(func(v *SourceLinkedinPages) SourceLinkedinPagesConfigurationOutput { return v.Configuration }).(SourceLinkedinPagesConfigurationOutput)
}

func (o SourceLinkedinPagesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinkedinPages) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceLinkedinPagesOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceLinkedinPages) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceLinkedinPagesOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinkedinPages) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceLinkedinPagesOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinkedinPages) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceLinkedinPagesOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceLinkedinPages) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceLinkedinPagesArrayOutput struct{ *pulumi.OutputState }

func (SourceLinkedinPagesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceLinkedinPages)(nil)).Elem()
}

func (o SourceLinkedinPagesArrayOutput) ToSourceLinkedinPagesArrayOutput() SourceLinkedinPagesArrayOutput {
	return o
}

func (o SourceLinkedinPagesArrayOutput) ToSourceLinkedinPagesArrayOutputWithContext(ctx context.Context) SourceLinkedinPagesArrayOutput {
	return o
}

func (o SourceLinkedinPagesArrayOutput) Index(i pulumi.IntInput) SourceLinkedinPagesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceLinkedinPages {
		return vs[0].([]*SourceLinkedinPages)[vs[1].(int)]
	}).(SourceLinkedinPagesOutput)
}

type SourceLinkedinPagesMapOutput struct{ *pulumi.OutputState }

func (SourceLinkedinPagesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceLinkedinPages)(nil)).Elem()
}

func (o SourceLinkedinPagesMapOutput) ToSourceLinkedinPagesMapOutput() SourceLinkedinPagesMapOutput {
	return o
}

func (o SourceLinkedinPagesMapOutput) ToSourceLinkedinPagesMapOutputWithContext(ctx context.Context) SourceLinkedinPagesMapOutput {
	return o
}

func (o SourceLinkedinPagesMapOutput) MapIndex(k pulumi.StringInput) SourceLinkedinPagesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceLinkedinPages {
		return vs[0].(map[string]*SourceLinkedinPages)[vs[1].(string)]
	}).(SourceLinkedinPagesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLinkedinPagesInput)(nil)).Elem(), &SourceLinkedinPages{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLinkedinPagesArrayInput)(nil)).Elem(), SourceLinkedinPagesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceLinkedinPagesMapInput)(nil)).Elem(), SourceLinkedinPagesMap{})
	pulumi.RegisterOutputType(SourceLinkedinPagesOutput{})
	pulumi.RegisterOutputType(SourceLinkedinPagesArrayOutput{})
	pulumi.RegisterOutputType(SourceLinkedinPagesMapOutput{})
}
