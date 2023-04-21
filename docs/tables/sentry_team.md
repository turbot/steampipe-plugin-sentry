# Table: sentry_team

Represent a list of teams bound to a organization.

## Examples

### Basic info

```sql
select
  id,
  name,
  has_access,
  slug,
  has_access,
  is_member,
  member_count,
  team_role
from
  sentry_team;
```

### List teams with admin access

```sql
select
  id,
  name,
  has_access,
  slug,
  has_access,
  is_member,
  member_count,
  team_role
from
  sentry_team
where
  team_role = 'admin';
```

### List your teams

```sql
select
  id,
  name,
  has_access,
  slug,
  has_access,
  is_member,
  member_count,
  team_role
from
  sentry_team
where
  is_member;
```

### List teams without any members

```sql
select
  id,
  name,
  has_access,
  slug,
  has_access,
  is_member,
  member_count,
  team_role
from
  sentry_team
where
  member_count = 0;
```

### List teams without role

```sql
select
  id,
  name,
  has_access,
  slug,
  has_access,
  is_member,
  member_count,
  team_role
from
  sentry_team
where
  team_role is null;
```

### List teams which are not assigned to any project

```sql
select
  id,
  name,
  has_access,
  slug,
  has_access,
  is_member,
  member_count,
  team_role
from
  sentry_team
where
  id not in
  (
    select
      t ->> 'id'
    from
      sentry_project,
      jsonb_array_elements(teams) as t
  );
```
