package circleci

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/turbot/steampipe-plugin-circleci/rest"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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

	workflow := h.Item.(map[string]string)
	logger.Debug("listCircleCIInsightsWorkflowRuns", "project workflow", workflow["project_slug"]+" "+workflow["workflow_name"])
	if projectSlug == "" {
		projectSlug = workflow["project_slug"]
	}
	if workflowName == "" {
		workflowName = workflow["workflow_name"]
	}
	_, err := listSingleWorkflowRuns(ctx, d, logger, projectSlug, workflowName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func listSingleWorkflowRuns(ctx context.Context, d *plugin.QueryData, logger hclog.Logger, projectSlug string, workflowName string) (interface{}, error) {
	branch := ""
	if d.EqualsQuals["branch"] != nil {
		branch = d.EqualsQualString("branch")
	}
	startDate, endDate := getStartDateAndEndDate(d)
	logger.Info("circleci_insights_workflow_run.listSingleWorkflowRuns", "branch", branch)

	if projectSlug == "" || workflowName == "" {
		return nil, nil
	}

	projectSlugSplit := strings.Split(projectSlug, "/")
	if len(projectSlugSplit) < 3 {
		err := errors.New("Malformed input for project_slug. Expected: {VCS}/{Org username}/{Repository name}")
		logger.Error("circleci_insights_workflow_run.listSingleWorkflowRuns", "malformed_input", err)
		return nil, err
	}

	client, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_insights_workflow_run.listSingleWorkflowRuns", "connect_error", err)
		return nil, err
	}

	// Generate a random sleep duration between 0 and 5 milliseconds to avoid hitting the rate limit
	rand.Seed(time.Now().UnixNano())
	sleepDuration := time.Duration(rand.Intn(5)) * time.Millisecond
	time.Sleep(sleepDuration)

	workflows, err := client.ListAllInsightsWorkflowRuns(projectSlug, workflowName, branch, startDate, endDate, logger)
	if err != nil {
		logger.Error("circleci_insights_workflow_run.listSingleWorkflowRuns", "list_insight_error", err)
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
		response := map[string]string{
			"project_slug":  projectSlugQual,
			"workflow_name": workflowNameQual,
		}
		d.StreamListItem(ctx, response)
		return nil, nil
	}

	clientV2, err := ConnectV2RestApi(ctx, d)
	if err != nil {
		logger.Error("circleci_insights_workflow_run.parentCircleCIWorkflows", "connect_v2_sdk_error", err)
		return nil, err
	}

	organizations, err := clientV2.ListOrganizations()
	if err != nil {
		logger.Error("circleci_insights_workflow_run.parentCircleCIWorkflows", "list_organizations_error", err)
		return nil, err
	}

	for _, organization := range *organizations {
		logger.Debug("circleci_insights_workflow_run.parentCircleCIWorkflows", "organization", organization.Slug)
		vcs := strings.Split(organization.Slug, "/")[0]

		pipelines, err := clientV2.ListPipelines(vcs, organization.Name)
		if err != nil {
			if strings.Contains(err.Error(), "Organization not found") {
				logger.Warn("circleci_insights_workflow_run.parentCircleCIWorkflows", "list_pipelines_error", fmt.Sprintf("Organization not found: %s", organization.Slug))
				continue
			}
			logger.Error("circleci_insights_workflow_run.parentCircleCIWorkflows", "list_pipelines_error", err)
			return nil, err
		}

		// a list of map[string]string{} to store the unique workflows
		// this is to avoid duplicate workflows
		var uniqueWorkflows []map[string]string

		// list all the pipelines for the organization
		for _, pipeline := range pipelines.Items {
			// if the project slug is provided, skip the other projects
			if projectSlugQual != "" && pipeline.ProjectSlug != projectSlugQual {
				continue
			}

			// Generate a random sleep duration between 0 and 3 milliseconds to avoid hitting the rate limit
			rand.Seed(time.Now().UnixNano())
			sleepDuration := time.Duration(rand.Intn(3)) * time.Millisecond
			time.Sleep(sleepDuration)

			// list all the workflows for the pipeline
			workflows, err := clientV2.ListPipelinesWorkflow(pipeline.ID)
			if err != nil {
				logger.Error("circleci_insights_workflow_run.parentCircleCIWorkflows", "list_workflows_error", err)
				return nil, err
			}
			for _, workflow := range workflows.Items {
				// if the workflow name is provided, skip the other workflows
				if workflowNameQual != "" && workflow.Name != workflowNameQual {
					continue
				}

				// adds the workflow to the uniqueWorkflows list
				// if the workflow is already in the list, it will be skipped
				if !isWorkflowInList(uniqueWorkflows, workflow) {
					uniqueWorkflows = append(uniqueWorkflows, map[string]string{
						"project_slug":  workflow.ProjectSlug,
						"workflow_name": workflow.Name,
					})
				}
			}
		}
		// streams the unique workflows
		for _, uniqueWorkflow := range uniqueWorkflows {
			d.StreamListItem(ctx, uniqueWorkflow)
		}

	}
	return nil, nil
}

func isWorkflowInList(workflows []map[string]string, workflow rest.Workflow) bool {
	for _, w := range workflows {
		if w["project_slug"] == workflow.ProjectSlug && w["workflow_name"] == workflow.Name {
			return true
		}
	}
	return false
}
