// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.dnsimple;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.dnsimple.EmailForwardArgs;
import com.pulumi.dnsimple.Utilities;
import com.pulumi.dnsimple.inputs.EmailForwardState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a DNSimple email forward resource.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.dnsimple.EmailForward;
 * import com.pulumi.dnsimple.EmailForwardArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         // Add an email forwarding rule to the domain
 *         var foobar = new EmailForward("foobar", EmailForwardArgs.builder()
 *             .domain(dnsimpleDomain)
 *             .aliasName("sales")
 *             .destinationEmail("jane.doe{@literal @}example.com")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="dnsimple:index/emailForward:EmailForward")
public class EmailForward extends com.pulumi.resources.CustomResource {
    /**
     * The source email address on the domain
     * 
     */
    @Export(name="aliasEmail", refs={String.class}, tree="[0]")
    private Output<String> aliasEmail;

    /**
     * @return The source email address on the domain
     * 
     */
    public Output<String> aliasEmail() {
        return this.aliasEmail;
    }
    /**
     * The name part (the part before the{@literal @}) of the source email address on the domain
     * 
     */
    @Export(name="aliasName", refs={String.class}, tree="[0]")
    private Output<String> aliasName;

    /**
     * @return The name part (the part before the{@literal @}) of the source email address on the domain
     * 
     */
    public Output<String> aliasName() {
        return this.aliasName;
    }
    /**
     * The destination email address on another domain
     * 
     */
    @Export(name="destinationEmail", refs={String.class}, tree="[0]")
    private Output<String> destinationEmail;

    /**
     * @return The destination email address on another domain
     * 
     */
    public Output<String> destinationEmail() {
        return this.destinationEmail;
    }
    /**
     * The domain to add the email forwarding rule to
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output<String> domain;

    /**
     * @return The domain to add the email forwarding rule to
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EmailForward(String name) {
        this(name, EmailForwardArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EmailForward(String name, EmailForwardArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EmailForward(String name, EmailForwardArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("dnsimple:index/emailForward:EmailForward", name, args == null ? EmailForwardArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EmailForward(String name, Output<String> id, @Nullable EmailForwardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("dnsimple:index/emailForward:EmailForward", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static EmailForward get(String name, Output<String> id, @Nullable EmailForwardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EmailForward(name, id, state, options);
    }
}
