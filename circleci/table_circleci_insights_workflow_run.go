package circleci

import (
	"context"
	"errors"
	"github.com/hashicorp/go-hclog"
	"github.com/turbot/steampipe-plugin-circleci/rest"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"strings"
	"time"
)

//// TABLE DEFINITION

func tableCircleCIInsightsWorkflowRun() *plugin.Table {
	return &plugin.Table{
		Name:        "circleci_insights_workflow_run",
		Description: "Get insights on project workflows runs.",
		List: &plugin.ListConfig{
			Hydrate:       listCircleCIInsightsWorkflowRuns,
			ParentHydrate: parentCircleCIWorkflows,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "project_slug", Require: plugin.Optional},
				{Name: "workflow_name", Require: plugin.Optional},
				{Name: "branch", Require: plugin.Optional},
				{Name: "created_at", Require: plugin.Optional},
			},
		},

		Columns: []*plugin.Column{
			{Name: "project_slug", Description: "A unique identification for the project in the form of: <vcs_type>/<org_name>/<repo_name>.", Type: proto.ColumnType_STRING},
			{Name: "workflow_name", Description: "The name of the workflow.", Type: proto.ColumnType_STRING},
			{Name: "id", Description: "Unique key for the workflow.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "branch", Description: "The VCS branch of a Workflow's trigger.", Type: proto.ColumnType_STRING},
			{Name: "duration", Description: "Duration of the workflow in seconds.", Type: proto.ColumnType_INT},
			{Name: "created_at", Description: "Timestamp of when the workflow was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "stopped_at", Description: "Timestamp of when workflow was stopped.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "credits_used", Description: "The number of credits used during execution.", Type: proto.ColumnType_INT},
			{Name: "status", Description: "Workflow status.", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listCircleCIInsightsWorkflowRuns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	projectSlug := d.EqualsQualString("project_slug")
	workflowName := d.EqualsQualString("workflow_name")
	if projectSlug != "" && workflowName != "" {
		return listSingleWorkflowRuns(ctx, d, logger, projectSlug, workflowName)
	}

	workflows := h.Item.(*rest.WorkflowResponse)
	for _, workflow := range workflows.Items {
		if projectSlug == "" {
			projectSlug = workflow.ProjectSlug
		}
		if workflowName == "" {
			workflowName = workflow.Name
		}
		_, err := listSingleWorkflowRuns(ctx, d, logger, projectSlug, workflowName)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func listSingleWorkflowRuns(ctx context.Context, d *plugin.QueryData, logger hclog.Logger, projectSlug string, workflowName string) (interface{}, error) {
	branch := ""
	if d.EqualsQuals["branch"] != nil {
		branch = d.EqualsQualString("branch")
	}
	startDate, endDate := getStartDateAndEndDate(d)
	logger.Info("circleci_insights_workflow_run.listCircleCIInsightsWorkflowRuns", "branch", branch)

	if projectSlug == "" || workflowName == "" {
		return nil, nil
	}

	projectSlugSplit := strings.Split(projectSlug, "/")
	if len(projectSlugSplit) < 3 {
		err := errors.New("Malformed input for project_slug. Expected: {VCS}/{Org username}/{Repository name}")
		logger.Error("circleci_insights_workflow_run.listCircleCIInsightsWorkflowRuns", "malformed_input", err)
		return nil, err
	}

	client, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_insights_workflow_run.listCircleCIInsightsWorkflowRuns", "connect_error", err)
		return nil, err
	}
	workflows, err := client.ListAllInsightsWorkflowRuns(projectSlug, workflowName, branch, startDate, endDate, logger)
	if err != nil {
		logger.Error("circleci_insights_workflow_run.listCircleCIInsightsWorkflowRuns", "list_insight_error", err)
		return nil, err
	}

	for _, workflow := range workflows {
		// These fields are not provided by the API, so we set them from the query arguments
		workflow.ProjectSlug = projectSlug
		workflow.WorkflowName = workflowName

		d.StreamListItem(ctx, workflow)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

func getStartDateAndEndDate(d *plugin.QueryData) (string, string) {
	startDate := ""
	endDate := ""
	if d.QueryContext.UnsafeQuals["created_at"] != nil {
		createdAtQuals := d.QueryContext.UnsafeQuals["created_at"].Quals
		if createdAtQuals != nil {
			for _, qual := range createdAtQuals {
				if _, ok := qual.GetOperator().(*proto.Qual_StringValue); ok {
					operator := qual.GetOperator().(*proto.Qual_StringValue).StringValue
					if operator == ">" || operator == ">=" {
						startDate = qual.Value.GetTimestampValue().AsTime().Format(time.RFC3339)
					}
					if operator == "<" || operator == "<=" {
						endDate = qual.Value.GetTimestampValue().AsTime().Format(time.RFC3339)
					}
				}
			}
		}
	}
	if startDate == "" {
		// end-date can be used only with start-date
		endDate = ""
	}
	return startDate, endDate
}

func parentCircleCIWorkflows(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	projectSlugQual := d.EqualsQualString("project_slug")
	workflowNameQual := d.EqualsQualString("workflow_name")
	if projectSlugQual != "" && workflowNameQual != "" {
		response := &rest.WorkflowResponse{
			Items: []rest.Workflow{{
				ProjectSlug: projectSlugQual,
				Name:        workflowNameQual,
			}},
		}
		d.StreamListItem(ctx, response)
		return nil, nil
	}
	clientV1, err := ConnectV1Sdk(ctx, d)
	if err != nil {
		logger.Error("circleci_insights_workflow_run.parentCircleCIWorkflows", "connect_v1_sdk_error", err)
		return nil, err
	}

	clientV2, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_insights_workflow_run.parentCircleCIWorkflows", "connect_v2_sdk_error", err)
		return nil, err
	}

	projects, err := clientV1.ListProjects()
	if err != nil {
		logger.Error("circleci_insights_workflow_run.parentCircleCIWorkflows", "list_projects_error", err)
		return nil, err
	}

	for _, project := range projects {
		_, projectSlug := Slugify(project.VCSURL, project.Username, project.Reponame)
		if projectSlug != projectSlugQual {
			continue
		}
		projectSlugSplit := strings.Split(projectSlug, "/")
		if len(projectSlugSplit) < 3 {
			err := errors.New("Malformed input for project_slug. Expected: {VCS}/{Org username}/{Repository name}")
			logger.Error("circleci_insights_workflow_run.parentCircleCIWorkflows", "malformed_input", err)
			return nil, err
		}
		vcs := projectSlugSplit[0]
		organization := projectSlugSplit[1]

		pipelines, err := clientV2.ListPipelines(vcs, organization)
		if err != nil {
			logger.Error("circleci_insights_workflow_run.parentCircleCIWorkflows", "list_pipelines_error", err)
			return nil, err
		}

		for _, pipeline := range pipelines.Items {
			workflows, err := clientV2.ListPipelinesWorkflow(pipeline.ID)
			if err != nil {
				logger.Error("circleci_insights_workflow_run.parentCircleCIWorkflows", "list_workflows_error", err)
				return nil, err
			}
			d.StreamListItem(ctx, workflows)
		}

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}
