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

// SourceSftp Resource
type SourceSftp struct {
	pulumi.CustomResourceState

	Configuration SourceSftpConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput           `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceSftp registers a new resource with the given unique name, arguments, and options.
func NewSourceSftp(ctx *pulumi.Context,
	name string, args *SourceSftpArgs, opts ...pulumi.ResourceOption) (*SourceSftp, error) {
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
	var resource SourceSftp
	err := ctx.RegisterResource("airbyte:index/sourceSftp:SourceSftp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSftp gets an existing SourceSftp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSftp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSftpState, opts ...pulumi.ResourceOption) (*SourceSftp, error) {
	var resource SourceSftp
	err := ctx.ReadResource("airbyte:index/sourceSftp:SourceSftp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSftp resources.
type sourceSftpState struct {
	Configuration *SourceSftpConfiguration `pulumi:"configuration"`
	Name          *string                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceSftpState struct {
	Configuration SourceSftpConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceSftpState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSftpState)(nil)).Elem()
}

type sourceSftpArgs struct {
	Configuration SourceSftpConfiguration `pulumi:"configuration"`
	Name          string                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSftp resource.
type SourceSftpArgs struct {
	Configuration SourceSftpConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSftpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSftpArgs)(nil)).Elem()
}

type SourceSftpInput interface {
	pulumi.Input

	ToSourceSftpOutput() SourceSftpOutput
	ToSourceSftpOutputWithContext(ctx context.Context) SourceSftpOutput
}

func (*SourceSftp) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSftp)(nil)).Elem()
}

func (i *SourceSftp) ToSourceSftpOutput() SourceSftpOutput {
	return i.ToSourceSftpOutputWithContext(context.Background())
}

func (i *SourceSftp) ToSourceSftpOutputWithContext(ctx context.Context) SourceSftpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSftpOutput)
}

// SourceSftpArrayInput is an input type that accepts SourceSftpArray and SourceSftpArrayOutput values.
// You can construct a concrete instance of `SourceSftpArrayInput` via:
//
//	SourceSftpArray{ SourceSftpArgs{...} }
type SourceSftpArrayInput interface {
	pulumi.Input

	ToSourceSftpArrayOutput() SourceSftpArrayOutput
	ToSourceSftpArrayOutputWithContext(context.Context) SourceSftpArrayOutput
}

type SourceSftpArray []SourceSftpInput

func (SourceSftpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSftp)(nil)).Elem()
}

func (i SourceSftpArray) ToSourceSftpArrayOutput() SourceSftpArrayOutput {
	return i.ToSourceSftpArrayOutputWithContext(context.Background())
}

func (i SourceSftpArray) ToSourceSftpArrayOutputWithContext(ctx context.Context) SourceSftpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSftpArrayOutput)
}

// SourceSftpMapInput is an input type that accepts SourceSftpMap and SourceSftpMapOutput values.
// You can construct a concrete instance of `SourceSftpMapInput` via:
//
//	SourceSftpMap{ "key": SourceSftpArgs{...} }
type SourceSftpMapInput interface {
	pulumi.Input

	ToSourceSftpMapOutput() SourceSftpMapOutput
	ToSourceSftpMapOutputWithContext(context.Context) SourceSftpMapOutput
}

type SourceSftpMap map[string]SourceSftpInput

func (SourceSftpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSftp)(nil)).Elem()
}

func (i SourceSftpMap) ToSourceSftpMapOutput() SourceSftpMapOutput {
	return i.ToSourceSftpMapOutputWithContext(context.Background())
}

func (i SourceSftpMap) ToSourceSftpMapOutputWithContext(ctx context.Context) SourceSftpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSftpMapOutput)
}

type SourceSftpOutput struct{ *pulumi.OutputState }

func (SourceSftpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSftp)(nil)).Elem()
}

func (o SourceSftpOutput) ToSourceSftpOutput() SourceSftpOutput {
	return o
}

func (o SourceSftpOutput) ToSourceSftpOutputWithContext(ctx context.Context) SourceSftpOutput {
	return o
}

func (o SourceSftpOutput) Configuration() SourceSftpConfigurationOutput {
	return o.ApplyT(func(v *SourceSftp) SourceSftpConfigurationOutput { return v.Configuration }).(SourceSftpConfigurationOutput)
}

func (o SourceSftpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSftp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceSftpOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSftp) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSftpOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSftp) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSftpOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSftp) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSftpOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSftp) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceSftpArrayOutput struct{ *pulumi.OutputState }

func (SourceSftpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSftp)(nil)).Elem()
}

func (o SourceSftpArrayOutput) ToSourceSftpArrayOutput() SourceSftpArrayOutput {
	return o
}

func (o SourceSftpArrayOutput) ToSourceSftpArrayOutputWithContext(ctx context.Context) SourceSftpArrayOutput {
	return o
}

func (o SourceSftpArrayOutput) Index(i pulumi.IntInput) SourceSftpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceSftp {
		return vs[0].([]*SourceSftp)[vs[1].(int)]
	}).(SourceSftpOutput)
}

type SourceSftpMapOutput struct{ *pulumi.OutputState }

func (SourceSftpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSftp)(nil)).Elem()
}

func (o SourceSftpMapOutput) ToSourceSftpMapOutput() SourceSftpMapOutput {
	return o
}

func (o SourceSftpMapOutput) ToSourceSftpMapOutputWithContext(ctx context.Context) SourceSftpMapOutput {
	return o
}

func (o SourceSftpMapOutput) MapIndex(k pulumi.StringInput) SourceSftpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceSftp {
		return vs[0].(map[string]*SourceSftp)[vs[1].(string)]
	}).(SourceSftpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSftpInput)(nil)).Elem(), &SourceSftp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSftpArrayInput)(nil)).Elem(), SourceSftpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSftpMapInput)(nil)).Elem(), SourceSftpMap{})
	pulumi.RegisterOutputType(SourceSftpOutput{})
	pulumi.RegisterOutputType(SourceSftpArrayOutput{})
	pulumi.RegisterOutputType(SourceSftpMapOutput{})
}
