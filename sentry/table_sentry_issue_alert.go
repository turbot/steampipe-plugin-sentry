package sentry

import (
	"context"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryIssueAlert(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_issue_alert",
		Description: "Retrieve information about your issue alerts.",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listIssueAlerts,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "project_slug",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"organization_slug", "project_slug", "id"}),
			Hydrate:    getIssueAlert,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the issue alert.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The issue alert name.",
			},
			{
				Name:        "owner",
				Type:        proto.ColumnType_STRING,
				Description: "The owner of the issue alert.",
			},
			{
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the organization the issue alert belongs to.",
			},
			{
				Name:        "project_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the project the issue alert belongs to.",
			},
			{
				Name:        "action_match",
				Type:        proto.ColumnType_STRING,
				Description: "Trigger actions when an event is captured by Sentry and any or all of the specified conditions happen.",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation timestamp of the issue alert.",
			},
			{
				Name:        "environment",
				Type:        proto.ColumnType_STRING,
				Description: "Perform issue alert in a specific environment.",
			},
			{
				Name:        "filter_match",
				Type:        proto.ColumnType_STRING,
				Description: "Trigger actions if all, any, or none of the specified filters match.",
			},
			{
				Name:        "frequency",
				Type:        proto.ColumnType_INT,
				Description: "Perform actions at most once every X minutes for this issue. Defaults to 30.",
			},
			{
				Name:        "task_uuid",
				Type:        proto.ColumnType_STRING,
				Description: "The UUID of the async task that can be spawned to create the issue alert.",
				Transform:   transform.FromField("TaskUUID"),
			},
			{
				Name:        "actions",
				Type:        proto.ColumnType_JSON,
				Description: "List of actions.",
			},
			{
				Name:        "conditions",
				Type:        proto.ColumnType_JSON,
				Description: "List of conditions.",
			},
			{
				Name:        "created_by",
				Type:        proto.ColumnType_JSON,
				Description: "The rule creator who created the issue alert.",
			},
			{
				Name:        "filters",
				Type:        proto.ColumnType_JSON,
				Description: "List of filters.",
			},
			{
				Name:        "projects",
				Type:        proto.ColumnType_JSON,
				Description: "The projects for which the issue alert is created.",
			},
		},
	}
}

type IssueAlert struct {
	sentry.IssueAlert
	OrganizationSlug string
	ProjectSlug      string
}

func listIssueAlerts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*sentry.Project)
	projectSlug := d.EqualsQuals["project_slug"].GetStringValue()

	// check if the provided projectSlug is not matching with the parentHydrate
	if projectSlug != "" && projectSlug != project.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listIssueAlerts", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListCursorParams{}
	for {
		issues, resp, err := conn.IssueAlerts.List(ctx, *project.Organization.Slug, project.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("listIssueAlerts", "api_error", err)
			return nil, err
		}
		for _, issue := range issues {
			d.StreamListItem(ctx, IssueAlert{*issue, *project.Organization.Slug, project.Slug})

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

func getIssueAlert(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	orgSlug := d.EqualsQuals["organization_slug"].GetStringValue()
	projectSlug := d.EqualsQuals["project_slug"].GetStringValue()
	id := d.EqualsQuals["id"].GetStringValue()

	// Check if orgSlug or projectSlug or id is empty.
	if orgSlug == "" || projectSlug == "" || id == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getIssueAlert", "connection_error", err)
		return nil, err
	}

	issue, _, err := conn.IssueAlerts.Get(ctx, orgSlug, projectSlug, id)
	if err != nil {
		plugin.Logger(ctx).Error("getIssueAlert", "api_error", err)
		return nil, err
	}

	return IssueAlert{*issue, orgSlug, projectSlug}, nil
}
