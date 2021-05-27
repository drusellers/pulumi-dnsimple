# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .email_forward import *
from .provider import *
from .record import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_dnsimple.config as config
else:
    config = _utilities.lazy_import('pulumi_dnsimple.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "dnsimple",
  "mod": "index/emailForward",
  "fqn": "pulumi_dnsimple",
  "classes": {
   "dnsimple:index/emailForward:EmailForward": "EmailForward"
  }
 },
 {
  "pkg": "dnsimple",
  "mod": "index/record",
  "fqn": "pulumi_dnsimple",
  "classes": {
   "dnsimple:index/record:Record": "Record"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "dnsimple",
  "token": "pulumi:providers:dnsimple",
  "fqn": "pulumi_dnsimple",
  "class": "Provider"
 }
]
"""
)
