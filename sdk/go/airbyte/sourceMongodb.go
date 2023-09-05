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

// SourceMongodb Resource
type SourceMongodb struct {
	pulumi.CustomResourceState

	Configuration SourceMongodbConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceMongodb registers a new resource with the given unique name, arguments, and options.
func NewSourceMongodb(ctx *pulumi.Context,
	name string, args *SourceMongodbArgs, opts ...pulumi.ResourceOption) (*SourceMongodb, error) {
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
	var resource SourceMongodb
	err := ctx.RegisterResource("airbyte:index/sourceMongodb:SourceMongodb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceMongodb gets an existing SourceMongodb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceMongodb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceMongodbState, opts ...pulumi.ResourceOption) (*SourceMongodb, error) {
	var resource SourceMongodb
	err := ctx.ReadResource("airbyte:index/sourceMongodb:SourceMongodb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceMongodb resources.
type sourceMongodbState struct {
	Configuration *SourceMongodbConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceMongodbState struct {
	Configuration SourceMongodbConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceMongodbState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMongodbState)(nil)).Elem()
}

type sourceMongodbArgs struct {
	Configuration SourceMongodbConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceMongodb resource.
type SourceMongodbArgs struct {
	Configuration SourceMongodbConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceMongodbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMongodbArgs)(nil)).Elem()
}

type SourceMongodbInput interface {
	pulumi.Input

	ToSourceMongodbOutput() SourceMongodbOutput
	ToSourceMongodbOutputWithContext(ctx context.Context) SourceMongodbOutput
}

func (*SourceMongodb) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMongodb)(nil)).Elem()
}

func (i *SourceMongodb) ToSourceMongodbOutput() SourceMongodbOutput {
	return i.ToSourceMongodbOutputWithContext(context.Background())
}

func (i *SourceMongodb) ToSourceMongodbOutputWithContext(ctx context.Context) SourceMongodbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMongodbOutput)
}

// SourceMongodbArrayInput is an input type that accepts SourceMongodbArray and SourceMongodbArrayOutput values.
// You can construct a concrete instance of `SourceMongodbArrayInput` via:
//
//	SourceMongodbArray{ SourceMongodbArgs{...} }
type SourceMongodbArrayInput interface {
	pulumi.Input

	ToSourceMongodbArrayOutput() SourceMongodbArrayOutput
	ToSourceMongodbArrayOutputWithContext(context.Context) SourceMongodbArrayOutput
}

type SourceMongodbArray []SourceMongodbInput

func (SourceMongodbArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceMongodb)(nil)).Elem()
}

func (i SourceMongodbArray) ToSourceMongodbArrayOutput() SourceMongodbArrayOutput {
	return i.ToSourceMongodbArrayOutputWithContext(context.Background())
}

func (i SourceMongodbArray) ToSourceMongodbArrayOutputWithContext(ctx context.Context) SourceMongodbArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMongodbArrayOutput)
}

// SourceMongodbMapInput is an input type that accepts SourceMongodbMap and SourceMongodbMapOutput values.
// You can construct a concrete instance of `SourceMongodbMapInput` via:
//
//	SourceMongodbMap{ "key": SourceMongodbArgs{...} }
type SourceMongodbMapInput interface {
	pulumi.Input

	ToSourceMongodbMapOutput() SourceMongodbMapOutput
	ToSourceMongodbMapOutputWithContext(context.Context) SourceMongodbMapOutput
}

type SourceMongodbMap map[string]SourceMongodbInput

func (SourceMongodbMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceMongodb)(nil)).Elem()
}

func (i SourceMongodbMap) ToSourceMongodbMapOutput() SourceMongodbMapOutput {
	return i.ToSourceMongodbMapOutputWithContext(context.Background())
}

func (i SourceMongodbMap) ToSourceMongodbMapOutputWithContext(ctx context.Context) SourceMongodbMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMongodbMapOutput)
}

type SourceMongodbOutput struct{ *pulumi.OutputState }

func (SourceMongodbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMongodb)(nil)).Elem()
}

func (o SourceMongodbOutput) ToSourceMongodbOutput() SourceMongodbOutput {
	return o
}

func (o SourceMongodbOutput) ToSourceMongodbOutputWithContext(ctx context.Context) SourceMongodbOutput {
	return o
}

func (o SourceMongodbOutput) Configuration() SourceMongodbConfigurationOutput {
	return o.ApplyT(func(v *SourceMongodb) SourceMongodbConfigurationOutput { return v.Configuration }).(SourceMongodbConfigurationOutput)
}

func (o SourceMongodbOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMongodb) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceMongodbOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceMongodb) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceMongodbOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMongodb) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceMongodbOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMongodb) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceMongodbOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMongodb) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceMongodbArrayOutput struct{ *pulumi.OutputState }

func (SourceMongodbArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceMongodb)(nil)).Elem()
}

func (o SourceMongodbArrayOutput) ToSourceMongodbArrayOutput() SourceMongodbArrayOutput {
	return o
}

func (o SourceMongodbArrayOutput) ToSourceMongodbArrayOutputWithContext(ctx context.Context) SourceMongodbArrayOutput {
	return o
}

func (o SourceMongodbArrayOutput) Index(i pulumi.IntInput) SourceMongodbOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceMongodb {
		return vs[0].([]*SourceMongodb)[vs[1].(int)]
	}).(SourceMongodbOutput)
}

type SourceMongodbMapOutput struct{ *pulumi.OutputState }

func (SourceMongodbMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceMongodb)(nil)).Elem()
}

func (o SourceMongodbMapOutput) ToSourceMongodbMapOutput() SourceMongodbMapOutput {
	return o
}

func (o SourceMongodbMapOutput) ToSourceMongodbMapOutputWithContext(ctx context.Context) SourceMongodbMapOutput {
	return o
}

func (o SourceMongodbMapOutput) MapIndex(k pulumi.StringInput) SourceMongodbOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceMongodb {
		return vs[0].(map[string]*SourceMongodb)[vs[1].(string)]
	}).(SourceMongodbOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMongodbInput)(nil)).Elem(), &SourceMongodb{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMongodbArrayInput)(nil)).Elem(), SourceMongodbArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMongodbMapInput)(nil)).Elem(), SourceMongodbMap{})
	pulumi.RegisterOutputType(SourceMongodbOutput{})
	pulumi.RegisterOutputType(SourceMongodbArrayOutput{})
	pulumi.RegisterOutputType(SourceMongodbMapOutput{})
}
