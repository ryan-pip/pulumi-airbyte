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

// SourceOutreach Resource
type SourceOutreach struct {
	pulumi.CustomResourceState

	Configuration SourceOutreachConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceOutreach registers a new resource with the given unique name, arguments, and options.
func NewSourceOutreach(ctx *pulumi.Context,
	name string, args *SourceOutreachArgs, opts ...pulumi.ResourceOption) (*SourceOutreach, error) {
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
	var resource SourceOutreach
	err := ctx.RegisterResource("airbyte:index/sourceOutreach:SourceOutreach", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceOutreach gets an existing SourceOutreach resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceOutreach(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceOutreachState, opts ...pulumi.ResourceOption) (*SourceOutreach, error) {
	var resource SourceOutreach
	err := ctx.ReadResource("airbyte:index/sourceOutreach:SourceOutreach", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceOutreach resources.
type sourceOutreachState struct {
	Configuration *SourceOutreachConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceOutreachState struct {
	Configuration SourceOutreachConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceOutreachState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOutreachState)(nil)).Elem()
}

type sourceOutreachArgs struct {
	Configuration SourceOutreachConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceOutreach resource.
type SourceOutreachArgs struct {
	Configuration SourceOutreachConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceOutreachArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOutreachArgs)(nil)).Elem()
}

type SourceOutreachInput interface {
	pulumi.Input

	ToSourceOutreachOutput() SourceOutreachOutput
	ToSourceOutreachOutputWithContext(ctx context.Context) SourceOutreachOutput
}

func (*SourceOutreach) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOutreach)(nil)).Elem()
}

func (i *SourceOutreach) ToSourceOutreachOutput() SourceOutreachOutput {
	return i.ToSourceOutreachOutputWithContext(context.Background())
}

func (i *SourceOutreach) ToSourceOutreachOutputWithContext(ctx context.Context) SourceOutreachOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutreachOutput)
}

// SourceOutreachArrayInput is an input type that accepts SourceOutreachArray and SourceOutreachArrayOutput values.
// You can construct a concrete instance of `SourceOutreachArrayInput` via:
//
//	SourceOutreachArray{ SourceOutreachArgs{...} }
type SourceOutreachArrayInput interface {
	pulumi.Input

	ToSourceOutreachArrayOutput() SourceOutreachArrayOutput
	ToSourceOutreachArrayOutputWithContext(context.Context) SourceOutreachArrayOutput
}

type SourceOutreachArray []SourceOutreachInput

func (SourceOutreachArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceOutreach)(nil)).Elem()
}

func (i SourceOutreachArray) ToSourceOutreachArrayOutput() SourceOutreachArrayOutput {
	return i.ToSourceOutreachArrayOutputWithContext(context.Background())
}

func (i SourceOutreachArray) ToSourceOutreachArrayOutputWithContext(ctx context.Context) SourceOutreachArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutreachArrayOutput)
}

// SourceOutreachMapInput is an input type that accepts SourceOutreachMap and SourceOutreachMapOutput values.
// You can construct a concrete instance of `SourceOutreachMapInput` via:
//
//	SourceOutreachMap{ "key": SourceOutreachArgs{...} }
type SourceOutreachMapInput interface {
	pulumi.Input

	ToSourceOutreachMapOutput() SourceOutreachMapOutput
	ToSourceOutreachMapOutputWithContext(context.Context) SourceOutreachMapOutput
}

type SourceOutreachMap map[string]SourceOutreachInput

func (SourceOutreachMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceOutreach)(nil)).Elem()
}

func (i SourceOutreachMap) ToSourceOutreachMapOutput() SourceOutreachMapOutput {
	return i.ToSourceOutreachMapOutputWithContext(context.Background())
}

func (i SourceOutreachMap) ToSourceOutreachMapOutputWithContext(ctx context.Context) SourceOutreachMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutreachMapOutput)
}

type SourceOutreachOutput struct{ *pulumi.OutputState }

func (SourceOutreachOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOutreach)(nil)).Elem()
}

func (o SourceOutreachOutput) ToSourceOutreachOutput() SourceOutreachOutput {
	return o
}

func (o SourceOutreachOutput) ToSourceOutreachOutputWithContext(ctx context.Context) SourceOutreachOutput {
	return o
}

func (o SourceOutreachOutput) Configuration() SourceOutreachConfigurationOutput {
	return o.ApplyT(func(v *SourceOutreach) SourceOutreachConfigurationOutput { return v.Configuration }).(SourceOutreachConfigurationOutput)
}

func (o SourceOutreachOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutreach) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceOutreachOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOutreach) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceOutreachOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutreach) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceOutreachOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutreach) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceOutreachOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutreach) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceOutreachArrayOutput struct{ *pulumi.OutputState }

func (SourceOutreachArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceOutreach)(nil)).Elem()
}

func (o SourceOutreachArrayOutput) ToSourceOutreachArrayOutput() SourceOutreachArrayOutput {
	return o
}

func (o SourceOutreachArrayOutput) ToSourceOutreachArrayOutputWithContext(ctx context.Context) SourceOutreachArrayOutput {
	return o
}

func (o SourceOutreachArrayOutput) Index(i pulumi.IntInput) SourceOutreachOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceOutreach {
		return vs[0].([]*SourceOutreach)[vs[1].(int)]
	}).(SourceOutreachOutput)
}

type SourceOutreachMapOutput struct{ *pulumi.OutputState }

func (SourceOutreachMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceOutreach)(nil)).Elem()
}

func (o SourceOutreachMapOutput) ToSourceOutreachMapOutput() SourceOutreachMapOutput {
	return o
}

func (o SourceOutreachMapOutput) ToSourceOutreachMapOutputWithContext(ctx context.Context) SourceOutreachMapOutput {
	return o
}

func (o SourceOutreachMapOutput) MapIndex(k pulumi.StringInput) SourceOutreachOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceOutreach {
		return vs[0].(map[string]*SourceOutreach)[vs[1].(string)]
	}).(SourceOutreachOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOutreachInput)(nil)).Elem(), &SourceOutreach{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOutreachArrayInput)(nil)).Elem(), SourceOutreachArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOutreachMapInput)(nil)).Elem(), SourceOutreachMap{})
	pulumi.RegisterOutputType(SourceOutreachOutput{})
	pulumi.RegisterOutputType(SourceOutreachArrayOutput{})
	pulumi.RegisterOutputType(SourceOutreachMapOutput{})
}
