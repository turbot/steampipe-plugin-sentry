# Table: sentry_key

Represents a list of client keys bound to a project.

## Examples

### Basic info

```sql
select
  id,
  name,
  is_active,
  secret,
  project_slug,
  public
from
  sentry_key;
```

### List keys for a particular project

```sql
select
  id,
  name,
  is_active,
  secret,
  project_slug,
  public
from
  sentry_key
where
  project_slug = 'go';
```

### List inactive keys

```sql
select
  id,
  name,
  is_active,
  secret,
  project_slug,
  public
from
  sentry_key
where
  not is_active;
```

### List keys older than 90 days

```sql
select
  id,
  name,
  is_active,
  secret,
  project_slug,
  public
from
  sentry_key
where
  date_created <= now() - interval '90 days';
```
