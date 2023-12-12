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
		Name:               pluginName,
		DefaultTransform:   transform.FromCamel(),
		DefaultRetryConfig: &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError([]string{"Rate Limit Exceeded"})},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"circleci_build":                        tableCircleCIBuild(),
			"circleci_context":                      tableCircleCIContext(),
			"circleci_context_environment_variable": tableCircleCIContextEnvironmentVariable(),
			"circleci_insights_workflow_run":        tableCircleCIInsightsWorkflowRun(),
			"circleci_organization":                 tableCircleCIOrganization(),
			"circleci_pipeline":                     tableCircleCIPipeline(),
			"circleci_project":                      tableCircleCIProject(),
			"circleci_workflow":                     tableCircleCIWorkflow(),
		},
	}

	return p
}
