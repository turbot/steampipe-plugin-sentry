# Table: sentry_metric_alert

Metric alerts trigger when a metric is breached for either error or transaction events. Use metric alerts to monitor a finite and known set of metrics and components you care about, such as error frequency or performance metrics in your entire project, on important pages, or with specific tags.

Create alerts to monitor metrics, such as:

- Total errors in your project
- Latency: min, max, average, percentile
- Failure rate
- Crash free session or user rate for monitoring release health
- Custom metrics

You can find a full list of available metric alerts in [Metric Alerts](https://docs.sentry.io/product/alerts/alert-types/#metric-alerts).

## Examples

### Basic info

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
  and t.name = 'test';
```

### Show list of triggers of a particular alert

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
