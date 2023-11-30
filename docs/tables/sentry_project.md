---
title: "Steampipe Table: sentry_project - Query Sentry Projects using SQL"
description: "Allows users to query Sentry Projects, providing insights into project details, including its id, name, platform, slug, and other related information."
---

# Table: sentry_project - Query Sentry Projects using SQL

Sentry is an open-source error tracking tool that helps developers monitor and fix crashes in real time. It provides complete visibility into your stack, enabling you to detect and fix issues as soon as they occur. Sentry supports all major languages and frameworks, and integrates with your existing workflow to identify, respond to, and resolve production software issues.

## Table Usage Guide

The `sentry_project` table provides insights into projects within Sentry. As a developer or DevOps engineer, explore project-specific details through this table, including project id, name, platform, slug, and other related information. Utilize it to manage and monitor your projects, understand their configuration, and quickly identify and resolve issues.

## Examples

### Basic info
Discover the segments that have access and are public within a certain platform. This is useful in assessing the status and understanding the accessibility of projects within a given platform.

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
Explore which projects are publicly accessible, enabling you to identify potential security risks or areas for collaboration.

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
Explore which projects are based on the 'Go' programming language. This is useful to identify and manage projects that are using this specific platform.

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
Determine the areas in which a specific organization has projects. This can be useful to gain insights into the project status, access details, and platform usage within the organization.

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
Discover the segments that are classified as internal projects within your organizational structure. This allows you to assess the elements within your organization that are not public-facing, aiding in strategic planning and resource allocation.

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
Determine the teams associated with a specific project to understand their roles and contributions. This is useful for project management and resource allocation.

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