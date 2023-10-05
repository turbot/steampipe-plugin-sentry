## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#12](https://github.com/turbot/steampipe-plugin-sentry/pull/12))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#9](https://github.com/turbot/steampipe-plugin-sentry/pull/9))
- Recompiled plugin with Go version `1.21`. ([#9](https://github.com/turbot/steampipe-plugin-sentry/pull/9))

## v0.0.1 [2023-05-04]

_What's new?_

- New tables added
  - [sentry_issue_alert](https://hub.steampipe.io/plugins/turbot/sentry/tables/sentry_issue_alert)
  - [sentry_key](https://hub.steampipe.io/plugins/turbot/sentry/tables/sentry_key)
  - [sentry_metric_alert](https://hub.steampipe.io/plugins/turbot/sentry/tables/sentry_metric_alert)
  - [sentry_organization](https://hub.steampipe.io/plugins/turbot/sentry/tables/sentry_organization)
  - [sentry_organization_integration](https://hub.steampipe.io/plugins/turbot/sentry/tables/sentry_organization_integration)
  - [sentry_organization_member](https://hub.steampipe.io/plugins/turbot/sentry/tables/sentry_organization_member)
  - [sentry_organization_repository](https://hub.steampipe.io/plugins/turbot/sentry/tables/sentry_organization_repository)
  - [sentry_project](https://hub.steampipe.io/plugins/turbot/sentry/tables/sentry_project)
  - [sentry_team](https://hub.steampipe.io/plugins/turbot/sentry/tables/sentry_team)
