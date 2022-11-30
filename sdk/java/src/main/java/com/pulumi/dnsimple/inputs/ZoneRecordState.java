// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.dnsimple.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ZoneRecordState extends com.pulumi.resources.ResourceArgs {

    public static final ZoneRecordState Empty = new ZoneRecordState();

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="priority")
    private @Nullable Output<String> priority;

    public Optional<Output<String>> priority() {
        return Optional.ofNullable(this.priority);
    }

    @Import(name="qualifiedName")
    private @Nullable Output<String> qualifiedName;

    public Optional<Output<String>> qualifiedName() {
        return Optional.ofNullable(this.qualifiedName);
    }

    @Import(name="ttl")
    private @Nullable Output<String> ttl;

    public Optional<Output<String>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    @Import(name="value")
    private @Nullable Output<String> value;

    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    @Import(name="zoneName")
    private @Nullable Output<String> zoneName;

    public Optional<Output<String>> zoneName() {
        return Optional.ofNullable(this.zoneName);
    }

    private ZoneRecordState() {}

    private ZoneRecordState(ZoneRecordState $) {
        this.name = $.name;
        this.priority = $.priority;
        this.qualifiedName = $.qualifiedName;
        this.ttl = $.ttl;
        this.type = $.type;
        this.value = $.value;
        this.zoneId = $.zoneId;
        this.zoneName = $.zoneName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ZoneRecordState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ZoneRecordState $;

        public Builder() {
            $ = new ZoneRecordState();
        }

        public Builder(ZoneRecordState defaults) {
            $ = new ZoneRecordState(Objects.requireNonNull(defaults));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder priority(@Nullable Output<String> priority) {
            $.priority = priority;
            return this;
        }

        public Builder priority(String priority) {
            return priority(Output.of(priority));
        }

        public Builder qualifiedName(@Nullable Output<String> qualifiedName) {
            $.qualifiedName = qualifiedName;
            return this;
        }

        public Builder qualifiedName(String qualifiedName) {
            return qualifiedName(Output.of(qualifiedName));
        }

        public Builder ttl(@Nullable Output<String> ttl) {
            $.ttl = ttl;
            return this;
        }

        public Builder ttl(String ttl) {
            return ttl(Output.of(ttl));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        public Builder value(String value) {
            return value(Output.of(value));
        }

        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public Builder zoneName(@Nullable Output<String> zoneName) {
            $.zoneName = zoneName;
            return this;
        }

        public Builder zoneName(String zoneName) {
            return zoneName(Output.of(zoneName));
        }

        public ZoneRecordState build() {
            return $;
        }
    }

}