package sentry

import (
	"context"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_team",
		Description: "Retrieve information about your teams.",
		List: &plugin.ListConfig{
			ParentHydrate: listOrganizations,
			Hydrate:       listTeams,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "organization_slug",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"organization_slug", "slug"}),
			Hydrate:    getTeam,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the team.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the team.",
			},
			{
				Name:        "has_access",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the team has access.",
			},
			{
				Name:        "slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the team.",
			},
			{
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the organization the team belongs to.",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation timestamp of the team.",
			},
			{
				Name:        "is_member",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the team is member.",
			},
			{
				Name:        "is_pending",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the team is pending.",
			},
			{
				Name:        "member_count",
				Type:        proto.ColumnType_INT,
				Description: "The total member count of the team.",
			},
			{
				Name:        "team_role",
				Type:        proto.ColumnType_STRING,
				Description: "The team role.",
			},
			{
				Name:        "avatar",
				Type:        proto.ColumnType_JSON,
				Description: "Represents an avatar.",
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

type Team struct {
	sentry.Team
	OrganizationSlug string
}

func listTeams(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	org := h.Item.(*sentry.Organization)
	orgSlug := d.EqualsQualString("organization_slug")

	// check if the provided orgSlug is not matching with the parentHydrate
	if orgSlug != "" && orgSlug != *org.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listTeams", "connection_error", err)
		return nil, err
	}

	teams, _, err := conn.Teams.List(ctx, *org.Slug)
	if err != nil {
		plugin.Logger(ctx).Error("listTeams", "api_error", err)
		return nil, err
	}

	for _, team := range teams {
		d.StreamListItem(ctx, Team{*team, *org.Slug})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getTeam(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	orgSlug := d.EqualsQualString("organization_slug")
	slug := d.EqualsQualString("slug")

	// Check if orgSlug or slug is empty.
	if orgSlug == "" || slug == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getTeam", "connection_error", err)
		return nil, err
	}

	team, _, err := conn.Teams.Get(ctx, orgSlug, slug)
	if err != nil {
		plugin.Logger(ctx).Error("getTeam", "api_error", err)
		return nil, err
	}

	return Team{*team, orgSlug}, nil
}
