# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('airbyte')


class _ExportableConfig(types.ModuleType):
    @property
    def bearer_auth(self) -> Optional[str]:
        return __config__.get('bearerAuth') or _utilities.get_env('AIRBYTE_BEARER_AUTH')

    @property
    def password(self) -> Optional[str]:
        return __config__.get('password') or _utilities.get_env('AIRBYTE_PASSWORD')

    @property
    def server_url(self) -> str:
        """
        Server URL (defaults to https://api.airbyte.com/v1)
        """
        return __config__.get('serverUrl') or (_utilities.get_env('AIRBYTE_SERVER_URL') or 'https://api.airbyte.com/v1')

    @property
    def username(self) -> Optional[str]:
        return __config__.get('username') or _utilities.get_env('AIRBYTE_USERNAME')

