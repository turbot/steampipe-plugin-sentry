# Table: sentry_issue_alert

Issue alerts trigger whenever any issue in a project matches specified criteria. You can create alerts for issue-level changes, such as:

- New issues
- Issue frequency increasing
- Resolved and ignored issues becoming unresolved

You can find a full list of issue alert triggers in [Issue Alert Configuration](https://docs.sentry.io/product/alerts/create-alerts/issue-alert-config/#when-conditions-triggers).

## Examples

### Basic info

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
