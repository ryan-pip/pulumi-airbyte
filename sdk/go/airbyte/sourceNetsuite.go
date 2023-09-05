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

// SourceNetsuite Resource
type SourceNetsuite struct {
	pulumi.CustomResourceState

	Configuration SourceNetsuiteConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput               `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceNetsuite registers a new resource with the given unique name, arguments, and options.
func NewSourceNetsuite(ctx *pulumi.Context,
	name string, args *SourceNetsuiteArgs, opts ...pulumi.ResourceOption) (*SourceNetsuite, error) {
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
	var resource SourceNetsuite
	err := ctx.RegisterResource("airbyte:index/sourceNetsuite:SourceNetsuite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceNetsuite gets an existing SourceNetsuite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceNetsuite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceNetsuiteState, opts ...pulumi.ResourceOption) (*SourceNetsuite, error) {
	var resource SourceNetsuite
	err := ctx.ReadResource("airbyte:index/sourceNetsuite:SourceNetsuite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceNetsuite resources.
type sourceNetsuiteState struct {
	Configuration *SourceNetsuiteConfiguration `pulumi:"configuration"`
	Name          *string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceNetsuiteState struct {
	Configuration SourceNetsuiteConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceNetsuiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceNetsuiteState)(nil)).Elem()
}

type sourceNetsuiteArgs struct {
	Configuration SourceNetsuiteConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceNetsuite resource.
type SourceNetsuiteArgs struct {
	Configuration SourceNetsuiteConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceNetsuiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceNetsuiteArgs)(nil)).Elem()
}

type SourceNetsuiteInput interface {
	pulumi.Input

	ToSourceNetsuiteOutput() SourceNetsuiteOutput
	ToSourceNetsuiteOutputWithContext(ctx context.Context) SourceNetsuiteOutput
}

func (*SourceNetsuite) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceNetsuite)(nil)).Elem()
}

func (i *SourceNetsuite) ToSourceNetsuiteOutput() SourceNetsuiteOutput {
	return i.ToSourceNetsuiteOutputWithContext(context.Background())
}

func (i *SourceNetsuite) ToSourceNetsuiteOutputWithContext(ctx context.Context) SourceNetsuiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceNetsuiteOutput)
}

type SourceNetsuiteOutput struct{ *pulumi.OutputState }

func (SourceNetsuiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceNetsuite)(nil)).Elem()
}

func (o SourceNetsuiteOutput) ToSourceNetsuiteOutput() SourceNetsuiteOutput {
	return o
}

func (o SourceNetsuiteOutput) ToSourceNetsuiteOutputWithContext(ctx context.Context) SourceNetsuiteOutput {
	return o
}

func (o SourceNetsuiteOutput) Configuration() SourceNetsuiteConfigurationOutput {
	return o.ApplyT(func(v *SourceNetsuite) SourceNetsuiteConfigurationOutput { return v.Configuration }).(SourceNetsuiteConfigurationOutput)
}

func (o SourceNetsuiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceNetsuiteOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceNetsuiteOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceNetsuiteOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceNetsuiteOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNetsuite) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceNetsuiteInput)(nil)).Elem(), &SourceNetsuite{})
	pulumi.RegisterOutputType(SourceNetsuiteOutput{})
}
