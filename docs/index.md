---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/sentry.svg"
brand_color: "#362D59"
display_name: "Sentry"
short_name: "sentry"
description: "Steampipe plugin to query organizations, projects, teams and more from Sentry."
og_description: "Query Sentry with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/sentry-social-graphic.png"
---

# Sentry + Steampipe

[Sentry](https://sentry.io) is a developer-first error tracking and performance monitoring platform that helps developers see what actually matters, solve quicker, and learn continuously about their applications.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List your Sentry organizations:

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

## Quick start

### Install

Download and install the latest Sentry plugin:

```sh
steampipe plugin install sentry
```

### Credentials

| Item        | Description                                                                                                                                                                    |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Credentials | Sentry requires an [Auth Token](https://docs.sentry.io/api/auth/) for all requests.                                                                                             |
| Permissions | Auth tokens have the same permissions as the user who creates them, and if the user permissions change, the Auth token permissions also change.                                |
| Radius      | Each connection represents a single Sentry Installation.                                                                                                                       |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/sentry.spc`)<br />2. Credentials specified in environment variables, e.g., `SENTRY_AUTH_TOKEN`. |

### Configuration

Installing the latest sentry plugin will create a config file (`~/.steampipe/config/sentry.spc`) with a single connection named `sentry`:

Configure your account details in `~/.steampipe/config/sentry.spc`:

```hcl
connection "sentry" {
  plugin = "sentry"

  # `auth_token` - The Auth Token for accessing the APIs of Sentry. (Required).
  # For more information on the Auth Token, please see https://docs.sentry.io/api/auth/.
  # Can also be set with the SENTRY_AUTH_TOKEN environment variable.
  # auth_token = "de70c93ecc594a0eb52463bd8f1e6d0b203615621e724762b3e5a9d82be291e9xfWdDNqwZPngS"

  # If no credentials are specified, the plugin will use Sentry CLI authentication.

  # `baseurl` - The base URL of your Sentry Instance.
  # Can also be set with the SENTRY_URL environment variable.
  # baseurl = "https://sentry.company.com/"
}
```

## Configuring Sentry Credentials

### Authentication Token Credentials

You may specify the Auth Token to authenticate:

- `auth_token`: Specify the authentication token.

```hcl
connection "sentry" {
  plugin   = "sentry"
  auth_token  = "de70c93ecc594a0eb52463bd8f1e6d0b203615621e724762b3e5a9d82be291e9xfWdDNqwZPngS"
}
```

### Credentials from Environment Variables

The Sentry plugin will use the Sentry environment variable to obtain credentials **only if the `auth_token` is not specified** in the connection:

```sh
export SENTRY_AUTH_TOKEN="de70c93ecc594a0eb52463bd8f1e6d0b203615621e724762b3e5a9d82be291e9xfWdDNqwZPngS"
export SENTRY_URL="https://sentry.company.com/"
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
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
