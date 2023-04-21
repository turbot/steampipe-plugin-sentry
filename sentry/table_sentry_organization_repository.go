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
				Description: "The ID of this repository.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the repository.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the repository.",
			},
			{
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the organization the repository belongs to.",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation timestamp of the repository.",
			},
			{
				Name:        "external_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the repository.",
			},
			{
				Name:        "integration_id",
				Type:        proto.ColumnType_STRING,
				Description: "The organization integration ID for repository.",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "The url of the repository.",
			},
			{
				Name:        "provider",
				Type:        proto.ColumnType_JSON,
				Description: "The repository provider.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
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
	orgSlug := d.EqualsQualString("organization_slug")

	// check if the provided orgSlug is not matching with the parentHydrate
	if orgSlug != "" && orgSlug != *org.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("sentry_organization_repository.listOrganizationRepositories", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListOrganizationRepositoriesParams{}
	if d.EqualsQuals["status"] != nil {
		params.Status = d.EqualsQualString("status")
	}

	for {
		repositoryList, resp, err := conn.OrganizationRepositories.List(ctx, *org.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("sentry_organization_repository.listOrganizationRepositories", "api_error", err)
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
