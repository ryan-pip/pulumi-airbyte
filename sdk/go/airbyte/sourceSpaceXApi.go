// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceSpaceXApi struct {
	pulumi.CustomResourceState

	Configuration SourceSpaceXApiConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceSpaceXApi registers a new resource with the given unique name, arguments, and options.
func NewSourceSpaceXApi(ctx *pulumi.Context,
	name string, args *SourceSpaceXApiArgs, opts ...pulumi.ResourceOption) (*SourceSpaceXApi, error) {
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
	opts = pkgResourceDefaultOpts(opts)
	var resource SourceSpaceXApi
	err := ctx.RegisterResource("airbyte:index/sourceSpaceXApi:SourceSpaceXApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSpaceXApi gets an existing SourceSpaceXApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSpaceXApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSpaceXApiState, opts ...pulumi.ResourceOption) (*SourceSpaceXApi, error) {
	var resource SourceSpaceXApi
	err := ctx.ReadResource("airbyte:index/sourceSpaceXApi:SourceSpaceXApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSpaceXApi resources.
type sourceSpaceXApiState struct {
	Configuration *SourceSpaceXApiConfiguration `pulumi:"configuration"`
	Name          *string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceSpaceXApiState struct {
	Configuration SourceSpaceXApiConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceSpaceXApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSpaceXApiState)(nil)).Elem()
}

type sourceSpaceXApiArgs struct {
	Configuration SourceSpaceXApiConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSpaceXApi resource.
type SourceSpaceXApiArgs struct {
	Configuration SourceSpaceXApiConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSpaceXApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSpaceXApiArgs)(nil)).Elem()
}

type SourceSpaceXApiInput interface {
	pulumi.Input

	ToSourceSpaceXApiOutput() SourceSpaceXApiOutput
	ToSourceSpaceXApiOutputWithContext(ctx context.Context) SourceSpaceXApiOutput
}

func (*SourceSpaceXApi) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSpaceXApi)(nil)).Elem()
}

func (i *SourceSpaceXApi) ToSourceSpaceXApiOutput() SourceSpaceXApiOutput {
	return i.ToSourceSpaceXApiOutputWithContext(context.Background())
}

func (i *SourceSpaceXApi) ToSourceSpaceXApiOutputWithContext(ctx context.Context) SourceSpaceXApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSpaceXApiOutput)
}

type SourceSpaceXApiOutput struct{ *pulumi.OutputState }

func (SourceSpaceXApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSpaceXApi)(nil)).Elem()
}

func (o SourceSpaceXApiOutput) ToSourceSpaceXApiOutput() SourceSpaceXApiOutput {
	return o
}

func (o SourceSpaceXApiOutput) ToSourceSpaceXApiOutputWithContext(ctx context.Context) SourceSpaceXApiOutput {
	return o
}

func (o SourceSpaceXApiOutput) Configuration() SourceSpaceXApiConfigurationOutput {
	return o.ApplyT(func(v *SourceSpaceXApi) SourceSpaceXApiConfigurationOutput { return v.Configuration }).(SourceSpaceXApiConfigurationOutput)
}

func (o SourceSpaceXApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSpaceXApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceSpaceXApiOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSpaceXApi) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSpaceXApiOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSpaceXApi) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSpaceXApiOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSpaceXApi) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSpaceXApiOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSpaceXApi) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSpaceXApiInput)(nil)).Elem(), &SourceSpaceXApi{})
	pulumi.RegisterOutputType(SourceSpaceXApiOutput{})
}
