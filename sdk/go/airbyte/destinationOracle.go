// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DestinationOracle struct {
	pulumi.CustomResourceState

	Configuration   DestinationOracleConfigurationOutput `pulumi:"configuration"`
	DestinationId   pulumi.StringOutput                  `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput                  `pulumi:"destinationType"`
	Name            pulumi.StringOutput                  `pulumi:"name"`
	WorkspaceId     pulumi.StringOutput                  `pulumi:"workspaceId"`
}

// NewDestinationOracle registers a new resource with the given unique name, arguments, and options.
func NewDestinationOracle(ctx *pulumi.Context,
	name string, args *DestinationOracleArgs, opts ...pulumi.ResourceOption) (*DestinationOracle, error) {
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
	var resource DestinationOracle
	err := ctx.RegisterResource("airbyte:index/destinationOracle:DestinationOracle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationOracle gets an existing DestinationOracle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationOracle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationOracleState, opts ...pulumi.ResourceOption) (*DestinationOracle, error) {
	var resource DestinationOracle
	err := ctx.ReadResource("airbyte:index/destinationOracle:DestinationOracle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationOracle resources.
type destinationOracleState struct {
	Configuration   *DestinationOracleConfiguration `pulumi:"configuration"`
	DestinationId   *string                         `pulumi:"destinationId"`
	DestinationType *string                         `pulumi:"destinationType"`
	Name            *string                         `pulumi:"name"`
	WorkspaceId     *string                         `pulumi:"workspaceId"`
}

type DestinationOracleState struct {
	Configuration   DestinationOracleConfigurationPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	WorkspaceId     pulumi.StringPtrInput
}

func (DestinationOracleState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationOracleState)(nil)).Elem()
}

type destinationOracleArgs struct {
	Configuration DestinationOracleConfiguration `pulumi:"configuration"`
	Name          string                         `pulumi:"name"`
	WorkspaceId   string                         `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationOracle resource.
type DestinationOracleArgs struct {
	Configuration DestinationOracleConfigurationInput
	Name          pulumi.StringInput
	WorkspaceId   pulumi.StringInput
}

func (DestinationOracleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationOracleArgs)(nil)).Elem()
}

type DestinationOracleInput interface {
	pulumi.Input

	ToDestinationOracleOutput() DestinationOracleOutput
	ToDestinationOracleOutputWithContext(ctx context.Context) DestinationOracleOutput
}

func (*DestinationOracle) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationOracle)(nil)).Elem()
}

func (i *DestinationOracle) ToDestinationOracleOutput() DestinationOracleOutput {
	return i.ToDestinationOracleOutputWithContext(context.Background())
}

func (i *DestinationOracle) ToDestinationOracleOutputWithContext(ctx context.Context) DestinationOracleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationOracleOutput)
}

type DestinationOracleOutput struct{ *pulumi.OutputState }

func (DestinationOracleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationOracle)(nil)).Elem()
}

func (o DestinationOracleOutput) ToDestinationOracleOutput() DestinationOracleOutput {
	return o
}

func (o DestinationOracleOutput) ToDestinationOracleOutputWithContext(ctx context.Context) DestinationOracleOutput {
	return o
}

func (o DestinationOracleOutput) Configuration() DestinationOracleConfigurationOutput {
	return o.ApplyT(func(v *DestinationOracle) DestinationOracleConfigurationOutput { return v.Configuration }).(DestinationOracleConfigurationOutput)
}

func (o DestinationOracleOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationOracle) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationOracleOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationOracle) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o DestinationOracleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationOracle) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationOracleOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationOracle) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationOracleInput)(nil)).Elem(), &DestinationOracle{})
	pulumi.RegisterOutputType(DestinationOracleOutput{})
}