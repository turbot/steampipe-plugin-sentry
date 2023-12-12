---
title: "Steampipe Table: sentry_organization - Query Sentry Organizations using SQL"
description: "Allows users to query Organizations within Sentry, specifically providing details about the organization's name, slug, status, onboarding tasks, and more."
---

# Table: sentry_organization - Query Sentry Organizations using SQL

Sentry is an open-source error tracking tool that helps developers monitor and fix crashes in real-time. The organization resource within Sentry represents a group of users and teams that have access to a set of projects. It provides details about the organization's name, slug, status, onboarding tasks, and more.

## Table Usage Guide

The `sentry_organization` table provides insights into Organizations within Sentry. As a developer or Sentry administrator, you can explore organization-specific details through this table, including name, slug, status, and onboarding tasks. Utilize it to uncover information about organizations, such as their status and details about onboarding tasks, helping to manage and monitor the organizations within Sentry.

## Examples

### Basic info
Explore which organizations within Sentry have open memberships and require two-factor authentication to understand the security measures in place. This can help in assessing the default role assigned to members for better user management.

```sql+postgres
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

```sql+sqlite
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
Discover organizations that have not enabled two-factor authentication, a feature crucial for enhancing security measures and preventing unauthorized access. This query is particularly useful for security audits and compliance checks.

```sql+postgres
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

```sql+sqlite
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
  require_2fa = 0;
```

### List organizations with open membership enabled
Explore which organizations have enabled open membership. This can be useful for assessing the security settings and understanding the access levels within each organization.

```sql+postgres
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

```sql+sqlite
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
  open_membership = 1;
```

### List organizations with default role admin
Explore which organizations have been set with 'admin' as the default role. This is useful for assessing the security posture of your organizations, as having 'admin' as the default role could potentially open up vulnerabilities.

```sql+postgres
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

```sql+sqlite
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
Explore which organizations are not currently active. This can be useful in identifying areas for potential re-engagement or cleanup.

```sql+postgres
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

```sql+sqlite
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
  json_extract(status, '$.id') <> 'active';
```

### List access details of a particular organization
Explore the access details of a specific organization to understand the permissions and roles associated with it. This can help in managing user access and ensuring appropriate security measures are in place.

```sql+postgres
select
  id,
  name,
  jsonb_pretty(access) as access
from
  sentry_organization
where
  slug = 'myorg';
```

```sql+sqlite
select
  id,
  name,
  access
from
  sentry_organization
where
  slug = 'myorg';
```