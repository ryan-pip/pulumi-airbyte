// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourcePokeapi struct {
	pulumi.CustomResourceState

	Configuration SourcePokeapiConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourcePokeapi registers a new resource with the given unique name, arguments, and options.
func NewSourcePokeapi(ctx *pulumi.Context,
	name string, args *SourcePokeapiArgs, opts ...pulumi.ResourceOption) (*SourcePokeapi, error) {
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
	var resource SourcePokeapi
	err := ctx.RegisterResource("airbyte:index/sourcePokeapi:SourcePokeapi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourcePokeapi gets an existing SourcePokeapi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourcePokeapi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourcePokeapiState, opts ...pulumi.ResourceOption) (*SourcePokeapi, error) {
	var resource SourcePokeapi
	err := ctx.ReadResource("airbyte:index/sourcePokeapi:SourcePokeapi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourcePokeapi resources.
type sourcePokeapiState struct {
	Configuration *SourcePokeapiConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourcePokeapiState struct {
	Configuration SourcePokeapiConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourcePokeapiState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePokeapiState)(nil)).Elem()
}

type sourcePokeapiArgs struct {
	Configuration SourcePokeapiConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourcePokeapi resource.
type SourcePokeapiArgs struct {
	Configuration SourcePokeapiConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourcePokeapiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePokeapiArgs)(nil)).Elem()
}

type SourcePokeapiInput interface {
	pulumi.Input

	ToSourcePokeapiOutput() SourcePokeapiOutput
	ToSourcePokeapiOutputWithContext(ctx context.Context) SourcePokeapiOutput
}

func (*SourcePokeapi) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePokeapi)(nil)).Elem()
}

func (i *SourcePokeapi) ToSourcePokeapiOutput() SourcePokeapiOutput {
	return i.ToSourcePokeapiOutputWithContext(context.Background())
}

func (i *SourcePokeapi) ToSourcePokeapiOutputWithContext(ctx context.Context) SourcePokeapiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePokeapiOutput)
}

type SourcePokeapiOutput struct{ *pulumi.OutputState }

func (SourcePokeapiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePokeapi)(nil)).Elem()
}

func (o SourcePokeapiOutput) ToSourcePokeapiOutput() SourcePokeapiOutput {
	return o
}

func (o SourcePokeapiOutput) ToSourcePokeapiOutputWithContext(ctx context.Context) SourcePokeapiOutput {
	return o
}

func (o SourcePokeapiOutput) Configuration() SourcePokeapiConfigurationOutput {
	return o.ApplyT(func(v *SourcePokeapi) SourcePokeapiConfigurationOutput { return v.Configuration }).(SourcePokeapiConfigurationOutput)
}

func (o SourcePokeapiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePokeapi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourcePokeapiOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePokeapi) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourcePokeapiOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePokeapi) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourcePokeapiOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePokeapi) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourcePokeapiOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePokeapi) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePokeapiInput)(nil)).Elem(), &SourcePokeapi{})
	pulumi.RegisterOutputType(SourcePokeapiOutput{})
}
