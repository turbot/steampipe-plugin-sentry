package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sentry/sentry"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: sentry.Plugin})

}
