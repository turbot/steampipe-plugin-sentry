---
title: "Steampipe Table: sentry_team - Query Sentry Teams using SQL"
description: "Allows users to query Teams in Sentry, specifically providing details on team id, name, slug, and associated project details."
---

# Table: sentry_team - Query Sentry Teams using SQL

A Sentry Team is a collection of members working on a set of projects. Teams are associated with organizations and are used to manage access to projects. Teams can be managed through the Sentry web UI, the Sentry CLI, or programmatically through the Sentry API.

## Table Usage Guide

The `sentry_team` table provides insights into teams within Sentry. As a DevOps engineer, explore team-specific details through this table, including team id, name, slug, and associated project details. Utilize it to manage access to projects and to understand the team structure within your organization.

## Examples

### Basic info
Assess the elements within your team by reviewing their access rights, roles, and membership status. This query can be used to gain insights into team dynamics and ensure appropriate access levels are maintained for each team member.

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
Explore which teams have administrative access to better manage security and user permissions within your organization. This can be particularly useful for auditing purposes or when planning access control strategies.

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
Explore which teams you are a member of, along with their respective roles and member count. This is useful to understand your team's dynamics and your role within it.

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
Determine the teams that lack members to assess potential areas for team reorganization or resource allocation. This query can be useful in identifying unused teams and optimizing team structures.

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

### List teams without an assigned role
Identify teams that have not been assigned a role. This is useful for ensuring all teams have the necessary permissions to perform their tasks.

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
Determine the teams that are currently unassigned to any projects. This query is useful in identifying potential resources that could be allocated to new or existing projects.

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