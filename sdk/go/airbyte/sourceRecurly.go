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

// SourceRecurly Resource
type SourceRecurly struct {
	pulumi.CustomResourceState

	Configuration SourceRecurlyConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceRecurly registers a new resource with the given unique name, arguments, and options.
func NewSourceRecurly(ctx *pulumi.Context,
	name string, args *SourceRecurlyArgs, opts ...pulumi.ResourceOption) (*SourceRecurly, error) {
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
	var resource SourceRecurly
	err := ctx.RegisterResource("airbyte:index/sourceRecurly:SourceRecurly", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceRecurly gets an existing SourceRecurly resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceRecurly(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceRecurlyState, opts ...pulumi.ResourceOption) (*SourceRecurly, error) {
	var resource SourceRecurly
	err := ctx.ReadResource("airbyte:index/sourceRecurly:SourceRecurly", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceRecurly resources.
type sourceRecurlyState struct {
	Configuration *SourceRecurlyConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceRecurlyState struct {
	Configuration SourceRecurlyConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceRecurlyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRecurlyState)(nil)).Elem()
}

type sourceRecurlyArgs struct {
	Configuration SourceRecurlyConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceRecurly resource.
type SourceRecurlyArgs struct {
	Configuration SourceRecurlyConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceRecurlyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRecurlyArgs)(nil)).Elem()
}

type SourceRecurlyInput interface {
	pulumi.Input

	ToSourceRecurlyOutput() SourceRecurlyOutput
	ToSourceRecurlyOutputWithContext(ctx context.Context) SourceRecurlyOutput
}

func (*SourceRecurly) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRecurly)(nil)).Elem()
}

func (i *SourceRecurly) ToSourceRecurlyOutput() SourceRecurlyOutput {
	return i.ToSourceRecurlyOutputWithContext(context.Background())
}

func (i *SourceRecurly) ToSourceRecurlyOutputWithContext(ctx context.Context) SourceRecurlyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRecurlyOutput)
}

// SourceRecurlyArrayInput is an input type that accepts SourceRecurlyArray and SourceRecurlyArrayOutput values.
// You can construct a concrete instance of `SourceRecurlyArrayInput` via:
//
//	SourceRecurlyArray{ SourceRecurlyArgs{...} }
type SourceRecurlyArrayInput interface {
	pulumi.Input

	ToSourceRecurlyArrayOutput() SourceRecurlyArrayOutput
	ToSourceRecurlyArrayOutputWithContext(context.Context) SourceRecurlyArrayOutput
}

type SourceRecurlyArray []SourceRecurlyInput

func (SourceRecurlyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceRecurly)(nil)).Elem()
}

func (i SourceRecurlyArray) ToSourceRecurlyArrayOutput() SourceRecurlyArrayOutput {
	return i.ToSourceRecurlyArrayOutputWithContext(context.Background())
}

func (i SourceRecurlyArray) ToSourceRecurlyArrayOutputWithContext(ctx context.Context) SourceRecurlyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRecurlyArrayOutput)
}

// SourceRecurlyMapInput is an input type that accepts SourceRecurlyMap and SourceRecurlyMapOutput values.
// You can construct a concrete instance of `SourceRecurlyMapInput` via:
//
//	SourceRecurlyMap{ "key": SourceRecurlyArgs{...} }
type SourceRecurlyMapInput interface {
	pulumi.Input

	ToSourceRecurlyMapOutput() SourceRecurlyMapOutput
	ToSourceRecurlyMapOutputWithContext(context.Context) SourceRecurlyMapOutput
}

type SourceRecurlyMap map[string]SourceRecurlyInput

func (SourceRecurlyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceRecurly)(nil)).Elem()
}

func (i SourceRecurlyMap) ToSourceRecurlyMapOutput() SourceRecurlyMapOutput {
	return i.ToSourceRecurlyMapOutputWithContext(context.Background())
}

func (i SourceRecurlyMap) ToSourceRecurlyMapOutputWithContext(ctx context.Context) SourceRecurlyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRecurlyMapOutput)
}

type SourceRecurlyOutput struct{ *pulumi.OutputState }

func (SourceRecurlyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRecurly)(nil)).Elem()
}

func (o SourceRecurlyOutput) ToSourceRecurlyOutput() SourceRecurlyOutput {
	return o
}

func (o SourceRecurlyOutput) ToSourceRecurlyOutputWithContext(ctx context.Context) SourceRecurlyOutput {
	return o
}

func (o SourceRecurlyOutput) Configuration() SourceRecurlyConfigurationOutput {
	return o.ApplyT(func(v *SourceRecurly) SourceRecurlyConfigurationOutput { return v.Configuration }).(SourceRecurlyConfigurationOutput)
}

func (o SourceRecurlyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRecurly) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceRecurlyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceRecurly) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceRecurlyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRecurly) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceRecurlyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRecurly) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceRecurlyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceRecurly) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceRecurlyArrayOutput struct{ *pulumi.OutputState }

func (SourceRecurlyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceRecurly)(nil)).Elem()
}

func (o SourceRecurlyArrayOutput) ToSourceRecurlyArrayOutput() SourceRecurlyArrayOutput {
	return o
}

func (o SourceRecurlyArrayOutput) ToSourceRecurlyArrayOutputWithContext(ctx context.Context) SourceRecurlyArrayOutput {
	return o
}

func (o SourceRecurlyArrayOutput) Index(i pulumi.IntInput) SourceRecurlyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceRecurly {
		return vs[0].([]*SourceRecurly)[vs[1].(int)]
	}).(SourceRecurlyOutput)
}

type SourceRecurlyMapOutput struct{ *pulumi.OutputState }

func (SourceRecurlyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceRecurly)(nil)).Elem()
}

func (o SourceRecurlyMapOutput) ToSourceRecurlyMapOutput() SourceRecurlyMapOutput {
	return o
}

func (o SourceRecurlyMapOutput) ToSourceRecurlyMapOutputWithContext(ctx context.Context) SourceRecurlyMapOutput {
	return o
}

func (o SourceRecurlyMapOutput) MapIndex(k pulumi.StringInput) SourceRecurlyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceRecurly {
		return vs[0].(map[string]*SourceRecurly)[vs[1].(string)]
	}).(SourceRecurlyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRecurlyInput)(nil)).Elem(), &SourceRecurly{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRecurlyArrayInput)(nil)).Elem(), SourceRecurlyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRecurlyMapInput)(nil)).Elem(), SourceRecurlyMap{})
	pulumi.RegisterOutputType(SourceRecurlyOutput{})
	pulumi.RegisterOutputType(SourceRecurlyArrayOutput{})
	pulumi.RegisterOutputType(SourceRecurlyMapOutput{})
}
