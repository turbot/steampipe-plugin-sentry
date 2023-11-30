---
title: "Steampipe Table: sentry_organization_member - Query Sentry Organization Members using SQL"
description: "Allows users to query Sentry Organization Members, providing specific member details within an organization like roles, email, and user status."
---

# Table: sentry_organization_member - Query Sentry Organization Members using SQL

Sentry is an open-source error tracking tool that helps developers monitor and fix crashes in real time. It provides full stack trace details for all programming languages in a single dashboard. This enables developers to replicate and resolve errors that occur in their applications, improving the overall performance and user experience.

## Table Usage Guide

The `sentry_organization_member` table provides insights into members within Sentry organizations. As a developer or team lead, you can use this table to explore specific member details, including their roles, email, and user status. This can be particularly useful for managing team access, understanding member roles, and tracking the status of each member within your organization.

## Examples

### Basic info
Gain insights into the members of an organization, including their roles and whether their access has expired. This can be useful for auditing purposes or to manage user access.

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
Discover the members of your organization whose memberships have expired. This can be useful in auditing and managing your organization's active and inactive members.

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
Discover the segments that have owner access within an organization. This can be useful to manage permissions and access rights, ensuring only the appropriate members have high-level control.

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
Explore which individuals are part of a specific organization, allowing you to understand their roles and contact details for better team management and communication.

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
  organization_slug = 'myorg';
```

### List members of a particular team
Explore which individuals are part of a specific team and gain insights into their roles, active status, and contact details. This is useful for understanding team composition and managing team-related operations effectively.

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
Identify members who have not enabled two-factor authentication. This can help enhance security measures by pinpointing potential vulnerabilities within your organization.

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
Discover the individuals tied to a specific project, enabling a comprehensive understanding of team composition and roles. This is useful for project management, offering insights into team structure and member responsibilities.

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
      name = 'project1'
  );
```