// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourcePostmarkapp struct {
	pulumi.CustomResourceState

	Configuration SourcePostmarkappConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                  `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourcePostmarkapp registers a new resource with the given unique name, arguments, and options.
func NewSourcePostmarkapp(ctx *pulumi.Context,
	name string, args *SourcePostmarkappArgs, opts ...pulumi.ResourceOption) (*SourcePostmarkapp, error) {
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
	var resource SourcePostmarkapp
	err := ctx.RegisterResource("airbyte:index/sourcePostmarkapp:SourcePostmarkapp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourcePostmarkapp gets an existing SourcePostmarkapp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourcePostmarkapp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourcePostmarkappState, opts ...pulumi.ResourceOption) (*SourcePostmarkapp, error) {
	var resource SourcePostmarkapp
	err := ctx.ReadResource("airbyte:index/sourcePostmarkapp:SourcePostmarkapp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourcePostmarkapp resources.
type sourcePostmarkappState struct {
	Configuration *SourcePostmarkappConfiguration `pulumi:"configuration"`
	Name          *string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourcePostmarkappState struct {
	Configuration SourcePostmarkappConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourcePostmarkappState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePostmarkappState)(nil)).Elem()
}

type sourcePostmarkappArgs struct {
	Configuration SourcePostmarkappConfiguration `pulumi:"configuration"`
	Name          string                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourcePostmarkapp resource.
type SourcePostmarkappArgs struct {
	Configuration SourcePostmarkappConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourcePostmarkappArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePostmarkappArgs)(nil)).Elem()
}

type SourcePostmarkappInput interface {
	pulumi.Input

	ToSourcePostmarkappOutput() SourcePostmarkappOutput
	ToSourcePostmarkappOutputWithContext(ctx context.Context) SourcePostmarkappOutput
}

func (*SourcePostmarkapp) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePostmarkapp)(nil)).Elem()
}

func (i *SourcePostmarkapp) ToSourcePostmarkappOutput() SourcePostmarkappOutput {
	return i.ToSourcePostmarkappOutputWithContext(context.Background())
}

func (i *SourcePostmarkapp) ToSourcePostmarkappOutputWithContext(ctx context.Context) SourcePostmarkappOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePostmarkappOutput)
}

type SourcePostmarkappOutput struct{ *pulumi.OutputState }

func (SourcePostmarkappOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePostmarkapp)(nil)).Elem()
}

func (o SourcePostmarkappOutput) ToSourcePostmarkappOutput() SourcePostmarkappOutput {
	return o
}

func (o SourcePostmarkappOutput) ToSourcePostmarkappOutputWithContext(ctx context.Context) SourcePostmarkappOutput {
	return o
}

func (o SourcePostmarkappOutput) Configuration() SourcePostmarkappConfigurationOutput {
	return o.ApplyT(func(v *SourcePostmarkapp) SourcePostmarkappConfigurationOutput { return v.Configuration }).(SourcePostmarkappConfigurationOutput)
}

func (o SourcePostmarkappOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePostmarkapp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourcePostmarkappOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePostmarkapp) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourcePostmarkappOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePostmarkapp) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourcePostmarkappOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePostmarkapp) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourcePostmarkappOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePostmarkapp) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePostmarkappInput)(nil)).Elem(), &SourcePostmarkapp{})
	pulumi.RegisterOutputType(SourcePostmarkappOutput{})
}
