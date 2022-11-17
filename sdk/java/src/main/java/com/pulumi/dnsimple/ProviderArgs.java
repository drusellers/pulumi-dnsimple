// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.dnsimple;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * The account for API operations.
     * 
     */
    @Import(name="account", required=true)
    private Output<String> account;

    /**
     * @return The account for API operations.
     * 
     */
    public Output<String> account() {
        return this.account;
    }

    /**
     * Flag to enable the sandbox API.
     * 
     */
    @Import(name="sandbox", json=true)
    private @Nullable Output<Boolean> sandbox;

    /**
     * @return Flag to enable the sandbox API.
     * 
     */
    public Optional<Output<Boolean>> sandbox() {
        return Optional.ofNullable(this.sandbox);
    }

    /**
     * The API v2 token for API operations.
     * 
     */
    @Import(name="token", required=true)
    private Output<String> token;

    /**
     * @return The API v2 token for API operations.
     * 
     */
    public Output<String> token() {
        return this.token;
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.account = $.account;
        this.sandbox = $.sandbox;
        this.token = $.token;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param account The account for API operations.
         * 
         * @return builder
         * 
         */
        public Builder account(Output<String> account) {
            $.account = account;
            return this;
        }

        /**
         * @param account The account for API operations.
         * 
         * @return builder
         * 
         */
        public Builder account(String account) {
            return account(Output.of(account));
        }

        /**
         * @param sandbox Flag to enable the sandbox API.
         * 
         * @return builder
         * 
         */
        public Builder sandbox(@Nullable Output<Boolean> sandbox) {
            $.sandbox = sandbox;
            return this;
        }

        /**
         * @param sandbox Flag to enable the sandbox API.
         * 
         * @return builder
         * 
         */
        public Builder sandbox(Boolean sandbox) {
            return sandbox(Output.of(sandbox));
        }

        /**
         * @param token The API v2 token for API operations.
         * 
         * @return builder
         * 
         */
        public Builder token(Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token The API v2 token for API operations.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        public ProviderArgs build() {
            $.account = Objects.requireNonNull($.account, "expected parameter 'account' to be non-null");
            $.token = Objects.requireNonNull($.token, "expected parameter 'token' to be non-null");
            return $;
        }
    }

}
