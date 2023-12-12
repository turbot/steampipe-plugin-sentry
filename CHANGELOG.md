## v0.3.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#](https://github.com/turbot/steampipe-plugin-sentry/pull/))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#](https://github.com/turbot/steampipe-plugin-sentry/pull/))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-sentry/blob/main/docs/LICENSE). ([#](https://github.com/turbot/steampipe-plugin-sentry/pull/))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#](https://github.com/turbot/steampipe-plugin-sentry/pull/))

## v0.2.0 [2023-10-20]

_What's new?_

- The Sentry base URL can now be set through the `base_url` config argument or `SENTRY_URL` environment variable. ([#11](https://github.com/turbot/steampipe-plugin-sentry/pull/11)) (Thanks [@beudbeud](https://github.com/beudbeud) for the contribution!)

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
