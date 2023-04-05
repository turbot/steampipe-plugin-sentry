package sentry

import (
	"context"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryOrganizationRepository(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_organization_repository",
		Description: "Retrieve information about your organization repositories.",
		List: &plugin.ListConfig{
			ParentHydrate: listOrganizations,
			Hydrate:       listOrganizationRepositories,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "organization_slug",
					Require: plugin.Optional,
				},
				{
					Name:    "status",
					Require: plugin.Optional,
				},
			},
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
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "external_slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "integration_id",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "provider",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
		},
	}
}

type OrganizationRepository struct {
	sentry.OrganizationRepository
	OrganizationSlug string
}

func listOrganizationRepositories(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	org := h.Item.(*sentry.Organization)
	orgSlug := d.EqualsQuals["organization_slug"].GetStringValue()

	// check if the provided orgSlug is not matching with the parentHydrate
	if orgSlug != "" && orgSlug != *org.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listOrganizationRepositories", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListOrganizationRepositoriesParams{}
	if d.EqualsQuals["status"] != nil {
		params.Status = d.EqualsQuals["status"].GetStringValue()
	}

	for {
		repositoryList, resp, err := conn.OrganizationRepositories.List(ctx, *org.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("listOrganizationRepositories", "api_error", err)
			return nil, err
		}
		for _, repository := range repositoryList {
			d.StreamListItem(ctx, OrganizationRepository{*repository, *org.Slug})

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if resp.Cursor != "" {
			params.Cursor = resp.Cursor
		} else {
			break
		}
	}

	return nil, nil
}
