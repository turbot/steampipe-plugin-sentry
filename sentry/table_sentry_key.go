package sentry

import (
	"context"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_key",
		Description: "Retrieve information about your project keys.",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listKeys,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "project_slug",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the key.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the key.",
			},
			{
				Name:        "is_active",
				Type:        proto.ColumnType_BOOL,
				Description: "Flag indicating the key is active.",
			},
			{
				Name:        "secret",
				Type:        proto.ColumnType_STRING,
				Description: "Secret key portion of the client key.",
			},
			{
				Name:        "project_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the project the key belongs to.",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation timestamp for the key.",
			},
			{
				Name:        "label",
				Type:        proto.ColumnType_STRING,
				Description: "The label of the key.",
			},
			{
				Name:        "project_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the project the keys belong to.",
				Transform:   transform.FromField("ProjectID"),
			},
			{
				Name:        "public",
				Type:        proto.ColumnType_STRING,
				Description: "Public key portion of the client key.",
			},
			{
				Name:        "dsn",
				Type:        proto.ColumnType_JSON,
				Description: "DSN for the key.",
				Transform:   transform.FromField("DSN"),
			},
			{
				Name:        "rate_limit",
				Type:        proto.ColumnType_JSON,
				Description: "Represents a project key's rate limit.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("name"),
			},
		},
	}
}

type ProjectKey struct {
	sentry.ProjectKey
	ProjectSlug string
}

func listKeys(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*sentry.Project)
	projectSlug := d.EqualsQuals["project_slug"].GetStringValue()

	// check if the provided projectSlug is not matching with the parentHydrate
	if projectSlug != "" && projectSlug != project.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("sentry_key.listKeys", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListCursorParams{}
	for {
		keys, resp, err := conn.ProjectKeys.List(ctx, *project.Organization.Slug, project.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("sentry_key.listKeys", "api_error", err)
			return nil, err
		}
		for _, key := range keys {
			d.StreamListItem(ctx, ProjectKey{*key, project.Slug})

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
