// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.dnsimple.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetCertificateArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCertificateArgs Empty = new GetCertificateArgs();

    /**
     * The ID of the SSL Certificate
     * 
     */
    @Import(name="certificateId", required=true)
    private Output<String> certificateId;

    /**
     * @return The ID of the SSL Certificate
     * 
     */
    public Output<String> certificateId() {
        return this.certificateId;
    }

    /**
     * The domain of the SSL Certificate
     * 
     */
    @Import(name="domain", required=true)
    private Output<String> domain;

    /**
     * @return The domain of the SSL Certificate
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }

    private GetCertificateArgs() {}

    private GetCertificateArgs(GetCertificateArgs $) {
        this.certificateId = $.certificateId;
        this.domain = $.domain;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCertificateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCertificateArgs $;

        public Builder() {
            $ = new GetCertificateArgs();
        }

        public Builder(GetCertificateArgs defaults) {
            $ = new GetCertificateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param certificateId The ID of the SSL Certificate
         * 
         * @return builder
         * 
         */
        public Builder certificateId(Output<String> certificateId) {
            $.certificateId = certificateId;
            return this;
        }

        /**
         * @param certificateId The ID of the SSL Certificate
         * 
         * @return builder
         * 
         */
        public Builder certificateId(String certificateId) {
            return certificateId(Output.of(certificateId));
        }

        /**
         * @param domain The domain of the SSL Certificate
         * 
         * @return builder
         * 
         */
        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The domain of the SSL Certificate
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        public GetCertificateArgs build() {
            if ($.certificateId == null) {
                throw new MissingRequiredPropertyException("GetCertificateArgs", "certificateId");
            }
            if ($.domain == null) {
                throw new MissingRequiredPropertyException("GetCertificateArgs", "domain");
            }
            return $;
        }
    }

}
