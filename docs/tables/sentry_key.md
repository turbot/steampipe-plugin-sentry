---
title: "Steampipe Table: sentry_key - Query Sentry API Keys using SQL"
description: "Allows users to query API Keys in Sentry, including the key's details, such as the name, date of creation, and associated project details."
---

# Table: sentry_key - Query Sentry API Keys using SQL

API Keys in Sentry are used to interact with the Sentry API on behalf of a user or an integration. They are a crucial part of the Sentry integration ecosystem, allowing for programmatic access to data and functionality within Sentry. API Keys can be associated with specific projects or be organization-wide, and can have varying levels of permissions based on their scope.

## Table Usage Guide

The `sentry_key` table provides insights into API Keys within Sentry. As a developer or a DevOps engineer, explore key-specific details through this table, including the name, date of creation, and associated project details. Utilize it to manage and audit your API keys, such as identifying keys with high permissions, keys that are no longer in use, or keys associated with specific projects.

## Examples

### Basic info
Explore which Sentry keys are active, their associated projects, and their public visibility status. This can help in managing access and maintaining security within your projects.

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
Discover the segments that are active within a specific project by identifying the unique keys associated with it. This can be useful for managing access and understanding the overall project configuration.

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
Discover the segments that consist of inactive keys to better manage and secure your data resources. This can help you maintain system integrity by identifying potential vulnerabilities or unused keys.

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
Discover the segments that have keys older than 90 days to assess the elements within your project that may need updating or removal for security reasons. This query is particularly useful for maintaining system security and preventing potential breaches.

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