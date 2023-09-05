// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

// SourceWikipediaPageviews Resource
type SourceWikipediaPageviews struct {
	pulumi.CustomResourceState

	Configuration SourceWikipediaPageviewsConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                         `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceWikipediaPageviews registers a new resource with the given unique name, arguments, and options.
func NewSourceWikipediaPageviews(ctx *pulumi.Context,
	name string, args *SourceWikipediaPageviewsArgs, opts ...pulumi.ResourceOption) (*SourceWikipediaPageviews, error) {
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
	var resource SourceWikipediaPageviews
	err := ctx.RegisterResource("airbyte:index/sourceWikipediaPageviews:SourceWikipediaPageviews", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceWikipediaPageviews gets an existing SourceWikipediaPageviews resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceWikipediaPageviews(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceWikipediaPageviewsState, opts ...pulumi.ResourceOption) (*SourceWikipediaPageviews, error) {
	var resource SourceWikipediaPageviews
	err := ctx.ReadResource("airbyte:index/sourceWikipediaPageviews:SourceWikipediaPageviews", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceWikipediaPageviews resources.
type sourceWikipediaPageviewsState struct {
	Configuration *SourceWikipediaPageviewsConfiguration `pulumi:"configuration"`
	Name          *string                                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceWikipediaPageviewsState struct {
	Configuration SourceWikipediaPageviewsConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceWikipediaPageviewsState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceWikipediaPageviewsState)(nil)).Elem()
}

type sourceWikipediaPageviewsArgs struct {
	Configuration SourceWikipediaPageviewsConfiguration `pulumi:"configuration"`
	Name          string                                `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceWikipediaPageviews resource.
type SourceWikipediaPageviewsArgs struct {
	Configuration SourceWikipediaPageviewsConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceWikipediaPageviewsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceWikipediaPageviewsArgs)(nil)).Elem()
}

type SourceWikipediaPageviewsInput interface {
	pulumi.Input

	ToSourceWikipediaPageviewsOutput() SourceWikipediaPageviewsOutput
	ToSourceWikipediaPageviewsOutputWithContext(ctx context.Context) SourceWikipediaPageviewsOutput
}

func (*SourceWikipediaPageviews) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceWikipediaPageviews)(nil)).Elem()
}

func (i *SourceWikipediaPageviews) ToSourceWikipediaPageviewsOutput() SourceWikipediaPageviewsOutput {
	return i.ToSourceWikipediaPageviewsOutputWithContext(context.Background())
}

func (i *SourceWikipediaPageviews) ToSourceWikipediaPageviewsOutputWithContext(ctx context.Context) SourceWikipediaPageviewsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceWikipediaPageviewsOutput)
}

type SourceWikipediaPageviewsOutput struct{ *pulumi.OutputState }

func (SourceWikipediaPageviewsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceWikipediaPageviews)(nil)).Elem()
}

func (o SourceWikipediaPageviewsOutput) ToSourceWikipediaPageviewsOutput() SourceWikipediaPageviewsOutput {
	return o
}

func (o SourceWikipediaPageviewsOutput) ToSourceWikipediaPageviewsOutputWithContext(ctx context.Context) SourceWikipediaPageviewsOutput {
	return o
}

func (o SourceWikipediaPageviewsOutput) Configuration() SourceWikipediaPageviewsConfigurationOutput {
	return o.ApplyT(func(v *SourceWikipediaPageviews) SourceWikipediaPageviewsConfigurationOutput { return v.Configuration }).(SourceWikipediaPageviewsConfigurationOutput)
}

func (o SourceWikipediaPageviewsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceWikipediaPageviews) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceWikipediaPageviewsOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceWikipediaPageviews) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceWikipediaPageviewsOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceWikipediaPageviews) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceWikipediaPageviewsOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceWikipediaPageviews) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceWikipediaPageviewsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceWikipediaPageviews) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceWikipediaPageviewsInput)(nil)).Elem(), &SourceWikipediaPageviews{})
	pulumi.RegisterOutputType(SourceWikipediaPageviewsOutput{})
}
