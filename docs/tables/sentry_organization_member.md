# Table: sentry_organization_member

Represent a list of users that belong to a given organization.

## Examples

### Basic info

```sql
select
  id,
  name,
  expired,
  role,
  organization_slug,
  email
from
  sentry_organization_member;
```

### List expired members

```sql
select
  id,
  name,
  expired,
  role,
  organization_slug,
  email
from
  sentry_organization_member
where
  expired;
```

### List members with owner access

```sql
select
  id,
  name,
  expired,
  role,
  organization_slug,
  email
from
  sentry_organization_member
where
  role = 'owner';
```

### List members of a particular organization

```sql
select
  id,
  name,
  expired,
  role,
  organization_slug,
  email
from
  sentry_organization_member
where
  organization_slug = 'test';
```

### List members of a particular team

```sql
select
  id,
  name,
  expired,
  role,
  organization_slug,
  email
from
  sentry_organization_member,
  jsonb_array_elements_text(teams) as team
where
  team = 'turbot';
```

### List members with 2FA disabled

```sql
select
  id,
  name,
  expired,
  role,
  organization_slug,
  email
from
  sentry_organization_member
where
  "user" ->> 'has2fa' = 'false';
```

### List members associated with a particular project

```sql
select
  id,
  name,
  expired,
  role,
  organization_slug,
  email
from
  sentry_organization_member,
  jsonb_array_elements_text(teams) as team
where
  team in
  (
    select
      t ->> 'name'
    from
      sentry_project,
      jsonb_array_elements(teams) as t
    where
      name = 'test-project'
  );
```
