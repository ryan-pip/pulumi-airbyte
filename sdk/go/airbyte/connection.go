// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-airbyte/sdk/go/airbyte/internal"
)

// Connection Resource
type Connection struct {
	pulumi.CustomResourceState

	// A list of configured stream options for a connection.
	Configurations ConnectionConfigurationsOutput `pulumi:"configurations"`
	ConnectionId   pulumi.StringOutput            `pulumi:"connectionId"`
	// must be one of ["auto", "us", "eu"]
	DataResidency pulumi.StringOutput `pulumi:"dataResidency"`
	DestinationId pulumi.StringOutput `pulumi:"destinationId"`
	// Optional name of the connection
	Name pulumi.StringOutput `pulumi:"name"`
	// must be one of ["source", "destination", "customFormat"]
	// Define the location where the data will be stored in the destination
	NamespaceDefinition pulumi.StringOutput `pulumi:"namespaceDefinition"`
	// Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat pulumi.StringOutput `pulumi:"namespaceFormat"`
	// must be one of ["ignore", "disable*connection", "propagate*columns", "propagateFully"]
	// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
	NonBreakingSchemaUpdatesBehavior pulumi.StringOutput `pulumi:"nonBreakingSchemaUpdatesBehavior"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
	Prefix pulumi.StringOutput `pulumi:"prefix"`
	// schedule for when the the connection should run, per the schedule type
	Schedule ConnectionScheduleOutput `pulumi:"schedule"`
	SourceId pulumi.StringOutput      `pulumi:"sourceId"`
	// must be one of ["active", "inactive", "deprecated"]
	Status      pulumi.StringOutput `pulumi:"status"`
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationId == nil {
		return nil, errors.New("invalid value for required argument 'DestinationId'")
	}
	if args.SourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Connection
	err := ctx.RegisterResource("airbyte:index/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("airbyte:index/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// A list of configured stream options for a connection.
	Configurations *ConnectionConfigurations `pulumi:"configurations"`
	ConnectionId   *string                   `pulumi:"connectionId"`
	// must be one of ["auto", "us", "eu"]
	DataResidency *string `pulumi:"dataResidency"`
	DestinationId *string `pulumi:"destinationId"`
	// Optional name of the connection
	Name *string `pulumi:"name"`
	// must be one of ["source", "destination", "customFormat"]
	// Define the location where the data will be stored in the destination
	NamespaceDefinition *string `pulumi:"namespaceDefinition"`
	// Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat *string `pulumi:"namespaceFormat"`
	// must be one of ["ignore", "disable*connection", "propagate*columns", "propagateFully"]
	// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
	NonBreakingSchemaUpdatesBehavior *string `pulumi:"nonBreakingSchemaUpdatesBehavior"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
	Prefix *string `pulumi:"prefix"`
	// schedule for when the the connection should run, per the schedule type
	Schedule *ConnectionSchedule `pulumi:"schedule"`
	SourceId *string             `pulumi:"sourceId"`
	// must be one of ["active", "inactive", "deprecated"]
	Status      *string `pulumi:"status"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type ConnectionState struct {
	// A list of configured stream options for a connection.
	Configurations ConnectionConfigurationsPtrInput
	ConnectionId   pulumi.StringPtrInput
	// must be one of ["auto", "us", "eu"]
	DataResidency pulumi.StringPtrInput
	DestinationId pulumi.StringPtrInput
	// Optional name of the connection
	Name pulumi.StringPtrInput
	// must be one of ["source", "destination", "customFormat"]
	// Define the location where the data will be stored in the destination
	NamespaceDefinition pulumi.StringPtrInput
	// Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat pulumi.StringPtrInput
	// must be one of ["ignore", "disable*connection", "propagate*columns", "propagateFully"]
	// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
	NonBreakingSchemaUpdatesBehavior pulumi.StringPtrInput
	// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
	Prefix pulumi.StringPtrInput
	// schedule for when the the connection should run, per the schedule type
	Schedule ConnectionSchedulePtrInput
	SourceId pulumi.StringPtrInput
	// must be one of ["active", "inactive", "deprecated"]
	Status      pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// A list of configured stream options for a connection.
	Configurations *ConnectionConfigurations `pulumi:"configurations"`
	// must be one of ["auto", "us", "eu"]
	DataResidency *string `pulumi:"dataResidency"`
	DestinationId string  `pulumi:"destinationId"`
	// Optional name of the connection
	Name *string `pulumi:"name"`
	// must be one of ["source", "destination", "customFormat"]
	// Define the location where the data will be stored in the destination
	NamespaceDefinition *string `pulumi:"namespaceDefinition"`
	// Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat *string `pulumi:"namespaceFormat"`
	// must be one of ["ignore", "disable*connection", "propagate*columns", "propagateFully"]
	// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
	NonBreakingSchemaUpdatesBehavior *string `pulumi:"nonBreakingSchemaUpdatesBehavior"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
	Prefix *string `pulumi:"prefix"`
	// schedule for when the the connection should run, per the schedule type
	Schedule *ConnectionSchedule `pulumi:"schedule"`
	SourceId string              `pulumi:"sourceId"`
	// must be one of ["active", "inactive", "deprecated"]
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// A list of configured stream options for a connection.
	Configurations ConnectionConfigurationsPtrInput
	// must be one of ["auto", "us", "eu"]
	DataResidency pulumi.StringPtrInput
	DestinationId pulumi.StringInput
	// Optional name of the connection
	Name pulumi.StringPtrInput
	// must be one of ["source", "destination", "customFormat"]
	// Define the location where the data will be stored in the destination
	NamespaceDefinition pulumi.StringPtrInput
	// Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat pulumi.StringPtrInput
	// must be one of ["ignore", "disable*connection", "propagate*columns", "propagateFully"]
	// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
	NonBreakingSchemaUpdatesBehavior pulumi.StringPtrInput
	// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
	Prefix pulumi.StringPtrInput
	// schedule for when the the connection should run, per the schedule type
	Schedule ConnectionSchedulePtrInput
	SourceId pulumi.StringInput
	// must be one of ["active", "inactive", "deprecated"]
	Status pulumi.StringPtrInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

// ConnectionArrayInput is an input type that accepts ConnectionArray and ConnectionArrayOutput values.
// You can construct a concrete instance of `ConnectionArrayInput` via:
//
//	ConnectionArray{ ConnectionArgs{...} }
type ConnectionArrayInput interface {
	pulumi.Input

	ToConnectionArrayOutput() ConnectionArrayOutput
	ToConnectionArrayOutputWithContext(context.Context) ConnectionArrayOutput
}

type ConnectionArray []ConnectionInput

func (ConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (i ConnectionArray) ToConnectionArrayOutput() ConnectionArrayOutput {
	return i.ToConnectionArrayOutputWithContext(context.Background())
}

func (i ConnectionArray) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionArrayOutput)
}

