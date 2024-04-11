package models

import (
	"context"

	"github.com/astronomer/astronomer-terraform-provider/internal/clients/platform"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// OrganizationDataSource describes the data source data model.
type OrganizationDataSource struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	CreatedAt types.String `tfsdk:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at"`
	CreatedBy types.Object `tfsdk:"created_by"`
	UpdatedBy types.Object `tfsdk:"updated_by"`
}

// OrganizationResource describes the resource data model.
type OrganizationResource struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	CreatedAt types.String `tfsdk:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at"`
	CreatedBy types.Object `tfsdk:"created_by"`
	UpdatedBy types.Object `tfsdk:"updated_by"`
}

func (data *OrganizationResource) ReadFromResponse(
	ctx context.Context,
	organization *platform.Organization,
) diag.Diagnostics {
	data.Id = types.StringValue(organization.Id)
	data.Name = types.StringValue(organization.Name)
	data.CreatedAt = types.StringValue(organization.CreatedAt.String())
	data.UpdatedAt = types.StringValue(organization.UpdatedAt.String())
	var diags diag.Diagnostics
	data.CreatedBy, diags = SubjectProfileTypesObject(ctx, organization.CreatedBy)
	if diags.HasError() {
		return diags
	}
	data.UpdatedBy, diags = SubjectProfileTypesObject(ctx, organization.CreatedBy)
	if diags.HasError() {
		return diags
	}

	return nil
}

func (data *OrganizationDataSource) ReadFromResponse(
	ctx context.Context,
	organization *platform.Organization,
) diag.Diagnostics {
	data.Id = types.StringValue(organization.Id)
	data.Name = types.StringValue(organization.Name)
	data.CreatedAt = types.StringValue(organization.CreatedAt.String())
	data.UpdatedAt = types.StringValue(organization.UpdatedAt.String())
	var diags diag.Diagnostics
	data.CreatedBy, diags = SubjectProfileTypesObject(ctx, organization.CreatedBy)
	if diags.HasError() {
		return diags
	}
	data.UpdatedBy, diags = SubjectProfileTypesObject(ctx, organization.UpdatedBy)
	if diags.HasError() {
		return diags
	}

	return nil
}