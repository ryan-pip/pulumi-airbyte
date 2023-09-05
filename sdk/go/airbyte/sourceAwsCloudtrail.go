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

// SourceAwsCloudtrail Resource
type SourceAwsCloudtrail struct {
	pulumi.CustomResourceState

	Configuration SourceAwsCloudtrailConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceAwsCloudtrail registers a new resource with the given unique name, arguments, and options.
func NewSourceAwsCloudtrail(ctx *pulumi.Context,
	name string, args *SourceAwsCloudtrailArgs, opts ...pulumi.ResourceOption) (*SourceAwsCloudtrail, error) {
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
	var resource SourceAwsCloudtrail
	err := ctx.RegisterResource("airbyte:index/sourceAwsCloudtrail:SourceAwsCloudtrail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceAwsCloudtrail gets an existing SourceAwsCloudtrail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceAwsCloudtrail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceAwsCloudtrailState, opts ...pulumi.ResourceOption) (*SourceAwsCloudtrail, error) {
	var resource SourceAwsCloudtrail
	err := ctx.ReadResource("airbyte:index/sourceAwsCloudtrail:SourceAwsCloudtrail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceAwsCloudtrail resources.
type sourceAwsCloudtrailState struct {
	Configuration *SourceAwsCloudtrailConfiguration `pulumi:"configuration"`
	Name          *string                           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceAwsCloudtrailState struct {
	Configuration SourceAwsCloudtrailConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceAwsCloudtrailState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAwsCloudtrailState)(nil)).Elem()
}

type sourceAwsCloudtrailArgs struct {
	Configuration SourceAwsCloudtrailConfiguration `pulumi:"configuration"`
	Name          string                           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceAwsCloudtrail resource.
type SourceAwsCloudtrailArgs struct {
	Configuration SourceAwsCloudtrailConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceAwsCloudtrailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAwsCloudtrailArgs)(nil)).Elem()
}

type SourceAwsCloudtrailInput interface {
	pulumi.Input

	ToSourceAwsCloudtrailOutput() SourceAwsCloudtrailOutput
	ToSourceAwsCloudtrailOutputWithContext(ctx context.Context) SourceAwsCloudtrailOutput
}

func (*SourceAwsCloudtrail) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAwsCloudtrail)(nil)).Elem()
}

func (i *SourceAwsCloudtrail) ToSourceAwsCloudtrailOutput() SourceAwsCloudtrailOutput {
	return i.ToSourceAwsCloudtrailOutputWithContext(context.Background())
}

func (i *SourceAwsCloudtrail) ToSourceAwsCloudtrailOutputWithContext(ctx context.Context) SourceAwsCloudtrailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAwsCloudtrailOutput)
}

// SourceAwsCloudtrailArrayInput is an input type that accepts SourceAwsCloudtrailArray and SourceAwsCloudtrailArrayOutput values.
// You can construct a concrete instance of `SourceAwsCloudtrailArrayInput` via:
//
//	SourceAwsCloudtrailArray{ SourceAwsCloudtrailArgs{...} }
type SourceAwsCloudtrailArrayInput interface {
	pulumi.Input

	ToSourceAwsCloudtrailArrayOutput() SourceAwsCloudtrailArrayOutput
	ToSourceAwsCloudtrailArrayOutputWithContext(context.Context) SourceAwsCloudtrailArrayOutput
}

type SourceAwsCloudtrailArray []SourceAwsCloudtrailInput

func (SourceAwsCloudtrailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAwsCloudtrail)(nil)).Elem()
}

func (i SourceAwsCloudtrailArray) ToSourceAwsCloudtrailArrayOutput() SourceAwsCloudtrailArrayOutput {
	return i.ToSourceAwsCloudtrailArrayOutputWithContext(context.Background())
}

func (i SourceAwsCloudtrailArray) ToSourceAwsCloudtrailArrayOutputWithContext(ctx context.Context) SourceAwsCloudtrailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAwsCloudtrailArrayOutput)
}

// SourceAwsCloudtrailMapInput is an input type that accepts SourceAwsCloudtrailMap and SourceAwsCloudtrailMapOutput values.
// You can construct a concrete instance of `SourceAwsCloudtrailMapInput` via:
//
//	SourceAwsCloudtrailMap{ "key": SourceAwsCloudtrailArgs{...} }
type SourceAwsCloudtrailMapInput interface {
	pulumi.Input

	ToSourceAwsCloudtrailMapOutput() SourceAwsCloudtrailMapOutput
	ToSourceAwsCloudtrailMapOutputWithContext(context.Context) SourceAwsCloudtrailMapOutput
}

