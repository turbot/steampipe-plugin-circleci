package circleci

import (
	"context"

	"github.com/turbot/steampipe-plugin-circleci/rest"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleCIOrganization() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_organization",
		Description: "CircleCI organization is a representation of a VCS account ownership.",
		List: &plugin.ListConfig{
			Hydrate: listCircleCIOrganizations,
		},

		Columns: []*plugin.Column{
			{Name: "id", Description: "Unique key for the organization.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "slug", Description: "A unique identification for the organization in the form of: <vcs_type>/<org_name>.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "The organization name.", Type: proto.ColumnType_STRING},
			{Name: "vcs_type", Description: "Version control system of the organization.", Type: proto.ColumnType_STRING},
			{Name: "avatar_url", Description: "Avatar icon of the organization.", Type: proto.ColumnType_STRING},
			{Name: "contexts", Description: "Contexts provide a mechanism for securing and sharing environment variables across projects.", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listCircleCIOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	client, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_organization.listCircleCIOrganizations", "connect_error", err)
		return nil, err
	}

	organizations, err := client.ListOrganizations()
	if err != nil {
		logger.Error("circleci_organization.listCircleCIOrganizations", "list_organizations_error", err)
		return nil, err
	}

	for _, organization := range *organizations {
		organization.Contexts, err = getOrganizationContexts(*client, organization.Slug)
		if err != nil {
			logger.Error("circleci_organization.listCircleCIContexts", "list_organizations_error", err)
			return nil, err
		}
		d.StreamListItem(ctx, organization)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}

func getOrganizationContexts(client rest.Client, organizationSlug string) ([]rest.Context, error) {
	var contexts []rest.Context
	var pageToken string
	for {
		contextResponses, err := client.ListContexts(organizationSlug, pageToken)
		if err != nil {
			return nil, err
		}
		contexts = append(contexts, contextResponses.Items...)
		if contextResponses.NextPageToken == "" {
			break
		}
		pageToken = contextResponses.NextPageToken
	}
	return contexts, nil
}
