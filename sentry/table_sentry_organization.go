package sentry

import (
	"context"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryOrganizations(ctx context.Context) *plugin.Table {
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
				Description: "",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "is_default",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "require_2fa",
				Type:        proto.ColumnType_BOOL,
				Description: "",
				Transform:   transform.FromField("Require2FA"),
			},
			{
				Name:        "role",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "alerts_member_write",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "allow_join_requests",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "allow_shared_issues",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "attachments_role",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "data_scrubber",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "data_scrubber_defaults",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "debug_files_role",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "default_role",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "enhanced_privacy",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "events_member_admin",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "is_early_adopter",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "open_membership",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "pending_access_request",
				Type:        proto.ColumnType_INT,
				Description: "",
			},
			{
				Name:        "relay_pii_config",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "require_email_verification",
				Type:        proto.ColumnType_BOOL,
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
				Name:        "store_crash_reports",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "access",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "available_roles",
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
				Name:        "quota",
				Type:        proto.ColumnType_JSON,
				Description: "",
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
				Name:        "status",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
		},
	}
}

func listOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listOrganizations", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListCursorParams{}
	for {
		orgList, resp, err := conn.Organizations.List(ctx, params)
		if err != nil {
			plugin.Logger(ctx).Error("listOrganizations", "api_error", err)
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
	slug := d.EqualsQuals["slug"].GetStringValue()

	// Check if slug is empty.
	if slug == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganization", "connection_error", err)
		return nil, err
	}

	org, _, err := conn.Organizations.Get(ctx, slug)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganization", "api_error", err)
		return nil, err
	}

	return org, nil
}