// ConnectionMapInput is an input type that accepts ConnectionMap and ConnectionMapOutput values.
// You can construct a concrete instance of `ConnectionMapInput` via:
//
//	ConnectionMap{ "key": ConnectionArgs{...} }
type ConnectionMapInput interface {
	pulumi.Input

	ToConnectionMapOutput() ConnectionMapOutput
	ToConnectionMapOutputWithContext(context.Context) ConnectionMapOutput
}

type ConnectionMap map[string]ConnectionInput

func (ConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (i ConnectionMap) ToConnectionMapOutput() ConnectionMapOutput {
	return i.ToConnectionMapOutputWithContext(context.Background())
}

func (i ConnectionMap) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMapOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

// A list of configured stream options for a connection.
func (o ConnectionOutput) Configurations() ConnectionConfigurationsOutput {
	return o.ApplyT(func(v *Connection) ConnectionConfigurationsOutput { return v.Configurations }).(ConnectionConfigurationsOutput)
}

func (o ConnectionOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ConnectionId }).(pulumi.StringOutput)
}

// must be one of ["auto", "us", "eu"]
func (o ConnectionOutput) DataResidency() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.DataResidency }).(pulumi.StringOutput)
}

func (o ConnectionOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

// Optional name of the connection
func (o ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// must be one of ["source", "destination", "customFormat"]
// Define the location where the data will be stored in the destination
func (o ConnectionOutput) NamespaceDefinition() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.NamespaceDefinition }).(pulumi.StringOutput)
}

// Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
func (o ConnectionOutput) NamespaceFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.NamespaceFormat }).(pulumi.StringOutput)
}

// must be one of ["ignore", "disable*connection", "propagate*columns", "propagateFully"]
// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
func (o ConnectionOutput) NonBreakingSchemaUpdatesBehavior() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.NonBreakingSchemaUpdatesBehavior }).(pulumi.StringOutput)
}

// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
func (o ConnectionOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Prefix }).(pulumi.StringOutput)
}

// schedule for when the the connection should run, per the schedule type
func (o ConnectionOutput) Schedule() ConnectionScheduleOutput {
	return o.ApplyT(func(v *Connection) ConnectionScheduleOutput { return v.Schedule }).(ConnectionScheduleOutput)
}

func (o ConnectionOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

// must be one of ["active", "inactive", "deprecated"]
func (o ConnectionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ConnectionOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type ConnectionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (o ConnectionArrayOutput) ToConnectionArrayOutput() ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) Index(i pulumi.IntInput) ConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].([]*Connection)[vs[1].(int)]
	}).(ConnectionOutput)
}

type ConnectionMapOutput struct{ *pulumi.OutputState }

func (ConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (o ConnectionMapOutput) ToConnectionMapOutput() ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) MapIndex(k pulumi.StringInput) ConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].(map[string]*Connection)[vs[1].(string)]
	}).(ConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionArrayInput)(nil)).Elem(), ConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionMapInput)(nil)).Elem(), ConnectionMap{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionMapOutput{})
}
