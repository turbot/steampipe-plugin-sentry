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
				Name:        "organization_slug",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "account_type",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "domain_name",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "external_id",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "grace_period_end",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "icon",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "organization_id",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "organization_integration_status",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "provider_key",
				Type:        proto.ColumnType_STRING,
				Description: "",
				Transform:   transform.FromField("Provider").Transform(fetchProviderKey),
			},
			{
				Name:        "provider",
				Type:        proto.ColumnType_JSON,
				Description: "",
			},
			{
				Name:        "scopes",
				Type:        proto.ColumnType_JSON,
				Description: "",
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
	orgSlug := d.EqualsQuals["organization_slug"].GetStringValue()

	// check if the provided orgSlug is not matching with the parentHydrate
	if orgSlug != "" && orgSlug != *org.Slug {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listOrganizationIntegrations", "connection_error", err)
		return nil, err
	}

	params := &sentry.ListOrganizationIntegrationsParams{}
	if d.EqualsQuals["provider_key"] != nil {
		params.ProviderKey = d.EqualsQuals["provider_key"].GetStringValue()
	}

	for {
		integrationList, resp, err := conn.OrganizationIntegrations.List(ctx, *org.Slug, params)
		if err != nil {
			plugin.Logger(ctx).Error("listOrganizationIntegrations", "api_error", err)
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
