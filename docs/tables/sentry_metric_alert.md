---
title: "Steampipe Table: sentry_metric_alert - Query Sentry Metric Alerts using SQL"
description: "Allows users to query Metric Alerts in Sentry, specifically the alert rules and their triggers, providing insights into application error tracking and notifications."
---

# Table: sentry_metric_alert - Query Sentry Metric Alerts using SQL

Sentry is an open-source error tracking system that helps developers monitor and fix crashes in real time. It provides complete stack traces, tells you how many times an error occurred, and shows the affected user count. A Metric Alert in Sentry is a rule that triggers notifications when a specific metric (like the count of a certain type of event) goes beyond a defined threshold.

## Table Usage Guide

The `sentry_metric_alert` table provides insights into Metric Alerts within Sentry's error tracking system. As a developer or DevOps engineer, explore alert-specific details through this table, including alert rules, their triggers, and associated metadata. Utilize it to uncover information about alerts, such as those triggered by certain types of events, the thresholds defined for alerts, and the actions taken when alerts are triggered.

## Examples

### Basic info
Explore which metric alerts have been created within your organization, allowing you to identify instances where specific alerts may need to be updated or adjusted. This is particularly useful for maintaining optimal performance and ensuring timely responses to any issues or anomalies.

```sql
select
  id,
  name,
  owner,
  organization_slug,
  project_slug,
  aggregate,
  data_set,
  date_created
from
  sentry_metric_alert;
```

### List alerts for a particular project
Explore which alerts are associated with a specific project to better manage and respond to issues. This can provide crucial insights to maintain project health and efficiency.

```sql
select
  id,
  name,
  owner,
  organization_slug,
  project_slug,
  aggregate,
  data_set,
  date_created
from
  sentry_metric_alert
where
  project_slug = 'go';
```

### List alerts owned by a particular team
Explore which alerts are managed by a specific team to better understand the distribution of responsibilities and ownership within your organization. This can be particularly useful in large organizations where multiple teams are managing different sets of alerts.

```sql
select
  a.id,
  a.name,
  a.owner,
  a.organization_slug,
  a.project_slug,
  a.aggregate,
  a.data_set,
  a.date_created
from
  sentry_metric_alert as a,
  sentry_team as t
where
  t.id = split_part(a.owner,':',2)
  and t.name = 'Team A';
```

### Show list of triggers of a particular alert
Explore the different triggers associated with a specific alert to understand their thresholds and actions. This can help in assessing the alert's sensitivity and response strategy.

```sql
select
  t ->> 'id' as id,
  t ->> 'alertRuleId' as alert_rule_id,
  t ->> 'label' as label,
  t ->> 'thresholdType' as threshold_type,
  t ->> 'alertThreshold' as alert_threshold,
  t ->> 'resolveThreshold' as resolve_threshold,
  t ->> 'dateCreated' as date_created,
  jsonb_pretty(t -> 'actions') as actions
from
  sentry_metric_alert,
  jsonb_array_elements(triggers) as t
where
  name = 'alert-metric';
```

### List alerts older than a month
Explore which alerts have been active for longer than a month to assess areas that may require attention or review. This can help identify lingering issues within your project or organization that have not been resolved.

```sql
select
  id,
  name,
  owner,
  organization_slug,
  project_slug,
  aggregate,
  data_set,
  date_created
from
  sentry_metric_alert
where
  date_created <= now() - interval '1 month';
```