package sentry

import (
	"context"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryOrganization(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_organization",
		Description: "Retrieve information about your organizations.",
		List: &plugin.ListConfig{
			Hydrate: listOrganizations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("slug"),
			Hydrate:    getOrganization,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the organization.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the organization.",
			},
			{
				Name:        "is_default",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization is the default.",
			},
			{
				Name:        "require_2fa",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization has 2FA enabled.",
				Transform:   transform.FromField("Require2FA"),
			},
			{
				Name:        "role",
				Type:        proto.ColumnType_STRING,
				Description: "The organization role.",
			},
			{
				Name:        "slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the organization.",
			},
			{
				Name:        "alerts_member_write",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization has alerts member write enabled.",
			},
			{
				Name:        "allow_join_requests",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization has allowed join requests.",
			},
			{
				Name:        "allow_shared_issues",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization has allowed shared issues.",
			},
			{
				Name:        "attachments_role",
				Type:        proto.ColumnType_STRING,
				Description: "The organization attachments role.",
			},
			{
				Name:        "data_scrubber",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization has data scrubber enabled.",
			},
			{
				Name:        "data_scrubber_defaults",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization has data scrubber defaults enabled.",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation timestamp of the organization.",
			},
			{
				Name:        "debug_files_role",
				Type:        proto.ColumnType_STRING,
				Description: "The organization debug files role.",
			},
			{
				Name:        "default_role",
				Type:        proto.ColumnType_STRING,
				Description: "The default role of the organization.",
			},
			{
				Name:        "enhanced_privacy",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization has enhanced privacy enabled.",
			},
			{
				Name:        "events_member_admin",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization has events member admin access.",
			},
			{
				Name:        "is_early_adopter",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization has early adopter enabled.",
			},
			{
				Name:        "open_membership",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization has open membership enabled.",
			},
			{
				Name:        "pending_access_request",
				Type:        proto.ColumnType_INT,
				Description: "The number of the pending access requests.",
			},
			{
				Name:        "relay_pii_config",
				Type:        proto.ColumnType_STRING,
				Description: "The relay pii config.",
			},
			{
				Name:        "require_email_verification",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization requires email verification.",
			},
			{
				Name:        "scrape_java_script",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization scrapes java script.",
			},
			{
				Name:        "scrub_ip_addresses",
				Type:        proto.ColumnType_BOOL,
				Description: "Check if the organization scrubs IP addresses.",
				Transform:   transform.FromField("ScrubIPAddresses"),
			},
			{
				Name:        "store_crash_reports",
				Type:        proto.ColumnType_STRING,
				Description: "The store crash reports.",
			},
			{
				Name:        "access",
				Type:        proto.ColumnType_JSON,
				Description: "The organization access.",
			},
			{
				Name:        "available_roles",
				Type:        proto.ColumnType_JSON,
				Description: "Represents a Sentry organization's available role.",
			},
			{
				Name:        "avatar",
				Type:        proto.ColumnType_JSON,
				Description: "Represents a Sentry organization's avatar.",
			},
			{
				Name:        "features",
				Type:        proto.ColumnType_JSON,
				Description: "The organization features.",
			},
			{
				Name:        "quota",
				Type:        proto.ColumnType_JSON,
				Description: "Represents a Sentry organization's quota.",
			},
			{
				Name:        "safe_fields",
				Type:        proto.ColumnType_JSON,
				Description: "The organization safe fields.",
			},
			{
				Name:        "sensitive_fields",
				Type:        proto.ColumnType_JSON,
				Description: "The organization sensitive fields.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_JSON,
				Description: "Represents a Sentry organization's status.",
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

func listOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("sentry_organization.listOrganizations", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListCursorParams{}
	for {
		orgList, resp, err := conn.Organizations.List(ctx, params)
		if err != nil {
			plugin.Logger(ctx).Error("sentry_organization.listOrganizations", "api_error", err)
			return nil, err
		}
		for _, org := range orgList {
			d.StreamListItem(ctx, org)

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

func getOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	slug := d.EqualsQualString("slug")

	// Check if slug is empty.
	if slug == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("sentry_organization.getOrganization", "connection_error", err)
		return nil, err
	}

	org, _, err := conn.Organizations.Get(ctx, slug)
	if err != nil {
		plugin.Logger(ctx).Error("sentry_organization.getOrganization", "api_error", err)
		return nil, err
	}

	return org, nil
}
