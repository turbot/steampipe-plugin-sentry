---
title: "Steampipe Table: sentry_organization_integration - Query Sentry Organization Integrations using SQL"
description: "Allows users to query Sentry Organization Integrations, providing insights into the integrations that are enabled for a Sentry organization."
---

# Table: sentry_organization_integration - Query Sentry Organization Integrations using SQL

Sentry Organization Integration is a feature within Sentry that allows organizations to integrate their Sentry projects with third-party services. It provides a unified way to manage and use integrations for various purposes, including issue tracking, alerting, data forwarding, and more. Sentry Organization Integration helps in enhancing the functionality of Sentry projects by leveraging the capabilities of other services.

## Table Usage Guide

The `sentry_organization_integration` table provides insights into the integrations enabled for a Sentry organization. As an engineer or administrator, explore integration-specific details through this table, including provider information, integration configuration, and associated metadata. Utilize it to uncover information about integrations, such as their status, the services they connect with, and the configuration settings applied.

## Examples

### Basic info
Explore which organizations have integrated with Sentry by checking their status and account type. This can help you understand the scope of Sentry's integration within your organization.

```sql+postgres
select
  id,
  name,
  status,
  organization_slug,
  account_type,
  domain_name
from
  sentry_organization_integration;
```

```sql+sqlite
select
  id,
  name,
  status,
  organization_slug,
  account_type,
  domain_name
from
  sentry_organization_integration;
```

### List integrations which are not active
Discover the segments that include inactive integrations within organizations to assess the areas requiring attention or potential updates. This can be beneficial in enhancing system efficiency and security by identifying and addressing inactive components.

```sql+postgres
select
  id,
  name,
  status,
  organization_slug,
  account_type,
  domain_name
from
  sentry_organization_integration
where
  status <> 'active';
```

```sql+sqlite
select
  id,
  name,
  status,
  organization_slug,
  account_type,
  domain_name
from
  sentry_organization_integration
where
  status != 'active';
```

### List github integrations
Explore which GitHub integrations are active within your organization by assessing their status, account types, and associated domain names. This can help identify any potential issues or areas for improvement in your organization's use of GitHub.

```sql+postgres
select
  id,
  name,
  status,
  organization_slug,
  account_type,
  domain_name
from
  sentry_organization_integration
where
  provider_key = 'github';
```

```sql+sqlite
select
  id,
  name,
  status,
  organization_slug,
  account_type,
  domain_name
from
  sentry_organization_integration
where
  provider_key = 'github';
```