// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.dnsimple.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.dnsimple.outputs.GetRegistrantChangeCheckExtendedAttribute;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRegistrantChangeCheckResult {
    private String contactId;
    private String domainId;
    private List<GetRegistrantChangeCheckExtendedAttribute> extendedAttributes;
    private String id;
    private Boolean registryOwnerChange;

    private GetRegistrantChangeCheckResult() {}
    public String contactId() {
        return this.contactId;
    }
    public String domainId() {
        return this.domainId;
    }
    public List<GetRegistrantChangeCheckExtendedAttribute> extendedAttributes() {
        return this.extendedAttributes;
    }
    public String id() {
        return this.id;
    }
    public Boolean registryOwnerChange() {
        return this.registryOwnerChange;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRegistrantChangeCheckResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String contactId;
        private String domainId;
        private List<GetRegistrantChangeCheckExtendedAttribute> extendedAttributes;
        private String id;
        private Boolean registryOwnerChange;
        public Builder() {}
        public Builder(GetRegistrantChangeCheckResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.contactId = defaults.contactId;
    	      this.domainId = defaults.domainId;
    	      this.extendedAttributes = defaults.extendedAttributes;
    	      this.id = defaults.id;
    	      this.registryOwnerChange = defaults.registryOwnerChange;
        }

        @CustomType.Setter
        public Builder contactId(String contactId) {
            this.contactId = Objects.requireNonNull(contactId);
            return this;
        }
        @CustomType.Setter
        public Builder domainId(String domainId) {
            this.domainId = Objects.requireNonNull(domainId);
            return this;
        }
        @CustomType.Setter
        public Builder extendedAttributes(List<GetRegistrantChangeCheckExtendedAttribute> extendedAttributes) {
            this.extendedAttributes = Objects.requireNonNull(extendedAttributes);
            return this;
        }
        public Builder extendedAttributes(GetRegistrantChangeCheckExtendedAttribute... extendedAttributes) {
            return extendedAttributes(List.of(extendedAttributes));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder registryOwnerChange(Boolean registryOwnerChange) {
            this.registryOwnerChange = Objects.requireNonNull(registryOwnerChange);
            return this;
        }
        public GetRegistrantChangeCheckResult build() {
            final var _resultValue = new GetRegistrantChangeCheckResult();
            _resultValue.contactId = contactId;
            _resultValue.domainId = domainId;
            _resultValue.extendedAttributes = extendedAttributes;
            _resultValue.id = id;
            _resultValue.registryOwnerChange = registryOwnerChange;
            return _resultValue;
        }
    }
}