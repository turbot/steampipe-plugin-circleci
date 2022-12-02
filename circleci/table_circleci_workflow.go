package circleci

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleciWorkflow() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_workflow",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate:    listCircleciWorkflows,
			KeyColumns: plugin.SingleColumn("pipeline_id"),
		},

		Columns: []*plugin.Column{
			{Name: "canceled_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "errored_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "pipeline_id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("PipelineID")},
			{Name: "pipeline_number", Description: "", Type: proto.ColumnType_INT},
			{Name: "project_slug", Description: "", Type: proto.ColumnType_STRING},
			{Name: "started_by", Description: "", Type: proto.ColumnType_STRING},
			{Name: "status", Description: "", Type: proto.ColumnType_STRING},
			{Name: "stopped_at", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "tag", Description: "", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listCircleciWorkflows(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	pipelineId := d.Table.Get.KeyColumns.Find("pipeline_id").String()
	logger.Info("pipelineId")

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
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
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
