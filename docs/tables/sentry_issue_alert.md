---
title: "Steampipe Table: sentry_issue_alert - Query Sentry Issue Alerts using SQL"
description: "Allows users to query Sentry Issue Alerts, providing information about each alert's details such as ID, alert rule, project, and user."
---

# Table: sentry_issue_alert - Query Sentry Issue Alerts using SQL

Sentry is an open-source application monitoring platform that helps uncover, triage, and prioritize errors in real-time. An Issue Alert in Sentry is a notification sent when certain conditions are met in a project, such as when an event occurs that matches the conditions defined in an alert rule. These alerts help to identify and manage issues in your applications and infrastructure.

## Table Usage Guide

The `sentry_issue_alert` table provides insights into Issue Alerts within Sentry's application monitoring platform. As a developer or DevOps engineer, explore alert-specific details through this table, including alert rules, projects, and associated metadata. Utilize it to uncover information about alerts, such as those related to specific projects or rules, helping you to triage and prioritize errors in your applications effectively.

## Examples

### Basic info
Explore which issues have triggered alerts in your Sentry application. This can help in identifying potential problematic areas, allowing for timely intervention and issue resolution.

```sql
select
  id,
  name,
  owner,
  organization_slug,
  project_slug,
  action_match,
  date_created
from
  sentry_issue_alert;
```

### List alerts for a particular project
Explore the alerts associated with a specific project to gain insights into issues that may need immediate attention. This can help in assessing the health and stability of the project.

```sql
select
  id,
  name,
  owner,
  organization_slug,
  project_slug,
  action_match,
  date_created
from
  sentry_issue_alert
where
  project_slug = 'go';
```

### List alerts owned by a particular team
Determine the areas in which specific alerts are owned by a certain team. This can help in understanding the distribution of responsibility and tracking the issues that a particular team needs to address.

```sql
select
  a.id,
  a.name,
  a.owner,
  a.organization_slug,
  a.project_slug,
  a.action_match,
  a.date_created
from
  sentry_issue_alert as a,
  sentry_team as t
where
  t.id = split_part(a.owner,':',2)
  and t.name = 'Team A';
```

### Show list of actions of a particular alert
Explore the actions associated with a specific alert to gain insights into its configuration and owner details. This can be useful for understanding the alert's role and impact within a project.

```sql
select
  id,
  name,
  owner,
  project_slug,
  jsonb_pretty(actions) as actions
from
  sentry_issue_alert
where
  name = 'alert1';
```

### List alerts older than a month
Analyze the settings to understand any alerts that have been active for over a month. This can help in identifying long-standing issues that may require immediate attention or escalation.

```sql
select
  id,
  name,
  owner,
  organization_slug,
  project_slug,
  action_match,
  date_created
from
  sentry_issue_alert
where
  date_created <= now() - interval '1 month';
```