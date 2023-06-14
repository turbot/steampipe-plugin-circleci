package circleci

import (
	"context"
	"errors"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"strings"
)

//// TABLE DEFINITION

func tableCircleCIInsightsWorkflow() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_insights_workflow",
		Description: "Get insights on project workflows.",
		List: &plugin.ListConfig{
			Hydrate: listCircleCIInsightsWorkflow,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "project_slug", Require: plugin.Required},
				{Name: "workflow_name", Require: plugin.Required},
				{Name: "branch", Require: plugin.Optional},
			},
		},

		Columns: []*plugin.Column{
			{Name: "project_slug", Description: "A unique identification for the project in the form of: <vcs_type>/<org_name>/<repo_name>.", Type: proto.ColumnType_STRING},
			{Name: "workflow_name", Description: "The name of the workflow.", Type: proto.ColumnType_STRING},
			{Name: "id", Description: "Unique key for the workflow.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "branch", Description: "The VCS branch of a Workflow's trigger.", Type: proto.ColumnType_STRING},
			{Name: "duration", Description: "Duration of the workflow in seconds.", Type: proto.ColumnType_INT},
			{Name: "created_at", Description: "Timestamp of when the workflow was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "stopped_at", Description: "Timestamp of when workflow was stopped.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "credits_used", Description: "The number of credits used during execution.", Type: proto.ColumnType_INT},
			{Name: "status", Description: "Workflow status.", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listCircleCIInsightsWorkflow(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	projectSlug := d.EqualsQualString("project_slug")
	workflowName := d.EqualsQualString("workflow_name")
	branch := ""
	if d.EqualsQuals["branch"] != nil {
		branch = d.EqualsQuals["branch"].GetStringValue()
	}
	logger.Info("circleci_insights_workflow.listCircleCIInsightsWorkflow", "branch", branch)

	if projectSlug == "" || workflowName == "" {
		return nil, nil
	}

	projectSlugSplit := strings.Split(projectSlug, "/")
	if len(projectSlugSplit) < 3 {
		err := errors.New("Malformed input for project_slug. Expected: {VCS}/{Org username}/{Repository name}")
		logger.Error("circleci_insights_workflow.listCircleCIInsightsWorkflow", "malformed_input", err)
		return nil, err
	}

	client, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_insights_workflow.listCircleCIInsightsWorkflow", "connect_error", err)
		return nil, err
	}
	workflows, err := client.ListAllInsightsWorkflows(projectSlug, workflowName, branch, logger)
	if err != nil {
		logger.Error("circleci_insights_workflow.listCircleCIInsightsWorkflow", "list_insight_error", err)
		return nil, err
	}

	for _, workflow := range workflows {
		// These fields are not provided by the API, so we set them from the query arguments
		workflow.ProjectSlug = projectSlug
		workflow.WorkflowName = workflowName

		d.StreamListItem(ctx, workflow)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
