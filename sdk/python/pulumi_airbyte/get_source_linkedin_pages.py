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
    'GetSourceLinkedinPagesResult',
    'AwaitableGetSourceLinkedinPagesResult',
    'get_source_linkedin_pages',
    'get_source_linkedin_pages_output',
]

@pulumi.output_type
class GetSourceLinkedinPagesResult:
    """
    A collection of values returned by getSourceLinkedinPages.
    """
    def __init__(__self__, configuration=None, id=None, name=None, secret_id=None, source_id=None, workspace_id=None):
        if configuration and not isinstance(configuration, dict):
            raise TypeError("Expected argument 'configuration' to be a dict")
        pulumi.set(__self__, "configuration", configuration)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if secret_id and not isinstance(secret_id, str):
            raise TypeError("Expected argument 'secret_id' to be a str")
        pulumi.set(__self__, "secret_id", secret_id)
        if source_id and not isinstance(source_id, str):
            raise TypeError("Expected argument 'source_id' to be a str")
        pulumi.set(__self__, "source_id", source_id)
        if workspace_id and not isinstance(workspace_id, str):
            raise TypeError("Expected argument 'workspace_id' to be a str")
        pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter
    def configuration(self) -> 'outputs.GetSourceLinkedinPagesConfigurationResult':
        return pulumi.get(self, "configuration")

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
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[str]:
        """
        Optional secretID obtained through the public API OAuth redirect flow.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> str:
        return pulumi.get(self, "source_id")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> str:
        return pulumi.get(self, "workspace_id")


class AwaitableGetSourceLinkedinPagesResult(GetSourceLinkedinPagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSourceLinkedinPagesResult(
            configuration=self.configuration,
            id=self.id,
            name=self.name,
            secret_id=self.secret_id,
            source_id=self.source_id,
            workspace_id=self.workspace_id)


def get_source_linkedin_pages(secret_id: Optional[str] = None,
                              source_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSourceLinkedinPagesResult:
    """
    SourceLinkedinPages DataSource

    ## Example Usage

    ```python
    import pulumi
    import pulumi_airbyte as airbyte

    my_source_linkedinpages = airbyte.get_source_linkedin_pages(secret_id="...my_secret_id...",
        source_id="...my_source_id...")
    ```


    :param str secret_id: Optional secretID obtained through the public API OAuth redirect flow.
    """
    __args__ = dict()
    __args__['secretId'] = secret_id
    __args__['sourceId'] = source_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('airbyte:index/getSourceLinkedinPages:getSourceLinkedinPages', __args__, opts=opts, typ=GetSourceLinkedinPagesResult).value

    return AwaitableGetSourceLinkedinPagesResult(
        configuration=pulumi.get(__ret__, 'configuration'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        secret_id=pulumi.get(__ret__, 'secret_id'),
        source_id=pulumi.get(__ret__, 'source_id'),
        workspace_id=pulumi.get(__ret__, 'workspace_id'))


@_utilities.lift_output_func(get_source_linkedin_pages)
def get_source_linkedin_pages_output(secret_id: Optional[pulumi.Input[Optional[str]]] = None,
                                     source_id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSourceLinkedinPagesResult]:
    """
    SourceLinkedinPages DataSource

    ## Example Usage

    ```python
    import pulumi
    import pulumi_airbyte as airbyte

    my_source_linkedinpages = airbyte.get_source_linkedin_pages(secret_id="...my_secret_id...",
        source_id="...my_source_id...")
    ```


    :param str secret_id: Optional secretID obtained through the public API OAuth redirect flow.
    """
    ...
