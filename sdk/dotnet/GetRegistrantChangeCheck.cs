// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DNSimple
{
    public static class GetRegistrantChangeCheck
    {
        /// <summary>
        /// Get information on the requirements of a registrant change.
        /// 
        /// &gt; **Note:** The registrant change API is currently in developer preview and is subject to change.
        /// 
        /// Get registrant change requirements for the `dnsimple.com` domain and the contact with ID `1234`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using DNSimple = Pulumi.DNSimple;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = DNSimple.GetRegistrantChangeCheck.Invoke(new()
        ///     {
        ///         ContactId = "1234",
        ///         DomainId = "dnsimple.com",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// The following arguments are supported:
        /// 
        /// * `domain_id` - (Required) The name or ID of the domain.
        /// * `contact_id` - (Required) The ID of the contact you are planning to change to.
        /// 
        /// The following additional attributes are exported:
        /// 
        /// * `contact_id` - The ID of the contact you are planning to change to.
        /// * `domain_id` - The name or ID of the domain.
        /// * `extended_attributes` - (List) A list of extended attributes that are required for the registrant change. (see below for nested schema)
        /// * `registry_owner_change` - (Boolean) Whether the registrant change is going to result in an owner change at the registry.
        /// 
        /// &lt;a id="nestedblock--extended_attributes"&gt;&lt;/a&gt;
        /// </summary>
        public static Task<GetRegistrantChangeCheckResult> InvokeAsync(GetRegistrantChangeCheckArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegistrantChangeCheckResult>("dnsimple:index/getRegistrantChangeCheck:getRegistrantChangeCheck", args ?? new GetRegistrantChangeCheckArgs(), options.WithDefaults());

        /// <summary>
        /// Get information on the requirements of a registrant change.
        /// 
        /// &gt; **Note:** The registrant change API is currently in developer preview and is subject to change.
        /// 
        /// Get registrant change requirements for the `dnsimple.com` domain and the contact with ID `1234`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using DNSimple = Pulumi.DNSimple;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = DNSimple.GetRegistrantChangeCheck.Invoke(new()
        ///     {
        ///         ContactId = "1234",
        ///         DomainId = "dnsimple.com",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// The following arguments are supported:
        /// 
        /// * `domain_id` - (Required) The name or ID of the domain.
        /// * `contact_id` - (Required) The ID of the contact you are planning to change to.
        /// 
        /// The following additional attributes are exported:
        /// 
        /// * `contact_id` - The ID of the contact you are planning to change to.
        /// * `domain_id` - The name or ID of the domain.
        /// * `extended_attributes` - (List) A list of extended attributes that are required for the registrant change. (see below for nested schema)
        /// * `registry_owner_change` - (Boolean) Whether the registrant change is going to result in an owner change at the registry.
        /// 
        /// &lt;a id="nestedblock--extended_attributes"&gt;&lt;/a&gt;
        /// </summary>
        public static Output<GetRegistrantChangeCheckResult> Invoke(GetRegistrantChangeCheckInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegistrantChangeCheckResult>("dnsimple:index/getRegistrantChangeCheck:getRegistrantChangeCheck", args ?? new GetRegistrantChangeCheckInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistrantChangeCheckArgs : global::Pulumi.InvokeArgs
    {
        [Input("contactId", required: true)]
        public string ContactId { get; set; } = null!;

        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        public GetRegistrantChangeCheckArgs()
        {
        }
        public static new GetRegistrantChangeCheckArgs Empty => new GetRegistrantChangeCheckArgs();
    }

    public sealed class GetRegistrantChangeCheckInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("contactId", required: true)]
        public Input<string> ContactId { get; set; } = null!;

        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        public GetRegistrantChangeCheckInvokeArgs()
        {
        }
        public static new GetRegistrantChangeCheckInvokeArgs Empty => new GetRegistrantChangeCheckInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegistrantChangeCheckResult
    {
        public readonly string ContactId;
        public readonly string DomainId;
        public readonly ImmutableArray<Outputs.GetRegistrantChangeCheckExtendedAttributeResult> ExtendedAttributes;
        public readonly string Id;
        public readonly bool RegistryOwnerChange;

        [OutputConstructor]
        private GetRegistrantChangeCheckResult(
            string contactId,

            string domainId,

            ImmutableArray<Outputs.GetRegistrantChangeCheckExtendedAttributeResult> extendedAttributes,

            string id,

            bool registryOwnerChange)
        {
            ContactId = contactId;
            DomainId = domainId;
            ExtendedAttributes = extendedAttributes;
            Id = id;
            RegistryOwnerChange = registryOwnerChange;
        }
    }
}