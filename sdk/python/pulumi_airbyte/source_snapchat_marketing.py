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

__all__ = ['SourceSnapchatMarketingArgs', 'SourceSnapchatMarketing']

@pulumi.input_type
class SourceSnapchatMarketingArgs:
    def __init__(__self__, *,
                 configuration: pulumi.Input['SourceSnapchatMarketingConfigurationArgs'],
                 name: pulumi.Input[str],
                 workspace_id: pulumi.Input[str],
                 secret_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SourceSnapchatMarketing resource.
        :param pulumi.Input[str] secret_id: Optional secretID obtained through the public API OAuth redirect flow.
        """
        pulumi.set(__self__, "configuration", configuration)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "workspace_id", workspace_id)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Input['SourceSnapchatMarketingConfigurationArgs']:
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: pulumi.Input['SourceSnapchatMarketingConfigurationArgs']):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workspace_id", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional secretID obtained through the public API OAuth redirect flow.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_id", value)


@pulumi.input_type
class _SourceSnapchatMarketingState:
    def __init__(__self__, *,
                 configuration: Optional[pulumi.Input['SourceSnapchatMarketingConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 source_id: Optional[pulumi.Input[str]] = None,
                 source_type: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SourceSnapchatMarketing resources.
        :param pulumi.Input[str] secret_id: Optional secretID obtained through the public API OAuth redirect flow.
        """
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)
        if source_id is not None:
            pulumi.set(__self__, "source_id", source_id)
        if source_type is not None:
            pulumi.set(__self__, "source_type", source_type)
        if workspace_id is not None:
            pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input['SourceSnapchatMarketingConfigurationArgs']]:
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input['SourceSnapchatMarketingConfigurationArgs']]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional secretID obtained through the public API OAuth redirect flow.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_id")

    @source_id.setter
    def source_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_id", value)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_type")

    @source_type.setter
    def source_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_type", value)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_id", value)


class SourceSnapchatMarketing(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['SourceSnapchatMarketingConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        SourceSnapchatMarketing Resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] secret_id: Optional secretID obtained through the public API OAuth redirect flow.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SourceSnapchatMarketingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        SourceSnapchatMarketing Resource

        :param str resource_name: The name of the resource.
        :param SourceSnapchatMarketingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SourceSnapchatMarketingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['SourceSnapchatMarketingConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SourceSnapchatMarketingArgs.__new__(SourceSnapchatMarketingArgs)

            if configuration is None and not opts.urn:
                raise TypeError("Missing required property 'configuration'")
            __props__.__dict__["configuration"] = configuration
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["secret_id"] = secret_id
            if workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_id'")
            __props__.__dict__["workspace_id"] = workspace_id
            __props__.__dict__["source_id"] = None
            __props__.__dict__["source_type"] = None
        super(SourceSnapchatMarketing, __self__).__init__(
            'airbyte:index/sourceSnapchatMarketing:SourceSnapchatMarketing',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            configuration: Optional[pulumi.Input[pulumi.InputType['SourceSnapchatMarketingConfigurationArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            secret_id: Optional[pulumi.Input[str]] = None,
            source_id: Optional[pulumi.Input[str]] = None,
            source_type: Optional[pulumi.Input[str]] = None,
            workspace_id: Optional[pulumi.Input[str]] = None) -> 'SourceSnapchatMarketing':
        """
        Get an existing SourceSnapchatMarketing resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] secret_id: Optional secretID obtained through the public API OAuth redirect flow.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SourceSnapchatMarketingState.__new__(_SourceSnapchatMarketingState)

        __props__.__dict__["configuration"] = configuration
        __props__.__dict__["name"] = name
        __props__.__dict__["secret_id"] = secret_id
        __props__.__dict__["source_id"] = source_id
        __props__.__dict__["source_type"] = source_type
        __props__.__dict__["workspace_id"] = workspace_id
        return SourceSnapchatMarketing(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output['outputs.SourceSnapchatMarketingConfiguration']:
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[Optional[str]]:
        """
        Optional secretID obtained through the public API OAuth redirect flow.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_id")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "workspace_id")

