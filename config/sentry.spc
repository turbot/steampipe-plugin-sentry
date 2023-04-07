connection "sentry" {
  plugin = "sentry"

  # `auth_token` - Create an authentication token in the Sentry Console for use.
  # Console path - settings -> account -> api -> auth-tokens
  # Can also be set with the SENTRY_AUTH_TOKEN environment variable.
  # auth_token = "de70c93ecc594a0eb52463bd8f1e6d0b203615621e724762b3e5a9d82be291e9xfWdDNqwZPngS"

  # If no credentials are specified, the plugin will use Sentry CLI authentication.
}
