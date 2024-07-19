// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.dnsimple.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCertificateTimeouts {
    /**
     * @return (String) - The timeout for the read operation e.g. `5m`
     * 
     */
    private @Nullable String read;

    private GetCertificateTimeouts() {}
    /**
     * @return (String) - The timeout for the read operation e.g. `5m`
     * 
     */
    public Optional<String> read() {
        return Optional.ofNullable(this.read);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCertificateTimeouts defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String read;
        public Builder() {}
        public Builder(GetCertificateTimeouts defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.read = defaults.read;
        }

        @CustomType.Setter
        public Builder read(@Nullable String read) {

            this.read = read;
            return this;
        }
        public GetCertificateTimeouts build() {
            final var _resultValue = new GetCertificateTimeouts();
            _resultValue.read = read;
            return _resultValue;
        }
    }
}
