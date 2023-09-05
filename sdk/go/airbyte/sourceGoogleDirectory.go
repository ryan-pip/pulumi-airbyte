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

// SourceGoogleDirectory Resource
type SourceGoogleDirectory struct {
	pulumi.CustomResourceState

	Configuration SourceGoogleDirectoryConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceGoogleDirectory registers a new resource with the given unique name, arguments, and options.
func NewSourceGoogleDirectory(ctx *pulumi.Context,
	name string, args *SourceGoogleDirectoryArgs, opts ...pulumi.ResourceOption) (*SourceGoogleDirectory, error) {
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
	var resource SourceGoogleDirectory
	err := ctx.RegisterResource("airbyte:index/sourceGoogleDirectory:SourceGoogleDirectory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceGoogleDirectory gets an existing SourceGoogleDirectory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceGoogleDirectory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceGoogleDirectoryState, opts ...pulumi.ResourceOption) (*SourceGoogleDirectory, error) {
	var resource SourceGoogleDirectory
	err := ctx.ReadResource("airbyte:index/sourceGoogleDirectory:SourceGoogleDirectory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceGoogleDirectory resources.
type sourceGoogleDirectoryState struct {
	Configuration *SourceGoogleDirectoryConfiguration `pulumi:"configuration"`
	Name          *string                             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceGoogleDirectoryState struct {
	Configuration SourceGoogleDirectoryConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceGoogleDirectoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGoogleDirectoryState)(nil)).Elem()
}

type sourceGoogleDirectoryArgs struct {
	Configuration SourceGoogleDirectoryConfiguration `pulumi:"configuration"`
	Name          string                             `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceGoogleDirectory resource.
type SourceGoogleDirectoryArgs struct {
	Configuration SourceGoogleDirectoryConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceGoogleDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGoogleDirectoryArgs)(nil)).Elem()
}

type SourceGoogleDirectoryInput interface {
	pulumi.Input

	ToSourceGoogleDirectoryOutput() SourceGoogleDirectoryOutput
	ToSourceGoogleDirectoryOutputWithContext(ctx context.Context) SourceGoogleDirectoryOutput
}

func (*SourceGoogleDirectory) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGoogleDirectory)(nil)).Elem()
}

func (i *SourceGoogleDirectory) ToSourceGoogleDirectoryOutput() SourceGoogleDirectoryOutput {
	return i.ToSourceGoogleDirectoryOutputWithContext(context.Background())
}

func (i *SourceGoogleDirectory) ToSourceGoogleDirectoryOutputWithContext(ctx context.Context) SourceGoogleDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleDirectoryOutput)
}

type SourceGoogleDirectoryOutput struct{ *pulumi.OutputState }

func (SourceGoogleDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGoogleDirectory)(nil)).Elem()
}

func (o SourceGoogleDirectoryOutput) ToSourceGoogleDirectoryOutput() SourceGoogleDirectoryOutput {
	return o
}

func (o SourceGoogleDirectoryOutput) ToSourceGoogleDirectoryOutputWithContext(ctx context.Context) SourceGoogleDirectoryOutput {
	return o
}

func (o SourceGoogleDirectoryOutput) Configuration() SourceGoogleDirectoryConfigurationOutput {
	return o.ApplyT(func(v *SourceGoogleDirectory) SourceGoogleDirectoryConfigurationOutput { return v.Configuration }).(SourceGoogleDirectoryConfigurationOutput)
}

func (o SourceGoogleDirectoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleDirectory) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceGoogleDirectoryOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGoogleDirectory) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceGoogleDirectoryOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleDirectory) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceGoogleDirectoryOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleDirectory) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceGoogleDirectoryOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleDirectory) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleDirectoryInput)(nil)).Elem(), &SourceGoogleDirectory{})
	pulumi.RegisterOutputType(SourceGoogleDirectoryOutput{})
}
