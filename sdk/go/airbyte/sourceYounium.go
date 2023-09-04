// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceYounium struct {
	pulumi.CustomResourceState

	Configuration SourceYouniumConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceYounium registers a new resource with the given unique name, arguments, and options.
func NewSourceYounium(ctx *pulumi.Context,
	name string, args *SourceYouniumArgs, opts ...pulumi.ResourceOption) (*SourceYounium, error) {
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
	var resource SourceYounium
	err := ctx.RegisterResource("airbyte:index/sourceYounium:SourceYounium", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceYounium gets an existing SourceYounium resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceYounium(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceYouniumState, opts ...pulumi.ResourceOption) (*SourceYounium, error) {
	var resource SourceYounium
	err := ctx.ReadResource("airbyte:index/sourceYounium:SourceYounium", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceYounium resources.
type sourceYouniumState struct {
	Configuration *SourceYouniumConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceYouniumState struct {
	Configuration SourceYouniumConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceYouniumState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceYouniumState)(nil)).Elem()
}

type sourceYouniumArgs struct {
	Configuration SourceYouniumConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceYounium resource.
type SourceYouniumArgs struct {
	Configuration SourceYouniumConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceYouniumArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceYouniumArgs)(nil)).Elem()
}

type SourceYouniumInput interface {
	pulumi.Input

	ToSourceYouniumOutput() SourceYouniumOutput
	ToSourceYouniumOutputWithContext(ctx context.Context) SourceYouniumOutput
}

func (*SourceYounium) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceYounium)(nil)).Elem()
}

func (i *SourceYounium) ToSourceYouniumOutput() SourceYouniumOutput {
	return i.ToSourceYouniumOutputWithContext(context.Background())
}

func (i *SourceYounium) ToSourceYouniumOutputWithContext(ctx context.Context) SourceYouniumOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceYouniumOutput)
}

type SourceYouniumOutput struct{ *pulumi.OutputState }

func (SourceYouniumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceYounium)(nil)).Elem()
}

func (o SourceYouniumOutput) ToSourceYouniumOutput() SourceYouniumOutput {
	return o
}

func (o SourceYouniumOutput) ToSourceYouniumOutputWithContext(ctx context.Context) SourceYouniumOutput {
	return o
}

func (o SourceYouniumOutput) Configuration() SourceYouniumConfigurationOutput {
	return o.ApplyT(func(v *SourceYounium) SourceYouniumConfigurationOutput { return v.Configuration }).(SourceYouniumConfigurationOutput)
}

func (o SourceYouniumOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYounium) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceYouniumOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceYounium) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceYouniumOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYounium) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceYouniumOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYounium) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceYouniumOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceYounium) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceYouniumInput)(nil)).Elem(), &SourceYounium{})
	pulumi.RegisterOutputType(SourceYouniumOutput{})
}
