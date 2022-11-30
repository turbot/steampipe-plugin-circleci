package circleci

import (
	"context"
	"net/http"
	"os"

	"github.com/CircleCI-Public/circleci-cli/api/graphql"
	"github.com/jszwedko/go-circleci"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func ConnectV1(ctx context.Context, d *plugin.QueryData) (*circleci.Client, error) {
	// have we already created and cached the session?
	sessionCacheKey := "CircleciSessionV1"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*circleci.Client), nil
	}

	circleciConfig := GetConfig(d.Connection)

	var apiToken string

	if circleciConfig.ApiToken != nil {
		apiToken = *circleciConfig.ApiToken
	} else {
		apiToken = os.Getenv("CIRCLECI_API_TOKEN")
	}

	client := &circleci.Client{Token: apiToken}
	// Save session into cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil

}

func ConnectV2(ctx context.Context, d *plugin.QueryData) (*graphql.Client, error) {
	// have we already created and cached the session?
	sessionCacheKey := "CircleciSessionV2"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*graphql.Client), nil
	}

	circleciConfig := GetConfig(d.Connection)

	var defaultEndpoint = "graphql-unstable"
	var defaultHost = "https://circleci.com"
	var apiToken string

	if circleciConfig.ApiToken != nil {
		apiToken = *circleciConfig.ApiToken
	} else {
		apiToken = os.Getenv("CIRCLECI_API_TOKEN")
	}

	client := graphql.NewClient(http.DefaultClient, defaultHost, defaultEndpoint, apiToken, false)
	// Save session into cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil

}
