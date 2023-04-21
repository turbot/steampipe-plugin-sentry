# Table: sentry_organization_integration

Sentry’s integration platform provides a way for external services to interact with Sentry using webhooks, UI components, and the REST API. Integrations using this platform are first-class actors within Sentry.

## Examples

### Basic info

```sql
select
  id,
  name,
  status,
  organization_slug,
  account_type,
  domain_name
from
  sentry_organization_integration;
```

### List integrations which are not active

```sql
select
  id,
  name,
  status,
  organization_slug,
  account_type,
  domain_name
from
  sentry_organization_integration
where
  status <> 'active';
```

### List github integrations

```sql
select
  id,
  name,
  status,
  organization_slug,
  account_type,
  domain_name
from
  sentry_organization_integration
where
  provider_key = 'github';
```
