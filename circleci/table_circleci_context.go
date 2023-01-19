package circleci

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleCIContext() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_context",
		Description: "CircleCI context provide a mechanism for securing and sharing environment variables across projects.",
		List: &plugin.ListConfig{
			Hydrate:    listCircleCIContexts,
			KeyColumns: plugin.SingleColumn("organization_slug"),
		},

		Columns: []*plugin.Column{
			{Name: "id", Description: "Unique key for the context.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "organization_slug", Description: "A unique identification for the organization in the form of: <vcs_type>/<org_name>.", Type: proto.ColumnType_STRING, Transform: transform.FromQual("organization_slug")},
			{Name: "name", Description: "The context name.", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "Timestamp of when context was created.", Type: proto.ColumnType_TIMESTAMP},
		},
	}
}

//// LIST FUNCTION

func listCircleCIContexts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	organizationSlug := d.EqualsQualString("organization_slug")

	// Check if the organizationSlug is empty
	if organizationSlug == "" {
		return nil, nil
	}

	client, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_context.listCircleCIContexts", "connect_error", err)
		return nil, err
	}

	var pageToken string
	for {
		contextResponses, err := client.ListContexts(organizationSlug, pageToken)
		if err != nil {
			logger.Error("circleci_context.listCircleCIContexts", "list_contexts_error", err)
			return nil, err
		}
		for _, context := range contextResponses.Items {
			if err != nil {
				logger.Error("circleci_context.listCircleCIContexts", "list_contexts_error", err)
				return nil, err
			}
			d.StreamListItem(ctx, context)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if contextResponses.NextPageToken == "" {
			break
		}
		pageToken = contextResponses.NextPageToken
	}

	return nil, err
}
