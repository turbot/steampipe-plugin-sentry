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
		Description: "Retrieve information about your keys.",
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
				Description: "",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "is_active",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "secret",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "project_slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "label",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "project_id",
				Type:        proto.ColumnType_STRING,
				Description: "",
				Transform:   transform.FromField("ProjectID"),
			},
			{
				Name:        "public",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "dsn",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Transform:   transform.FromField("DSN"),
			},
			{
				Name:        "rate_limit",
				Type:        proto.ColumnType_JSON,
				Description: "",
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
	projSlug := d.EqualsQuals["project_slug"].GetStringValue()

	// check if the provided projSlug is not matching with the parentHydrate
	if projSlug != "" && projSlug != project.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listKeys", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListCursorParams{}
	for {
		keys, resp, err := conn.ProjectKeys.List(ctx, *project.Organization.Slug, project.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("listKeys", "api_error", err)
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
