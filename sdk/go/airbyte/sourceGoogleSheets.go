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

// SourceGoogleSheets Resource
type SourceGoogleSheets struct {
	pulumi.CustomResourceState

	Configuration SourceGoogleSheetsConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceGoogleSheets registers a new resource with the given unique name, arguments, and options.
func NewSourceGoogleSheets(ctx *pulumi.Context,
	name string, args *SourceGoogleSheetsArgs, opts ...pulumi.ResourceOption) (*SourceGoogleSheets, error) {
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
	var resource SourceGoogleSheets
	err := ctx.RegisterResource("airbyte:index/sourceGoogleSheets:SourceGoogleSheets", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceGoogleSheets gets an existing SourceGoogleSheets resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceGoogleSheets(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceGoogleSheetsState, opts ...pulumi.ResourceOption) (*SourceGoogleSheets, error) {
	var resource SourceGoogleSheets
	err := ctx.ReadResource("airbyte:index/sourceGoogleSheets:SourceGoogleSheets", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceGoogleSheets resources.
type sourceGoogleSheetsState struct {
	Configuration *SourceGoogleSheetsConfiguration `pulumi:"configuration"`
	Name          *string                          `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceGoogleSheetsState struct {
	Configuration SourceGoogleSheetsConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceGoogleSheetsState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGoogleSheetsState)(nil)).Elem()
}

type sourceGoogleSheetsArgs struct {
	Configuration SourceGoogleSheetsConfiguration `pulumi:"configuration"`
	Name          string                          `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceGoogleSheets resource.
type SourceGoogleSheetsArgs struct {
	Configuration SourceGoogleSheetsConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceGoogleSheetsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGoogleSheetsArgs)(nil)).Elem()
}

type SourceGoogleSheetsInput interface {
	pulumi.Input

	ToSourceGoogleSheetsOutput() SourceGoogleSheetsOutput
	ToSourceGoogleSheetsOutputWithContext(ctx context.Context) SourceGoogleSheetsOutput
}

func (*SourceGoogleSheets) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGoogleSheets)(nil)).Elem()
}

func (i *SourceGoogleSheets) ToSourceGoogleSheetsOutput() SourceGoogleSheetsOutput {
	return i.ToSourceGoogleSheetsOutputWithContext(context.Background())
}

func (i *SourceGoogleSheets) ToSourceGoogleSheetsOutputWithContext(ctx context.Context) SourceGoogleSheetsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleSheetsOutput)
}

// SourceGoogleSheetsArrayInput is an input type that accepts SourceGoogleSheetsArray and SourceGoogleSheetsArrayOutput values.
// You can construct a concrete instance of `SourceGoogleSheetsArrayInput` via:
//
//	SourceGoogleSheetsArray{ SourceGoogleSheetsArgs{...} }
type SourceGoogleSheetsArrayInput interface {
	pulumi.Input

	ToSourceGoogleSheetsArrayOutput() SourceGoogleSheetsArrayOutput
	ToSourceGoogleSheetsArrayOutputWithContext(context.Context) SourceGoogleSheetsArrayOutput
}

type SourceGoogleSheetsArray []SourceGoogleSheetsInput

func (SourceGoogleSheetsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGoogleSheets)(nil)).Elem()
}

func (i SourceGoogleSheetsArray) ToSourceGoogleSheetsArrayOutput() SourceGoogleSheetsArrayOutput {
	return i.ToSourceGoogleSheetsArrayOutputWithContext(context.Background())
}

func (i SourceGoogleSheetsArray) ToSourceGoogleSheetsArrayOutputWithContext(ctx context.Context) SourceGoogleSheetsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleSheetsArrayOutput)
}

// SourceGoogleSheetsMapInput is an input type that accepts SourceGoogleSheetsMap and SourceGoogleSheetsMapOutput values.
// You can construct a concrete instance of `SourceGoogleSheetsMapInput` via:
//
//	SourceGoogleSheetsMap{ "key": SourceGoogleSheetsArgs{...} }
type SourceGoogleSheetsMapInput interface {
	pulumi.Input

	ToSourceGoogleSheetsMapOutput() SourceGoogleSheetsMapOutput
	ToSourceGoogleSheetsMapOutputWithContext(context.Context) SourceGoogleSheetsMapOutput
}

