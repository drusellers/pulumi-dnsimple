// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnsimple

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-dnsimple/sdk/v4/go/dnsimple/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DNSimple contact resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-dnsimple/sdk/v4/go/dnsimple"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dnsimple.NewContact(ctx, "me", &dnsimple.ContactArgs{
//				Address1:         pulumi.String("Level 1, 2 Main St"),
//				Address2:         pulumi.String("Marsfield"),
//				City:             pulumi.String("San Francisco"),
//				Country:          pulumi.String("US"),
//				Email:            pulumi.String("apple@contoso.com"),
//				Fax:              pulumi.String("+1849491024"),
//				FirstName:        pulumi.String("Apple"),
//				JobTitle:         pulumi.String("Manager"),
//				Label:            pulumi.String("Apple Appleseed"),
//				LastName:         pulumi.String("Appleseed"),
//				OrganizationName: pulumi.String("Contoso"),
//				Phone:            pulumi.String("+1401239523"),
//				PostalCode:       pulumi.String("90210"),
//				StateProvince:    pulumi.String("California"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// DNSimple contacts can be imported using their numeric ID. bash
//
// ```sh
//
//	$ pulumi import dnsimple:index/contact:Contact resource_name 5678
//
// ```
//
//	The ID can be found within [DNSimple Contacts API](https://developer.dnsimple.com/v2/contacts/#listContacts). Check out [Authentication](https://developer.dnsimple.com/v2/#authentication) in API Overview for available options. bash curl -u 'EMAIL:PASSWORD' https://api.dnsimple.com/v2/1234/contacts?label_like=example.com | jq {
//
//	"data"[
//
//	{
//
//	"id"1,
//
//	"account_id"1010,
//
//	"label""Default",
//
//	"first_name""First",
//
//	"last_name""User",
//
//	"job_title""CEO",
//
//	"organization_name""Awesome Company",
//
//	"email""first@example.com",
//
//	"phone""+18001234567",
//
//	"fax""+18011234567",
//
//	"address1""Italian Street, 10",
//
//	"address2""",
//
//	"city""Roma",
//
//	"state_province""RM",
//
//	"postal_code""00100",
//
//	"country""IT",
//
//	"created_at""2013-11-08T17:23:15Z",
//
//	"updated_at""2015-01-08T21:30:50Z"
//
//	},
//
//	{
//
//	"id"2,
//
//	"account_id"1010,
//
//	"label""",
//
//	"first_name""Second",
//
//	"last_name""User",
//
//	"job_title""",
//
//	"organization_name""",
//
//	"email""second@example.com",
//
//	"phone""+18881234567",
//
//	"fax""",
//
//	"address1""French Street",
//
//	"address2""c/o Someone",
//
//	"city""Paris",
//
//	"state_province""XY",
//
//	"postal_code""00200",
//
//	"country""FR",
//
//	"created_at""2014-12-06T15:46:18Z",
//
//	"updated_at""2014-12-06T15:46:18Z"
//
//	}
//
//	],
//
//	"pagination"{
//
//	"current_page"1,
//
//	"per_page"30,
//
//	"total_entries"2,
//
//	"total_pages"1
//
//	} }
type Contact struct {
	pulumi.CustomResourceState

	// The account ID for the contact.
	AccountId pulumi.IntOutput `pulumi:"accountId"`
	// Address line 1
	Address1 pulumi.StringOutput `pulumi:"address1"`
	// Address line 2
	Address2 pulumi.StringOutput `pulumi:"address2"`
	// City
	City pulumi.StringOutput `pulumi:"city"`
	// Country
	Country pulumi.StringOutput `pulumi:"country"`
	// Timestamp representing when this contact was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Email
	//
	// # Attributes Reference
	Email pulumi.StringOutput `pulumi:"email"`
	// Fax
	Fax pulumi.StringOutput `pulumi:"fax"`
	// The fax number, normalized.
	FaxNormalized pulumi.StringOutput `pulumi:"faxNormalized"`
	// First name
	FirstName pulumi.StringOutput `pulumi:"firstName"`
	// Job title
	JobTitle pulumi.StringOutput `pulumi:"jobTitle"`
	// Label
	Label pulumi.StringOutput `pulumi:"label"`
	// Last name
	LastName pulumi.StringOutput `pulumi:"lastName"`
	// Organization name
	OrganizationName pulumi.StringOutput `pulumi:"organizationName"`
	// Phone
	Phone pulumi.StringOutput `pulumi:"phone"`
	// The phone number, normalized.
	PhoneNormalized pulumi.StringOutput `pulumi:"phoneNormalized"`
	// Postal code
	PostalCode pulumi.StringOutput `pulumi:"postalCode"`
	// State province
	StateProvince pulumi.StringOutput `pulumi:"stateProvince"`
	// Timestamp representing when this contact was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewContact registers a new resource with the given unique name, arguments, and options.
func NewContact(ctx *pulumi.Context,
	name string, args *ContactArgs, opts ...pulumi.ResourceOption) (*Contact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address1 == nil {
		return nil, errors.New("invalid value for required argument 'Address1'")
	}
	if args.City == nil {
		return nil, errors.New("invalid value for required argument 'City'")
	}
	if args.Country == nil {
		return nil, errors.New("invalid value for required argument 'Country'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.FirstName == nil {
		return nil, errors.New("invalid value for required argument 'FirstName'")
	}
	if args.LastName == nil {
		return nil, errors.New("invalid value for required argument 'LastName'")
	}
	if args.Phone == nil {
		return nil, errors.New("invalid value for required argument 'Phone'")
	}
	if args.PostalCode == nil {
		return nil, errors.New("invalid value for required argument 'PostalCode'")
	}
	if args.StateProvince == nil {
		return nil, errors.New("invalid value for required argument 'StateProvince'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Contact
	err := ctx.RegisterResource("dnsimple:index/contact:Contact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContact gets an existing Contact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactState, opts ...pulumi.ResourceOption) (*Contact, error) {
	var resource Contact
	err := ctx.ReadResource("dnsimple:index/contact:Contact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Contact resources.
type contactState struct {
	// The account ID for the contact.
	AccountId *int `pulumi:"accountId"`
	// Address line 1
	Address1 *string `pulumi:"address1"`
	// Address line 2
	Address2 *string `pulumi:"address2"`
	// City
	City *string `pulumi:"city"`
	// Country
	Country *string `pulumi:"country"`
	// Timestamp representing when this contact was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Email
	//
	// # Attributes Reference
	Email *string `pulumi:"email"`
	// Fax
	Fax *string `pulumi:"fax"`
	// The fax number, normalized.
	FaxNormalized *string `pulumi:"faxNormalized"`
	// First name
	FirstName *string `pulumi:"firstName"`
	// Job title
	JobTitle *string `pulumi:"jobTitle"`
	// Label
	Label *string `pulumi:"label"`
	// Last name
	LastName *string `pulumi:"lastName"`
	// Organization name
	OrganizationName *string `pulumi:"organizationName"`
	// Phone
	Phone *string `pulumi:"phone"`
	// The phone number, normalized.
	PhoneNormalized *string `pulumi:"phoneNormalized"`
	// Postal code
	PostalCode *string `pulumi:"postalCode"`
	// State province
	StateProvince *string `pulumi:"stateProvince"`
	// Timestamp representing when this contact was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ContactState struct {
	// The account ID for the contact.
	AccountId pulumi.IntPtrInput
	// Address line 1
	Address1 pulumi.StringPtrInput
	// Address line 2
	Address2 pulumi.StringPtrInput
	// City
	City pulumi.StringPtrInput
	// Country
	Country pulumi.StringPtrInput
	// Timestamp representing when this contact was created.
	CreatedAt pulumi.StringPtrInput
	// Email
	//
	// # Attributes Reference
	Email pulumi.StringPtrInput
	// Fax
	Fax pulumi.StringPtrInput
	// The fax number, normalized.
	FaxNormalized pulumi.StringPtrInput
	// First name
	FirstName pulumi.StringPtrInput
	// Job title
	JobTitle pulumi.StringPtrInput
	// Label
	Label pulumi.StringPtrInput
	// Last name
	LastName pulumi.StringPtrInput
	// Organization name
	OrganizationName pulumi.StringPtrInput
	// Phone
	Phone pulumi.StringPtrInput
	// The phone number, normalized.
	PhoneNormalized pulumi.StringPtrInput
	// Postal code
	PostalCode pulumi.StringPtrInput
	// State province
	StateProvince pulumi.StringPtrInput
	// Timestamp representing when this contact was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (ContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactState)(nil)).Elem()
}

type contactArgs struct {
	// Address line 1
	Address1 string `pulumi:"address1"`
	// Address line 2
	Address2 *string `pulumi:"address2"`
	// City
	City string `pulumi:"city"`
	// Country
	Country string `pulumi:"country"`
	// Email
	//
	// # Attributes Reference
	Email string `pulumi:"email"`
	// Fax
	Fax *string `pulumi:"fax"`
	// First name
	FirstName string `pulumi:"firstName"`
	// Job title
	JobTitle *string `pulumi:"jobTitle"`
	// Label
	Label *string `pulumi:"label"`
	// Last name
	LastName string `pulumi:"lastName"`
	// Organization name
	OrganizationName *string `pulumi:"organizationName"`
	// Phone
	Phone string `pulumi:"phone"`
	// Postal code
	PostalCode string `pulumi:"postalCode"`
	// State province
	StateProvince string `pulumi:"stateProvince"`
}

// The set of arguments for constructing a Contact resource.
type ContactArgs struct {
	// Address line 1
	Address1 pulumi.StringInput
	// Address line 2
	Address2 pulumi.StringPtrInput
	// City
	City pulumi.StringInput
	// Country
	Country pulumi.StringInput
	// Email
	//
	// # Attributes Reference
	Email pulumi.StringInput
	// Fax
	Fax pulumi.StringPtrInput
	// First name
	FirstName pulumi.StringInput
	// Job title
	JobTitle pulumi.StringPtrInput
	// Label
	Label pulumi.StringPtrInput
	// Last name
	LastName pulumi.StringInput
	// Organization name
	OrganizationName pulumi.StringPtrInput
	// Phone
	Phone pulumi.StringInput
	// Postal code
	PostalCode pulumi.StringInput
	// State province
	StateProvince pulumi.StringInput
}

func (ContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactArgs)(nil)).Elem()
}

type ContactInput interface {
	pulumi.Input

	ToContactOutput() ContactOutput
	ToContactOutputWithContext(ctx context.Context) ContactOutput
}

func (*Contact) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (i *Contact) ToContactOutput() ContactOutput {
	return i.ToContactOutputWithContext(context.Background())
}

func (i *Contact) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput)
}

// ContactArrayInput is an input type that accepts ContactArray and ContactArrayOutput values.
// You can construct a concrete instance of `ContactArrayInput` via:
//
//	ContactArray{ ContactArgs{...} }
type ContactArrayInput interface {
	pulumi.Input

	ToContactArrayOutput() ContactArrayOutput
	ToContactArrayOutputWithContext(context.Context) ContactArrayOutput
}

type ContactArray []ContactInput

func (ContactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Contact)(nil)).Elem()
}

func (i ContactArray) ToContactArrayOutput() ContactArrayOutput {
	return i.ToContactArrayOutputWithContext(context.Background())
}

func (i ContactArray) ToContactArrayOutputWithContext(ctx context.Context) ContactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactArrayOutput)
}

