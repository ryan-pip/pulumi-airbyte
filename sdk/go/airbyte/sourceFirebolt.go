// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceFirebolt struct {
	pulumi.CustomResourceState

	Configuration SourceFireboltConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceFirebolt registers a new resource with the given unique name, arguments, and options.
func NewSourceFirebolt(ctx *pulumi.Context,
	name string, args *SourceFireboltArgs, opts ...pulumi.ResourceOption) (*SourceFirebolt, error) {
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
	var resource SourceFirebolt
	err := ctx.RegisterResource("airbyte:index/sourceFirebolt:SourceFirebolt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceFirebolt gets an existing SourceFirebolt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceFirebolt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceFireboltState, opts ...pulumi.ResourceOption) (*SourceFirebolt, error) {
	var resource SourceFirebolt
	err := ctx.ReadResource("airbyte:index/sourceFirebolt:SourceFirebolt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceFirebolt resources.
type sourceFireboltState struct {
	Configuration *SourceFireboltConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceFireboltState struct {
	Configuration SourceFireboltConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceFireboltState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFireboltState)(nil)).Elem()
}

type sourceFireboltArgs struct {
	Configuration SourceFireboltConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceFirebolt resource.
type SourceFireboltArgs struct {
	Configuration SourceFireboltConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceFireboltArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFireboltArgs)(nil)).Elem()
}

type SourceFireboltInput interface {
	pulumi.Input

	ToSourceFireboltOutput() SourceFireboltOutput
	ToSourceFireboltOutputWithContext(ctx context.Context) SourceFireboltOutput
}

func (*SourceFirebolt) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFirebolt)(nil)).Elem()
}

func (i *SourceFirebolt) ToSourceFireboltOutput() SourceFireboltOutput {
	return i.ToSourceFireboltOutputWithContext(context.Background())
}

func (i *SourceFirebolt) ToSourceFireboltOutputWithContext(ctx context.Context) SourceFireboltOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFireboltOutput)
}

type SourceFireboltOutput struct{ *pulumi.OutputState }

func (SourceFireboltOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFirebolt)(nil)).Elem()
}

func (o SourceFireboltOutput) ToSourceFireboltOutput() SourceFireboltOutput {
	return o
}

func (o SourceFireboltOutput) ToSourceFireboltOutputWithContext(ctx context.Context) SourceFireboltOutput {
	return o
}

func (o SourceFireboltOutput) Configuration() SourceFireboltConfigurationOutput {
	return o.ApplyT(func(v *SourceFirebolt) SourceFireboltConfigurationOutput { return v.Configuration }).(SourceFireboltConfigurationOutput)
}

func (o SourceFireboltOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFirebolt) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceFireboltOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceFirebolt) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceFireboltOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFirebolt) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceFireboltOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFirebolt) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceFireboltOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFirebolt) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFireboltInput)(nil)).Elem(), &SourceFirebolt{})
	pulumi.RegisterOutputType(SourceFireboltOutput{})
}
