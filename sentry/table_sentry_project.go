package sentry

import (
	"context"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_project",
		Description: "Retrieve information about your projects.",
		List: &plugin.ListConfig{
			Hydrate: listProjects,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"slug", "organization_slug"}),
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
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
				Transform:   transform.FromField("Organization").Transform(orgToSlug),
			},
			{
				Name:        "color",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "data_scrubber",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "data_scrubber_defaults",
				Type:        proto.ColumnType_STRING,
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
				Name:        "fingerprinting_rules",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "first_event",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "grouping_enhancements",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "has_access",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "is_bookmarked",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "is_internal",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "is_member",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "is_public",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "platform",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "processing_issues",
				Type:        proto.ColumnType_INT,
				Description: "",
			},
			{
				Name:        "resolve_age",
				Type:        proto.ColumnType_INT,
				Description: "",
			},
			{
				Name:        "scrape_java_script",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "scrub_ip_addresses",
				Type:        proto.ColumnType_BOOL,
				Description: "",
				Transform:   transform.FromField("ScrubIPAddresses"),
			},
			{
				Name:        "security_token",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "security_token_header",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "subject_prefix",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "subject_template",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "verify_ssl",
				Type:        proto.ColumnType_BOOL,
				Description: "",
				Transform:   transform.FromField("VerifySSL"),
			},
			{
				Name:        "allowed_domains",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "avatar",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "features",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "filters",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Hydrate:     getProjectFilters,
				Transform:   transform.FromValue(),
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
				Name:        "ownership",
				Type:        proto.ColumnType_JSON,
				Description: "",
				Hydrate:     getProjectOwnership,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "safe_fields",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "sensitive_fields",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "team",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "teams",
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

	projects, _, err := conn.Projects.List(ctx)
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
	organization_slug := d.EqualsQuals["organization_slug"].GetStringValue()

	// Check if slug or organization_slug is empty.
	if slug == "" || organization_slug == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getProject", "connection_error", err)
		return nil, err
	}

	org, _, err := conn.Projects.Get(ctx, organization_slug, slug)
	if err != nil {
		plugin.Logger(ctx).Error("getProject", "api_error", err)
		return nil, err
	}

	return org, nil
}

func getProjectFilters(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*sentry.Project)

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getProjectFilters", "connection_error", err)
		return nil, err
	}

	filters, _, err := conn.ProjectFilter.Get(ctx, *project.Organization.Slug, project.Slug)
	if err != nil {
		plugin.Logger(ctx).Error("getProjectFilters", "api_error", err)
		return nil, err
	}

	return filters, nil
}

func getProjectOwnership(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*sentry.Project)

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getProjectOwnership", "connection_error", err)
		return nil, err
	}

	ownership, _, err := conn.ProjectOwnerships.Get(ctx, *project.Organization.Slug, project.Slug)
	if err != nil {
		plugin.Logger(ctx).Error("getProjectOwnership", "api_error", err)
		return nil, err
	}

	return ownership, nil
}

func orgToSlug(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	return d.Value.(sentry.Organization).Slug, nil
}
