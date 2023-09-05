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

// SourcePostgres Resource
type SourcePostgres struct {
	pulumi.CustomResourceState

	Configuration SourcePostgresConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourcePostgres registers a new resource with the given unique name, arguments, and options.
func NewSourcePostgres(ctx *pulumi.Context,
	name string, args *SourcePostgresArgs, opts ...pulumi.ResourceOption) (*SourcePostgres, error) {
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
	var resource SourcePostgres
	err := ctx.RegisterResource("airbyte:index/sourcePostgres:SourcePostgres", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourcePostgres gets an existing SourcePostgres resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourcePostgres(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourcePostgresState, opts ...pulumi.ResourceOption) (*SourcePostgres, error) {
	var resource SourcePostgres
	err := ctx.ReadResource("airbyte:index/sourcePostgres:SourcePostgres", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourcePostgres resources.
type sourcePostgresState struct {
	Configuration *SourcePostgresConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourcePostgresState struct {
	Configuration SourcePostgresConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourcePostgresState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePostgresState)(nil)).Elem()
}

type sourcePostgresArgs struct {
	Configuration SourcePostgresConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourcePostgres resource.
type SourcePostgresArgs struct {
	Configuration SourcePostgresConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourcePostgresArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePostgresArgs)(nil)).Elem()
}

type SourcePostgresInput interface {
	pulumi.Input

	ToSourcePostgresOutput() SourcePostgresOutput
	ToSourcePostgresOutputWithContext(ctx context.Context) SourcePostgresOutput
}

func (*SourcePostgres) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePostgres)(nil)).Elem()
}

func (i *SourcePostgres) ToSourcePostgresOutput() SourcePostgresOutput {
	return i.ToSourcePostgresOutputWithContext(context.Background())
}

func (i *SourcePostgres) ToSourcePostgresOutputWithContext(ctx context.Context) SourcePostgresOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePostgresOutput)
}

// SourcePostgresArrayInput is an input type that accepts SourcePostgresArray and SourcePostgresArrayOutput values.
// You can construct a concrete instance of `SourcePostgresArrayInput` via:
//
//	SourcePostgresArray{ SourcePostgresArgs{...} }
type SourcePostgresArrayInput interface {
	pulumi.Input

	ToSourcePostgresArrayOutput() SourcePostgresArrayOutput
	ToSourcePostgresArrayOutputWithContext(context.Context) SourcePostgresArrayOutput
}

type SourcePostgresArray []SourcePostgresInput

func (SourcePostgresArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourcePostgres)(nil)).Elem()
}

func (i SourcePostgresArray) ToSourcePostgresArrayOutput() SourcePostgresArrayOutput {
	return i.ToSourcePostgresArrayOutputWithContext(context.Background())
}

func (i SourcePostgresArray) ToSourcePostgresArrayOutputWithContext(ctx context.Context) SourcePostgresArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePostgresArrayOutput)
}

// SourcePostgresMapInput is an input type that accepts SourcePostgresMap and SourcePostgresMapOutput values.
// You can construct a concrete instance of `SourcePostgresMapInput` via:
//
//	SourcePostgresMap{ "key": SourcePostgresArgs{...} }
type SourcePostgresMapInput interface {
	pulumi.Input

	ToSourcePostgresMapOutput() SourcePostgresMapOutput
	ToSourcePostgresMapOutputWithContext(context.Context) SourcePostgresMapOutput
}

type SourcePostgresMap map[string]SourcePostgresInput

func (SourcePostgresMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourcePostgres)(nil)).Elem()
}

func (i SourcePostgresMap) ToSourcePostgresMapOutput() SourcePostgresMapOutput {
	return i.ToSourcePostgresMapOutputWithContext(context.Background())
}

func (i SourcePostgresMap) ToSourcePostgresMapOutputWithContext(ctx context.Context) SourcePostgresMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePostgresMapOutput)
}

type SourcePostgresOutput struct{ *pulumi.OutputState }

func (SourcePostgresOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePostgres)(nil)).Elem()
}

func (o SourcePostgresOutput) ToSourcePostgresOutput() SourcePostgresOutput {
	return o
}

func (o SourcePostgresOutput) ToSourcePostgresOutputWithContext(ctx context.Context) SourcePostgresOutput {
	return o
}

func (o SourcePostgresOutput) Configuration() SourcePostgresConfigurationOutput {
	return o.ApplyT(func(v *SourcePostgres) SourcePostgresConfigurationOutput { return v.Configuration }).(SourcePostgresConfigurationOutput)
}

func (o SourcePostgresOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePostgres) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourcePostgresOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePostgres) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourcePostgresOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePostgres) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourcePostgresOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePostgres) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourcePostgresOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePostgres) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourcePostgresArrayOutput struct{ *pulumi.OutputState }

func (SourcePostgresArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourcePostgres)(nil)).Elem()
}

func (o SourcePostgresArrayOutput) ToSourcePostgresArrayOutput() SourcePostgresArrayOutput {
	return o
}

func (o SourcePostgresArrayOutput) ToSourcePostgresArrayOutputWithContext(ctx context.Context) SourcePostgresArrayOutput {
	return o
}

func (o SourcePostgresArrayOutput) Index(i pulumi.IntInput) SourcePostgresOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourcePostgres {
		return vs[0].([]*SourcePostgres)[vs[1].(int)]
	}).(SourcePostgresOutput)
}

type SourcePostgresMapOutput struct{ *pulumi.OutputState }

func (SourcePostgresMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourcePostgres)(nil)).Elem()
}

func (o SourcePostgresMapOutput) ToSourcePostgresMapOutput() SourcePostgresMapOutput {
	return o
}

func (o SourcePostgresMapOutput) ToSourcePostgresMapOutputWithContext(ctx context.Context) SourcePostgresMapOutput {
	return o
}

func (o SourcePostgresMapOutput) MapIndex(k pulumi.StringInput) SourcePostgresOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourcePostgres {
		return vs[0].(map[string]*SourcePostgres)[vs[1].(string)]
	}).(SourcePostgresOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePostgresInput)(nil)).Elem(), &SourcePostgres{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePostgresArrayInput)(nil)).Elem(), SourcePostgresArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePostgresMapInput)(nil)).Elem(), SourcePostgresMap{})
	pulumi.RegisterOutputType(SourcePostgresOutput{})
	pulumi.RegisterOutputType(SourcePostgresArrayOutput{})
	pulumi.RegisterOutputType(SourcePostgresMapOutput{})
}
