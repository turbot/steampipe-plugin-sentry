package sentry

import (
	"context"
	"errors"

	"github.com/atlassian/go-sentry-api"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type sentryConfig struct {
	AuthToken *string `cty:"auth_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"auth_token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &sentryConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) sentryConfig {
	if connection == nil || connection.Config == nil {
		return sentryConfig{}
	}
	config, _ := connection.Config.(sentryConfig)
	return config
}

func getClient(ctx context.Context, d *plugin.QueryData) (*sentry.Client, error) {
	sentryConfig := GetConfig(d.Connection)

	authToken := ""
	// endpoint := ""
	// timeout :=

	if sentryConfig.AuthToken != nil {
		authToken = *sentryConfig.AuthToken
	}

	if authToken != "" { // Authenticate with auth_token
		client, err := sentry.NewClient(authToken, nil, nil)
		if err != nil {
			return nil, err
		}
		return client, nil
	}

	return nil, errors.New("'auth_token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
}
