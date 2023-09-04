// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceSurveymonkey struct {
	pulumi.CustomResourceState

	Configuration SourceSurveymonkeyConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                   `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceSurveymonkey registers a new resource with the given unique name, arguments, and options.
func NewSourceSurveymonkey(ctx *pulumi.Context,
	name string, args *SourceSurveymonkeyArgs, opts ...pulumi.ResourceOption) (*SourceSurveymonkey, error) {
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
	var resource SourceSurveymonkey
	err := ctx.RegisterResource("airbyte:index/sourceSurveymonkey:SourceSurveymonkey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSurveymonkey gets an existing SourceSurveymonkey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSurveymonkey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSurveymonkeyState, opts ...pulumi.ResourceOption) (*SourceSurveymonkey, error) {
	var resource SourceSurveymonkey
	err := ctx.ReadResource("airbyte:index/sourceSurveymonkey:SourceSurveymonkey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSurveymonkey resources.
type sourceSurveymonkeyState struct {
	Configuration *SourceSurveymonkeyConfiguration `pulumi:"configuration"`
	Name          *string                          `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceSurveymonkeyState struct {
	Configuration SourceSurveymonkeyConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceSurveymonkeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSurveymonkeyState)(nil)).Elem()
}

type sourceSurveymonkeyArgs struct {
	Configuration SourceSurveymonkeyConfiguration `pulumi:"configuration"`
	Name          string                          `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSurveymonkey resource.
type SourceSurveymonkeyArgs struct {
	Configuration SourceSurveymonkeyConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSurveymonkeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSurveymonkeyArgs)(nil)).Elem()
}

type SourceSurveymonkeyInput interface {
	pulumi.Input

	ToSourceSurveymonkeyOutput() SourceSurveymonkeyOutput
	ToSourceSurveymonkeyOutputWithContext(ctx context.Context) SourceSurveymonkeyOutput
}

func (*SourceSurveymonkey) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSurveymonkey)(nil)).Elem()
}

func (i *SourceSurveymonkey) ToSourceSurveymonkeyOutput() SourceSurveymonkeyOutput {
	return i.ToSourceSurveymonkeyOutputWithContext(context.Background())
}

func (i *SourceSurveymonkey) ToSourceSurveymonkeyOutputWithContext(ctx context.Context) SourceSurveymonkeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSurveymonkeyOutput)
}

type SourceSurveymonkeyOutput struct{ *pulumi.OutputState }

func (SourceSurveymonkeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSurveymonkey)(nil)).Elem()
}

func (o SourceSurveymonkeyOutput) ToSourceSurveymonkeyOutput() SourceSurveymonkeyOutput {
	return o
}

func (o SourceSurveymonkeyOutput) ToSourceSurveymonkeyOutputWithContext(ctx context.Context) SourceSurveymonkeyOutput {
	return o
}

func (o SourceSurveymonkeyOutput) Configuration() SourceSurveymonkeyConfigurationOutput {
	return o.ApplyT(func(v *SourceSurveymonkey) SourceSurveymonkeyConfigurationOutput { return v.Configuration }).(SourceSurveymonkeyConfigurationOutput)
}

func (o SourceSurveymonkeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSurveymonkey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceSurveymonkeyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSurveymonkey) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSurveymonkeyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSurveymonkey) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSurveymonkeyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSurveymonkey) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSurveymonkeyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSurveymonkey) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSurveymonkeyInput)(nil)).Elem(), &SourceSurveymonkey{})
	pulumi.RegisterOutputType(SourceSurveymonkeyOutput{})
}
