package circleci

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type circleciConfig struct {
	ApiToken *string `hcl:"api_token"`
}

func ConfigInstance() interface{} {
	return &circleciConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) circleciConfig {
	if connection == nil || connection.GetConfig() == nil {
		return circleciConfig{}
	}
	config, _ := connection.GetConfig().(circleciConfig)
	return config
}
