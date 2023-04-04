package sentry

import (
	"context"
	"errors"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
	"golang.org/x/oauth2"
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

	if sentryConfig.AuthToken != nil {
		authToken = *sentryConfig.AuthToken
	}
	if authToken != "" {
		tokenSrc := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: authToken},
		)
		httpClient := oauth2.NewClient(ctx, tokenSrc)

		client := sentry.NewClient(httpClient)

		return client, nil
	}

	return nil, errors.New("'auth_token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
}
