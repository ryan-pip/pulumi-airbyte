// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DestinationBigquery struct {
	pulumi.CustomResourceState

	Configuration   DestinationBigqueryConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                    `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                    `pulumi:"destinationType"`
	Name            pulumi.StringOutput                    `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                    `pulumi:"workspaceId"`
}

// NewDestinationBigquery registers a new resource with the given unique name, arguments, and options.
func NewDestinationBigquery(ctx *pulumi.Context,
	name string, args *DestinationBigqueryArgs, opts ...pulumi.ResourceOption) (*DestinationBigquery, error) {
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
	var resource DestinationBigquery
	err := ctx.RegisterResource("airbyte:index/destinationBigquery:DestinationBigquery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationBigquery gets an existing DestinationBigquery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationBigquery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationBigqueryState, opts ...pulumi.ResourceOption) (*DestinationBigquery, error) {
	var resource DestinationBigquery
	err := ctx.ReadResource("airbyte:index/destinationBigquery:DestinationBigquery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationBigquery resources.
type destinationBigqueryState struct {
	Configuration   *DestinationBigqueryConfiguration `pulumi:"configuration"`
	DestinationId   *string                           `pulumi:"destinationId"`
	DestinationType *string                           `pulumi:"destinationType"`
	Name            *string                           `pulumi:"name"`
	WorkspaceId     *string                           `pulumi:"workspaceId"`
}

type DestinationBigqueryState struct {
	Configuration   DestinationBigqueryConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationBigqueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationBigqueryState)(nil)).Elem()
}

type destinationBigqueryArgs struct {
	Configuration DestinationBigqueryConfiguration `pulumi:"configuration"`
	Name          string                           `pulumi:"name"`
	WorkspaceId   string                           `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationBigquery resource.
type DestinationBigqueryArgs struct {
	Configuration DestinationBigqueryConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationBigqueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationBigqueryArgs)(nil)).Elem()
}

type DestinationBigqueryInput interface {
	pulumi.Input

	ToDestinationBigqueryOutput() DestinationBigqueryOutput
	ToDestinationBigqueryOutputWithContext(ctx context.Context) DestinationBigqueryOutput
}

func (*DestinationBigquery) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationBigquery)(nil)).Elem()
}

func (i *DestinationBigquery) ToDestinationBigqueryOutput() DestinationBigqueryOutput {
	return i.ToDestinationBigqueryOutputWithContext(context.Background())
}

func (i *DestinationBigquery) ToDestinationBigqueryOutputWithContext(ctx context.Context) DestinationBigqueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationBigqueryOutput)
}

type DestinationBigqueryOutput struct{ *pulumi.OutputState }

func (DestinationBigqueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationBigquery)(nil)).Elem()
}

func (o DestinationBigqueryOutput) ToDestinationBigqueryOutput() DestinationBigqueryOutput {
	return o
}

func (o DestinationBigqueryOutput) ToDestinationBigqueryOutputWithContext(ctx context.Context) DestinationBigqueryOutput {
	return o
}

func (o DestinationBigqueryOutput) Configuration() DestinationBigqueryConfigurationOutput {
	return o.ApplyT(func(v *DestinationBigquery) DestinationBigqueryConfigurationOutput { return v.Configuration }).(DestinationBigqueryConfigurationOutput)
}

func (o DestinationBigqueryOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationBigquery) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationBigqueryOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationBigquery) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationBigqueryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationBigquery) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationBigqueryOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationBigquery) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationBigqueryInput)(nil)).Elem(), &DestinationBigquery{})
	pulumi.RegisterOutputType(DestinationBigqueryOutput{})
}