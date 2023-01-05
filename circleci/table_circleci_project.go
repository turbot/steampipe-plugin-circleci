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
			{Name: "reponame", Description: "Name of the repository the project represent", Type: proto.ColumnType_STRING},
			{Name: "vcs_url", Description: "URL to versioning code source.", Type: proto.ColumnType_STRING, Transform: transform.FromField("VCSURL")},
			{Name: "branches", Description: "Branches of the repository the project represents.", Type: proto.ColumnType_JSON},
			// {Name: "campfire_notify_prefs", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "campfire_room", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "campfire_subdomain", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "campfire_token", Description: "", Type: proto.ColumnType_STRING},
			{Name: "compile", Description: "", Type: proto.ColumnType_STRING},
			{Name: "default_branch", Description: "Default branch name of the repository the project represents.", Type: proto.ColumnType_STRING},
			{Name: "dependencies", Description: "Other projects this depends on.", Type: proto.ColumnType_STRING},
			{Name: "extra", Description: "", Type: proto.ColumnType_STRING},
			{Name: "feature_flags", Description: "", Type: proto.ColumnType_JSON},
			// {Name: "flowdock_api_token", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("FlowdockAPIToken")},
			{Name: "followed", Description: "Whether you follow this project.", Type: proto.ColumnType_BOOL},
			// {Name: "hall_notify_prefs", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "hall_room_api_token", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("HallRoomAPIToken")},
			{Name: "has_usable_key", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "heroku_deploy_user", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "hipchat_api_token", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("HipchatAPIToken")},
			// {Name: "hipchat_notify_prefs", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "hipchat_notify", Description: "", Type: proto.ColumnType_BOOL},
			// {Name: "hipchat_room", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "irc_channel", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "irc_keyword", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "irc_notify_prefs", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "irc_password", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "irc_server", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "irc_username", Description: "", Type: proto.ColumnType_STRING},
			{Name: "parallel", Description: "Number of parallel execution.", Type: proto.ColumnType_INT},
			{Name: "setup", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "slack_api_token", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SlackAPIToken")},
			// {Name: "slack_channel", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "slack_notify_prefs", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "slack_subdomain", Description: "", Type: proto.ColumnType_STRING},
			// {Name: "slack_webhook_url", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SlackWebhookURL")},
			{Name: "ssh_keys", Description: "", Type: proto.ColumnType_JSON, Transform: transform.FromField("SSHKeys")},
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

		projectMap := map[string]interface{}{
			"Branches":            project.Branches,
			"CampfireNotifyPrefs": project.CampfireNotifyPrefs,
			"CampfireRoom":        project.CampfireRoom,
			"CampfireSubdomain":   project.CampfireSubdomain,
			"CampfireToken":       project.CampfireToken,
			"Compile":             project.Compile,
			"DefaultBranch":       project.DefaultBranch,
			"Dependencies":        project.Dependencies,
			"Extra":               project.Extra,
			"FeatureFlags":        project.FeatureFlags,
			"FlowdockAPIToken":    project.FlowdockAPIToken,
			"Followed":            project.Followed,
			"HallNotifyPrefs":     project.HallNotifyPrefs,
			"HallRoomAPIToken":    project.HallRoomAPIToken,
			"HasUsableKey":        project.HasUsableKey,
			"HerokuDeployUser":    project.HerokuDeployUser,
			"HipchatAPIToken":     project.HipchatAPIToken,
			"HipchatNotifyPrefs":  project.HipchatNotifyPrefs,
			"HipchatNotify":       project.HipchatNotify,
			"HipchatRoom":         project.HipchatRoom,
			"IrcChannel":          project.IrcChannel,
			"IrcKeyword":          project.IrcKeyword,
			"IrcNotifyPrefs":      project.IrcNotifyPrefs,
			"IrcPassword":         project.IrcPassword,
			"IrcServer":           project.IrcServer,
			"IrcUsername":         project.IrcUsername,
			"Parallel":            project.Parallel,
			"Reponame":            project.Reponame,
			"Setup":               project.Setup,
			"SlackAPIToken":       project.SlackAPIToken,
			"SlackChannel":        project.SlackChannel,
			"SlackNotifyPrefs":    project.SlackNotifyPrefs,
			"SlackSubdomain":      project.SlackSubdomain,
			"SlackWebhookURL":     project.SlackWebhookURL,
			"SSHKeys":             project.SSHKeys,
			"Username":            project.Username,
			"OrganizationSlug":    organizationSlug,
			"Slug":                projectSlug,
			"VCSURL":              project.VCSURL,
		}
		d.StreamListItem(ctx, projectMap)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
