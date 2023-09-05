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

// SourceTwitter Resource
type SourceTwitter struct {
	pulumi.CustomResourceState

	Configuration SourceTwitterConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceTwitter registers a new resource with the given unique name, arguments, and options.
func NewSourceTwitter(ctx *pulumi.Context,
	name string, args *SourceTwitterArgs, opts ...pulumi.ResourceOption) (*SourceTwitter, error) {
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
	var resource SourceTwitter
	err := ctx.RegisterResource("airbyte:index/sourceTwitter:SourceTwitter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceTwitter gets an existing SourceTwitter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceTwitter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceTwitterState, opts ...pulumi.ResourceOption) (*SourceTwitter, error) {
	var resource SourceTwitter
	err := ctx.ReadResource("airbyte:index/sourceTwitter:SourceTwitter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceTwitter resources.
type sourceTwitterState struct {
	Configuration *SourceTwitterConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceTwitterState struct {
	Configuration SourceTwitterConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceTwitterState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceTwitterState)(nil)).Elem()
}

type sourceTwitterArgs struct {
	Configuration SourceTwitterConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceTwitter resource.
type SourceTwitterArgs struct {
	Configuration SourceTwitterConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceTwitterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceTwitterArgs)(nil)).Elem()
}

type SourceTwitterInput interface {
	pulumi.Input

	ToSourceTwitterOutput() SourceTwitterOutput
	ToSourceTwitterOutputWithContext(ctx context.Context) SourceTwitterOutput
}

func (*SourceTwitter) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTwitter)(nil)).Elem()
}

func (i *SourceTwitter) ToSourceTwitterOutput() SourceTwitterOutput {
	return i.ToSourceTwitterOutputWithContext(context.Background())
}

func (i *SourceTwitter) ToSourceTwitterOutputWithContext(ctx context.Context) SourceTwitterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTwitterOutput)
}

// SourceTwitterArrayInput is an input type that accepts SourceTwitterArray and SourceTwitterArrayOutput values.
// You can construct a concrete instance of `SourceTwitterArrayInput` via:
//
//	SourceTwitterArray{ SourceTwitterArgs{...} }
type SourceTwitterArrayInput interface {
	pulumi.Input

	ToSourceTwitterArrayOutput() SourceTwitterArrayOutput
	ToSourceTwitterArrayOutputWithContext(context.Context) SourceTwitterArrayOutput
}

type SourceTwitterArray []SourceTwitterInput

func (SourceTwitterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceTwitter)(nil)).Elem()
}

func (i SourceTwitterArray) ToSourceTwitterArrayOutput() SourceTwitterArrayOutput {
	return i.ToSourceTwitterArrayOutputWithContext(context.Background())
}

func (i SourceTwitterArray) ToSourceTwitterArrayOutputWithContext(ctx context.Context) SourceTwitterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTwitterArrayOutput)
}

// SourceTwitterMapInput is an input type that accepts SourceTwitterMap and SourceTwitterMapOutput values.
// You can construct a concrete instance of `SourceTwitterMapInput` via:
//
//	SourceTwitterMap{ "key": SourceTwitterArgs{...} }
type SourceTwitterMapInput interface {
	pulumi.Input

	ToSourceTwitterMapOutput() SourceTwitterMapOutput
	ToSourceTwitterMapOutputWithContext(context.Context) SourceTwitterMapOutput
}

type SourceTwitterMap map[string]SourceTwitterInput

func (SourceTwitterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceTwitter)(nil)).Elem()
}

func (i SourceTwitterMap) ToSourceTwitterMapOutput() SourceTwitterMapOutput {
	return i.ToSourceTwitterMapOutputWithContext(context.Background())
}

func (i SourceTwitterMap) ToSourceTwitterMapOutputWithContext(ctx context.Context) SourceTwitterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTwitterMapOutput)
}

type SourceTwitterOutput struct{ *pulumi.OutputState }

func (SourceTwitterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTwitter)(nil)).Elem()
}

func (o SourceTwitterOutput) ToSourceTwitterOutput() SourceTwitterOutput {
	return o
}

func (o SourceTwitterOutput) ToSourceTwitterOutputWithContext(ctx context.Context) SourceTwitterOutput {
	return o
}

func (o SourceTwitterOutput) Configuration() SourceTwitterConfigurationOutput {
	return o.ApplyT(func(v *SourceTwitter) SourceTwitterConfigurationOutput { return v.Configuration }).(SourceTwitterConfigurationOutput)
}

func (o SourceTwitterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTwitter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceTwitterOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTwitter) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceTwitterOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTwitter) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceTwitterOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTwitter) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceTwitterOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTwitter) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceTwitterArrayOutput struct{ *pulumi.OutputState }

func (SourceTwitterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceTwitter)(nil)).Elem()
}

func (o SourceTwitterArrayOutput) ToSourceTwitterArrayOutput() SourceTwitterArrayOutput {
	return o
}

func (o SourceTwitterArrayOutput) ToSourceTwitterArrayOutputWithContext(ctx context.Context) SourceTwitterArrayOutput {
	return o
}

func (o SourceTwitterArrayOutput) Index(i pulumi.IntInput) SourceTwitterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceTwitter {
		return vs[0].([]*SourceTwitter)[vs[1].(int)]
	}).(SourceTwitterOutput)
}

type SourceTwitterMapOutput struct{ *pulumi.OutputState }

func (SourceTwitterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceTwitter)(nil)).Elem()
}

func (o SourceTwitterMapOutput) ToSourceTwitterMapOutput() SourceTwitterMapOutput {
	return o
}

func (o SourceTwitterMapOutput) ToSourceTwitterMapOutputWithContext(ctx context.Context) SourceTwitterMapOutput {
	return o
}

func (o SourceTwitterMapOutput) MapIndex(k pulumi.StringInput) SourceTwitterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceTwitter {
		return vs[0].(map[string]*SourceTwitter)[vs[1].(string)]
	}).(SourceTwitterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTwitterInput)(nil)).Elem(), &SourceTwitter{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTwitterArrayInput)(nil)).Elem(), SourceTwitterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTwitterMapInput)(nil)).Elem(), SourceTwitterMap{})
	pulumi.RegisterOutputType(SourceTwitterOutput{})
	pulumi.RegisterOutputType(SourceTwitterArrayOutput{})
	pulumi.RegisterOutputType(SourceTwitterMapOutput{})
}
