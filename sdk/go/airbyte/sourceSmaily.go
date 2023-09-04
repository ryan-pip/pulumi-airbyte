// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceSmaily struct {
	pulumi.CustomResourceState

	Configuration SourceSmailyConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceSmaily registers a new resource with the given unique name, arguments, and options.
func NewSourceSmaily(ctx *pulumi.Context,
	name string, args *SourceSmailyArgs, opts ...pulumi.ResourceOption) (*SourceSmaily, error) {
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
	var resource SourceSmaily
	err := ctx.RegisterResource("airbyte:index/sourceSmaily:SourceSmaily", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSmaily gets an existing SourceSmaily resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSmaily(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSmailyState, opts ...pulumi.ResourceOption) (*SourceSmaily, error) {
	var resource SourceSmaily
	err := ctx.ReadResource("airbyte:index/sourceSmaily:SourceSmaily", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSmaily resources.
type sourceSmailyState struct {
	Configuration *SourceSmailyConfiguration `pulumi:"configuration"`
	Name          *string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceSmailyState struct {
	Configuration SourceSmailyConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceSmailyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSmailyState)(nil)).Elem()
}

type sourceSmailyArgs struct {
	Configuration SourceSmailyConfiguration `pulumi:"configuration"`
	Name          string                    `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSmaily resource.
type SourceSmailyArgs struct {
	Configuration SourceSmailyConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSmailyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSmailyArgs)(nil)).Elem()
}

type SourceSmailyInput interface {
	pulumi.Input

	ToSourceSmailyOutput() SourceSmailyOutput
	ToSourceSmailyOutputWithContext(ctx context.Context) SourceSmailyOutput
}

func (*SourceSmaily) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSmaily)(nil)).Elem()
}

func (i *SourceSmaily) ToSourceSmailyOutput() SourceSmailyOutput {
	return i.ToSourceSmailyOutputWithContext(context.Background())
}

func (i *SourceSmaily) ToSourceSmailyOutputWithContext(ctx context.Context) SourceSmailyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSmailyOutput)
}

type SourceSmailyOutput struct{ *pulumi.OutputState }

func (SourceSmailyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSmaily)(nil)).Elem()
}

func (o SourceSmailyOutput) ToSourceSmailyOutput() SourceSmailyOutput {
	return o
}

func (o SourceSmailyOutput) ToSourceSmailyOutputWithContext(ctx context.Context) SourceSmailyOutput {
	return o
}

func (o SourceSmailyOutput) Configuration() SourceSmailyConfigurationOutput {
	return o.ApplyT(func(v *SourceSmaily) SourceSmailyConfigurationOutput { return v.Configuration }).(SourceSmailyConfigurationOutput)
}

func (o SourceSmailyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceSmailyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSmailyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSmailyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSmailyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSmaily) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSmailyInput)(nil)).Elem(), &SourceSmaily{})
	pulumi.RegisterOutputType(SourceSmailyOutput{})
}
