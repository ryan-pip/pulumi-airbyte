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

// SourceMicrosoftTeams Resource
type SourceMicrosoftTeams struct {
	pulumi.CustomResourceState

	Configuration SourceMicrosoftTeamsConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceMicrosoftTeams registers a new resource with the given unique name, arguments, and options.
func NewSourceMicrosoftTeams(ctx *pulumi.Context,
	name string, args *SourceMicrosoftTeamsArgs, opts ...pulumi.ResourceOption) (*SourceMicrosoftTeams, error) {
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
	var resource SourceMicrosoftTeams
	err := ctx.RegisterResource("airbyte:index/sourceMicrosoftTeams:SourceMicrosoftTeams", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceMicrosoftTeams gets an existing SourceMicrosoftTeams resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceMicrosoftTeams(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceMicrosoftTeamsState, opts ...pulumi.ResourceOption) (*SourceMicrosoftTeams, error) {
	var resource SourceMicrosoftTeams
	err := ctx.ReadResource("airbyte:index/sourceMicrosoftTeams:SourceMicrosoftTeams", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceMicrosoftTeams resources.
type sourceMicrosoftTeamsState struct {
	Configuration *SourceMicrosoftTeamsConfiguration `pulumi:"configuration"`
	Name          *string                            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceMicrosoftTeamsState struct {
	Configuration SourceMicrosoftTeamsConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceMicrosoftTeamsState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMicrosoftTeamsState)(nil)).Elem()
}

type sourceMicrosoftTeamsArgs struct {
	Configuration SourceMicrosoftTeamsConfiguration `pulumi:"configuration"`
	Name          string                            `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceMicrosoftTeams resource.
type SourceMicrosoftTeamsArgs struct {
	Configuration SourceMicrosoftTeamsConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceMicrosoftTeamsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMicrosoftTeamsArgs)(nil)).Elem()
}

type SourceMicrosoftTeamsInput interface {
	pulumi.Input

	ToSourceMicrosoftTeamsOutput() SourceMicrosoftTeamsOutput
	ToSourceMicrosoftTeamsOutputWithContext(ctx context.Context) SourceMicrosoftTeamsOutput
}

func (*SourceMicrosoftTeams) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMicrosoftTeams)(nil)).Elem()
}

func (i *SourceMicrosoftTeams) ToSourceMicrosoftTeamsOutput() SourceMicrosoftTeamsOutput {
	return i.ToSourceMicrosoftTeamsOutputWithContext(context.Background())
}

func (i *SourceMicrosoftTeams) ToSourceMicrosoftTeamsOutputWithContext(ctx context.Context) SourceMicrosoftTeamsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMicrosoftTeamsOutput)
}

type SourceMicrosoftTeamsOutput struct{ *pulumi.OutputState }

func (SourceMicrosoftTeamsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMicrosoftTeams)(nil)).Elem()
}

func (o SourceMicrosoftTeamsOutput) ToSourceMicrosoftTeamsOutput() SourceMicrosoftTeamsOutput {
	return o
}

func (o SourceMicrosoftTeamsOutput) ToSourceMicrosoftTeamsOutputWithContext(ctx context.Context) SourceMicrosoftTeamsOutput {
	return o
}

func (o SourceMicrosoftTeamsOutput) Configuration() SourceMicrosoftTeamsConfigurationOutput {
	return o.ApplyT(func(v *SourceMicrosoftTeams) SourceMicrosoftTeamsConfigurationOutput { return v.Configuration }).(SourceMicrosoftTeamsConfigurationOutput)
}

func (o SourceMicrosoftTeamsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftTeams) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceMicrosoftTeamsOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceMicrosoftTeams) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceMicrosoftTeamsOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftTeams) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceMicrosoftTeamsOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftTeams) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceMicrosoftTeamsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftTeams) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMicrosoftTeamsInput)(nil)).Elem(), &SourceMicrosoftTeams{})
	pulumi.RegisterOutputType(SourceMicrosoftTeamsOutput{})
}
