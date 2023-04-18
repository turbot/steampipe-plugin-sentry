package sentry

import (
	"context"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSentryOrganizationIntegration(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "sentry_organization_integration",
		Description: "Retrieve information about your organization integrations.",
		List: &plugin.ListConfig{
			ParentHydrate: listOrganizations,
			Hydrate:       listOrganizationIntegrations,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "organization_slug",
					Require: plugin.Optional,
				},
				{
					Name:    "provider_key",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the integration.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the integration.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the integration.",
			},
			{
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug of the organization the integration belongs to.",
			},
			{
				Name:        "account_type",
				Type:        proto.ColumnType_STRING,
				Description: "The account type of the integration.",
			},
			{
				Name:        "domain_name",
				Type:        proto.ColumnType_STRING,
				Description: "The domain name of the integration.",
			},
			{
				Name:        "external_id",
				Type:        proto.ColumnType_STRING,
				Description: "The external id of the integration.",
			},
			{
				Name:        "grace_period_end",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The grace period end timestamp.",
			},
			{
				Name:        "icon",
				Type:        proto.ColumnType_STRING,
				Description: "The icon of the integration.",
			},
			{
				Name:        "organization_id",
				Type:        proto.ColumnType_STRING,
				Description: "The organization ID.",
			},
			{
				Name:        "organization_integration_status",
				Type:        proto.ColumnType_STRING,
				Description: "The organization integration status.",
			},
			{
				Name:        "provider_key",
				Type:        proto.ColumnType_STRING,
				Description: "The integration provider key.",
				Transform:   transform.FromField("Provider").Transform(fetchProviderKey),
			},
			{
				Name:        "provider",
				Type:        proto.ColumnType_JSON,
				Description: "The integration provider.",
			},
			{
				Name:        "scopes",
				Type:        proto.ColumnType_JSON,
				Description: "The integration scopes.",
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

type OrganizationIntegration struct {
	sentry.OrganizationIntegration
	OrganizationSlug string
}

func listOrganizationIntegrations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	org := h.Item.(*sentry.Organization)
	orgSlug := d.EqualsQualString("organization_slug")

	// check if the provided orgSlug is not matching with the parentHydrate
	if orgSlug != "" && orgSlug != *org.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("sentry_organization_integration.listOrganizationIntegrations", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListOrganizationIntegrationsParams{}
	if d.EqualsQuals["provider_key"] != nil {
		params.ProviderKey = d.EqualsQualString("provider_key")
	}

	for {
		integrationList, resp, err := conn.OrganizationIntegrations.List(ctx, *org.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("sentry_organization_integration.listOrganizationIntegrations", "api_error", err)
			return nil, err
		}
		for _, integration := range integrationList {
			d.StreamListItem(ctx, OrganizationIntegration{*integration, *org.Slug})

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

func fetchProviderKey(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	return d.Value.(sentry.OrganizationIntegrationProvider).Key, nil
}
