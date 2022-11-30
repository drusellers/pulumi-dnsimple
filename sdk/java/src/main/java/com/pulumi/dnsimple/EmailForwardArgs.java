// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.dnsimple;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class EmailForwardArgs extends com.pulumi.resources.ResourceArgs {

    public static final EmailForwardArgs Empty = new EmailForwardArgs();

    @Import(name="aliasName", required=true)
    private Output<String> aliasName;

    public Output<String> aliasName() {
        return this.aliasName;
    }

    @Import(name="destinationEmail", required=true)
    private Output<String> destinationEmail;

    public Output<String> destinationEmail() {
        return this.destinationEmail;
    }

    @Import(name="domain", required=true)
    private Output<String> domain;

    public Output<String> domain() {
        return this.domain;
    }

    private EmailForwardArgs() {}

    private EmailForwardArgs(EmailForwardArgs $) {
        this.aliasName = $.aliasName;
        this.destinationEmail = $.destinationEmail;
        this.domain = $.domain;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EmailForwardArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EmailForwardArgs $;

        public Builder() {
            $ = new EmailForwardArgs();
        }

        public Builder(EmailForwardArgs defaults) {
            $ = new EmailForwardArgs(Objects.requireNonNull(defaults));
        }

        public Builder aliasName(Output<String> aliasName) {
            $.aliasName = aliasName;
            return this;
        }

        public Builder aliasName(String aliasName) {
            return aliasName(Output.of(aliasName));
        }

        public Builder destinationEmail(Output<String> destinationEmail) {
            $.destinationEmail = destinationEmail;
            return this;
        }

        public Builder destinationEmail(String destinationEmail) {
            return destinationEmail(Output.of(destinationEmail));
        }

        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        public EmailForwardArgs build() {
            $.aliasName = Objects.requireNonNull($.aliasName, "expected parameter 'aliasName' to be non-null");
            $.destinationEmail = Objects.requireNonNull($.destinationEmail, "expected parameter 'destinationEmail' to be non-null");
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            return $;
        }
    }

}