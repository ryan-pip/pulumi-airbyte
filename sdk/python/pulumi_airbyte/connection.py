# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ConnectionArgs', 'Connection']

@pulumi.input_type
class ConnectionArgs:
    def __init__(__self__, *,
                 destination_id: pulumi.Input[str],
                 source_id: pulumi.Input[str],
                 configurations: Optional[pulumi.Input['ConnectionConfigurationsArgs']] = None,
                 data_residency: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_definition: Optional[pulumi.Input[str]] = None,
                 namespace_format: Optional[pulumi.Input[str]] = None,
                 non_breaking_schema_updates_behavior: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input['ConnectionScheduleArgs']] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Connection resource.
        :param pulumi.Input['ConnectionConfigurationsArgs'] configurations: A list of configured stream options for a connection.
        :param pulumi.Input[str] data_residency: must be one of ["auto", "us", "eu"]
        :param pulumi.Input[str] name: Optional name of the connection
        :param pulumi.Input[str] namespace_definition: must be one of ["source", "destination", "custom_format"]
               Define the location where the data will be stored in the destination
        :param pulumi.Input[str] namespace_format: Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
        :param pulumi.Input[str] non_breaking_schema_updates_behavior: must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]
               Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
        :param pulumi.Input[str] prefix: Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
        :param pulumi.Input['ConnectionScheduleArgs'] schedule: schedule for when the the connection should run, per the schedule type
        :param pulumi.Input[str] status: must be one of ["active", "inactive", "deprecated"]
        """
        pulumi.set(__self__, "destination_id", destination_id)
        pulumi.set(__self__, "source_id", source_id)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)
        if data_residency is not None:
            pulumi.set(__self__, "data_residency", data_residency)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace_definition is not None:
            pulumi.set(__self__, "namespace_definition", namespace_definition)
        if namespace_format is not None:
            pulumi.set(__self__, "namespace_format", namespace_format)
        if non_breaking_schema_updates_behavior is not None:
            pulumi.set(__self__, "non_breaking_schema_updates_behavior", non_breaking_schema_updates_behavior)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="destinationId")
    def destination_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "destination_id")

    @destination_id.setter
    def destination_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_id", value)

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source_id")

    @source_id.setter
    def source_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_id", value)

    @property
    @pulumi.getter
    def configurations(self) -> Optional[pulumi.Input['ConnectionConfigurationsArgs']]:
        """
        A list of configured stream options for a connection.
        """
        return pulumi.get(self, "configurations")

    @configurations.setter
    def configurations(self, value: Optional[pulumi.Input['ConnectionConfigurationsArgs']]):
        pulumi.set(self, "configurations", value)

    @property
    @pulumi.getter(name="dataResidency")
    def data_residency(self) -> Optional[pulumi.Input[str]]:
        """
        must be one of ["auto", "us", "eu"]
        """
        return pulumi.get(self, "data_residency")

    @data_residency.setter
    def data_residency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_residency", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional name of the connection
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namespaceDefinition")
    def namespace_definition(self) -> Optional[pulumi.Input[str]]:
        """
        must be one of ["source", "destination", "custom_format"]
        Define the location where the data will be stored in the destination
        """
        return pulumi.get(self, "namespace_definition")

    @namespace_definition.setter
    def namespace_definition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_definition", value)

    @property
    @pulumi.getter(name="namespaceFormat")
    def namespace_format(self) -> Optional[pulumi.Input[str]]:
        """
        Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
        """
        return pulumi.get(self, "namespace_format")

    @namespace_format.setter
    def namespace_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_format", value)

    @property
    @pulumi.getter(name="nonBreakingSchemaUpdatesBehavior")
    def non_breaking_schema_updates_behavior(self) -> Optional[pulumi.Input[str]]:
        """
        must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]
        Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
        """
        return pulumi.get(self, "non_breaking_schema_updates_behavior")

    @non_breaking_schema_updates_behavior.setter
    def non_breaking_schema_updates_behavior(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "non_breaking_schema_updates_behavior", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input['ConnectionScheduleArgs']]:
        """
        schedule for when the the connection should run, per the schedule type
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input['ConnectionScheduleArgs']]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        must be one of ["active", "inactive", "deprecated"]
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _ConnectionState:
    def __init__(__self__, *,
                 configurations: Optional[pulumi.Input['ConnectionConfigurationsArgs']] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 data_residency: Optional[pulumi.Input[str]] = None,
                 destination_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_definition: Optional[pulumi.Input[str]] = None,
                 namespace_format: Optional[pulumi.Input[str]] = None,
                 non_breaking_schema_updates_behavior: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input['ConnectionScheduleArgs']] = None,
                 source_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Connection resources.
        :param pulumi.Input['ConnectionConfigurationsArgs'] configurations: A list of configured stream options for a connection.
        :param pulumi.Input[str] data_residency: must be one of ["auto", "us", "eu"]
        :param pulumi.Input[str] name: Optional name of the connection
        :param pulumi.Input[str] namespace_definition: must be one of ["source", "destination", "custom_format"]
               Define the location where the data will be stored in the destination
        :param pulumi.Input[str] namespace_format: Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
        :param pulumi.Input[str] non_breaking_schema_updates_behavior: must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]
               Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
        :param pulumi.Input[str] prefix: Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
        :param pulumi.Input['ConnectionScheduleArgs'] schedule: schedule for when the the connection should run, per the schedule type
        :param pulumi.Input[str] status: must be one of ["active", "inactive", "deprecated"]
        """
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)
        if connection_id is not None:
            pulumi.set(__self__, "connection_id", connection_id)
        if data_residency is not None:
            pulumi.set(__self__, "data_residency", data_residency)
        if destination_id is not None:
            pulumi.set(__self__, "destination_id", destination_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace_definition is not None:
            pulumi.set(__self__, "namespace_definition", namespace_definition)
        if namespace_format is not None:
            pulumi.set(__self__, "namespace_format", namespace_format)
        if non_breaking_schema_updates_behavior is not None:
            pulumi.set(__self__, "non_breaking_schema_updates_behavior", non_breaking_schema_updates_behavior)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if source_id is not None:
            pulumi.set(__self__, "source_id", source_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if workspace_id is not None:
            pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter
    def configurations(self) -> Optional[pulumi.Input['ConnectionConfigurationsArgs']]:
        """
        A list of configured stream options for a connection.
        """
        return pulumi.get(self, "configurations")

    @configurations.setter
    def configurations(self, value: Optional[pulumi.Input['ConnectionConfigurationsArgs']]):
        pulumi.set(self, "configurations", value)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "connection_id")

    @connection_id.setter
    def connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_id", value)

    @property
    @pulumi.getter(name="dataResidency")
    def data_residency(self) -> Optional[pulumi.Input[str]]:
        """
        must be one of ["auto", "us", "eu"]
        """
        return pulumi.get(self, "data_residency")

    @data_residency.setter
    def data_residency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_residency", value)

    @property
    @pulumi.getter(name="destinationId")
    def destination_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination_id")

    @destination_id.setter
    def destination_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional name of the connection
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namespaceDefinition")
    def namespace_definition(self) -> Optional[pulumi.Input[str]]:
        """
        must be one of ["source", "destination", "custom_format"]
        Define the location where the data will be stored in the destination
        """
        return pulumi.get(self, "namespace_definition")

    @namespace_definition.setter
    def namespace_definition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_definition", value)

    @property
    @pulumi.getter(name="namespaceFormat")
    def namespace_format(self) -> Optional[pulumi.Input[str]]:
        """
        Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
        """
        return pulumi.get(self, "namespace_format")

    @namespace_format.setter
    def namespace_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_format", value)

    @property
    @pulumi.getter(name="nonBreakingSchemaUpdatesBehavior")
    def non_breaking_schema_updates_behavior(self) -> Optional[pulumi.Input[str]]:
        """
        must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]
        Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
        """
        return pulumi.get(self, "non_breaking_schema_updates_behavior")

    @non_breaking_schema_updates_behavior.setter
    def non_breaking_schema_updates_behavior(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "non_breaking_schema_updates_behavior", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input['ConnectionScheduleArgs']]:
        """
        schedule for when the the connection should run, per the schedule type
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input['ConnectionScheduleArgs']]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_id")

    @source_id.setter
    def source_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        must be one of ["active", "inactive", "deprecated"]
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_id", value)


