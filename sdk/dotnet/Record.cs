// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Dnsimple
{
    /// <summary>
    /// Provides a DNSimple record resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-dnsimple/blob/master/website/docs/r/record.html.markdown.
    /// </summary>
    public partial class Record : Pulumi.CustomResource
    {
        /// <summary>
        /// The domain to add the record to
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The domain ID of the record
        /// </summary>
        [Output("domainId")]
        public Output<string> DomainId { get; private set; } = null!;

        /// <summary>
        /// The FQDN of the record
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The name of the record
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The priority of the record - only useful for some record types
        /// </summary>
        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

        /// <summary>
        /// The TTL of the record
        /// </summary>
        [Output("ttl")]
        public Output<string?> Ttl { get; private set; } = null!;

        /// <summary>
        /// The type of the record
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The value of the record
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Record resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Record(string name, RecordArgs args, CustomResourceOptions? options = null)
            : base("dnsimple:index/record:Record", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Record(string name, Input<string> id, RecordState? state = null, CustomResourceOptions? options = null)
            : base("dnsimple:index/record:Record", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Record resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Record Get(string name, Input<string> id, RecordState? state = null, CustomResourceOptions? options = null)
        {
            return new Record(name, id, state, options);
        }
    }

    public sealed class RecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to add the record to
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The name of the record
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The priority of the record - only useful for some record types
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// The TTL of the record
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// The type of the record
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The value of the record
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public RecordArgs()
        {
        }
    }

    public sealed class RecordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to add the record to
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The domain ID of the record
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// The FQDN of the record
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The name of the record
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority of the record - only useful for some record types
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// The TTL of the record
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// The type of the record
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The value of the record
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public RecordState()
        {
        }
    }
}
