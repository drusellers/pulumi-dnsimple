// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnsimple

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ZoneRecord struct {
	pulumi.CustomResourceState

	Name          pulumi.StringOutput    `pulumi:"name"`
	Priority      pulumi.StringOutput    `pulumi:"priority"`
	QualifiedName pulumi.StringOutput    `pulumi:"qualifiedName"`
	Ttl           pulumi.StringPtrOutput `pulumi:"ttl"`
	Type          pulumi.StringOutput    `pulumi:"type"`
	Value         pulumi.StringOutput    `pulumi:"value"`
	ZoneId        pulumi.StringOutput    `pulumi:"zoneId"`
	ZoneName      pulumi.StringOutput    `pulumi:"zoneName"`
}

// NewZoneRecord registers a new resource with the given unique name, arguments, and options.
func NewZoneRecord(ctx *pulumi.Context,
	name string, args *ZoneRecordArgs, opts ...pulumi.ResourceOption) (*ZoneRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.ZoneName == nil {
		return nil, errors.New("invalid value for required argument 'ZoneName'")
	}
	var resource ZoneRecord
	err := ctx.RegisterResource("dnsimple:index/zoneRecord:ZoneRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneRecord gets an existing ZoneRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneRecordState, opts ...pulumi.ResourceOption) (*ZoneRecord, error) {
	var resource ZoneRecord
	err := ctx.ReadResource("dnsimple:index/zoneRecord:ZoneRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneRecord resources.
type zoneRecordState struct {
	Name          *string `pulumi:"name"`
	Priority      *string `pulumi:"priority"`
	QualifiedName *string `pulumi:"qualifiedName"`
	Ttl           *string `pulumi:"ttl"`
	Type          *string `pulumi:"type"`
	Value         *string `pulumi:"value"`
	ZoneId        *string `pulumi:"zoneId"`
	ZoneName      *string `pulumi:"zoneName"`
}

type ZoneRecordState struct {
	Name          pulumi.StringPtrInput
	Priority      pulumi.StringPtrInput
	QualifiedName pulumi.StringPtrInput
	Ttl           pulumi.StringPtrInput
	Type          pulumi.StringPtrInput
	Value         pulumi.StringPtrInput
	ZoneId        pulumi.StringPtrInput
	ZoneName      pulumi.StringPtrInput
}

func (ZoneRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneRecordState)(nil)).Elem()
}

type zoneRecordArgs struct {
	Name     string  `pulumi:"name"`
	Priority *string `pulumi:"priority"`
	Ttl      *string `pulumi:"ttl"`
	Type     string  `pulumi:"type"`
	Value    string  `pulumi:"value"`
	ZoneName string  `pulumi:"zoneName"`
}

// The set of arguments for constructing a ZoneRecord resource.
type ZoneRecordArgs struct {
	Name     pulumi.StringInput
	Priority pulumi.StringPtrInput
	Ttl      pulumi.StringPtrInput
	Type     pulumi.StringInput
	Value    pulumi.StringInput
	ZoneName pulumi.StringInput
}

func (ZoneRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneRecordArgs)(nil)).Elem()
}

type ZoneRecordInput interface {
	pulumi.Input

	ToZoneRecordOutput() ZoneRecordOutput
	ToZoneRecordOutputWithContext(ctx context.Context) ZoneRecordOutput
}

func (*ZoneRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneRecord)(nil)).Elem()
}

func (i *ZoneRecord) ToZoneRecordOutput() ZoneRecordOutput {
	return i.ToZoneRecordOutputWithContext(context.Background())
}

func (i *ZoneRecord) ToZoneRecordOutputWithContext(ctx context.Context) ZoneRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordOutput)
}

// ZoneRecordArrayInput is an input type that accepts ZoneRecordArray and ZoneRecordArrayOutput values.
// You can construct a concrete instance of `ZoneRecordArrayInput` via:
//
//	ZoneRecordArray{ ZoneRecordArgs{...} }
type ZoneRecordArrayInput interface {
	pulumi.Input

	ToZoneRecordArrayOutput() ZoneRecordArrayOutput
	ToZoneRecordArrayOutputWithContext(context.Context) ZoneRecordArrayOutput
}

