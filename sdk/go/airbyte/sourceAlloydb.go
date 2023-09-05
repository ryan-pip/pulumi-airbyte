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

// SourceAlloydb Resource
type SourceAlloydb struct {
	pulumi.CustomResourceState

	Configuration SourceAlloydbConfigurationOutput `pulumi:"configuration"`
	Name          pulumi.StringOutput              `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceAlloydb registers a new resource with the given unique name, arguments, and options.
func NewSourceAlloydb(ctx *pulumi.Context,
	name string, args *SourceAlloydbArgs, opts ...pulumi.ResourceOption) (*SourceAlloydb, error) {
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
	var resource SourceAlloydb
	err := ctx.RegisterResource("airbyte:index/sourceAlloydb:SourceAlloydb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceAlloydb gets an existing SourceAlloydb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceAlloydb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceAlloydbState, opts ...pulumi.ResourceOption) (*SourceAlloydb, error) {
	var resource SourceAlloydb
	err := ctx.ReadResource("airbyte:index/sourceAlloydb:SourceAlloydb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceAlloydb resources.
type sourceAlloydbState struct {
	Configuration *SourceAlloydbConfiguration `pulumi:"configuration"`
	Name          *string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceAlloydbState struct {
	Configuration SourceAlloydbConfigurationPtrInput
	Name          pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceAlloydbState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAlloydbState)(nil)).Elem()
}

type sourceAlloydbArgs struct {
	Configuration SourceAlloydbConfiguration `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceAlloydb resource.
type SourceAlloydbArgs struct {
	Configuration SourceAlloydbConfigurationInput
	Name          pulumi.StringInput
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceAlloydbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAlloydbArgs)(nil)).Elem()
}

type SourceAlloydbInput interface {
	pulumi.Input

	ToSourceAlloydbOutput() SourceAlloydbOutput
	ToSourceAlloydbOutputWithContext(ctx context.Context) SourceAlloydbOutput
}

func (*SourceAlloydb) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAlloydb)(nil)).Elem()
}

func (i *SourceAlloydb) ToSourceAlloydbOutput() SourceAlloydbOutput {
	return i.ToSourceAlloydbOutputWithContext(context.Background())
}

func (i *SourceAlloydb) ToSourceAlloydbOutputWithContext(ctx context.Context) SourceAlloydbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAlloydbOutput)
}

type SourceAlloydbOutput struct{ *pulumi.OutputState }

func (SourceAlloydbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAlloydb)(nil)).Elem()
}

func (o SourceAlloydbOutput) ToSourceAlloydbOutput() SourceAlloydbOutput {
	return o
}

func (o SourceAlloydbOutput) ToSourceAlloydbOutputWithContext(ctx context.Context) SourceAlloydbOutput {
	return o
}

func (o SourceAlloydbOutput) Configuration() SourceAlloydbConfigurationOutput {
	return o.ApplyT(func(v *SourceAlloydb) SourceAlloydbConfigurationOutput { return v.Configuration }).(SourceAlloydbConfigurationOutput)
}

func (o SourceAlloydbOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAlloydb) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow.
func (o SourceAlloydbOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAlloydb) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceAlloydbOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAlloydb) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceAlloydbOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAlloydb) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceAlloydbOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAlloydb) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAlloydbInput)(nil)).Elem(), &SourceAlloydb{})
	pulumi.RegisterOutputType(SourceAlloydbOutput{})
}
