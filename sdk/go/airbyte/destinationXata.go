// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DestinationXata struct {
	pulumi.CustomResourceState

	Configuration   DestinationXataConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                `pulumi:"destinationType"`
	Name            pulumi.StringOutput                `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                `pulumi:"workspaceId"`
}

// NewDestinationXata registers a new resource with the given unique name, arguments, and options.
func NewDestinationXata(ctx *pulumi.Context,
	name string, args *DestinationXataArgs, opts ...pulumi.ResourceOption) (*DestinationXata, error) {
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
	var resource DestinationXata
	err := ctx.RegisterResource("airbyte:index/destinationXata:DestinationXata", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationXata gets an existing DestinationXata resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationXata(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationXataState, opts ...pulumi.ResourceOption) (*DestinationXata, error) {
	var resource DestinationXata
	err := ctx.ReadResource("airbyte:index/destinationXata:DestinationXata", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationXata resources.
type destinationXataState struct {
	Configuration   *DestinationXataConfiguration `pulumi:"configuration"`
	DestinationId   *string                       `pulumi:"destinationId"`
	DestinationType *string                       `pulumi:"destinationType"`
	Name            *string                       `pulumi:"name"`
	WorkspaceId     *string                       `pulumi:"workspaceId"`
}

type DestinationXataState struct {
	Configuration   DestinationXataConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationXataState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationXataState)(nil)).Elem()
}

type destinationXataArgs struct {
	Configuration DestinationXataConfiguration `pulumi:"configuration"`
	Name          string                       `pulumi:"name"`
	WorkspaceId   string                       `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationXata resource.
type DestinationXataArgs struct {
	Configuration DestinationXataConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationXataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationXataArgs)(nil)).Elem()
}

type DestinationXataInput interface {
	pulumi.Input

	ToDestinationXataOutput() DestinationXataOutput
	ToDestinationXataOutputWithContext(ctx context.Context) DestinationXataOutput
}

func (*DestinationXata) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationXata)(nil)).Elem()
}

func (i *DestinationXata) ToDestinationXataOutput() DestinationXataOutput {
	return i.ToDestinationXataOutputWithContext(context.Background())
}

func (i *DestinationXata) ToDestinationXataOutputWithContext(ctx context.Context) DestinationXataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationXataOutput)
}

type DestinationXataOutput struct{ *pulumi.OutputState }

func (DestinationXataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationXata)(nil)).Elem()
}

func (o DestinationXataOutput) ToDestinationXataOutput() DestinationXataOutput {
	return o
}

func (o DestinationXataOutput) ToDestinationXataOutputWithContext(ctx context.Context) DestinationXataOutput {
	return o
}

func (o DestinationXataOutput) Configuration() DestinationXataConfigurationOutput {
	return o.ApplyT(func(v *DestinationXata) DestinationXataConfigurationOutput { return v.Configuration }).(DestinationXataConfigurationOutput)
}

func (o DestinationXataOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationXata) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationXataOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationXata) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationXataOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationXata) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationXataOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationXata) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationXataInput)(nil)).Elem(), &DestinationXata{})
	pulumi.RegisterOutputType(DestinationXataOutput{})
}