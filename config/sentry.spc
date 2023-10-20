connection "sentry" {
  plugin = "sentry"

  # `auth_token` - The Auth Token for accessing the APIs of Sentry. (Required)
  # For more information on the Auth Token, please see https://docs.sentry.io/api/auth/.
  # Can also be set with the SENTRY_AUTH_TOKEN environment variable.
  # auth_token = "de70c93ecc594a0eb52463bd8f1e6d0b203615621e724762b3e5a9d82be291e9xfWdDNqwZPngS"

  # If no credentials are specified, the plugin will use Sentry CLI authentication.

  # `base_url` - The base URL of your Sentry Instance.
  # Can also be set with the SENTRY_URL environment variable.
  # base_url = "https://sentry.company.com/"
}
