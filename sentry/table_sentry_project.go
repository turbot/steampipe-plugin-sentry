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
				Description: "The ID of the project.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the project.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the project.",
			},
			{
				Name:        "slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the project.",
			},
			{
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the organization the project belongs to.",
				Transform:   transform.FromField("Organization").Transform(orgToSlug),
			},
			{
				Name:        "color",
				Type:        proto.ColumnType_STRING,
				Description: "The project color.",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation timestamp of the project.",
			},
			{
				Name:        "data_scrubber",
				Type:        proto.ColumnType_STRING,
				Description: "The project data scrubber.",
			},
			{
				Name:        "data_scrubber_defaults",
				Type:        proto.ColumnType_STRING,
				Description: "The project default data scrubber.",
			},
			{
				Name:        "digest_min_delay",
				Type:        proto.ColumnType_INT,
				Description: "The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.",
			},
			{
				Name:        "digest_max_delay",
				Type:        proto.ColumnType_INT,
				Description: "The maximum amount of time (in seconds) to wait between scheduling digests for delivery.",
			},
			{
				Name:        "fingerprinting_rules",
				Type:        proto.ColumnType_STRING,
				Description: "The fingerprinting rules of the project.",
			},
			{
				Name:        "first_event",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The first event timestamp.",
			},
			{
				Name:        "grouping_enhancements",
				Type:        proto.ColumnType_STRING,
				Description: "The grouping enhancements of the project.",
			},
			{
				Name:        "has_access",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the project has access.",
			},
			{
				Name:        "is_bookmarked",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the project is bookmarked.",
			},
			{
				Name:        "is_internal",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the project is internal.",
			},
			{
				Name:        "is_member",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the project is member.",
			},
			{
				Name:        "is_public",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the project is public.",
			},
			{
				Name:        "platform",
				Type:        proto.ColumnType_STRING,
				Description: "The optional platform for this project.",
			},
			{
				Name:        "processing_issues",
				Type:        proto.ColumnType_INT,
				Description: "The number of processing issues of the project.",
			},
			{
				Name:        "resolve_age",
				Type:        proto.ColumnType_INT,
				Description: "Hours in which an issue is automatically resolve if not seen after this amount of time.",
			},
			{
				Name:        "scrape_java_script",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the project has scrape java script enabled.",
			},
			{
				Name:        "scrub_ip_addresses",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the project has scrub IP addresses enabled.",
				Transform:   transform.FromField("ScrubIPAddresses"),
			},
			{
				Name:        "security_token",
				Type:        proto.ColumnType_STRING,
				Description: "The security token of the project.",
			},
			{
				Name:        "security_token_header",
				Type:        proto.ColumnType_STRING,
				Description: "The security token header of the project.",
			},
			{
				Name:        "subject_prefix",
				Type:        proto.ColumnType_STRING,
				Description: "The subject prefix of the project.",
			},
			{
				Name:        "subject_template",
				Type:        proto.ColumnType_STRING,
				Description: "The subject template of the project.",
			},
			{
				Name:        "verify_ssl",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the project is verified with SSL.",
				Transform:   transform.FromField("VerifySSL"),
			},
			{
				Name:        "allowed_domains",
				Type:        proto.ColumnType_JSON,
				Description: "The allowed domains of the project.",
			},
			{
				Name:        "avatar",
				Type:        proto.ColumnType_JSON,
				Description: "Represents an avatar.",
			},
			{
				Name:        "features",
				Type:        proto.ColumnType_JSON,
				Description: "The features of the project.",
			},
			{
				Name:        "filters",
				Type:        proto.ColumnType_JSON,
				Description: "The project filters.",
				Hydrate:     getProjectFilters,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "options",
				Type:        proto.ColumnType_JSON,
				Description: "The project options.",
			},
			{
				Name:        "organization",
				Type:        proto.ColumnType_JSON,
				Description: "The organization to which the project belongs to.",
			},
			{
				Name:        "safe_fields",
				Type:        proto.ColumnType_JSON,
				Description: "The project safe fields.",
			},
			{
				Name:        "sensitive_fields",
				Type:        proto.ColumnType_JSON,
				Description: "The project sensitive fields.",
			},
			{
				Name:        "teams",
				Type:        proto.ColumnType_JSON,
				Description: "The project teams.",
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

func orgToSlug(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	return d.Value.(sentry.Organization).Slug, nil
}
