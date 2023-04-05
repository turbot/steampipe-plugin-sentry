package sentry

import (
	"context"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryMetricAlert(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_metric_alert",
		Description: "Retrieve information about your metric alert rules.",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listMetricAlerts,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "project_slug",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"organization_slug", "project_slug", "id"}),
			Hydrate:    getMetricAlert,
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
				Name:        "owner",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "project_slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "aggregate",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "data_set",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "environment",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "query",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "resolve_threshold",
				Type:        proto.ColumnType_DOUBLE,
				Description: "",
			},
			{
				Name:        "task_uuid",
				Type:        proto.ColumnType_STRING,
				Description: "",
				Transform:   transform.FromField("TaskUUID"),
			},
			{
				Name:        "threshold_type",
				Type:        proto.ColumnType_INT,
				Description: "",
			},
			{
				Name:        "time_window",
				Type:        proto.ColumnType_DOUBLE,
				Description: "",
			},
			{
				Name:        "event_types",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "projects",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "triggers",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
		},
	}
}

type MetricAlert struct {
	sentry.MetricAlert
	OrganizationSlug string
	ProjectSlug      string
}

func listMetricAlerts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*sentry.Project)
	projSlug := d.EqualsQuals["project_slug"].GetStringValue()

	// check if the provided projSlug is not matching with the parentHydrate
	if projSlug != "" && projSlug != project.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listMetricAlerts", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListCursorParams{}
	for {
		metrics, resp, err := conn.MetricAlerts.List(ctx, *project.Organization.Slug, project.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("listMetricAlerts", "api_error", err)
			return nil, err
		}
		for _, metric := range metrics {
			d.StreamListItem(ctx, MetricAlert{*metric, *project.Organization.Slug, project.Slug})

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

func getMetricAlert(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	orgSlug := d.EqualsQuals["organization_slug"].GetStringValue()
	projSlug := d.EqualsQuals["project_slug"].GetStringValue()
	id := d.EqualsQuals["id"].GetStringValue()

	// Check if orgSlug or projSlug or id is empty.
	if orgSlug == "" || projSlug == "" || id == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getMetricAlert", "connection_error", err)
		return nil, err
	}

	metric, _, err := conn.MetricAlerts.Get(ctx, orgSlug, projSlug, id)
	if err != nil {
		plugin.Logger(ctx).Error("getMetricAlert", "api_error", err)
		return nil, err
	}

	return MetricAlert{*metric, orgSlug, projSlug}, nil
}