// ContactMapInput is an input type that accepts ContactMap and ContactMapOutput values.
// You can construct a concrete instance of `ContactMapInput` via:
//
//	ContactMap{ "key": ContactArgs{...} }
type ContactMapInput interface {
	pulumi.Input

	ToContactMapOutput() ContactMapOutput
	ToContactMapOutputWithContext(context.Context) ContactMapOutput
}

type ContactMap map[string]ContactInput

func (ContactMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Contact)(nil)).Elem()
}

func (i ContactMap) ToContactMapOutput() ContactMapOutput {
	return i.ToContactMapOutputWithContext(context.Background())
}

func (i ContactMap) ToContactMapOutputWithContext(ctx context.Context) ContactMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactMapOutput)
}

type ContactOutput struct{ *pulumi.OutputState }

func (ContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (o ContactOutput) ToContactOutput() ContactOutput {
	return o
}

func (o ContactOutput) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return o
}

// The account ID for the contact.
func (o ContactOutput) AccountId() pulumi.IntOutput {
	return o.ApplyT(func(v *Contact) pulumi.IntOutput { return v.AccountId }).(pulumi.IntOutput)
}

// Address line 1
func (o ContactOutput) Address1() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Address1 }).(pulumi.StringOutput)
}

// Address line 2
func (o ContactOutput) Address2() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Address2 }).(pulumi.StringOutput)
}

