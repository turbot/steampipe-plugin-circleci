package circleci

import (
	"context"
	"errors"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleCIPipeline() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_pipeline",
		Description: "CircleCI pipelines are the highest-level unit of work, encompassing a project’s full .circleci/config.yml file.",
		List: &plugin.ListConfig{
			Hydrate:    listCircleCIPipelines,
			KeyColumns: plugin.SingleColumn("project_slug"),
		},

		Columns: commonColumns([]*plugin.Column{
			{Name: "project_slug", Description: "A unique identification for the project in the form of: <vcs_type>/<org_name>/<repo_name>.", Type: proto.ColumnType_STRING},
			{Name: "id", Description: "Unique key for the pipeline.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "number", Description: "A second identifier for the pipeline.", Type: proto.ColumnType_INT},
			{Name: "created_at", Description: "Timestamp of when the pipeline was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "errors", Description: "A list of errors while executing pipeline's jobs.", Type: proto.ColumnType_JSON},
			{Name: "state", Description: "The state of the pipeline.", Type: proto.ColumnType_STRING},
			{Name: "trigger_parameters", Description: "Any parameter for pipeline triggering.", Type: proto.ColumnType_JSON},
			{Name: "trigger", Description: "What triggers the pipeline to run.", Type: proto.ColumnType_JSON},
			{Name: "updated_at", Description: "Timestamp of when pipeline was updated.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "vcs", Description: "Version control system of the pipeline", Type: proto.ColumnType_JSON},
		}),
	}
}

//// LIST FUNCTION

func listCircleCIPipelines(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	projectSlug := d.EqualsQualString("project_slug")

	// Check if the projectSlug is empty
	if projectSlug == "" {
		return nil, nil
	}
	projectSlugSplit := strings.Split(projectSlug, "/")
	if len(projectSlugSplit) < 3 {
		err := errors.New("Malformed input for project_slug. Expected: {VCS}/{Org username}/{Repository name}")
		logger.Error("circleci_pipeline.listCircleCIPipelines", "malformed_input", err)
		return nil, err
	}
	vcs := projectSlugSplit[0]
	organization := projectSlugSplit[1]

	client, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_pipeline.listCircleCIPipelines", "connect_error", err)
		return nil, err
	}

	pipelines, err := client.ListPipelines(vcs, organization)
	if err != nil {
		logger.Error("circleci_pipeline.listCircleCIPipelines", "list_pipelines_error", err)
		return nil, err
	}

	for _, pipeline := range pipelines.Items {

		// For projects that use GitLab or GitHub App, use circleci as the vcs-slug, replace org-name with the organization ID (found in Organization Settings), and replace repo-name with the project ID (found in Project Settings).
		// If we pass the `project_slug` value as "select * from circleci_pipeline where project_slug = 'circleci/5ad8293d-30b0-4594-be94-0b28fb727eef/b20334e0-c813-4c0f-a27d-9a1ca2488548'",
		// the API returns the value as "circleci/CDdrPhWsW7Pa5PNcj4Y7en/NywbLpDu1Eb1Pobxj3Nphu".
		// Due to Steampipe's filtering mechanism, this results in empty rows.
		// Since "project_slug" is required to query this table, we assign the "project_slug" from the WHERE clause directly.
		pipeline.ProjectSlug = projectSlug
		
		d.StreamListItem(ctx, pipeline)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
