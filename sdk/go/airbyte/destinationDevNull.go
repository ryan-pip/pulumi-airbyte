// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DestinationDevNull struct {
	pulumi.CustomResourceState

	Configuration   DestinationDevNullConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                   `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                   `pulumi:"destinationType"`
	Name            pulumi.StringOutput                   `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                   `pulumi:"workspaceId"`
}

// NewDestinationDevNull registers a new resource with the given unique name, arguments, and options.
func NewDestinationDevNull(ctx *pulumi.Context,
	name string, args *DestinationDevNullArgs, opts ...pulumi.ResourceOption) (*DestinationDevNull, error) {
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
	var resource DestinationDevNull
	err := ctx.RegisterResource("airbyte:index/destinationDevNull:DestinationDevNull", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationDevNull gets an existing DestinationDevNull resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationDevNull(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationDevNullState, opts ...pulumi.ResourceOption) (*DestinationDevNull, error) {
	var resource DestinationDevNull
	err := ctx.ReadResource("airbyte:index/destinationDevNull:DestinationDevNull", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationDevNull resources.
type destinationDevNullState struct {
	Configuration   *DestinationDevNullConfiguration `pulumi:"configuration"`
	DestinationId   *string                          `pulumi:"destinationId"`
	DestinationType *string                          `pulumi:"destinationType"`
	Name            *string                          `pulumi:"name"`
	WorkspaceId     *string                          `pulumi:"workspaceId"`
}

type DestinationDevNullState struct {
	Configuration   DestinationDevNullConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationDevNullState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationDevNullState)(nil)).Elem()
}

type destinationDevNullArgs struct {
	Configuration DestinationDevNullConfiguration `pulumi:"configuration"`
	Name          string                          `pulumi:"name"`
	WorkspaceId   string                          `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationDevNull resource.
type DestinationDevNullArgs struct {
	Configuration DestinationDevNullConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationDevNullArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationDevNullArgs)(nil)).Elem()
}

type DestinationDevNullInput interface {
	pulumi.Input

	ToDestinationDevNullOutput() DestinationDevNullOutput
	ToDestinationDevNullOutputWithContext(ctx context.Context) DestinationDevNullOutput
}

func (*DestinationDevNull) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationDevNull)(nil)).Elem()
}

func (i *DestinationDevNull) ToDestinationDevNullOutput() DestinationDevNullOutput {
	return i.ToDestinationDevNullOutputWithContext(context.Background())
}

func (i *DestinationDevNull) ToDestinationDevNullOutputWithContext(ctx context.Context) DestinationDevNullOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationDevNullOutput)
}

type DestinationDevNullOutput struct{ *pulumi.OutputState }

func (DestinationDevNullOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationDevNull)(nil)).Elem()
}

func (o DestinationDevNullOutput) ToDestinationDevNullOutput() DestinationDevNullOutput {
	return o
}

func (o DestinationDevNullOutput) ToDestinationDevNullOutputWithContext(ctx context.Context) DestinationDevNullOutput {
	return o
}

func (o DestinationDevNullOutput) Configuration() DestinationDevNullConfigurationOutput {
	return o.ApplyT(func(v *DestinationDevNull) DestinationDevNullConfigurationOutput { return v.Configuration }).(DestinationDevNullConfigurationOutput)
}

func (o DestinationDevNullOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationDevNull) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationDevNullOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationDevNull) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationDevNullOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationDevNull) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationDevNullOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationDevNull) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationDevNullInput)(nil)).Elem(), &DestinationDevNull{})
	pulumi.RegisterOutputType(DestinationDevNullOutput{})
}
