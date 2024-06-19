package circleci

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableCircleCIBuild() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_build",
		Description: "A CircleCI build is a result of a single execution of a workflow.",
		List: &plugin.ListConfig{
			Hydrate: listCircleCIBuilds,
		},

		Columns: commonColumns([]*plugin.Column{
			{Name: "build_num", Description: "Sequential number of build.", Type: proto.ColumnType_INT},
			{Name: "project_slug", Description: "A unique identification for the project in the form of: <vcs_type>/<org_name>/<repo_name>.", Type: proto.ColumnType_STRING},
			{Name: "organization_slug", Description: "Organization that pipeline belongs to, in the form of: <vcs_type>/<org_name>.", Type: proto.ColumnType_STRING},
			{Name: "reponame", Description: "Repository name.", Type: proto.ColumnType_STRING},
			{Name: "username", Description: "Organization username.", Type: proto.ColumnType_STRING},
			{Name: "branch", Description: "Branch used to build the code.", Type: proto.ColumnType_STRING},
			{Name: "build_url", Description: "Build URL.", Type: proto.ColumnType_STRING, Transform: transform.FromField("BuildURL")},
			{Name: "status", Description: "Status of the build.", Type: proto.ColumnType_STRING},
			{Name: "all_commit_details", Description: "Commit details.", Type: proto.ColumnType_JSON},
			{Name: "author_email", Description: "Author email.", Type: proto.ColumnType_STRING},
			{Name: "author_name", Description: "Author name.", Type: proto.ColumnType_STRING},
			{Name: "build_parameters", Description: "Any parameter optional or required to build.", Type: proto.ColumnType_JSON},
			{Name: "build_time_millis", Description: "Duration of the build.", Type: proto.ColumnType_INT},
			{Name: "canceled", Description: "Indicates if the build was canceled.", Type: proto.ColumnType_BOOL},
			{Name: "committer_date", Description: "Committer Date.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "committer_email", Description: "Committer email.", Type: proto.ColumnType_STRING},
			{Name: "committer_name", Description: "Committer name.", Type: proto.ColumnType_STRING},
			{Name: "failed", Description: "Indicates if the build failed.", Type: proto.ColumnType_BOOL},
			{Name: "infrastructure_fail", Description: "Indicates if the build failed due to infrastructure.", Type: proto.ColumnType_BOOL},
			{Name: "is_first_green_build", Description: "Indicates if this is the first succeeded build of the project.", Type: proto.ColumnType_BOOL},
			{Name: "lifecycle", Description: "Build lifecycle.", Type: proto.ColumnType_STRING},
			{Name: "outcome", Description: "Result of the build.", Type: proto.ColumnType_STRING},
			{Name: "parallel", Description: "Number of parallel executions.", Type: proto.ColumnType_INT},
			{Name: "platform", Description: "Platform version where build ran.", Type: proto.ColumnType_STRING},
			{Name: "previous", Description: "Previous build.", Type: proto.ColumnType_JSON},
			{Name: "queued_at", Description: "Timestamp of when the build was queued.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "retries", Description: "Number of build retries.", Type: proto.ColumnType_JSON},
			{Name: "retry_of", Description: "Precedent build of the retry.", Type: proto.ColumnType_INT},
			{Name: "ssh_users", Description: "SSH users with access to the build, if any.", Type: proto.ColumnType_JSON, Transform: transform.FromField("SSHUsers")},
			{Name: "start_time", Description: "Start time of the build.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "stop_time", Description: "Stop time of the build.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "subject", Description: "Commit message that triggered the build.", Type: proto.ColumnType_STRING},
			{Name: "timed_out", Description: "Indicates if the build got timed out.", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Timedout")},
			{Name: "usage_queued_at", Description: "Usage queued time.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "user", Description: "User who triggered the build to run.", Type: proto.ColumnType_JSON},
			{Name: "vcs_revision", Description: "VCS Revision.", Type: proto.ColumnType_STRING},
			{Name: "vcs_tag", Description: "VCS Tag.", Type: proto.ColumnType_STRING},
			{Name: "vcs_url", Description: "VCS URL.", Type: proto.ColumnType_STRING, Transform: transform.FromField("VCSURL")},
			{Name: "workflows", Description: "Workflow which ran the build.", Type: proto.ColumnType_JSON, Transform: transform.FromField("workflows")},
		}),
	}
}

//// LIST FUNCTION

func listCircleCIBuilds(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := ConnectV1Sdk(ctx, d)
	if err != nil {
		logger.Error("circleci_build.listCircleCIBuilds", "connect_error", err)
		return nil, err
	}

	// limit -1 means it will use the SDK built-in pagination
	limit := -1
	offset := 0

	builds, err := client.ListRecentBuilds(limit, offset)
	if err != nil {
		logger.Error("circleci_build.listCircleciBuilds", "query_error", err)
		return nil, err
	}

	for _, build := range builds {
		organizationSlug, projectSlug := Slugify(build.VCSURL, build.Username, build.Reponame)
		buildMap := map[string]interface{}{
			"OrganizationSlug":   organizationSlug,
			"ProjectSlug":        projectSlug,
			"BuildNum":           build.BuildNum,
			"Reponame":           build.Reponame,
			"Username":           build.Username,
			"Branch":             build.Branch,
			"AllCommitDetails":   build.AllCommitDetails,
			"AuthorEmail":        build.AuthorEmail,
			"AuthorName":         build.AuthorName,
			"BuildParameters":    build.BuildParameters,
			"BuildTimeMillis":    build.BuildTimeMillis,
			"BuildURL":           build.BuildURL,
			"Canceled":           build.Canceled,
			"CommitterDate":      build.CommitterDate,
			"CommitterEmail":     build.CommitterEmail,
			"CommitterName":      build.CommitterName,
			"Failed":             build.Failed,
			"InfrastructureFail": build.InfrastructureFail,
			"IsFirstGreenBuild":  build.IsFirstGreenBuild,
			"Lifecycle":          build.Lifecycle,
			"Outcome":            build.Outcome,
			"Parallel":           build.Parallel,
			"Platform":           build.Platform,
			"Previous":           build.Previous,
			"QueuedAt":           build.QueuedAt,
			"Retries":            build.Retries,
			"RetryOf":            build.RetryOf,
			"SSHUsers":           build.SSHUsers,
			"StartTime":          build.StartTime,
			"Status":             build.Status,
			"StopTime":           build.StopTime,
			"Subject":            build.Subject,
			"Timedout":           build.Timedout,
			"UsageQueuedAt":      build.UsageQueuedAt,
			"User":               build.User,
			"VcsRevision":        build.VcsRevision,
			"VcsTag":             build.VcsTag,
			"VCSURL":             build.VCSURL,
			"Workflows":          build.Workflows,
		}
		d.StreamListItem(ctx, buildMap)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
