package sentry

import (
	"context"

	"github.com/atlassian/go-sentry-api"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryProjects(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_project",
		Description: "Retrieve information about your projects.",
		List: &plugin.ListConfig{
			Hydrate: listProjects,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"slug", "org_slug"}),
			Hydrate:    getProject,
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
				Name:        "slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "org_slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
				Transform:   transform.FromField("Organization").Transform(orgToSlug),
			},
			{
				Name:        "default_environment",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "color",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "is_public",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "call_sign",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "first_event",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "is_bookmarked",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "call_sign_reviewed",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "digest_min_delay",
				Type:        proto.ColumnType_INT,
				Description: "",
			},
			{
				Name:        "digest_max_delay",
				Type:        proto.ColumnType_INT,
				Description: "",
			},
			{
				Name:        "verify_ssl",
				Type:        proto.ColumnType_BOOL,
				Description: "",
				Transform:   transform.FromField("VerifySSL"),
			},
			{
				Name:        "options",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "organization",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "platforms",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "plugins",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "team",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
		},
	}
}

func listProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listProjects", "connection_error", err)
		return nil, err
	}

	projects, _, err := conn.GetProjects()
	if err != nil {
		plugin.Logger(ctx).Error("listProjects", "api_error", err)
		return nil, err
	}

	for _, project := range projects {
		d.StreamListItem(ctx, project)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	slug := d.EqualsQuals["slug"].GetStringValue()
	org_slug := d.EqualsQuals["org_slug"].GetStringValue()

	// Check if slug or org_slug is empty.
	if slug == "" || org_slug == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getProject", "connection_error", err)
		return nil, err
	}

	input := sentry.Organization{}
	input.Slug = types.String(org_slug)

	org, err := conn.GetProject(input, slug)
	if err != nil {
		plugin.Logger(ctx).Error("getProject", "api_error", err)
		return nil, err
	}

	return org, nil
}

func orgToSlug(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	return d.Value.(*sentry.Organization).Slug, nil
}