class Connection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configurations: Optional[pulumi.Input[pulumi.InputType['ConnectionConfigurationsArgs']]] = None,
                 data_residency: Optional[pulumi.Input[str]] = None,
                 destination_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_definition: Optional[pulumi.Input[str]] = None,
                 namespace_format: Optional[pulumi.Input[str]] = None,
                 non_breaking_schema_updates_behavior: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['ConnectionScheduleArgs']]] = None,
                 source_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Connection Resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ConnectionConfigurationsArgs']] configurations: A list of configured stream options for a connection.
        :param pulumi.Input[str] data_residency: must be one of ["auto", "us", "eu"]
        :param pulumi.Input[str] name: Optional name of the connection
        :param pulumi.Input[str] namespace_definition: must be one of ["source", "destination", "custom_format"]
               Define the location where the data will be stored in the destination
        :param pulumi.Input[str] namespace_format: Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
        :param pulumi.Input[str] non_breaking_schema_updates_behavior: must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]
               Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
        :param pulumi.Input[str] prefix: Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
        :param pulumi.Input[pulumi.InputType['ConnectionScheduleArgs']] schedule: schedule for when the the connection should run, per the schedule type
        :param pulumi.Input[str] status: must be one of ["active", "inactive", "deprecated"]
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Connection Resource

        :param str resource_name: The name of the resource.
        :param ConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configurations: Optional[pulumi.Input[pulumi.InputType['ConnectionConfigurationsArgs']]] = None,
                 data_residency: Optional[pulumi.Input[str]] = None,
                 destination_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_definition: Optional[pulumi.Input[str]] = None,
                 namespace_format: Optional[pulumi.Input[str]] = None,
                 non_breaking_schema_updates_behavior: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['ConnectionScheduleArgs']]] = None,
                 source_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConnectionArgs.__new__(ConnectionArgs)

            __props__.__dict__["configurations"] = configurations
            __props__.__dict__["data_residency"] = data_residency
            if destination_id is None and not opts.urn:
                raise TypeError("Missing required property 'destination_id'")
            __props__.__dict__["destination_id"] = destination_id
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace_definition"] = namespace_definition
            __props__.__dict__["namespace_format"] = namespace_format
            __props__.__dict__["non_breaking_schema_updates_behavior"] = non_breaking_schema_updates_behavior
            __props__.__dict__["prefix"] = prefix
            __props__.__dict__["schedule"] = schedule
            if source_id is None and not opts.urn:
                raise TypeError("Missing required property 'source_id'")
            __props__.__dict__["source_id"] = source_id
            __props__.__dict__["status"] = status
            __props__.__dict__["connection_id"] = None
            __props__.__dict__["workspace_id"] = None
        super(Connection, __self__).__init__(
            'airbyte:index/connection:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            configurations: Optional[pulumi.Input[pulumi.InputType['ConnectionConfigurationsArgs']]] = None,
            connection_id: Optional[pulumi.Input[str]] = None,
            data_residency: Optional[pulumi.Input[str]] = None,
            destination_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_definition: Optional[pulumi.Input[str]] = None,
            namespace_format: Optional[pulumi.Input[str]] = None,
            non_breaking_schema_updates_behavior: Optional[pulumi.Input[str]] = None,
            prefix: Optional[pulumi.Input[str]] = None,
            schedule: Optional[pulumi.Input[pulumi.InputType['ConnectionScheduleArgs']]] = None,
            source_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            workspace_id: Optional[pulumi.Input[str]] = None) -> 'Connection':
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ConnectionConfigurationsArgs']] configurations: A list of configured stream options for a connection.
        :param pulumi.Input[str] data_residency: must be one of ["auto", "us", "eu"]
        :param pulumi.Input[str] name: Optional name of the connection
        :param pulumi.Input[str] namespace_definition: must be one of ["source", "destination", "custom_format"]
               Define the location where the data will be stored in the destination
        :param pulumi.Input[str] namespace_format: Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
        :param pulumi.Input[str] non_breaking_schema_updates_behavior: must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]
               Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
        :param pulumi.Input[str] prefix: Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
        :param pulumi.Input[pulumi.InputType['ConnectionScheduleArgs']] schedule: schedule for when the the connection should run, per the schedule type
        :param pulumi.Input[str] status: must be one of ["active", "inactive", "deprecated"]
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConnectionState.__new__(_ConnectionState)

        __props__.__dict__["configurations"] = configurations
        __props__.__dict__["connection_id"] = connection_id
        __props__.__dict__["data_residency"] = data_residency
        __props__.__dict__["destination_id"] = destination_id
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace_definition"] = namespace_definition
        __props__.__dict__["namespace_format"] = namespace_format
        __props__.__dict__["non_breaking_schema_updates_behavior"] = non_breaking_schema_updates_behavior
        __props__.__dict__["prefix"] = prefix
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["source_id"] = source_id
        __props__.__dict__["status"] = status
        __props__.__dict__["workspace_id"] = workspace_id
        return Connection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def configurations(self) -> pulumi.Output['outputs.ConnectionConfigurations']:
        """
        A list of configured stream options for a connection.
        """
        return pulumi.get(self, "configurations")

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "connection_id")

    @property
    @pulumi.getter(name="dataResidency")
    def data_residency(self) -> pulumi.Output[str]:
        """
        must be one of ["auto", "us", "eu"]
        """
        return pulumi.get(self, "data_residency")

    @property
    @pulumi.getter(name="destinationId")
    def destination_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "destination_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Optional name of the connection
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceDefinition")
    def namespace_definition(self) -> pulumi.Output[str]:
        """
        must be one of ["source", "destination", "custom_format"]
        Define the location where the data will be stored in the destination
        """
        return pulumi.get(self, "namespace_definition")

    @property
    @pulumi.getter(name="namespaceFormat")
    def namespace_format(self) -> pulumi.Output[str]:
        """
        Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
        """
        return pulumi.get(self, "namespace_format")

    @property
    @pulumi.getter(name="nonBreakingSchemaUpdatesBehavior")
    def non_breaking_schema_updates_behavior(self) -> pulumi.Output[str]:
        """
        must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]
        Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
        """
        return pulumi.get(self, "non_breaking_schema_updates_behavior")

    @property
    @pulumi.getter
    def prefix(self) -> pulumi.Output[str]:
        """
        Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output['outputs.ConnectionSchedule']:
        """
        schedule for when the the connection should run, per the schedule type
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        must be one of ["active", "inactive", "deprecated"]
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "workspace_id")

