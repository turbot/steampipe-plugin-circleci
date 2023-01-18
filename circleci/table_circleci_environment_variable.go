package circleci

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleCIEnvironmentVariable() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_environment_variable",
		Description: "CircleCI environment variables store customer data that is used by projects.",
		List: &plugin.ListConfig{
			Hydrate:    listCircleCIEnvironmentVariables,
			KeyColumns: plugin.SingleColumn("context_id"),
		},

		Columns: []*plugin.Column{
			{Name: "context_id", Description: "Unique key for the context.", Transform: transform.FromField("ContextID"), Type: proto.ColumnType_STRING},
			{Name: "variable", Description: "Variable name.", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "Timestamp of when pipeline was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "updated_at", Description: "Timestamp of when variable was updated.", Type: proto.ColumnType_TIMESTAMP},
		},
	}
}

//// LIST FUNCTION

func listCircleCIEnvironmentVariables(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	contextId := d.EqualsQualString("context_id")

	// Check if the contextId is empty
	if contextId == "" {
		return nil, nil
	}

	client, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_environment_variable.listCircleCIEnvironmentVariables", "connect_error", err)
		return nil, err
	}

	var pageToken string
	for {
		envVarResponses, err := client.ListEnvironmentVariable(contextId, pageToken)
		if err != nil {
			logger.Error("circleci_environment_variable.listCircleCIEnvironmentVariables", "list_environment_variables_error", err)
			return nil, err
		}
		for _, pipeline := range envVarResponses.Items {
			d.StreamListItem(ctx, pipeline)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if envVarResponses.NextPageToken == "" {
			break
		}
		pageToken = envVarResponses.NextPageToken
	}

	return nil, err
}
