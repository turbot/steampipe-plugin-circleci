package circleci

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-circleci"

// Plugin creates this (circleci) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"circleci_build":    tableCircleciBuild(),
			"circleci_orb":      tableCircleciOrb(),
			"circleci_pipeline": tableCircleciPipeline(),
			"circleci_project":  tableCircleciProject(),
			"circleci_workflow": tableCircleciWorkflow(),
		},
	}

	return p
}
