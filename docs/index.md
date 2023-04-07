---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/sentry.svg"
brand_color: "#FAF9F6"
display_name: "Sentry"
short_name: "sentry"
description: "Steampipe plugin for Sentry."
og_description: "Query Sentry with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/sentry-social-graphic.png"
---

# Sentry + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Sentry](https://sentry.io) is a developer-first error tracking and performance monitoring platform that helps developers see what actually matters, solve quicker, and learn continuously about their applications.

For example:

```sql
select
  id,
  name,
  is_default,
  require_2fa,
  open_membership,
  default_role
from
  sentry_organization;
```

```
+------------------+--------+------------+-------------+-----------------+--------------+
| id               | name   | is_default | require_2fa | open_membership | default_role |
+------------------+--------+------------+-------------+-----------------+--------------+
| 4504948474773504 | Turbot | false      | false       | true            | admin        |
+------------------+--------+------------+-------------+-----------------+--------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/sentry/tables)**

## Get started

### Install

Download and install the latest Sentry plugin:

```bash
steampipe plugin install sentry
```

### Configuration

Installing the latest sentry plugin will create a config file (`~/.steampipe/config/sentry.spc`) with a single connection named `sentry`:

```hcl
connection "sentry" {
  plugin = "sentry"

  # `auth_token` - Create an authentication token in the Sentry Console for use.
  # Console path - settings -> account -> api -> auth-tokens
  # Can also be set with the SENTRY_AUTH_TOKEN environment variable.
  # auth_token = "de70c93ecc594a0eb52463bd8f1e6d0b203615621e724762b3e5a9d82be291e9xfWdDNqwZPngS"

  # If no credentials are specified, the plugin will use Sentry CLI authentication.
}
```

### Authentication Token Credentials

You may specify the Auth Token to authenticate:

- `auth_token`: Specify the authentication token.

```hcl
connection "sentry_via_auth_token" {
  plugin   = "sentry"
  auth_token  = "de70c93ecc594a0eb52463bd8f1e6d0b203615621e724762b3e5a9d82be291e9xfWdDNqwZPngS"
}
```

### Credentials from Environment Variables

The Sentry plugin will use the Sentry environment variable to obtain credentials **only if the `auth_token` is not specified** in the connection:

```sh
export SENTRY_AUTH_TOKEN="de70c93ecc594a0eb52463bd8f1e6d0b203615621e724762b3e5a9d82be291e9xfWdDNqwZPngS"
```

```hcl
connection "sentry" {
  plugin = "sentry"
}
```

### Sentry CLI

If no credentials are specified and the environment variables are not set, the plugin will use the active credentials from the Sentry CLI. You can run `sentry-cli login` to set up these credentials.

```hcl
connection "sentry" {
  plugin = "sentry"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-sentry
- Community: [Slack Channel](https://steampipe.io/community/join)
