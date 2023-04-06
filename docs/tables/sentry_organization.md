# Table: sentry_organization

Represent a list of organizations available to the authenticated session.

## Examples

### Basic info

```sql
select
  id,
  name,
  is_default,
  require_2fa,
  open_membership,
  default_role
from
  sentry_organization;
```

### List organizations with 2FA disabled

```sql
select
  id,
  name,
  is_default,
  require_2fa,
  open_membership,
  default_role
from
  sentry_organization
where
  not require_2fa;
```

### List organizations with open membership enabled

```sql
select
  id,
  name,
  is_default,
  require_2fa,
  open_membership,
  default_role
from
  sentry_organization
where
  open_membership;
```

### List organizations with default role admin

```sql
select
  id,
  name,
  is_default,
  require_2fa,
  open_membership,
  default_role
from
  sentry_organization
where
  default_role = 'admin';
```

### List inactive organizations

```sql
select
  id,
  name,
  is_default,
  require_2fa,
  open_membership,
  default_role
from
  sentry_organization
where
  status ->> 'id' <> 'active';
```

### List access details of a particular organization

```sql
select
  id,
  name,
  jsonb_pretty(access) as access
from
  sentry_organization
where
  slug = 'test';
```
