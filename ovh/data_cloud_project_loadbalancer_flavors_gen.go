// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ovhtypes "github.com/ovh/terraform-provider-ovh/v2/ovh/types"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudProjectLoadbalancerFlavorsDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"flavors": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Flavor id",
						MarkdownDescription: "Flavor id",
					},
					"name": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Flavor name",
						MarkdownDescription: "Flavor name",
					},
					"region": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Region name",
						MarkdownDescription: "Region name",
					},
				},
				CustomType: CloudProjectLoadbalancerFlavorsType{
					ObjectType: types.ObjectType{
						AttrTypes: CloudProjectLoadbalancerFlavorsValue{}.AttributeTypes(ctx),
					},
				},
			},
			CustomType: ovhtypes.NewTfListNestedType[CloudProjectLoadbalancerFlavorsValue](ctx),
			Computed:   true,
		},
		"region_name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Region name",
			MarkdownDescription: "Region name",
		},
		"service_name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Service name",
			MarkdownDescription: "Service name",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type CloudProjectLoadbalancerFlavorsModel struct {
	Flavors     ovhtypes.TfListNestedValue[CloudProjectLoadbalancerFlavorsValue] `tfsdk:"flavors" json:"flavors"`
	RegionName  ovhtypes.TfStringValue                                           `tfsdk:"region_name" json:"regionName"`
	ServiceName ovhtypes.TfStringValue                                           `tfsdk:"service_name" json:"serviceName"`
}

func (v *CloudProjectLoadbalancerFlavorsModel) MergeWith(other *CloudProjectLoadbalancerFlavorsModel) {
	if (v.Flavors.IsUnknown() || v.Flavors.IsNull()) && !other.Flavors.IsUnknown() {
		v.Flavors = other.Flavors
	}

	if (v.RegionName.IsUnknown() || v.RegionName.IsNull()) && !other.RegionName.IsUnknown() {
		v.RegionName = other.RegionName
	}

	if (v.ServiceName.IsUnknown() || v.ServiceName.IsNull()) && !other.ServiceName.IsUnknown() {
		v.ServiceName = other.ServiceName
	}
}

var _ basetypes.ObjectTypable = CloudProjectLoadbalancerFlavorsType{}

type CloudProjectLoadbalancerFlavorsType struct {
	basetypes.ObjectType
}

