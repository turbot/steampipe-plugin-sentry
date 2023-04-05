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
		Description: "Retrieve information about your issue alert rules.",
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
				Name:        "action_match",
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
				Name:        "filter_match",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "frequency",
				Type:        proto.ColumnType_INT,
				Description: "",
			},
			{
				Name:        "task_uuid",
				Type:        proto.ColumnType_STRING,
				Description: "",
				Transform:   transform.FromField("TaskUUID"),
			},
			{
				Name:        "actions",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "conditions",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "created_by",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "filters",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "projects",
				Type:        proto.ColumnType_JSON,
				Description: "",
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
	projSlug := d.EqualsQuals["project_slug"].GetStringValue()

	// check if the provided projSlug is not matching with the parentHydrate
	if projSlug != "" && projSlug != project.Slug {
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
	projSlug := d.EqualsQuals["project_slug"].GetStringValue()
	id := d.EqualsQuals["id"].GetStringValue()

	// Check if orgSlug or projSlug or id is empty.
	if orgSlug == "" || projSlug == "" || id == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getIssueAlert", "connection_error", err)
		return nil, err
	}

	issue, _, err := conn.IssueAlerts.Get(ctx, orgSlug, projSlug, id)
	if err != nil {
		plugin.Logger(ctx).Error("getIssueAlert", "api_error", err)
		return nil, err
	}

	return IssueAlert{*issue, orgSlug, projSlug}, nil
}
