package sentry

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Plugin creates this (sentry) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-sentry",
		DefaultTransform: transform.FromCamel(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"sentry_issue_alert":              tableSentryIssueAlert(ctx),
			"sentry_key":                      tableSentryKey(ctx),
			"sentry_metric_alert":             tableSentryMetricAlert(ctx),
			"sentry_organization":             tableSentryOrganization(ctx),
			"sentry_organization_integration": tableSentryOrganizationIntegration(ctx),
			"sentry_organization_member":      tableSentryOrganizationMember(ctx),
			"sentry_organization_repository":  tableSentryOrganizationRepository(ctx),
			"sentry_project":                  tableSentryProject(ctx),
			"sentry_team":                     tableSentryTeam(ctx),
		},
	}
	return p
}
