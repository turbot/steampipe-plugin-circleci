package circleci

import (
	"context"
	"net/http"
	"os"

	"github.com/CircleCI-Public/circleci-cli/api/graphql"
	"github.com/jszwedko/go-circleci"
	"github.com/turbot/steampipe-plugin-circleci/circleci/rest"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func ConnectV1Sdk(ctx context.Context, d *plugin.QueryData) (*circleci.Client, error) {
	// have we already created and cached the session?
	sessionCacheKey := "CircleciSessionV1Sdk"
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

func ConnectV2Sdk(ctx context.Context, d *plugin.QueryData) (*graphql.Client, error) {
	// have we already created and cached the session?
	sessionCacheKey := "CircleciSessionV2Sdk"
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

func ConnectV2RestApi(ctx context.Context, d *plugin.QueryData) (*rest.Client, error) {
	// have we already created and cached the session?
	sessionCacheKey := "CircleciSessionV2RestApi"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*rest.Client), nil
	}

	circleciConfig := GetConfig(d.Connection)

	var defaultHost = "https://circleci.com/api/v2/"
	var apiToken string

	if circleciConfig.ApiToken != nil {
		apiToken = *circleciConfig.ApiToken
	} else {
		apiToken = os.Getenv("CIRCLECI_API_TOKEN")
	}

	client := rest.New(rest.Config{
		Token: apiToken,
		URL:   defaultHost,
	})
	// Save session into cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil

}
