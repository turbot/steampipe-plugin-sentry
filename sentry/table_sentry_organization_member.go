package sentry

import (
	"context"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryOrganizationMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_organization_member",
		Description: "Retrieve information about your organization members.",
		List: &plugin.ListConfig{
			ParentHydrate: listOrganizations,
			Hydrate:       listOrganizationMembers,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "organization_slug",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"organization_slug", "id"}),
			Hydrate:    getOrganizationMember,
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
				Name:        "expired",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "role",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "invite_status",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "inviter_name",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "pending",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "role_name",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "flags",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "teams",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "user",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
		},
	}
}

type OrganizationMember struct {
	sentry.OrganizationMember
	OrganizationSlug string
}

func listOrganizationMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	org := h.Item.(*sentry.Organization)
	orgSlug := d.EqualsQuals["organization_slug"].GetStringValue()

	// check if the provided orgSlug is not matching with the parentHydrate
	if orgSlug != "" && orgSlug != *org.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listOrganizationMembers", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListCursorParams{}
	for {
		memberList, resp, err := conn.OrganizationMembers.List(ctx, *org.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("listOrganizationMembers", "api_error", err)
			return nil, err
		}
		for _, member := range memberList {
			d.StreamListItem(ctx, OrganizationMember{*member, *org.Slug})

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

func getOrganizationMember(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	orgSlug := d.EqualsQuals["organization_slug"].GetStringValue()
	memberId := d.EqualsQuals["id"].GetStringValue()

	// Check if orgSlug or memberId is empty.
	if orgSlug == "" || memberId == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganizationMember", "connection_error", err)
		return nil, err
	}

	member, _, err := conn.OrganizationMembers.Get(ctx, orgSlug, memberId)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganizationMember", "api_error", err)
		return nil, err
	}

	return OrganizationMember{*member, orgSlug}, nil
}
