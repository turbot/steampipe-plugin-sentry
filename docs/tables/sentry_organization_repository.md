---
title: "Steampipe Table: sentry_organization_repository - Query Sentry Organization Repositories using SQL"
description: "Allows users to query Organization Repositories in Sentry, specifically providing information about the repositories used by the organizations, including their name, ID, status, and other relevant details."
---

# Table: sentry_organization_repository - Query Sentry Organization Repositories using SQL

Sentry is an open-source error tracking system that helps developers monitor and fix crashes in real time. It offers the ability to track errors in real-time, with detailed reports and alerts. An Organization in Sentry represents a group of users and encapsulates a set of projects. Repositories in Sentry are references to the remote version control repositories that are connected to your Sentry organization.

## Table Usage Guide

The `sentry_organization_repository` table provides insights into repositories in Sentry that are used by organizations. As a developer or DevOps engineer, you can explore repository-specific details through this table, including repository names, IDs, status, and other related information. This table is beneficial for managing and auditing the repositories associated with your Sentry organization, helping you to monitor and troubleshoot issues more effectively.

## Examples

### Basic info
Explore which organizations have created repositories, along with their status and creation date. This can be useful to determine the areas in which new repositories are being added and by which organizations.

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
Explore which repositories are inactive. This is useful to identify potential areas for clean-up or archival in your organization's repository storage.

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
Explore which GitHub repositories are linked to your organization on Sentry. This is useful to understand the status and integration details of your repositories, helping you manage your codebase more effectively.

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