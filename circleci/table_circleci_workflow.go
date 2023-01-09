package circleci

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleciWorkflow() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_workflow",
		Description: "Workflows define a list of jobs and their run order.",
		List: &plugin.ListConfig{
			Hydrate:    listCircleciWorkflows,
			KeyColumns: plugin.SingleColumn("pipeline_id"),
		},

		Columns: []*plugin.Column{
			{Name: "canceled_by", Description: "ID of the user who canceled the workflow.", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "Timestamp of when workflow was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "errored_by", Description: "ID of the user who caused the workflow to error.", Type: proto.ColumnType_STRING},
			{Name: "id", Description: "Unique key for the workflow.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "Human readable name of the workflow.", Type: proto.ColumnType_STRING},
			{Name: "pipeline_id", Description: "Unique key for the pipeline.", Type: proto.ColumnType_STRING, Transform: transform.FromField("PipelineID")},
			{Name: "pipeline_number", Description: "A second identifier for the pipeline.", Type: proto.ColumnType_INT},
			{Name: "project_slug", Description: "A unique identification for the project in the form of: <vcs_type>/<org_name>/<repo_name> .", Type: proto.ColumnType_STRING},
			{Name: "started_by", Description: "Id of the user who started the workflow.", Type: proto.ColumnType_STRING},
			{Name: "status", Description: "Workflow status", Type: proto.ColumnType_STRING},
			{Name: "stopped_at", Description: "Timestamp of when workflow was stopped.", Type: proto.ColumnType_TIMESTAMP},
		},
	}
}

//// LIST FUNCTION

func listCircleciWorkflows(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	pipelineId := d.EqualsQualString("pipeline_id")

	// Empty check for pipelineId
	if pipelineId == "" {
		return nil, nil
	}

	client, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_workflow.listCircleciWorkflows", "connect_error", err)
		return nil, err
	}

	workflows, err := client.ListPipelinesWorkflow(pipelineId)
	if err != nil {
		logger.Error("circleci_workflow.listCircleciWorkflows", "list_workflows_error", err)
		return nil, err
	}

	for _, workflow := range workflows.Items {
		d.StreamListItem(ctx, workflow)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
