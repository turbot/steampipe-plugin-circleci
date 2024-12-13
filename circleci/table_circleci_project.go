package circleci

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleCIProject() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_project",
		Description: "A CircleCI project shares the name of the code repository for which it automates workflows, tests, and deployment.",
		List: &plugin.ListConfig{
			Hydrate: listCircleCIProjects,
		},

		Columns: commonColumns([]*plugin.Column{
			{Name: "slug", Description: "A unique identification for the project in the form of: <vcs_type>/<org_name>/<repo_name>.", Type: proto.ColumnType_STRING},
			{Name: "organization_slug", Description: "Organization that pipeline belongs to, in the form of: <vcs_type>/<org_name>.", Type: proto.ColumnType_STRING},
			{Name: "username", Description: "Organization or person's username who owns the repository.", Type: proto.ColumnType_STRING},
			{Name: "reponame", Description: "Name of the repository the project represents.", Type: proto.ColumnType_STRING},
			{Name: "vcs_url", Description: "VCS URL.", Type: proto.ColumnType_STRING, Transform: transform.FromField("VCSURL")},
			{Name: "default_branch", Description: "Default branch name of the repository the project represents.", Type: proto.ColumnType_STRING},
			{Name: "env_vars", Description: "Environment variables set on the project.", Type: proto.ColumnType_JSON},
			{Name: "checkout_keys", Description: "Keys used to checkout the code from VCS.", Type: proto.ColumnType_JSON},
			{Name: "branches", Description: "Branches of the repository the project represents.", Type: proto.ColumnType_JSON},
		}),
	}
}

//// LIST FUNCTION

func listCircleCIProjects(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := ConnectV1Sdk(ctx, d)
	if err != nil {
		logger.Error("circleci_project.listCircleCIProjects", "connect_error", err)
		return nil, err
	}

	projects, err := client.ListProjects()
	if err != nil {
		logger.Error("circleci_project.listCircleCIProjects", "list_projects_error", err)
		return nil, err
	}

	// Get project followed by user
	// The List/Get project CLI call does not include slug information if the project is not associated with Bitbucket or GitHub on CircleCI.
	privateClient, err := ConnectPrivateRestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_project.ConnectPrivateRestApi", "connect_error", err)
		return nil, err
	}

	privateProfileInfo, err := privateClient.GetCurrentUserFollowedProject()
	if err != nil {
		logger.Error("circleci_project.GetCurrentUserFollowedProject", "api_error", err)
		return nil, err
	}

	for _, project := range projects {

		organizationSlug, projectSlug := "", ""
		for _, followedProject := range privateProfileInfo.FollowedProjects {
			if project.Reponame == followedProject.Name {
				projectSlug = followedProject.Slug
				organizationSlug = strings.Join(strings.Split(followedProject.Slug, "/")[:2], "/")
			}
		}
		envVars, _ := client.ListEnvVars(project.Username, project.Reponame)
		checkoutKeys, _ := client.ListCheckoutKeys(project.Username, project.Reponame)

		projectMap := map[string]interface{}{
			"Branches":         project.Branches,
			"DefaultBranch":    project.DefaultBranch,
			"Reponame":         project.Reponame,
			"Username":         project.Username,
			"OrganizationSlug": organizationSlug,
			"Slug":             projectSlug,
			"VCSURL":           project.VCSURL,
			"EnvVars":          envVars,
			"CheckoutKeys":     checkoutKeys,
		}
		d.StreamListItem(ctx, projectMap)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
