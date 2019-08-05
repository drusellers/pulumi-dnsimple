# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Record(pulumi.CustomResource):
    domain: pulumi.Output[str]
    """
    The domain to add the record to
    """
    domain_id: pulumi.Output[str]
    """
    The domain ID of the record
    """
    hostname: pulumi.Output[str]
    """
    The FQDN of the record
    """
    name: pulumi.Output[str]
    """
    The name of the record
    """
    priority: pulumi.Output[str]
    """
    The priority of the record - only useful for some record types
    """
    ttl: pulumi.Output[str]
    """
    The TTL of the record
    """
    type: pulumi.Output[str]
    """
    The type of the record
    """
    value: pulumi.Output[str]
    """
    The value of the record
    """
    def __init__(__self__, resource_name, opts=None, domain=None, name=None, priority=None, ttl=None, type=None, value=None, __name__=None, __opts__=None):
        """
        Provides a DNSimple record resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: The domain to add the record to
        :param pulumi.Input[str] name: The name of the record
        :param pulumi.Input[str] priority: The priority of the record - only useful for some record types
        :param pulumi.Input[str] ttl: The TTL of the record
        :param pulumi.Input[str] type: The type of the record
        :param pulumi.Input[str] value: The value of the record

        > This content is derived from https://github.com/terraform-providers/terraform-provider-dnsimple/blob/master/website/docs/r/record.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if domain is None:
            raise TypeError("Missing required property 'domain'")
        __props__['domain'] = domain

        if name is None:
            raise TypeError("Missing required property 'name'")
        __props__['name'] = name

        __props__['priority'] = priority

        __props__['ttl'] = ttl

        if type is None:
            raise TypeError("Missing required property 'type'")
        __props__['type'] = type

        if value is None:
            raise TypeError("Missing required property 'value'")
        __props__['value'] = value

        __props__['domain_id'] = None
        __props__['hostname'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Record, __self__).__init__(
            'dnsimple:index/record:Record',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