type ZoneRecordArray []ZoneRecordInput

func (ZoneRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneRecord)(nil)).Elem()
}

func (i ZoneRecordArray) ToZoneRecordArrayOutput() ZoneRecordArrayOutput {
	return i.ToZoneRecordArrayOutputWithContext(context.Background())
}

func (i ZoneRecordArray) ToZoneRecordArrayOutputWithContext(ctx context.Context) ZoneRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordArrayOutput)
}

// ZoneRecordMapInput is an input type that accepts ZoneRecordMap and ZoneRecordMapOutput values.
// You can construct a concrete instance of `ZoneRecordMapInput` via:
//
//	ZoneRecordMap{ "key": ZoneRecordArgs{...} }
type ZoneRecordMapInput interface {
	pulumi.Input

	ToZoneRecordMapOutput() ZoneRecordMapOutput
	ToZoneRecordMapOutputWithContext(context.Context) ZoneRecordMapOutput
}

type ZoneRecordMap map[string]ZoneRecordInput

func (ZoneRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneRecord)(nil)).Elem()
}

func (i ZoneRecordMap) ToZoneRecordMapOutput() ZoneRecordMapOutput {
	return i.ToZoneRecordMapOutputWithContext(context.Background())
}

func (i ZoneRecordMap) ToZoneRecordMapOutputWithContext(ctx context.Context) ZoneRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordMapOutput)
}

type ZoneRecordOutput struct{ *pulumi.OutputState }

func (ZoneRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneRecord)(nil)).Elem()
}

func (o ZoneRecordOutput) ToZoneRecordOutput() ZoneRecordOutput {
	return o
}

func (o ZoneRecordOutput) ToZoneRecordOutputWithContext(ctx context.Context) ZoneRecordOutput {
	return o
}

func (o ZoneRecordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ZoneRecordOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

func (o ZoneRecordOutput) QualifiedName() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringOutput { return v.QualifiedName }).(pulumi.StringOutput)
}

func (o ZoneRecordOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

func (o ZoneRecordOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ZoneRecordOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func (o ZoneRecordOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

func (o ZoneRecordOutput) ZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringOutput { return v.ZoneName }).(pulumi.StringOutput)
}

type ZoneRecordArrayOutput struct{ *pulumi.OutputState }

func (ZoneRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneRecord)(nil)).Elem()
}

func (o ZoneRecordArrayOutput) ToZoneRecordArrayOutput() ZoneRecordArrayOutput {
	return o
}

func (o ZoneRecordArrayOutput) ToZoneRecordArrayOutputWithContext(ctx context.Context) ZoneRecordArrayOutput {
	return o
}

func (o ZoneRecordArrayOutput) Index(i pulumi.IntInput) ZoneRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZoneRecord {
		return vs[0].([]*ZoneRecord)[vs[1].(int)]
	}).(ZoneRecordOutput)
}

type ZoneRecordMapOutput struct{ *pulumi.OutputState }

func (ZoneRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneRecord)(nil)).Elem()
}

func (o ZoneRecordMapOutput) ToZoneRecordMapOutput() ZoneRecordMapOutput {
	return o
}

func (o ZoneRecordMapOutput) ToZoneRecordMapOutputWithContext(ctx context.Context) ZoneRecordMapOutput {
	return o
}

func (o ZoneRecordMapOutput) MapIndex(k pulumi.StringInput) ZoneRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZoneRecord {
		return vs[0].(map[string]*ZoneRecord)[vs[1].(string)]
	}).(ZoneRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneRecordInput)(nil)).Elem(), &ZoneRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneRecordArrayInput)(nil)).Elem(), ZoneRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneRecordMapInput)(nil)).Elem(), ZoneRecordMap{})
	pulumi.RegisterOutputType(ZoneRecordOutput{})
	pulumi.RegisterOutputType(ZoneRecordArrayOutput{})
	pulumi.RegisterOutputType(ZoneRecordMapOutput{})
}
