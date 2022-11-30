// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("dnsimple");

/**
 * The account for API operations.
 */
export declare const account: string | undefined;
Object.defineProperty(exports, "account", {
    get() {
        return __config.get("account");
    },
    enumerable: true,
});

/**
 * Flag to enable the prefetch of zone records.
 */
export declare const prefetch: boolean | undefined;
Object.defineProperty(exports, "prefetch", {
    get() {
        return __config.getObject<boolean>("prefetch");
    },
    enumerable: true,
});

/**
 * Flag to enable the sandbox API.
 */
export declare const sandbox: boolean | undefined;
Object.defineProperty(exports, "sandbox", {
    get() {
        return __config.getObject<boolean>("sandbox");
    },
    enumerable: true,
});

/**
 * The API v2 token for API operations.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token");
    },
    enumerable: true,
});

/**
 * Custom string to append to the user agent used for sending HTTP requests to the API.
 */
export declare const userAgent: string | undefined;
Object.defineProperty(exports, "userAgent", {
    get() {
        return __config.get("userAgent");
    },
    enumerable: true,
});

