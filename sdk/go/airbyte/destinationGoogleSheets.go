// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DestinationGoogleSheets struct {
	pulumi.CustomResourceState

	Configuration   DestinationGoogleSheetsConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                        `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                        `pulumi:"destinationType"`
	Name            pulumi.StringOutput                        `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                        `pulumi:"workspaceId"`
}

// NewDestinationGoogleSheets registers a new resource with the given unique name, arguments, and options.
func NewDestinationGoogleSheets(ctx *pulumi.Context,
	name string, args *DestinationGoogleSheetsArgs, opts ...pulumi.ResourceOption) (*DestinationGoogleSheets, error) {
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
	var resource DestinationGoogleSheets
	err := ctx.RegisterResource("airbyte:index/destinationGoogleSheets:DestinationGoogleSheets", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationGoogleSheets gets an existing DestinationGoogleSheets resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationGoogleSheets(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationGoogleSheetsState, opts ...pulumi.ResourceOption) (*DestinationGoogleSheets, error) {
	var resource DestinationGoogleSheets
	err := ctx.ReadResource("airbyte:index/destinationGoogleSheets:DestinationGoogleSheets", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationGoogleSheets resources.
type destinationGoogleSheetsState struct {
	Configuration   *DestinationGoogleSheetsConfiguration `pulumi:"configuration"`
	DestinationId   *string                               `pulumi:"destinationId"`
	DestinationType *string                               `pulumi:"destinationType"`
	Name            *string                               `pulumi:"name"`
	WorkspaceId     *string                               `pulumi:"workspaceId"`
}

type DestinationGoogleSheetsState struct {
	Configuration   DestinationGoogleSheetsConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationGoogleSheetsState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationGoogleSheetsState)(nil)).Elem()
}

type destinationGoogleSheetsArgs struct {
	Configuration DestinationGoogleSheetsConfiguration `pulumi:"configuration"`
	Name          string                               `pulumi:"name"`
	WorkspaceId   string                               `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationGoogleSheets resource.
type DestinationGoogleSheetsArgs struct {
	Configuration DestinationGoogleSheetsConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationGoogleSheetsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationGoogleSheetsArgs)(nil)).Elem()
}

type DestinationGoogleSheetsInput interface {
	pulumi.Input

	ToDestinationGoogleSheetsOutput() DestinationGoogleSheetsOutput
	ToDestinationGoogleSheetsOutputWithContext(ctx context.Context) DestinationGoogleSheetsOutput
}

func (*DestinationGoogleSheets) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationGoogleSheets)(nil)).Elem()
}

func (i *DestinationGoogleSheets) ToDestinationGoogleSheetsOutput() DestinationGoogleSheetsOutput {
	return i.ToDestinationGoogleSheetsOutputWithContext(context.Background())
}

func (i *DestinationGoogleSheets) ToDestinationGoogleSheetsOutputWithContext(ctx context.Context) DestinationGoogleSheetsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationGoogleSheetsOutput)
}

type DestinationGoogleSheetsOutput struct{ *pulumi.OutputState }

func (DestinationGoogleSheetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationGoogleSheets)(nil)).Elem()
}

func (o DestinationGoogleSheetsOutput) ToDestinationGoogleSheetsOutput() DestinationGoogleSheetsOutput {
	return o
}

func (o DestinationGoogleSheetsOutput) ToDestinationGoogleSheetsOutputWithContext(ctx context.Context) DestinationGoogleSheetsOutput {
	return o
}

func (o DestinationGoogleSheetsOutput) Configuration() DestinationGoogleSheetsConfigurationOutput {
	return o.ApplyT(func(v *DestinationGoogleSheets) DestinationGoogleSheetsConfigurationOutput { return v.Configuration }).(DestinationGoogleSheetsConfigurationOutput)
}

func (o DestinationGoogleSheetsOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationGoogleSheets) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationGoogleSheetsOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationGoogleSheets) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationGoogleSheetsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationGoogleSheets) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationGoogleSheetsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationGoogleSheets) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationGoogleSheetsInput)(nil)).Elem(), &DestinationGoogleSheets{})
	pulumi.RegisterOutputType(DestinationGoogleSheetsOutput{})
}
