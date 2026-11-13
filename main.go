package main

import (
	"github.com/turbot/steampipe-plugin-circleci/circleci"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: circleci.Plugin})
}
