package circleci

import (
	"context"
	"regexp"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleciProject() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_project",
		Description: "A CircleCI project shares the name of the code repository for which it automates workflows, tests, and deployment.",
		List: &plugin.ListConfig{
			Hydrate: listCircleciProjects,
		},

		Columns: []*plugin.Column{
			{Name: "slug", Description: "A unique identification for the project in the form of: <vcs_type>/<org_name>/<repo_name> .", Type: proto.ColumnType_STRING},
			{Name: "organization_slug", Description: "Organization that pipeline belongs to, in the form of: <vcs_type>/<org_name> .", Type: proto.ColumnType_STRING},
			{Name: "username", Description: "Organization or person's username who owns the repository.", Type: proto.ColumnType_STRING},
			{Name: "reponame", Description: "Name of the repository the project represents.", Type: proto.ColumnType_STRING},
			{Name: "vcs_url", Description: "URL to versioning code source.", Type: proto.ColumnType_STRING, Transform: transform.FromField("VCSURL")},
			{Name: "default_branch", Description: "Default branch name of the repository the project represents.", Type: proto.ColumnType_STRING},
			{Name: "env_vars", Description: "Environment variables set on the project.", Type: proto.ColumnType_JSON},
			{Name: "branches", Description: "Branches of the repository the project represents.", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listCircleciProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := ConnectV1Sdk(ctx, d)
	if err != nil {
		logger.Error("circleci_project.listCircleciProjects", "connect_error", err)
		return nil, err
	}

	projects, err := client.ListProjects()
	if err != nil {
		logger.Error("circleci_project.listCircleciProjects", "list_projects_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	githubRegex, _ := regexp.Compile("^https://github")
	bitbucketRegex, _ := regexp.Compile("^https://bitbucket")

	for _, project := range projects {

		var vcsSlug string
		githubMatch := githubRegex.MatchString(project.VCSURL)
		if githubMatch {
			vcsSlug = "gh"
		} else {
			bitbucketMatch := bitbucketRegex.MatchString(project.VCSURL)
			if bitbucketMatch {
				vcsSlug = "bb"
			}
		}

		var organizationSlug string
		if vcsSlug != "" {
			organizationSlug = vcsSlug + "/" + project.Username
		}

		var projectSlug string
		if organizationSlug != "" {
			projectSlug = organizationSlug + "/" + project.Reponame
		}

		envVars, _ := client.ListEnvVars(project.Username, project.Reponame)

		projectMap := map[string]interface{}{
			"Branches":         project.Branches,
			"DefaultBranch":    project.DefaultBranch,
			"Reponame":         project.Reponame,
			"Username":         project.Username,
			"OrganizationSlug": organizationSlug,
			"Slug":             projectSlug,
			"VCSURL":           project.VCSURL,
			"EnvVars":          envVars,
		}
		d.StreamListItem(ctx, projectMap)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
