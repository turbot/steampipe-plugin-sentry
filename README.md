![image](https://hub.steampipe.io/images/plugins/turbot/sentry-social-graphic.png)

# Sentry Plugin for Steampipe

Use SQL to query organizations, projects, teams and more from Sentry.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/sentry)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/sentry/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-sentry/issues)

## Quick start

Download and install the latest Sentry plugin:

```bash
steampipe plugin install sentry
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/sentry#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/sentry#configuration).

### Configuring Sentry Credentials

Configure your account details in `~/.steampipe/config/sentry.spc`:

You may specify the Auth Token to authenticate:

- `auth_token`: Specify the authentication token.

```hcl
connection "sentry" {
  plugin   = "sentry"
  auth_token  = "de70c93ecc594a0eb52463bd8f1e6d0b203615621e724762b3e5a9d82be291e9xfWdDNqwZPngS"
}
```

or through environment variables

```sh
export SENTRY_AUTH_TOKEN="de70c93ecc594a0eb52463bd8f1e6d0b203615621e724762b3e5a9d82be291e9xfWdDNqwZPngS"
```

or through the active credentials from the Sentry CLI. You can run `sentry-cli login` to set up these credentials.

```hcl
connection "sentry" {
  plugin = "sentry"
}
```

Run steampipe:

```shell
steampipe query
```

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-sentry.git
cd steampipe-plugin-sentry
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/sentry.spc
```

Try it!

```
steampipe query
> .inspect sentry
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-sentry/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-sentry/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Sentry Plugin](https://github.com/turbot/steampipe-plugin-sentry/labels/help%20wanted)
