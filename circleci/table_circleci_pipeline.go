package circleci

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleciPipeline() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_pipeline",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listCircleciPipelines,
		},

		Columns: []*plugin.Column{
			{Name: "created_at", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "errors", Description: "", Type: proto.ColumnType_JSON},
			{Name: "id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "number", Description: "", Type: proto.ColumnType_INT},
			{Name: "project_slug", Description: "", Type: proto.ColumnType_STRING},
			{Name: "state", Description: "", Type: proto.ColumnType_STRING},
			{Name: "trigger_parameters", Description: "", Type: proto.ColumnType_JSON},
			{Name: "trigger", Description: "", Type: proto.ColumnType_JSON},
			{Name: "updated_at", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "vcs", Description: "", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listCircleciPipelines(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_pipeline.listCircleciPipelines", "connect_error", err)
		return nil, err
	}

	pipelines, err := client.ListPipelines("gh", "fluent-cattle")
	if err != nil {
		logger.Error("circleci_pipeline.listCircleciPipelines", "list_pipelines_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, pipeline := range pipelines.Items {
		d.StreamListItem(ctx, pipeline)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
