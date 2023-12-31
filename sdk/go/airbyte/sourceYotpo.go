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

// SourceYotpo Resource
type SourceYotpo struct {
	pulumi.CustomResourceState

	Configuration SourceYotpoConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceYotpo registers a new resource with the given unique name, arguments, and options.
func NewSourceYotpo(ctx *pulumi.Context,
	name string, args *SourceYotpoArgs, opts ...pulumi.ResourceOption) (*SourceYotpo, error) {
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
	var resource SourceYotpo
	err := ctx.RegisterResource("airbyte:index/sourceYotpo:SourceYotpo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceYotpo gets an existing SourceYotpo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceYotpo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceYotpoState, opts ...pulumi.ResourceOption) (*SourceYotpo, error) {
	var resource SourceYotpo
	err := ctx.ReadResource("airbyte:index/sourceYotpo:SourceYotpo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceYotpo resources.
type sourceYotpoState struct {
	Configuration *SourceYotpoConfiguration `pulumi:"configuration"`
	Name          *string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceYotpoState struct {
	Configuration SourceYotpoConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceYotpoState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceYotpoState)(nil)).Elem()
}

type sourceYotpoArgs struct {
	Configuration SourceYotpoConfiguration `pulumi:"configuration"`
	Name          string                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceYotpo resource.
type SourceYotpoArgs struct {
	Configuration SourceYotpoConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceYotpoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceYotpoArgs)(nil)).Elem()
}

type SourceYotpoInput interface {
	pulumi.Input

	ToSourceYotpoOutput() SourceYotpoOutput
	ToSourceYotpoOutputWithContext(ctx context.Context) SourceYotpoOutput
}

func (*SourceYotpo) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceYotpo)(nil)).Elem()
}

func (i *SourceYotpo) ToSourceYotpoOutput() SourceYotpoOutput {
	return i.ToSourceYotpoOutputWithContext(context.Background())
}

func (i *SourceYotpo) ToSourceYotpoOutputWithContext(ctx context.Context) SourceYotpoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceYotpoOutput)
}

// SourceYotpoArrayInput is an input type that accepts SourceYotpoArray and SourceYotpoArrayOutput values.
// You can construct a concrete instance of `SourceYotpoArrayInput` via:
//
//	SourceYotpoArray{ SourceYotpoArgs{...} }
type SourceYotpoArrayInput interface {
	pulumi.Input

	ToSourceYotpoArrayOutput() SourceYotpoArrayOutput
	ToSourceYotpoArrayOutputWithContext(context.Context) SourceYotpoArrayOutput
}

type SourceYotpoArray []SourceYotpoInput

func (SourceYotpoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceYotpo)(nil)).Elem()
}

func (i SourceYotpoArray) ToSourceYotpoArrayOutput() SourceYotpoArrayOutput {
	return i.ToSourceYotpoArrayOutputWithContext(context.Background())
}

func (i SourceYotpoArray) ToSourceYotpoArrayOutputWithContext(ctx context.Context) SourceYotpoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceYotpoArrayOutput)
}

// SourceYotpoMapInput is an input type that accepts SourceYotpoMap and SourceYotpoMapOutput values.
// You can construct a concrete instance of `SourceYotpoMapInput` via:
//
//	SourceYotpoMap{ "key": SourceYotpoArgs{...} }
type SourceYotpoMapInput interface {
	pulumi.Input

	ToSourceYotpoMapOutput() SourceYotpoMapOutput
	ToSourceYotpoMapOutputWithContext(context.Context) SourceYotpoMapOutput
}

type SourceYotpoMap map[string]SourceYotpoInput

func (SourceYotpoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceYotpo)(nil)).Elem()
}

func (i SourceYotpoMap) ToSourceYotpoMapOutput() SourceYotpoMapOutput {
	return i.ToSourceYotpoMapOutputWithContext(context.Background())
}

func (i SourceYotpoMap) ToSourceYotpoMapOutputWithContext(ctx context.Context) SourceYotpoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceYotpoMapOutput)
}

type SourceYotpoOutput struct{ *pulumi.OutputState }

func (SourceYotpoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceYotpo)(nil)).Elem()
}

func (o SourceYotpoOutput) ToSourceYotpoOutput() SourceYotpoOutput {
	return o
}

func (o SourceYotpoOutput) ToSourceYotpoOutputWithContext(ctx context.Context) SourceYotpoOutput {
	return o
}

func (o SourceYotpoOutput) Configuration() SourceYotpoConfigurationOutput {
	return o.ApplyT(func(v *SourceYotpo) SourceYotpoConfigurationOutput { return v.Configuration }).(SourceYotpoConfigurationOutput)
}

func (o SourceYotpoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYotpo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceYotpoOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceYotpo) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceYotpoOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYotpo) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceYotpoOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYotpo) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceYotpoOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYotpo) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceYotpoArrayOutput struct{ *pulumi.OutputState }

func (SourceYotpoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceYotpo)(nil)).Elem()
}

func (o SourceYotpoArrayOutput) ToSourceYotpoArrayOutput() SourceYotpoArrayOutput {
	return o
}

func (o SourceYotpoArrayOutput) ToSourceYotpoArrayOutputWithContext(ctx context.Context) SourceYotpoArrayOutput {
	return o
}

func (o SourceYotpoArrayOutput) Index(i pulumi.IntInput) SourceYotpoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceYotpo {
		return vs[0].([]*SourceYotpo)[vs[1].(int)]
	}).(SourceYotpoOutput)
}

type SourceYotpoMapOutput struct{ *pulumi.OutputState }

func (SourceYotpoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceYotpo)(nil)).Elem()
}

func (o SourceYotpoMapOutput) ToSourceYotpoMapOutput() SourceYotpoMapOutput {
	return o
}

func (o SourceYotpoMapOutput) ToSourceYotpoMapOutputWithContext(ctx context.Context) SourceYotpoMapOutput {
	return o
}

func (o SourceYotpoMapOutput) MapIndex(k pulumi.StringInput) SourceYotpoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceYotpo {
		return vs[0].(map[string]*SourceYotpo)[vs[1].(string)]
	}).(SourceYotpoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceYotpoInput)(nil)).Elem(), &SourceYotpo{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceYotpoArrayInput)(nil)).Elem(), SourceYotpoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceYotpoMapInput)(nil)).Elem(), SourceYotpoMap{})
	pulumi.RegisterOutputType(SourceYotpoOutput{})
	pulumi.RegisterOutputType(SourceYotpoArrayOutput{})
	pulumi.RegisterOutputType(SourceYotpoMapOutput{})
}
