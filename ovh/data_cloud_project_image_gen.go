// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"
	ovhtypes "github.com/ovh/terraform-provider-ovh/v2/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudProjectImageDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"creation_date": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Image creation date",
			MarkdownDescription: "Image creation date",
		},
		"flavor_type": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Image usable only for this type of flavor if not null",
			MarkdownDescription: "Image usable only for this type of flavor if not null",
		},
		"id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Image id",
			MarkdownDescription: "Image id",
		},
		"image_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Image ID",
			MarkdownDescription: "Image ID",
		},
		"min_disk": schema.Int64Attribute{
			CustomType:          ovhtypes.TfInt64Type{},
			Computed:            true,
			Description:         "Minimum disks required to use image",
			MarkdownDescription: "Minimum disks required to use image",
		},
		"min_ram": schema.Int64Attribute{
			CustomType:          ovhtypes.TfInt64Type{},
			Computed:            true,
			Description:         "Minimum RAM required to use image",
			MarkdownDescription: "Minimum RAM required to use image",
		},
		"name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Image name",
			MarkdownDescription: "Image name",
		},
		"plan_code": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Order plan code",
			MarkdownDescription: "Order plan code",
		},
		"region": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Image region",
			MarkdownDescription: "Image region",
		},
		"service_name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Service name",
			MarkdownDescription: "Service name",
		},
		"size": schema.NumberAttribute{
			CustomType:          ovhtypes.TfNumberType{},
			Computed:            true,
			Description:         "Image size (in GiB)",
			MarkdownDescription: "Image size (in GiB)",
		},
		"status": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Image status",
			MarkdownDescription: "Image status",
		},
		"tags": schema.ListAttribute{
			CustomType:          ovhtypes.NewTfListNestedType[ovhtypes.TfStringValue](ctx),
			Computed:            true,
			Description:         "Tags about the image",
			MarkdownDescription: "Tags about the image",
		},
		"type": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Image type",
			MarkdownDescription: "Image type",
		},
		"user": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "User to connect with",
			MarkdownDescription: "User to connect with",
		},
		"visibility": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Image visibility",
			MarkdownDescription: "Image visibility",
		},
	}

	return schema.Schema{
		Description: "Get image",
		Attributes:  attrs,
	}
}

type CloudProjectImageModel struct {
	CreationDate ovhtypes.TfStringValue                             `tfsdk:"creation_date" json:"creationDate"`
	FlavorType   ovhtypes.TfStringValue                             `tfsdk:"flavor_type" json:"flavorType"`
	Id           ovhtypes.TfStringValue                             `tfsdk:"id" json:"id"`
	ImageId      ovhtypes.TfStringValue                             `tfsdk:"image_id" json:"imageId"`
	MinDisk      ovhtypes.TfInt64Value                              `tfsdk:"min_disk" json:"minDisk"`
	MinRam       ovhtypes.TfInt64Value                              `tfsdk:"min_ram" json:"minRam"`
	Name         ovhtypes.TfStringValue                             `tfsdk:"name" json:"name"`
	PlanCode     ovhtypes.TfStringValue                             `tfsdk:"plan_code" json:"planCode"`
	Region       ovhtypes.TfStringValue                             `tfsdk:"region" json:"region"`
	ServiceName  ovhtypes.TfStringValue                             `tfsdk:"service_name" json:"serviceName"`
	Size         ovhtypes.TfNumberValue                             `tfsdk:"size" json:"size"`
	Status       ovhtypes.TfStringValue                             `tfsdk:"status" json:"status"`
	Tags         ovhtypes.TfListNestedValue[ovhtypes.TfStringValue] `tfsdk:"tags" json:"tags"`
	Type         ovhtypes.TfStringValue                             `tfsdk:"type" json:"type"`
	User         ovhtypes.TfStringValue                             `tfsdk:"user" json:"user"`
	Visibility   ovhtypes.TfStringValue                             `tfsdk:"visibility" json:"visibility"`
}

func (v *CloudProjectImageModel) MergeWith(other *CloudProjectImageModel) {

	if (v.CreationDate.IsUnknown() || v.CreationDate.IsNull()) && !other.CreationDate.IsUnknown() {
		v.CreationDate = other.CreationDate
	}

	if (v.FlavorType.IsUnknown() || v.FlavorType.IsNull()) && !other.FlavorType.IsUnknown() {
		v.FlavorType = other.FlavorType
	}

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	if (v.ImageId.IsUnknown() || v.ImageId.IsNull()) && !other.ImageId.IsUnknown() {
		v.ImageId = other.ImageId
	}

	if (v.MinDisk.IsUnknown() || v.MinDisk.IsNull()) && !other.MinDisk.IsUnknown() {
		v.MinDisk = other.MinDisk
	}

	if (v.MinRam.IsUnknown() || v.MinRam.IsNull()) && !other.MinRam.IsUnknown() {
		v.MinRam = other.MinRam
	}

	if (v.Name.IsUnknown() || v.Name.IsNull()) && !other.Name.IsUnknown() {
		v.Name = other.Name
	}

	if (v.PlanCode.IsUnknown() || v.PlanCode.IsNull()) && !other.PlanCode.IsUnknown() {
		v.PlanCode = other.PlanCode
	}

	if (v.Region.IsUnknown() || v.Region.IsNull()) && !other.Region.IsUnknown() {
		v.Region = other.Region
	}

	if (v.ServiceName.IsUnknown() || v.ServiceName.IsNull()) && !other.ServiceName.IsUnknown() {
		v.ServiceName = other.ServiceName
	}

	if (v.Size.IsUnknown() || v.Size.IsNull()) && !other.Size.IsUnknown() {
		v.Size = other.Size
	}

	if (v.Status.IsUnknown() || v.Status.IsNull()) && !other.Status.IsUnknown() {
		v.Status = other.Status
	}

	if (v.Tags.IsUnknown() || v.Tags.IsNull()) && !other.Tags.IsUnknown() {
		v.Tags = other.Tags
	}

	if (v.Type.IsUnknown() || v.Type.IsNull()) && !other.Type.IsUnknown() {
		v.Type = other.Type
	}

	if (v.User.IsUnknown() || v.User.IsNull()) && !other.User.IsUnknown() {
		v.User = other.User
	}

	if (v.Visibility.IsUnknown() || v.Visibility.IsNull()) && !other.Visibility.IsUnknown() {
		v.Visibility = other.Visibility
	}

}
