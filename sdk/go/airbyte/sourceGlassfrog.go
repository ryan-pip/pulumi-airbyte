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

// SourceGlassfrog Resource
type SourceGlassfrog struct {
	pulumi.CustomResourceState

	Configuration SourceGlassfrogConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceGlassfrog registers a new resource with the given unique name, arguments, and options.
func NewSourceGlassfrog(ctx *pulumi.Context,
	name string, args *SourceGlassfrogArgs, opts ...pulumi.ResourceOption) (*SourceGlassfrog, error) {
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
	var resource SourceGlassfrog
	err := ctx.RegisterResource("airbyte:index/sourceGlassfrog:SourceGlassfrog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceGlassfrog gets an existing SourceGlassfrog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceGlassfrog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceGlassfrogState, opts ...pulumi.ResourceOption) (*SourceGlassfrog, error) {
	var resource SourceGlassfrog
	err := ctx.ReadResource("airbyte:index/sourceGlassfrog:SourceGlassfrog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceGlassfrog resources.
type sourceGlassfrogState struct {
	Configuration *SourceGlassfrogConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceGlassfrogState struct {
	Configuration SourceGlassfrogConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceGlassfrogState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGlassfrogState)(nil)).Elem()
}

type sourceGlassfrogArgs struct {
	Configuration SourceGlassfrogConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceGlassfrog resource.
type SourceGlassfrogArgs struct {
	Configuration SourceGlassfrogConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceGlassfrogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGlassfrogArgs)(nil)).Elem()
}

type SourceGlassfrogInput interface {
	pulumi.Input

	ToSourceGlassfrogOutput() SourceGlassfrogOutput
	ToSourceGlassfrogOutputWithContext(ctx context.Context) SourceGlassfrogOutput
}

func (*SourceGlassfrog) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGlassfrog)(nil)).Elem()
}

func (i *SourceGlassfrog) ToSourceGlassfrogOutput() SourceGlassfrogOutput {
	return i.ToSourceGlassfrogOutputWithContext(context.Background())
}

func (i *SourceGlassfrog) ToSourceGlassfrogOutputWithContext(ctx context.Context) SourceGlassfrogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGlassfrogOutput)
}

// SourceGlassfrogArrayInput is an input type that accepts SourceGlassfrogArray and SourceGlassfrogArrayOutput values.
// You can construct a concrete instance of `SourceGlassfrogArrayInput` via:
//
//	SourceGlassfrogArray{ SourceGlassfrogArgs{...} }
type SourceGlassfrogArrayInput interface {
	pulumi.Input

	ToSourceGlassfrogArrayOutput() SourceGlassfrogArrayOutput
	ToSourceGlassfrogArrayOutputWithContext(context.Context) SourceGlassfrogArrayOutput
}

type SourceGlassfrogArray []SourceGlassfrogInput

func (SourceGlassfrogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGlassfrog)(nil)).Elem()
}

func (i SourceGlassfrogArray) ToSourceGlassfrogArrayOutput() SourceGlassfrogArrayOutput {
	return i.ToSourceGlassfrogArrayOutputWithContext(context.Background())
}

func (i SourceGlassfrogArray) ToSourceGlassfrogArrayOutputWithContext(ctx context.Context) SourceGlassfrogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGlassfrogArrayOutput)
}

// SourceGlassfrogMapInput is an input type that accepts SourceGlassfrogMap and SourceGlassfrogMapOutput values.
// You can construct a concrete instance of `SourceGlassfrogMapInput` via:
//
//	SourceGlassfrogMap{ "key": SourceGlassfrogArgs{...} }
type SourceGlassfrogMapInput interface {
	pulumi.Input

	ToSourceGlassfrogMapOutput() SourceGlassfrogMapOutput
	ToSourceGlassfrogMapOutputWithContext(context.Context) SourceGlassfrogMapOutput
}

type SourceGlassfrogMap map[string]SourceGlassfrogInput

func (SourceGlassfrogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGlassfrog)(nil)).Elem()
}

func (i SourceGlassfrogMap) ToSourceGlassfrogMapOutput() SourceGlassfrogMapOutput {
	return i.ToSourceGlassfrogMapOutputWithContext(context.Background())
}

func (i SourceGlassfrogMap) ToSourceGlassfrogMapOutputWithContext(ctx context.Context) SourceGlassfrogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGlassfrogMapOutput)
}

type SourceGlassfrogOutput struct{ *pulumi.OutputState }

func (SourceGlassfrogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGlassfrog)(nil)).Elem()
}

func (o SourceGlassfrogOutput) ToSourceGlassfrogOutput() SourceGlassfrogOutput {
	return o
}

func (o SourceGlassfrogOutput) ToSourceGlassfrogOutputWithContext(ctx context.Context) SourceGlassfrogOutput {
	return o
}

func (o SourceGlassfrogOutput) Configuration() SourceGlassfrogConfigurationOutput {
	return o.ApplyT(func(v *SourceGlassfrog) SourceGlassfrogConfigurationOutput { return v.Configuration }).(SourceGlassfrogConfigurationOutput)
}

func (o SourceGlassfrogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGlassfrog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceGlassfrogOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGlassfrog) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceGlassfrogOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGlassfrog) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceGlassfrogOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGlassfrog) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceGlassfrogOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGlassfrog) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceGlassfrogArrayOutput struct{ *pulumi.OutputState }

func (SourceGlassfrogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGlassfrog)(nil)).Elem()
}

func (o SourceGlassfrogArrayOutput) ToSourceGlassfrogArrayOutput() SourceGlassfrogArrayOutput {
	return o
}

func (o SourceGlassfrogArrayOutput) ToSourceGlassfrogArrayOutputWithContext(ctx context.Context) SourceGlassfrogArrayOutput {
	return o
}

func (o SourceGlassfrogArrayOutput) Index(i pulumi.IntInput) SourceGlassfrogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceGlassfrog {
		return vs[0].([]*SourceGlassfrog)[vs[1].(int)]
	}).(SourceGlassfrogOutput)
}

type SourceGlassfrogMapOutput struct{ *pulumi.OutputState }

func (SourceGlassfrogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGlassfrog)(nil)).Elem()
}

func (o SourceGlassfrogMapOutput) ToSourceGlassfrogMapOutput() SourceGlassfrogMapOutput {
	return o
}

func (o SourceGlassfrogMapOutput) ToSourceGlassfrogMapOutputWithContext(ctx context.Context) SourceGlassfrogMapOutput {
	return o
}

func (o SourceGlassfrogMapOutput) MapIndex(k pulumi.StringInput) SourceGlassfrogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceGlassfrog {
		return vs[0].(map[string]*SourceGlassfrog)[vs[1].(string)]
	}).(SourceGlassfrogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGlassfrogInput)(nil)).Elem(), &SourceGlassfrog{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGlassfrogArrayInput)(nil)).Elem(), SourceGlassfrogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGlassfrogMapInput)(nil)).Elem(), SourceGlassfrogMap{})
	pulumi.RegisterOutputType(SourceGlassfrogOutput{})
	pulumi.RegisterOutputType(SourceGlassfrogArrayOutput{})
	pulumi.RegisterOutputType(SourceGlassfrogMapOutput{})
}
