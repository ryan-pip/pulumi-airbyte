// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceTheGuardianApi struct {
	pulumi.CustomResourceState

	Configuration SourceTheGuardianApiConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceTheGuardianApi registers a new resource with the given unique name, arguments, and options.
func NewSourceTheGuardianApi(ctx *pulumi.Context,
	name string, args *SourceTheGuardianApiArgs, opts ...pulumi.ResourceOption) (*SourceTheGuardianApi, error) {
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
	var resource SourceTheGuardianApi
	err := ctx.RegisterResource("airbyte:index/sourceTheGuardianApi:SourceTheGuardianApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceTheGuardianApi gets an existing SourceTheGuardianApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceTheGuardianApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceTheGuardianApiState, opts ...pulumi.ResourceOption) (*SourceTheGuardianApi, error) {
	var resource SourceTheGuardianApi
	err := ctx.ReadResource("airbyte:index/sourceTheGuardianApi:SourceTheGuardianApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceTheGuardianApi resources.
type sourceTheGuardianApiState struct {
	Configuration *SourceTheGuardianApiConfiguration `pulumi:"configuration"`
	Name          *string                            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceTheGuardianApiState struct {
	Configuration SourceTheGuardianApiConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceTheGuardianApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceTheGuardianApiState)(nil)).Elem()
}

type sourceTheGuardianApiArgs struct {
	Configuration SourceTheGuardianApiConfiguration `pulumi:"configuration"`
	Name          string                            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceTheGuardianApi resource.
type SourceTheGuardianApiArgs struct {
	Configuration SourceTheGuardianApiConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceTheGuardianApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceTheGuardianApiArgs)(nil)).Elem()
}

type SourceTheGuardianApiInput interface {
	pulumi.Input

	ToSourceTheGuardianApiOutput() SourceTheGuardianApiOutput
	ToSourceTheGuardianApiOutputWithContext(ctx context.Context) SourceTheGuardianApiOutput
}

func (*SourceTheGuardianApi) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTheGuardianApi)(nil)).Elem()
}

func (i *SourceTheGuardianApi) ToSourceTheGuardianApiOutput() SourceTheGuardianApiOutput {
	return i.ToSourceTheGuardianApiOutputWithContext(context.Background())
}

func (i *SourceTheGuardianApi) ToSourceTheGuardianApiOutputWithContext(ctx context.Context) SourceTheGuardianApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTheGuardianApiOutput)
}

type SourceTheGuardianApiOutput struct{ *pulumi.OutputState }

func (SourceTheGuardianApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTheGuardianApi)(nil)).Elem()
}

func (o SourceTheGuardianApiOutput) ToSourceTheGuardianApiOutput() SourceTheGuardianApiOutput {
	return o
}

func (o SourceTheGuardianApiOutput) ToSourceTheGuardianApiOutputWithContext(ctx context.Context) SourceTheGuardianApiOutput {
	return o
}

func (o SourceTheGuardianApiOutput) Configuration() SourceTheGuardianApiConfigurationOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) SourceTheGuardianApiConfigurationOutput { return v.Configuration }).(SourceTheGuardianApiConfigurationOutput)
}

func (o SourceTheGuardianApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceTheGuardianApiOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceTheGuardianApiOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceTheGuardianApiOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceTheGuardianApiOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTheGuardianApiInput)(nil)).Elem(), &SourceTheGuardianApi{})
	pulumi.RegisterOutputType(SourceTheGuardianApiOutput{})
}