// City
func (o ContactOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.City }).(pulumi.StringOutput)
}

// Country
func (o ContactOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Country }).(pulumi.StringOutput)
}

// Timestamp representing when this contact was created.
func (o ContactOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Email
//
// # Attributes Reference
func (o ContactOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Fax
func (o ContactOutput) Fax() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Fax }).(pulumi.StringOutput)
}

// The fax number, normalized.
func (o ContactOutput) FaxNormalized() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.FaxNormalized }).(pulumi.StringOutput)
}

// First name
func (o ContactOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.FirstName }).(pulumi.StringOutput)
}

// Job title
func (o ContactOutput) JobTitle() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.JobTitle }).(pulumi.StringOutput)
}

// Label
func (o ContactOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Label }).(pulumi.StringOutput)
}

// Last name
func (o ContactOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.LastName }).(pulumi.StringOutput)
}

// Organization name
func (o ContactOutput) OrganizationName() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.OrganizationName }).(pulumi.StringOutput)
}

// Phone
func (o ContactOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Phone }).(pulumi.StringOutput)
}

// The phone number, normalized.
func (o ContactOutput) PhoneNormalized() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.PhoneNormalized }).(pulumi.StringOutput)
}

// Postal code
func (o ContactOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.PostalCode }).(pulumi.StringOutput)
}

