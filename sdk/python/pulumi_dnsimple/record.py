# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RecordArgs', 'Record']

@pulumi.input_type
class RecordArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 value: pulumi.Input[str],
                 priority: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Record resource.
        :param pulumi.Input[str] domain: The domain to add the record to
        :param pulumi.Input[str] name: The name of the record
        :param pulumi.Input[str] type: The type of the record
        :param pulumi.Input[str] value: The value of the record
        :param pulumi.Input[str] priority: The priority of the record - only useful for some record types
        :param pulumi.Input[str] ttl: The TTL of the record
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The domain to add the record to
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the record
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the record
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value of the record
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[str]]:
        """
        The priority of the record - only useful for some record types
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        The TTL of the record
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)


@pulumi.input_type
class _RecordState:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Record resources.
        :param pulumi.Input[str] domain: The domain to add the record to
        :param pulumi.Input[str] domain_id: The domain ID of the record
        :param pulumi.Input[str] hostname: The FQDN of the record
        :param pulumi.Input[str] name: The name of the record
        :param pulumi.Input[str] priority: The priority of the record - only useful for some record types
        :param pulumi.Input[str] ttl: The TTL of the record
        :param pulumi.Input[str] type: The type of the record
        :param pulumi.Input[str] value: The value of the record
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain to add the record to
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[pulumi.Input[str]]:
        """
        The domain ID of the record
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The FQDN of the record
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the record
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[str]]:
        """
        The priority of the record - only useful for some record types
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        The TTL of the record
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the record
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the record
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class Record(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a DNSimple record resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_dnsimple as dnsimple

        # Add a record to the root domain
        foobar = dnsimple.Record("foobar",
            domain=var["dnsimple_domain"],
            name="",
            ttl="3600",
            type="A",
            value="192.168.0.11")
        ```

        ```python
        import pulumi
        import pulumi_dnsimple as dnsimple

        # Add a record to a sub-domain
        foobar = dnsimple.Record("foobar",
            domain=var["dnsimple_domain"],
            name="terraform",
            ttl="3600",
            type="A",
            value="192.168.0.11")
        ```

        ## Import

        DNSimple resources can be imported using their parent zone name (domain name) and numeric record ID. **Importing record example.com with record ID 1234**

        ```sh
         $ pulumi import dnsimple:index/record:Record resource_name example.com_1234
        ```

         **Importing record www.example.com with record ID 1234**

        ```sh
         $ pulumi import dnsimple:index/record:Record resource_name example.com_1234
        ```

         The record ID can be found in the URL when editing a record on the DNSimple web dashboard.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: The domain to add the record to
        :param pulumi.Input[str] name: The name of the record
        :param pulumi.Input[str] priority: The priority of the record - only useful for some record types
        :param pulumi.Input[str] ttl: The TTL of the record
        :param pulumi.Input[str] type: The type of the record
        :param pulumi.Input[str] value: The value of the record
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RecordArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a DNSimple record resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_dnsimple as dnsimple

        # Add a record to the root domain
        foobar = dnsimple.Record("foobar",
            domain=var["dnsimple_domain"],
            name="",
            ttl="3600",
            type="A",
            value="192.168.0.11")
        ```

        ```python
        import pulumi
        import pulumi_dnsimple as dnsimple

        # Add a record to a sub-domain
        foobar = dnsimple.Record("foobar",
            domain=var["dnsimple_domain"],
            name="terraform",
            ttl="3600",
            type="A",
            value="192.168.0.11")
        ```

        ## Import

        DNSimple resources can be imported using their parent zone name (domain name) and numeric record ID. **Importing record example.com with record ID 1234**

        ```sh
         $ pulumi import dnsimple:index/record:Record resource_name example.com_1234
        ```

         **Importing record www.example.com with record ID 1234**

        ```sh
         $ pulumi import dnsimple:index/record:Record resource_name example.com_1234
        ```

         The record ID can be found in the URL when editing a record on the DNSimple web dashboard.

        :param str resource_name: The name of the resource.
        :param RecordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RecordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RecordArgs.__new__(RecordArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["priority"] = priority
            __props__.__dict__["ttl"] = ttl
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            __props__.__dict__["domain_id"] = None
            __props__.__dict__["hostname"] = None
        super(Record, __self__).__init__(
            'dnsimple:index/record:Record',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain: Optional[pulumi.Input[str]] = None,
            domain_id: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'Record':
        """
        Get an existing Record resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: The domain to add the record to
        :param pulumi.Input[str] domain_id: The domain ID of the record
        :param pulumi.Input[str] hostname: The FQDN of the record
        :param pulumi.Input[str] name: The name of the record
        :param pulumi.Input[str] priority: The priority of the record - only useful for some record types
        :param pulumi.Input[str] ttl: The TTL of the record
        :param pulumi.Input[str] type: The type of the record
        :param pulumi.Input[str] value: The value of the record
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RecordState.__new__(_RecordState)

        __props__.__dict__["domain"] = domain
        __props__.__dict__["domain_id"] = domain_id
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["name"] = name
        __props__.__dict__["priority"] = priority
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["type"] = type
        __props__.__dict__["value"] = value
        return Record(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The domain to add the record to
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Output[str]:
        """
        The domain ID of the record
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        The FQDN of the record
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the record
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[str]:
        """
        The priority of the record - only useful for some record types
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[str]]:
        """
        The TTL of the record
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the record
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The value of the record
        """
        return pulumi.get(self, "value")

