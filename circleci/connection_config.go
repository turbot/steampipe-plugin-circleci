package circleci

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type circleciConfig struct {
	ApiToken *string `hcl:"api_token"`
}

func ConfigInstance() interface{} {
	return &circleciConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) circleciConfig {
	if connection == nil || connection.Config == nil {
		return circleciConfig{}
	}
	config, _ := connection.Config.(circleciConfig)
	return config
}
