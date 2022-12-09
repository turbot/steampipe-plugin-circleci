package circleci

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleciBuild() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_build",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listCircleciBuilds,
		},

		Columns: []*plugin.Column{
			{Name: "all_commit_details", Description: "", Type: proto.ColumnType_JSON},
			{Name: "author_date", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "author_email", Description: "", Type: proto.ColumnType_STRING},
			{Name: "author_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "body", Description: "", Type: proto.ColumnType_STRING},
			{Name: "branch", Description: "", Type: proto.ColumnType_STRING},
			{Name: "build_num", Description: "", Type: proto.ColumnType_INT},
			{Name: "build_parameters", Description: "", Type: proto.ColumnType_JSON},
			{Name: "build_time_millis", Description: "", Type: proto.ColumnType_INT},
			{Name: "build_url", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("BuildURL")},
			{Name: "canceled", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "circle_yml", Description: "", Type: proto.ColumnType_JSON, Transform: transform.FromField("CircleYML")},
			{Name: "committer_date", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "committer_email", Description: "", Type: proto.ColumnType_STRING},
			{Name: "committer_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "compare", Description: "", Type: proto.ColumnType_STRING},
			{Name: "dont_build", Description: "", Type: proto.ColumnType_STRING},
			{Name: "failed", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "feature_flags", Description: "", Type: proto.ColumnType_JSON},
			{Name: "infrastructure_fail", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "is_first_green_build", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "job_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "lifecycle", Description: "", Type: proto.ColumnType_STRING},
			{Name: "messages", Description: "", Type: proto.ColumnType_JSON},
			{Name: "node", Description: "", Type: proto.ColumnType_JSON},
			{Name: "oss", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "outcome", Description: "", Type: proto.ColumnType_STRING},
			{Name: "parallel", Description: "", Type: proto.ColumnType_INT},
			{Name: "picard", Description: "", Type: proto.ColumnType_JSON},
			{Name: "platform", Description: "", Type: proto.ColumnType_STRING},
			{Name: "previous_successful_build", Description: "", Type: proto.ColumnType_JSON},
			{Name: "previous", Description: "", Type: proto.ColumnType_JSON},
			{Name: "pull_requests", Description: "", Type: proto.ColumnType_JSON},
			{Name: "queued_at", Description: "", Type: proto.ColumnType_STRING},
			{Name: "reponame", Description: "", Type: proto.ColumnType_STRING},
			{Name: "retries", Description: "", Type: proto.ColumnType_JSON},
			{Name: "retry_of", Description: "", Type: proto.ColumnType_INT},
			{Name: "ssh_enabled", Description: "", Type: proto.ColumnType_BOOL, Transform: transform.FromField("SSHEnabled")},
			{Name: "ssh_users", Description: "", Type: proto.ColumnType_JSON, Transform: transform.FromField("SSHUsers")},
			{Name: "start_time", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "status", Description: "", Type: proto.ColumnType_STRING},
			{Name: "steps", Description: "", Type: proto.ColumnType_JSON},
			{Name: "stop_time", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "subject", Description: "", Type: proto.ColumnType_STRING},
			{Name: "timedout", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "usage_queued_at", Description: "", Type: proto.ColumnType_STRING},
			{Name: "user", Description: "", Type: proto.ColumnType_JSON},
			{Name: "username", Description: "", Type: proto.ColumnType_STRING},
			{Name: "vcs_revision", Description: "", Type: proto.ColumnType_STRING},
			{Name: "vcs_tag", Description: "", Type: proto.ColumnType_STRING},
			{Name: "vcs_url", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("VCSURL")},
			{Name: "why", Description: "", Type: proto.ColumnType_STRING},
			{Name: "workflows", Description: "", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listCircleciBuilds(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := ConnectV1Sdk(ctx, d)
	if err != nil {
		logger.Error("circleci_build.listCircleciBuilds", "connect_error", err)
		return nil, err
	}

	limit := -1
	offset := 0

	builds, err := client.ListRecentBuilds(limit, offset)
	if err != nil {
		logger.Error("circleci_build.listCircleciBuilds", "list_builds_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, build := range builds {
		d.StreamListItem(ctx, build)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