type SourceGoogleSheetsMap map[string]SourceGoogleSheetsInput

func (SourceGoogleSheetsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGoogleSheets)(nil)).Elem()
}

func (i SourceGoogleSheetsMap) ToSourceGoogleSheetsMapOutput() SourceGoogleSheetsMapOutput {
	return i.ToSourceGoogleSheetsMapOutputWithContext(context.Background())
}

func (i SourceGoogleSheetsMap) ToSourceGoogleSheetsMapOutputWithContext(ctx context.Context) SourceGoogleSheetsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleSheetsMapOutput)
}

type SourceGoogleSheetsOutput struct{ *pulumi.OutputState }

func (SourceGoogleSheetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGoogleSheets)(nil)).Elem()
}

func (o SourceGoogleSheetsOutput) ToSourceGoogleSheetsOutput() SourceGoogleSheetsOutput {
	return o
}

func (o SourceGoogleSheetsOutput) ToSourceGoogleSheetsOutputWithContext(ctx context.Context) SourceGoogleSheetsOutput {
	return o
}

func (o SourceGoogleSheetsOutput) Configuration() SourceGoogleSheetsConfigurationOutput {
	return o.ApplyT(func(v *SourceGoogleSheets) SourceGoogleSheetsConfigurationOutput { return v.Configuration }).(SourceGoogleSheetsConfigurationOutput)
}

func (o SourceGoogleSheetsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleSheets) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceGoogleSheetsOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGoogleSheets) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceGoogleSheetsOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleSheets) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceGoogleSheetsOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleSheets) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceGoogleSheetsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleSheets) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceGoogleSheetsArrayOutput struct{ *pulumi.OutputState }

func (SourceGoogleSheetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGoogleSheets)(nil)).Elem()
}

func (o SourceGoogleSheetsArrayOutput) ToSourceGoogleSheetsArrayOutput() SourceGoogleSheetsArrayOutput {
	return o
}

func (o SourceGoogleSheetsArrayOutput) ToSourceGoogleSheetsArrayOutputWithContext(ctx context.Context) SourceGoogleSheetsArrayOutput {
	return o
}

func (o SourceGoogleSheetsArrayOutput) Index(i pulumi.IntInput) SourceGoogleSheetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceGoogleSheets {
		return vs[0].([]*SourceGoogleSheets)[vs[1].(int)]
	}).(SourceGoogleSheetsOutput)
}

type SourceGoogleSheetsMapOutput struct{ *pulumi.OutputState }

func (SourceGoogleSheetsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGoogleSheets)(nil)).Elem()
}

func (o SourceGoogleSheetsMapOutput) ToSourceGoogleSheetsMapOutput() SourceGoogleSheetsMapOutput {
	return o
}

func (o SourceGoogleSheetsMapOutput) ToSourceGoogleSheetsMapOutputWithContext(ctx context.Context) SourceGoogleSheetsMapOutput {
	return o
}

func (o SourceGoogleSheetsMapOutput) MapIndex(k pulumi.StringInput) SourceGoogleSheetsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceGoogleSheets {
		return vs[0].(map[string]*SourceGoogleSheets)[vs[1].(string)]
	}).(SourceGoogleSheetsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleSheetsInput)(nil)).Elem(), &SourceGoogleSheets{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleSheetsArrayInput)(nil)).Elem(), SourceGoogleSheetsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleSheetsMapInput)(nil)).Elem(), SourceGoogleSheetsMap{})
	pulumi.RegisterOutputType(SourceGoogleSheetsOutput{})
	pulumi.RegisterOutputType(SourceGoogleSheetsArrayOutput{})
	pulumi.RegisterOutputType(SourceGoogleSheetsMapOutput{})
}
