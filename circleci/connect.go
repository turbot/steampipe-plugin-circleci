package circleci

import (
	"context"
	"os"

	"github.com/jszwedko/go-circleci"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func Connect(ctx context.Context, d *plugin.QueryData) (*circleci.Client, error) {
	// have we already created and cached the session?
	sessionCacheKey := "CircleciSession"
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
