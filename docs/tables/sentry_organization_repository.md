# Table: sentry_organization_repository

Represent a list of version control repositories for a given organization.

## Examples

### Basic info

```sql
select
  id,
  name,
  status,
  organization_slug,
  date_created,
  external_slug
from
  sentry_organization_repository;
```

### List repositories which are not active

```sql
select
  id,
  name,
  status,
  organization_slug,
  date_created,
  external_slug
from
  sentry_organization_repository
where
  status <> 'active';
```

### List github repositories

```sql
select
  r.id as repository_id,
  r.name as repository_name,
  r.status,
  r.organization_slug,
  r.date_created,
  external_slug
from
  sentry_organization_repository as r,
  sentry_organization_integration as i
where
  r.integration_id = i.id
  and i.provider_key = 'github';
```
