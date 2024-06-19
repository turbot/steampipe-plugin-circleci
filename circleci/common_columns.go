package circleci

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Login ID is unique per connection.
// We can have multiple organization, projects, context, etc.. under an account.
func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "login_id",
			Description: "Unique identifier for the account login.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getLoginId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getLoginIdMemoized = plugin.HydrateFunc(getLoginIdUncached).Memoize(memoize.WithCacheKeyFunction(getLoginIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getLoginId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getLoginIdMemoized(ctx, d, h)
}

// Build a cache key for the call to getLoginIdCacheKey.
func getLoginIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getLoginId"
	return key, nil
}

func getLoginIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	client, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getLoginIdUncached", "connection_error", err)
		return nil, err
	}

	login, err := client.GetCurrentLogin()
	if err != nil {
		plugin.Logger(ctx).Error("getLoginIdUncached", "api_error", err)
		return nil, err
	}

	return login.ID, nil
}
