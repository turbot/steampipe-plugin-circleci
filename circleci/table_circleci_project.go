package circleci

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleciProject() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_project",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listCircleciProjects,
		},

		Columns: []*plugin.Column{
			{Name: "branches", Description: "", Type: proto.ColumnType_JSON},
			{Name: "campfire_notify_prefs", Description: "", Type: proto.ColumnType_STRING},
			{Name: "campfire_room", Description: "", Type: proto.ColumnType_STRING},
			{Name: "campfire_subdomain", Description: "", Type: proto.ColumnType_STRING},
			{Name: "campfire_token", Description: "", Type: proto.ColumnType_STRING},
			{Name: "compile", Description: "", Type: proto.ColumnType_STRING},
			{Name: "default_branch", Description: "", Type: proto.ColumnType_STRING},
			{Name: "dependencies", Description: "", Type: proto.ColumnType_STRING},
			{Name: "extra", Description: "", Type: proto.ColumnType_STRING},
			{Name: "feature_flags", Description: "", Type: proto.ColumnType_JSON},
			{Name: "flowdock_api_token", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("FlowdockAPIToken")},
			{Name: "followed", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "hall_notify_prefs", Description: "", Type: proto.ColumnType_STRING},
			{Name: "hall_room_api_token", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("HallRoomAPIToken")},
			{Name: "has_usable_key", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "heroku_deploy_user", Description: "", Type: proto.ColumnType_STRING},
			{Name: "hipchat_api_token", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("HipchatAPIToken")},
			{Name: "hipchat_notify_prefs", Description: "", Type: proto.ColumnType_STRING},
			{Name: "hipchat_notify", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "hipchat_room", Description: "", Type: proto.ColumnType_STRING},
			{Name: "irc_channel", Description: "", Type: proto.ColumnType_STRING},
			{Name: "irc_keyword", Description: "", Type: proto.ColumnType_STRING},
			{Name: "irc_notify_prefs", Description: "", Type: proto.ColumnType_STRING},
			{Name: "irc_password", Description: "", Type: proto.ColumnType_STRING},
			{Name: "irc_server", Description: "", Type: proto.ColumnType_STRING},
			{Name: "irc_username", Description: "", Type: proto.ColumnType_STRING},
			{Name: "parallel", Description: "", Type: proto.ColumnType_INT},
			{Name: "reponame", Description: "", Type: proto.ColumnType_STRING},
			{Name: "setup", Description: "", Type: proto.ColumnType_STRING},
			{Name: "slack_api_token", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SlackAPIToken")},
			{Name: "slack_channel", Description: "", Type: proto.ColumnType_STRING},
			{Name: "slack_notify_prefs", Description: "", Type: proto.ColumnType_STRING},
			{Name: "slack_subdomain", Description: "", Type: proto.ColumnType_STRING},
			{Name: "slack_webhook_url", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("SlackWebhookURL")},
			{Name: "ssh_keys", Description: "", Type: proto.ColumnType_JSON, Transform: transform.FromField("SSHKeys")},
			{Name: "test", Description: "", Type: proto.ColumnType_STRING},
			{Name: "username", Description: "", Type: proto.ColumnType_STRING},
			{Name: "vcs_url", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("VCSURL")},
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

	for _, project := range projects {
		d.StreamListItem(ctx, project)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
