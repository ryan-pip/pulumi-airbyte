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

__all__ = [
    'GetConnectionResult',
    'AwaitableGetConnectionResult',
    'get_connection',
    'get_connection_output',
]

@pulumi.output_type
class GetConnectionResult:
    """
    A collection of values returned by getConnection.
    """
    def __init__(__self__, configurations=None, connection_id=None, data_residency=None, destination_id=None, id=None, name=None, namespace_definition=None, namespace_format=None, non_breaking_schema_updates_behavior=None, prefix=None, schedule=None, source_id=None, status=None, workspace_id=None):
        if configurations and not isinstance(configurations, dict):
            raise TypeError("Expected argument 'configurations' to be a dict")
        pulumi.set(__self__, "configurations", configurations)
        if connection_id and not isinstance(connection_id, str):
            raise TypeError("Expected argument 'connection_id' to be a str")
        pulumi.set(__self__, "connection_id", connection_id)
        if data_residency and not isinstance(data_residency, str):
            raise TypeError("Expected argument 'data_residency' to be a str")
        pulumi.set(__self__, "data_residency", data_residency)
        if destination_id and not isinstance(destination_id, str):
            raise TypeError("Expected argument 'destination_id' to be a str")
        pulumi.set(__self__, "destination_id", destination_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace_definition and not isinstance(namespace_definition, str):
            raise TypeError("Expected argument 'namespace_definition' to be a str")
        pulumi.set(__self__, "namespace_definition", namespace_definition)
        if namespace_format and not isinstance(namespace_format, str):
            raise TypeError("Expected argument 'namespace_format' to be a str")
        pulumi.set(__self__, "namespace_format", namespace_format)
        if non_breaking_schema_updates_behavior and not isinstance(non_breaking_schema_updates_behavior, str):
            raise TypeError("Expected argument 'non_breaking_schema_updates_behavior' to be a str")
        pulumi.set(__self__, "non_breaking_schema_updates_behavior", non_breaking_schema_updates_behavior)
        if prefix and not isinstance(prefix, str):
            raise TypeError("Expected argument 'prefix' to be a str")
        pulumi.set(__self__, "prefix", prefix)
        if schedule and not isinstance(schedule, dict):
            raise TypeError("Expected argument 'schedule' to be a dict")
        pulumi.set(__self__, "schedule", schedule)
        if source_id and not isinstance(source_id, str):
            raise TypeError("Expected argument 'source_id' to be a str")
        pulumi.set(__self__, "source_id", source_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if workspace_id and not isinstance(workspace_id, str):
            raise TypeError("Expected argument 'workspace_id' to be a str")
        pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter
    def configurations(self) -> 'outputs.GetConnectionConfigurationsResult':
        """
        A list of configured stream options for a connection.
        """
        return pulumi.get(self, "configurations")

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> str:
        return pulumi.get(self, "connection_id")

    @property
    @pulumi.getter(name="dataResidency")
    def data_residency(self) -> str:
        """
        must be one of ["auto", "us", "eu"]
        """
        return pulumi.get(self, "data_residency")

    @property
    @pulumi.getter(name="destinationId")
    def destination_id(self) -> str:
        return pulumi.get(self, "destination_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Optional name of the connection
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceDefinition")
    def namespace_definition(self) -> str:
        """
        must be one of ["source", "destination", "custom_format"]
        Define the location where the data will be stored in the destination
        """
        return pulumi.get(self, "namespace_definition")

    @property
    @pulumi.getter(name="namespaceFormat")
    def namespace_format(self) -> str:
        """
        Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'.
        """
        return pulumi.get(self, "namespace_format")

    @property
    @pulumi.getter(name="nonBreakingSchemaUpdatesBehavior")
    def non_breaking_schema_updates_behavior(self) -> str:
        """
        must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]
        Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
        """
        return pulumi.get(self, "non_breaking_schema_updates_behavior")

    @property
    @pulumi.getter
    def prefix(self) -> str:
        """
        Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” => “airbyte*projects”).
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter
    def schedule(self) -> 'outputs.GetConnectionScheduleResult':
        """
        schedule for when the the connection should run, per the schedule type
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> str:
        return pulumi.get(self, "source_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        must be one of ["active", "inactive", "deprecated"]
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> str:
        return pulumi.get(self, "workspace_id")


class AwaitableGetConnectionResult(GetConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectionResult(
            configurations=self.configurations,
            connection_id=self.connection_id,
            data_residency=self.data_residency,
            destination_id=self.destination_id,
            id=self.id,
            name=self.name,
            namespace_definition=self.namespace_definition,
            namespace_format=self.namespace_format,
            non_breaking_schema_updates_behavior=self.non_breaking_schema_updates_behavior,
            prefix=self.prefix,
            schedule=self.schedule,
            source_id=self.source_id,
            status=self.status,
            workspace_id=self.workspace_id)


def get_connection(connection_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectionResult:
    """
    Connection DataSource

    ## Example Usage

    ```python
    import pulumi
    import pulumi_airbyte as airbyte

    my_connection = airbyte.get_connection(connection_id="...my_connection_id...")
    ```
    """
    __args__ = dict()
    __args__['connectionId'] = connection_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('airbyte:index/getConnection:getConnection', __args__, opts=opts, typ=GetConnectionResult).value

    return AwaitableGetConnectionResult(
        configurations=pulumi.get(__ret__, 'configurations'),
        connection_id=pulumi.get(__ret__, 'connection_id'),
        data_residency=pulumi.get(__ret__, 'data_residency'),
        destination_id=pulumi.get(__ret__, 'destination_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        namespace_definition=pulumi.get(__ret__, 'namespace_definition'),
        namespace_format=pulumi.get(__ret__, 'namespace_format'),
        non_breaking_schema_updates_behavior=pulumi.get(__ret__, 'non_breaking_schema_updates_behavior'),
        prefix=pulumi.get(__ret__, 'prefix'),
        schedule=pulumi.get(__ret__, 'schedule'),
        source_id=pulumi.get(__ret__, 'source_id'),
        status=pulumi.get(__ret__, 'status'),
        workspace_id=pulumi.get(__ret__, 'workspace_id'))


@_utilities.lift_output_func(get_connection)
def get_connection_output(connection_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConnectionResult]:
    """
    Connection DataSource

    ## Example Usage

    ```python
    import pulumi
    import pulumi_airbyte as airbyte

    my_connection = airbyte.get_connection(connection_id="...my_connection_id...")
    ```
    """
    ...
