// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("dnsimple");

/**
 * The account for API operations.
 */
export let account: string | undefined = __config.get("account");
/**
 * Flag to enable the sandbox API.
 */
export let sandbox: boolean | undefined = __config.getObject<boolean>("sandbox");
/**
 * The API v2 token for API operations.
 */
export let token: string | undefined = __config.get("token");