func (t CloudProjectLoadbalancerFlavorsType) Equal(o attr.Type) bool {
	other, ok := o.(CloudProjectLoadbalancerFlavorsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t CloudProjectLoadbalancerFlavorsType) String() string {
	return "CloudProjectLoadbalancerFlavorsType"
}

func (t CloudProjectLoadbalancerFlavorsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return nil, diags
	}

	idVal, ok := idAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be ovhtypes.TfStringValue, was: %T`, idAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be ovhtypes.TfStringValue, was: %T`, nameAttribute))
	}

	regionAttribute, ok := attributes["region"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`region is missing from object`)

		return nil, diags
	}

	regionVal, ok := regionAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`region expected to be ovhtypes.TfStringValue, was: %T`, regionAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return CloudProjectLoadbalancerFlavorsValue{
		Id:     idVal,
		Name:   nameVal,
		Region: regionVal,
		state:  attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectLoadbalancerFlavorsValueNull() CloudProjectLoadbalancerFlavorsValue {
	return CloudProjectLoadbalancerFlavorsValue{
		state: attr.ValueStateNull,
	}
}

func NewCloudProjectLoadbalancerFlavorsValueUnknown() CloudProjectLoadbalancerFlavorsValue {
	return CloudProjectLoadbalancerFlavorsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewCloudProjectLoadbalancerFlavorsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (CloudProjectLoadbalancerFlavorsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing CloudProjectLoadbalancerFlavorsValue Attribute Value",
				"While creating a CloudProjectLoadbalancerFlavorsValue value, a missing attribute value was detected. "+
					"A CloudProjectLoadbalancerFlavorsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectLoadbalancerFlavorsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid CloudProjectLoadbalancerFlavorsValue Attribute Type",
				"While creating a CloudProjectLoadbalancerFlavorsValue value, an invalid attribute value was detected. "+
					"A CloudProjectLoadbalancerFlavorsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectLoadbalancerFlavorsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("CloudProjectLoadbalancerFlavorsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra CloudProjectLoadbalancerFlavorsValue Attribute Value",
				"While creating a CloudProjectLoadbalancerFlavorsValue value, an extra attribute value was detected. "+
					"A CloudProjectLoadbalancerFlavorsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra CloudProjectLoadbalancerFlavorsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewCloudProjectLoadbalancerFlavorsValueUnknown(), diags
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewCloudProjectLoadbalancerFlavorsValueUnknown(), diags
	}

	idVal, ok := idAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be ovhtypes.TfStringValue, was: %T`, idAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewCloudProjectLoadbalancerFlavorsValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be ovhtypes.TfStringValue, was: %T`, nameAttribute))
	}

	regionAttribute, ok := attributes["region"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`region is missing from object`)

		return NewCloudProjectLoadbalancerFlavorsValueUnknown(), diags
	}

	regionVal, ok := regionAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`region expected to be ovhtypes.TfStringValue, was: %T`, regionAttribute))
	}

	if diags.HasError() {
		return NewCloudProjectLoadbalancerFlavorsValueUnknown(), diags
	}

	return CloudProjectLoadbalancerFlavorsValue{
		Id:     idVal,
		Name:   nameVal,
		Region: regionVal,
		state:  attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectLoadbalancerFlavorsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) CloudProjectLoadbalancerFlavorsValue {
	object, diags := NewCloudProjectLoadbalancerFlavorsValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewCloudProjectLoadbalancerFlavorsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t CloudProjectLoadbalancerFlavorsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewCloudProjectLoadbalancerFlavorsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewCloudProjectLoadbalancerFlavorsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewCloudProjectLoadbalancerFlavorsValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewCloudProjectLoadbalancerFlavorsValueMust(CloudProjectLoadbalancerFlavorsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t CloudProjectLoadbalancerFlavorsType) ValueType(ctx context.Context) attr.Value {
	return CloudProjectLoadbalancerFlavorsValue{}
}

var _ basetypes.ObjectValuable = CloudProjectLoadbalancerFlavorsValue{}

type CloudProjectLoadbalancerFlavorsValue struct {
	Id     ovhtypes.TfStringValue `tfsdk:"id" json:"id"`
	Name   ovhtypes.TfStringValue `tfsdk:"name" json:"name"`
	Region ovhtypes.TfStringValue `tfsdk:"region" json:"region"`
	state  attr.ValueState
}

func (v *CloudProjectLoadbalancerFlavorsValue) UnmarshalJSON(data []byte) error {
	type JsonCloudProjectLoadbalancerFlavorsValue CloudProjectLoadbalancerFlavorsValue

	var tmp JsonCloudProjectLoadbalancerFlavorsValue
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v.Id = tmp.Id
	v.Name = tmp.Name
	v.Region = tmp.Region

	v.state = attr.ValueStateKnown

	return nil
}

func (v *CloudProjectLoadbalancerFlavorsValue) MergeWith(other *CloudProjectLoadbalancerFlavorsValue) {

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	if (v.Name.IsUnknown() || v.Name.IsNull()) && !other.Name.IsUnknown() {
		v.Name = other.Name
	}

	if (v.Region.IsUnknown() || v.Region.IsNull()) && !other.Region.IsUnknown() {
		v.Region = other.Region
	}

	if (v.state == attr.ValueStateUnknown || v.state == attr.ValueStateNull) && other.state != attr.ValueStateUnknown {
		v.state = other.state
	}
}

func (v CloudProjectLoadbalancerFlavorsValue) Attributes() map[string]attr.Value {
	return map[string]attr.Value{
		"id":     v.Id,
		"name":   v.Name,
		"region": v.Region,
	}
}
func (v CloudProjectLoadbalancerFlavorsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["region"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.Region.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["region"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v CloudProjectLoadbalancerFlavorsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v CloudProjectLoadbalancerFlavorsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v CloudProjectLoadbalancerFlavorsValue) String() string {
	return "CloudProjectLoadbalancerFlavorsValue"
}

func (v CloudProjectLoadbalancerFlavorsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"id":     ovhtypes.TfStringType{},
			"name":   ovhtypes.TfStringType{},
			"region": ovhtypes.TfStringType{},
		},
		map[string]attr.Value{
			"id":     v.Id,
			"name":   v.Name,
			"region": v.Region,
		})

	return objVal, diags
}

func (v CloudProjectLoadbalancerFlavorsValue) Equal(o attr.Value) bool {
	other, ok := o.(CloudProjectLoadbalancerFlavorsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.Region.Equal(other.Region) {
		return false
	}

	return true
}

func (v CloudProjectLoadbalancerFlavorsValue) Type(ctx context.Context) attr.Type {
	return CloudProjectLoadbalancerFlavorsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v CloudProjectLoadbalancerFlavorsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"id":     ovhtypes.TfStringType{},
		"name":   ovhtypes.TfStringType{},
		"region": ovhtypes.TfStringType{},
	}
}