type SourceAwsCloudtrailMap map[string]SourceAwsCloudtrailInput

func (SourceAwsCloudtrailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAwsCloudtrail)(nil)).Elem()
}

func (i SourceAwsCloudtrailMap) ToSourceAwsCloudtrailMapOutput() SourceAwsCloudtrailMapOutput {
	return i.ToSourceAwsCloudtrailMapOutputWithContext(context.Background())
}

func (i SourceAwsCloudtrailMap) ToSourceAwsCloudtrailMapOutputWithContext(ctx context.Context) SourceAwsCloudtrailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAwsCloudtrailMapOutput)
}

type SourceAwsCloudtrailOutput struct{ *pulumi.OutputState }

func (SourceAwsCloudtrailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAwsCloudtrail)(nil)).Elem()
}

func (o SourceAwsCloudtrailOutput) ToSourceAwsCloudtrailOutput() SourceAwsCloudtrailOutput {
	return o
}

func (o SourceAwsCloudtrailOutput) ToSourceAwsCloudtrailOutputWithContext(ctx context.Context) SourceAwsCloudtrailOutput {
	return o
}

func (o SourceAwsCloudtrailOutput) Configuration() SourceAwsCloudtrailConfigurationOutput {
	return o.ApplyT(func(v *SourceAwsCloudtrail) SourceAwsCloudtrailConfigurationOutput { return v.Configuration }).(SourceAwsCloudtrailConfigurationOutput)
}

func (o SourceAwsCloudtrailOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAwsCloudtrail) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceAwsCloudtrailOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAwsCloudtrail) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceAwsCloudtrailOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAwsCloudtrail) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceAwsCloudtrailOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAwsCloudtrail) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceAwsCloudtrailOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAwsCloudtrail) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceAwsCloudtrailArrayOutput struct{ *pulumi.OutputState }

func (SourceAwsCloudtrailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAwsCloudtrail)(nil)).Elem()
}

func (o SourceAwsCloudtrailArrayOutput) ToSourceAwsCloudtrailArrayOutput() SourceAwsCloudtrailArrayOutput {
	return o
}

func (o SourceAwsCloudtrailArrayOutput) ToSourceAwsCloudtrailArrayOutputWithContext(ctx context.Context) SourceAwsCloudtrailArrayOutput {
	return o
}

func (o SourceAwsCloudtrailArrayOutput) Index(i pulumi.IntInput) SourceAwsCloudtrailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceAwsCloudtrail {
		return vs[0].([]*SourceAwsCloudtrail)[vs[1].(int)]
	}).(SourceAwsCloudtrailOutput)
}

type SourceAwsCloudtrailMapOutput struct{ *pulumi.OutputState }

func (SourceAwsCloudtrailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAwsCloudtrail)(nil)).Elem()
}

func (o SourceAwsCloudtrailMapOutput) ToSourceAwsCloudtrailMapOutput() SourceAwsCloudtrailMapOutput {
	return o
}

func (o SourceAwsCloudtrailMapOutput) ToSourceAwsCloudtrailMapOutputWithContext(ctx context.Context) SourceAwsCloudtrailMapOutput {
	return o
}

func (o SourceAwsCloudtrailMapOutput) MapIndex(k pulumi.StringInput) SourceAwsCloudtrailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceAwsCloudtrail {
		return vs[0].(map[string]*SourceAwsCloudtrail)[vs[1].(string)]
	}).(SourceAwsCloudtrailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAwsCloudtrailInput)(nil)).Elem(), &SourceAwsCloudtrail{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAwsCloudtrailArrayInput)(nil)).Elem(), SourceAwsCloudtrailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAwsCloudtrailMapInput)(nil)).Elem(), SourceAwsCloudtrailMap{})
	pulumi.RegisterOutputType(SourceAwsCloudtrailOutput{})
	pulumi.RegisterOutputType(SourceAwsCloudtrailArrayOutput{})
	pulumi.RegisterOutputType(SourceAwsCloudtrailMapOutput{})
}
