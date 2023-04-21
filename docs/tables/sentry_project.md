# Table: sentry_project

A Project represents your service or application in Sentry. For example, you might have separate projects for your API server and frontend client.

## Examples

### Basic info

```sql
select
  id,
  name,
  status,
  slug,
  has_access,
  is_public,
  platform
from
  sentry_project;
```

### List public projects

```sql
select
  id,
  name,
  status,
  slug,
  has_access,
  is_public,
  platform
from
  sentry_project
where
  is_public;
```

### List go based projects

```sql
select
  id,
  name,
  status,
  slug,
  has_access,
  is_public,
  platform
from
  sentry_project
where
  platform = 'go';
```

### List projects of a particular organization

```sql
select
  id,
  name,
  status,
  slug,
  has_access,
  is_public,
  platform
from
  sentry_project
where
  organization_slug = 'myorg';
```

### List internal projects

```sql
select
  id,
  name,
  status,
  slug,
  has_access,
  is_public,
  platform
from
  sentry_project
where
  is_internal;
```

### List teams of a particular project

```sql
select
  id,
  name,
  t ->> 'id' as team_id,
  t ->> 'name' as team_name,
  t ->> 'slug' as team_slug
from
  sentry_project,
  jsonb_array_elements(teams) as t
where
  slug = 'myproj';
```
