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
				Description: "The ID of this member.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the member.",
			},
			{
				Name:        "expired",
				Type:        proto.ColumnType_BOOL,
				Description: "The invite has expired.",
			},
			{
				Name:        "role",
				Type:        proto.ColumnType_STRING,
				Description: "The role of the organization member.",
			},
			{
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the organization the member belongs to.",
			},
			{
				Name:        "date_created",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation timestamp of the member.",
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "The email of the organization member.",
			},
			{
				Name:        "invite_status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the invite.",
			},
			{
				Name:        "inviter_name",
				Type:        proto.ColumnType_STRING,
				Description: "The inviter name who invited the member.",
			},
			{
				Name:        "pending",
				Type:        proto.ColumnType_BOOL,
				Description: "The invite is pending.",
			},
			{
				Name:        "role_name",
				Type:        proto.ColumnType_STRING,
				Description: "The role name of the organization member.",
			},
			{
				Name:        "flags",
				Type:        proto.ColumnType_JSON,
				Description: "The member flags",
			},
			{
				Name:        "teams",
				Type:        proto.ColumnType_JSON,
				Description: "The teams the organization member should be added to.",
			},
			{
				Name:        "user",
				Type:        proto.ColumnType_JSON,
				Description: "Represents a Sentry user.",
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
		plugin.Logger(ctx).Error("sentry_organization_member.listOrganizationMembers", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListCursorParams{}
	for {
		memberList, resp, err := conn.OrganizationMembers.List(ctx, *org.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("sentry_organization_member.listOrganizationMembers", "api_error", err)
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
		plugin.Logger(ctx).Error("sentry_organization_member.getOrganizationMember", "connection_error", err)
		return nil, err
	}

	member, _, err := conn.OrganizationMembers.Get(ctx, orgSlug, memberId)
	if err != nil {
		plugin.Logger(ctx).Error("sentry_organization_member.getOrganizationMember", "api_error", err)
		return nil, err
	}

	return OrganizationMember{*member, orgSlug}, nil
}
