package ovh

import (
	"context"
	ovhtypes "github.com/ovh/terraform-provider-ovh/v2/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func OkmsResourceDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "OKMS ID",
			MarkdownDescription: "OKMS ID",
		},
		"kmip_endpoint": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "KMS kmip API endpoint",
			MarkdownDescription: "KMS kmip API endpoint",
		},
		"public_ca": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "KMS public CA (Certificate Authority)",
			MarkdownDescription: "KMS public CA (Certificate Authority)",
		},
		"region": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Region",
			MarkdownDescription: "Region",
		},
		"rest_endpoint": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "KMS rest API endpoint",
			MarkdownDescription: "KMS rest API endpoint",
		},
		"swagger_endpoint": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "KMS rest API swagger UI",
			MarkdownDescription: "KMS rest API swagger UI",
		},
	}

	AppendIamDatasourceSchema(attrs, ctx)
	return schema.Schema{
		Attributes:  attrs,
		Description: "Use this data source to retrieve information about a KMS associated with this account",
	}
}

type OkmsResourceModel struct {
	Iam             IamValue               `tfsdk:"iam" json:"iam"`
	Id              ovhtypes.TfStringValue `tfsdk:"id" json:"id"`
	KmipEndpoint    ovhtypes.TfStringValue `tfsdk:"kmip_endpoint" json:"kmipEndpoint"`
	PublicCa        ovhtypes.TfStringValue `tfsdk:"public_ca" json:"publicCa"`
	Region          ovhtypes.TfStringValue `tfsdk:"region" json:"region"`
	RestEndpoint    ovhtypes.TfStringValue `tfsdk:"rest_endpoint" json:"restEndpoint"`
	SwaggerEndpoint ovhtypes.TfStringValue `tfsdk:"swagger_endpoint" json:"swaggerEndpoint"`
}

func (v *OkmsResourceModel) MergeWith(other *OkmsResourceModel) {

	if (v.Iam.IsUnknown() || v.Iam.IsNull()) && !other.Iam.IsUnknown() {
		v.Iam = other.Iam
	}

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	if (v.KmipEndpoint.IsUnknown() || v.KmipEndpoint.IsNull()) && !other.KmipEndpoint.IsUnknown() {
		v.KmipEndpoint = other.KmipEndpoint
	}

	if (v.PublicCa.IsUnknown() || v.PublicCa.IsNull()) && !other.PublicCa.IsUnknown() {
		v.PublicCa = other.PublicCa
	}

	if (v.Region.IsUnknown() || v.Region.IsNull()) && !other.Region.IsUnknown() {
		v.Region = other.Region
	}

	if (v.RestEndpoint.IsUnknown() || v.RestEndpoint.IsNull()) && !other.RestEndpoint.IsUnknown() {
		v.RestEndpoint = other.RestEndpoint
	}

	if (v.SwaggerEndpoint.IsUnknown() || v.SwaggerEndpoint.IsNull()) && !other.SwaggerEndpoint.IsUnknown() {
		v.SwaggerEndpoint = other.SwaggerEndpoint
	}

}
