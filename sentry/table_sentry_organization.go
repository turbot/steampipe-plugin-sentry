package sentry

import (
	"context"

	"github.com/atlassian/go-sentry-api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryOrganizations(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_organization",
		Description: "Retrieve information about your organizations.",
		List: &plugin.ListConfig{
			Hydrate: listOrganizations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("slug"),
			Hydrate:    getOrganization,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "pending_access_request",
				Type:        proto.ColumnType_INT,
				Description: "",
			},
			{
				Name:        "slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "is_early_adopter",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "features",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "quota",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "teams",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "users",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Hydrate:     listOrganizationUsers,
				Transform:   transform.FromValue(),
			},
		},
	}
}

func listOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listOrganizations", "connection_error", err)
		return nil, err
	}

	orgList, _, err := conn.GetOrganizations()
	if err != nil {
		plugin.Logger(ctx).Error("listOrganizations", "api_error", err)
		return nil, err
	}

	for _, org := range orgList {
		d.StreamListItem(ctx, org)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	slug := d.EqualsQuals["slug"].GetStringValue()

	// Check if slug is empty.
	if slug == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganization", "connection_error", err)
		return nil, err
	}

	org, err := conn.GetOrganization(slug)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganization", "api_error", err)
		return nil, err
	}

	return org, nil
}

func listOrganizationUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	slug := *h.Item.(sentry.Organization).Slug

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listOrganizationUsers", "connection_error", err)
		return nil, err
	}

	users, err := conn.ListOrganizationUsers(slug)
	if err != nil {
		plugin.Logger(ctx).Error("listOrganizationUsers", "api_error", err)
		return nil, err
	}

	return users, nil
}
