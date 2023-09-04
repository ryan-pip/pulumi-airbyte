// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DestinationGcs struct {
	pulumi.CustomResourceState

	Configuration   DestinationGcsConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput               `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput               `pulumi:"destinationType"`
	Name            pulumi.StringOutput               `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput               `pulumi:"workspaceId"`
}

// NewDestinationGcs registers a new resource with the given unique name, arguments, and options.
func NewDestinationGcs(ctx *pulumi.Context,
	name string, args *DestinationGcsArgs, opts ...pulumi.ResourceOption) (*DestinationGcs, error) {
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
	var resource DestinationGcs
	err := ctx.RegisterResource("airbyte:index/destinationGcs:DestinationGcs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationGcs gets an existing DestinationGcs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationGcs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationGcsState, opts ...pulumi.ResourceOption) (*DestinationGcs, error) {
	var resource DestinationGcs
	err := ctx.ReadResource("airbyte:index/destinationGcs:DestinationGcs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationGcs resources.
type destinationGcsState struct {
	Configuration   *DestinationGcsConfiguration `pulumi:"configuration"`
	DestinationId   *string                      `pulumi:"destinationId"`
	DestinationType *string                      `pulumi:"destinationType"`
	Name            *string                      `pulumi:"name"`
	WorkspaceId     *string                      `pulumi:"workspaceId"`
}

type DestinationGcsState struct {
	Configuration   DestinationGcsConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationGcsState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationGcsState)(nil)).Elem()
}

type destinationGcsArgs struct {
	Configuration DestinationGcsConfiguration `pulumi:"configuration"`
	Name          string                      `pulumi:"name"`
	WorkspaceId   string                      `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationGcs resource.
type DestinationGcsArgs struct {
	Configuration DestinationGcsConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationGcsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationGcsArgs)(nil)).Elem()
}

type DestinationGcsInput interface {
	pulumi.Input

	ToDestinationGcsOutput() DestinationGcsOutput
	ToDestinationGcsOutputWithContext(ctx context.Context) DestinationGcsOutput
}

func (*DestinationGcs) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationGcs)(nil)).Elem()
}

func (i *DestinationGcs) ToDestinationGcsOutput() DestinationGcsOutput {
	return i.ToDestinationGcsOutputWithContext(context.Background())
}

func (i *DestinationGcs) ToDestinationGcsOutputWithContext(ctx context.Context) DestinationGcsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationGcsOutput)
}

type DestinationGcsOutput struct{ *pulumi.OutputState }

func (DestinationGcsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationGcs)(nil)).Elem()
}

func (o DestinationGcsOutput) ToDestinationGcsOutput() DestinationGcsOutput {
	return o
}

func (o DestinationGcsOutput) ToDestinationGcsOutputWithContext(ctx context.Context) DestinationGcsOutput {
	return o
}

func (o DestinationGcsOutput) Configuration() DestinationGcsConfigurationOutput {
	return o.ApplyT(func(v *DestinationGcs) DestinationGcsConfigurationOutput { return v.Configuration }).(DestinationGcsConfigurationOutput)
}

func (o DestinationGcsOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationGcs) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationGcsOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationGcs) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationGcsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationGcs) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationGcsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationGcs) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationGcsInput)(nil)).Elem(), &DestinationGcs{})
	pulumi.RegisterOutputType(DestinationGcsOutput{})
}
