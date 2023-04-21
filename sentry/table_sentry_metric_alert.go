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
		Description: "Retrieve information about your metric alerts.",
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
				Description: "The ID of this metric alert.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The metric alert name.",
			},
			{
				Name:        "owner",
				Type:        proto.ColumnType_STRING,
				Description: "Specifies the owner ID of this metric alert rule.",
			},
			{
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the organization the metric alert belongs to.",
			},
			{
				Name:        "project_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the project the metric alert belongs to.",
			},
			{
				Name:        "aggregate",
				Type:        proto.ColumnType_STRING,
				Description: "The aggregation criteria to apply.",
			},
			{
				Name:        "data_set",
				Type:        proto.ColumnType_STRING,
				Description: "The Sentry metric alert category.",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation timestamp of the metric alert.",
			},
			{
				Name:        "environment",
				Type:        proto.ColumnType_STRING,
				Description: "Perform metric alert rule in a specific environment.",
			},
			{
				Name:        "query",
				Type:        proto.ColumnType_STRING,
				Description: "The query filter to apply.",
			},
			{
				Name:        "resolve_threshold",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The value at which the metric alert rule resolves.",
			},
			{
				Name:        "task_uuid",
				Type:        proto.ColumnType_STRING,
				Description: "The UUID of the async task that can be spawned to create the metric alert.",
				Transform:   transform.FromField("TaskUUID"),
			},
			{
				Name:        "threshold_type",
				Type:        proto.ColumnType_INT,
				Description: "The type of threshold.",
			},
			{
				Name:        "time_window",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The period to evaluate the metric alert rule in minutes.",
			},
			{
				Name:        "event_types",
				Type:        proto.ColumnType_JSON,
				Description: "The events type of dataset.",
			},
			{
				Name:        "projects",
				Type:        proto.ColumnType_JSON,
				Description: "The projects for which the metric alert is created.",
			},
			{
				Name:        "triggers",
				Type:        proto.ColumnType_JSON,
				Description: "Represents a metric alert trigger.",
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

type MetricAlert struct {
	sentry.MetricAlert
	OrganizationSlug string
	ProjectSlug      string
}

func listMetricAlerts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*sentry.Project)
	projectSlug := d.EqualsQualString("project_slug")

	// check if the provided projectSlug is not matching with the parentHydrate
	if projectSlug != "" && projectSlug != project.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("sentry_metric_alert.listMetricAlerts", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListCursorParams{}
	for {
		metrics, resp, err := conn.MetricAlerts.List(ctx, *project.Organization.Slug, project.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("sentry_metric_alert.listMetricAlerts", "api_error", err)
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
	orgSlug := d.EqualsQualString("organization_slug")
	projectSlug := d.EqualsQualString("project_slug")
	id := d.EqualsQualString("id")

	// Check if orgSlug or projectSlug or id is empty.
	if orgSlug == "" || projectSlug == "" || id == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("sentry_metric_alert.getMetricAlert", "connection_error", err)
		return nil, err
	}

	metric, _, err := conn.MetricAlerts.Get(ctx, orgSlug, projectSlug, id)
	if err != nil {
		plugin.Logger(ctx).Error("sentry_metric_alert.getMetricAlert", "api_error", err)
		return nil, err
	}

	return MetricAlert{*metric, orgSlug, projectSlug}, nil
}
