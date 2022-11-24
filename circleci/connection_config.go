package circleci

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type circleciConfig struct {
	ApiToken *string `cty:"api_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_token": {
		Type: schema.TypeString,
	},
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
