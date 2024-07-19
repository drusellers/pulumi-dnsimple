// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.dnsimple.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCertificateTimeouts extends com.pulumi.resources.InvokeArgs {

    public static final GetCertificateTimeouts Empty = new GetCertificateTimeouts();

    /**
     * (String) - The timeout for the read operation e.g. `5m`
     * 
     */
    @Import(name="read")
    private @Nullable String read;

    /**
     * @return (String) - The timeout for the read operation e.g. `5m`
     * 
     */
    public Optional<String> read() {
        return Optional.ofNullable(this.read);
    }

    private GetCertificateTimeouts() {}

    private GetCertificateTimeouts(GetCertificateTimeouts $) {
        this.read = $.read;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCertificateTimeouts defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCertificateTimeouts $;

        public Builder() {
            $ = new GetCertificateTimeouts();
        }

        public Builder(GetCertificateTimeouts defaults) {
            $ = new GetCertificateTimeouts(Objects.requireNonNull(defaults));
        }

        /**
         * @param read (String) - The timeout for the read operation e.g. `5m`
         * 
         * @return builder
         * 
         */
        public Builder read(@Nullable String read) {
            $.read = read;
            return this;
        }

        public GetCertificateTimeouts build() {
            return $;
        }
    }

}
