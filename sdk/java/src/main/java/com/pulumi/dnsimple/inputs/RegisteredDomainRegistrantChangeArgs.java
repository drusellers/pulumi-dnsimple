// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.dnsimple.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegisteredDomainRegistrantChangeArgs extends com.pulumi.resources.ResourceArgs {

    public static final RegisteredDomainRegistrantChangeArgs Empty = new RegisteredDomainRegistrantChangeArgs();

    @Import(name="accountId")
    private @Nullable Output<Integer> accountId;

    public Optional<Output<Integer>> accountId() {
        return Optional.ofNullable(this.accountId);
    }

    /**
     * The ID of the contact to be used for the domain registration. The contact ID can be changed after the domain has been registered. The change will result in a new registrant change this may result in a [60-day lock](https://support.dnsimple.com/articles/icann-60-day-lock-registrant-change/).
     * 
     */
    @Import(name="contactId")
    private @Nullable Output<Integer> contactId;

    /**
     * @return The ID of the contact to be used for the domain registration. The contact ID can be changed after the domain has been registered. The change will result in a new registrant change this may result in a [60-day lock](https://support.dnsimple.com/articles/icann-60-day-lock-registrant-change/).
     * 
     */
    public Optional<Output<Integer>> contactId() {
        return Optional.ofNullable(this.contactId);
    }

    @Import(name="domainId")
    private @Nullable Output<String> domainId;

    public Optional<Output<String>> domainId() {
        return Optional.ofNullable(this.domainId);
    }

    /**
     * A map of extended attributes to be set for the domain registration. To see if there are any required extended attributes for any TLD use our [Lists the TLD Extended Attributes API](https://developer.dnsimple.com/v2/tlds/#getTldExtendedAttributes). The values provided in the `extended_attributes` will also be sent when a registrant change is initiated as part of changing the `contact_id`.
     * 
     */
    @Import(name="extendedAttributes")
    private @Nullable Output<Map<String,String>> extendedAttributes;

    /**
     * @return A map of extended attributes to be set for the domain registration. To see if there are any required extended attributes for any TLD use our [Lists the TLD Extended Attributes API](https://developer.dnsimple.com/v2/tlds/#getTldExtendedAttributes). The values provided in the `extended_attributes` will also be sent when a registrant change is initiated as part of changing the `contact_id`.
     * 
     */
    public Optional<Output<Map<String,String>>> extendedAttributes() {
        return Optional.ofNullable(this.extendedAttributes);
    }

    /**
     * The ID of this resource.
     * 
     */
    @Import(name="id")
    private @Nullable Output<Integer> id;

    /**
     * @return The ID of this resource.
     * 
     */
    public Optional<Output<Integer>> id() {
        return Optional.ofNullable(this.id);
    }

    @Import(name="irtLockLiftedBy")
    private @Nullable Output<String> irtLockLiftedBy;

    public Optional<Output<String>> irtLockLiftedBy() {
        return Optional.ofNullable(this.irtLockLiftedBy);
    }

    @Import(name="registryOwnerChange")
    private @Nullable Output<Boolean> registryOwnerChange;

    public Optional<Output<Boolean>> registryOwnerChange() {
        return Optional.ofNullable(this.registryOwnerChange);
    }

    /**
     * The state of the domain.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return The state of the domain.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    private RegisteredDomainRegistrantChangeArgs() {}

    private RegisteredDomainRegistrantChangeArgs(RegisteredDomainRegistrantChangeArgs $) {
        this.accountId = $.accountId;
        this.contactId = $.contactId;
        this.domainId = $.domainId;
        this.extendedAttributes = $.extendedAttributes;
        this.id = $.id;
        this.irtLockLiftedBy = $.irtLockLiftedBy;
        this.registryOwnerChange = $.registryOwnerChange;
        this.state = $.state;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegisteredDomainRegistrantChangeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegisteredDomainRegistrantChangeArgs $;

        public Builder() {
            $ = new RegisteredDomainRegistrantChangeArgs();
        }

        public Builder(RegisteredDomainRegistrantChangeArgs defaults) {
            $ = new RegisteredDomainRegistrantChangeArgs(Objects.requireNonNull(defaults));
        }

        public Builder accountId(@Nullable Output<Integer> accountId) {
            $.accountId = accountId;
            return this;
        }

        public Builder accountId(Integer accountId) {
            return accountId(Output.of(accountId));
        }

        /**
         * @param contactId The ID of the contact to be used for the domain registration. The contact ID can be changed after the domain has been registered. The change will result in a new registrant change this may result in a [60-day lock](https://support.dnsimple.com/articles/icann-60-day-lock-registrant-change/).
         * 
         * @return builder
         * 
         */
        public Builder contactId(@Nullable Output<Integer> contactId) {
            $.contactId = contactId;
            return this;
        }

        /**
         * @param contactId The ID of the contact to be used for the domain registration. The contact ID can be changed after the domain has been registered. The change will result in a new registrant change this may result in a [60-day lock](https://support.dnsimple.com/articles/icann-60-day-lock-registrant-change/).
         * 
         * @return builder
         * 
         */
        public Builder contactId(Integer contactId) {
            return contactId(Output.of(contactId));
        }

        public Builder domainId(@Nullable Output<String> domainId) {
            $.domainId = domainId;
            return this;
        }

        public Builder domainId(String domainId) {
            return domainId(Output.of(domainId));
        }

        /**
         * @param extendedAttributes A map of extended attributes to be set for the domain registration. To see if there are any required extended attributes for any TLD use our [Lists the TLD Extended Attributes API](https://developer.dnsimple.com/v2/tlds/#getTldExtendedAttributes). The values provided in the `extended_attributes` will also be sent when a registrant change is initiated as part of changing the `contact_id`.
         * 
         * @return builder
         * 
         */
        public Builder extendedAttributes(@Nullable Output<Map<String,String>> extendedAttributes) {
            $.extendedAttributes = extendedAttributes;
            return this;
        }

        /**
         * @param extendedAttributes A map of extended attributes to be set for the domain registration. To see if there are any required extended attributes for any TLD use our [Lists the TLD Extended Attributes API](https://developer.dnsimple.com/v2/tlds/#getTldExtendedAttributes). The values provided in the `extended_attributes` will also be sent when a registrant change is initiated as part of changing the `contact_id`.
         * 
         * @return builder
         * 
         */
        public Builder extendedAttributes(Map<String,String> extendedAttributes) {
            return extendedAttributes(Output.of(extendedAttributes));
        }

        /**
         * @param id The ID of this resource.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<Integer> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The ID of this resource.
         * 
         * @return builder
         * 
         */
        public Builder id(Integer id) {
            return id(Output.of(id));
        }

        public Builder irtLockLiftedBy(@Nullable Output<String> irtLockLiftedBy) {
            $.irtLockLiftedBy = irtLockLiftedBy;
            return this;
        }

        public Builder irtLockLiftedBy(String irtLockLiftedBy) {
            return irtLockLiftedBy(Output.of(irtLockLiftedBy));
        }

        public Builder registryOwnerChange(@Nullable Output<Boolean> registryOwnerChange) {
            $.registryOwnerChange = registryOwnerChange;
            return this;
        }

        public Builder registryOwnerChange(Boolean registryOwnerChange) {
            return registryOwnerChange(Output.of(registryOwnerChange));
        }

        /**
         * @param state The state of the domain.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state The state of the domain.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        public RegisteredDomainRegistrantChangeArgs build() {
            return $;
        }
    }

}