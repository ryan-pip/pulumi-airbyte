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

// DestinationRedshift Resource
type DestinationRedshift struct {
	pulumi.CustomResourceState

	Configuration   DestinationRedshiftConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                    `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                    `pulumi:"destinationType"`
	Name            pulumi.StringOutput                    `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                    `pulumi:"workspaceId"`
}

// NewDestinationRedshift registers a new resource with the given unique name, arguments, and options.
func NewDestinationRedshift(ctx *pulumi.Context,
	name string, args *DestinationRedshiftArgs, opts ...pulumi.ResourceOption) (*DestinationRedshift, error) {
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
	var resource DestinationRedshift
	err := ctx.RegisterResource("airbyte:index/destinationRedshift:DestinationRedshift", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationRedshift gets an existing DestinationRedshift resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationRedshift(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationRedshiftState, opts ...pulumi.ResourceOption) (*DestinationRedshift, error) {
	var resource DestinationRedshift
	err := ctx.ReadResource("airbyte:index/destinationRedshift:DestinationRedshift", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationRedshift resources.
type destinationRedshiftState struct {
	Configuration   *DestinationRedshiftConfiguration `pulumi:"configuration"`
	DestinationId   *string                           `pulumi:"destinationId"`
	DestinationType *string                           `pulumi:"destinationType"`
	Name            *string                           `pulumi:"name"`
	WorkspaceId     *string                           `pulumi:"workspaceId"`
}

type DestinationRedshiftState struct {
	Configuration   DestinationRedshiftConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationRedshiftState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationRedshiftState)(nil)).Elem()
}

type destinationRedshiftArgs struct {
	Configuration DestinationRedshiftConfiguration `pulumi:"configuration"`
	Name          string                           `pulumi:"name"`
	WorkspaceId   string                           `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationRedshift resource.
type DestinationRedshiftArgs struct {
	Configuration DestinationRedshiftConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationRedshiftArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationRedshiftArgs)(nil)).Elem()
}

type DestinationRedshiftInput interface {
	pulumi.Input

	ToDestinationRedshiftOutput() DestinationRedshiftOutput
	ToDestinationRedshiftOutputWithContext(ctx context.Context) DestinationRedshiftOutput
}

func (*DestinationRedshift) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationRedshift)(nil)).Elem()
}

func (i *DestinationRedshift) ToDestinationRedshiftOutput() DestinationRedshiftOutput {
	return i.ToDestinationRedshiftOutputWithContext(context.Background())
}

func (i *DestinationRedshift) ToDestinationRedshiftOutputWithContext(ctx context.Context) DestinationRedshiftOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationRedshiftOutput)
}

type DestinationRedshiftOutput struct{ *pulumi.OutputState }

func (DestinationRedshiftOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationRedshift)(nil)).Elem()
}

func (o DestinationRedshiftOutput) ToDestinationRedshiftOutput() DestinationRedshiftOutput {
	return o
}

func (o DestinationRedshiftOutput) ToDestinationRedshiftOutputWithContext(ctx context.Context) DestinationRedshiftOutput {
	return o
}

func (o DestinationRedshiftOutput) Configuration() DestinationRedshiftConfigurationOutput {
	return o.ApplyT(func(v *DestinationRedshift) DestinationRedshiftConfigurationOutput { return v.Configuration }).(DestinationRedshiftConfigurationOutput)
}

func (o DestinationRedshiftOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationRedshift) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationRedshiftOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationRedshift) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationRedshiftOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationRedshift) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationRedshiftOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationRedshift) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationRedshiftInput)(nil)).Elem(), &DestinationRedshift{})
	pulumi.RegisterOutputType(DestinationRedshiftOutput{})
}