// State province
func (o ContactOutput) StateProvince() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.StateProvince }).(pulumi.StringOutput)
}

// Timestamp representing when this contact was updated.
func (o ContactOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ContactArrayOutput struct{ *pulumi.OutputState }

func (ContactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Contact)(nil)).Elem()
}

func (o ContactArrayOutput) ToContactArrayOutput() ContactArrayOutput {
	return o
}

func (o ContactArrayOutput) ToContactArrayOutputWithContext(ctx context.Context) ContactArrayOutput {
	return o
}

func (o ContactArrayOutput) Index(i pulumi.IntInput) ContactOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Contact {
		return vs[0].([]*Contact)[vs[1].(int)]
	}).(ContactOutput)
}

type ContactMapOutput struct{ *pulumi.OutputState }

func (ContactMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Contact)(nil)).Elem()
}

func (o ContactMapOutput) ToContactMapOutput() ContactMapOutput {
	return o
}

func (o ContactMapOutput) ToContactMapOutputWithContext(ctx context.Context) ContactMapOutput {
	return o
}

func (o ContactMapOutput) MapIndex(k pulumi.StringInput) ContactOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Contact {
		return vs[0].(map[string]*Contact)[vs[1].(string)]
	}).(ContactOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContactInput)(nil)).Elem(), &Contact{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactArrayInput)(nil)).Elem(), ContactArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactMapInput)(nil)).Elem(), ContactMap{})
	pulumi.RegisterOutputType(ContactOutput{})
	pulumi.RegisterOutputType(ContactArrayOutput{})
	pulumi.RegisterOutputType(ContactMapOutput{})
}