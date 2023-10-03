package sentry

import (
	"bufio"
	"context"
	"errors"
	"os"
	"strings"

	"github.com/jianyuan/go-sentry/v2/sentry"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
	"golang.org/x/oauth2"
)

type sentryConfig struct {
	AuthToken *string `cty:"auth_token"`
	BaseUrl *string `cty:"baseurl"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"auth_token": {
		Type: schema.TypeString,
	},
	"baseurl": {
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

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "sentry"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*sentry.Client), nil
	}

	sentryConfig := GetConfig(d.Connection)

	authToken := os.Getenv("SENTRY_AUTH_TOKEN")
	baseUrl := os.Getenv("SENTRY_URL")

	if sentryConfig.BaseUrl != nil {
		baseUrl= *sentryConfig.BaseUrl
	}

	if sentryConfig.AuthToken != nil {
		authToken = *sentryConfig.AuthToken
	}

	if baseUrl != "" {
		tokenSrc := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: authToken},
		)
		httpClient := oauth2.NewClient(ctx, tokenSrc)

		client, nil := sentry.NewOnPremiseClient(baseUrl, httpClient)

		// Save to cache
		d.ConnectionManager.Cache.Set(cacheKey, client)

		return client, nil
		// Authenticate with AuthToken
	}

	if authToken != "" { // Authenticate with AuthToken
		tokenSrc := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: authToken},
		)
		httpClient := oauth2.NewClient(ctx, tokenSrc)

		client := sentry.NewClient(httpClient)

		// Save to cache
		d.ConnectionManager.Cache.Set(cacheKey, client)

		return client, nil
	} else { // Authenticate with CLI
		home, _ := os.UserHomeDir()
		file, _ := os.Open(home + "/.sentryclirc")
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if equal := strings.Index(line, "="); equal >= 0 {
				if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
					authToken = ""
					if len(line) > equal {
						authToken = strings.TrimSpace(line[equal+1:])
					}
				}
			}
		}

		if authToken != "" {
			tokenSrc := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: authToken},
			)
			httpClient := oauth2.NewClient(ctx, tokenSrc)

			client := sentry.NewClient(httpClient)

			// Save to cache
			d.ConnectionManager.Cache.Set(cacheKey, client)

			return client, nil
		}
	}

	return nil, errors.New("'auth_token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
}
