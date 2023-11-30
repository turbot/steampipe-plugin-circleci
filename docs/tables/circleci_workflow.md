---
title: "Steampipe Table: circleci_workflow - Query CircleCI Workflows using SQL"
description: "Allows users to query CircleCI Workflows, specifically the details of each workflow run, providing insights into build status, project details, pipeline information, and more."
---

# Table: circleci_workflow - Query CircleCI Workflows using SQL

CircleCI is a Continuous Integration and Continuous Deployment (CI/CD) platform that automates the build, test, and deployment of applications. A CircleCI Workflow is a set of rules for defining a collection of jobs and their run order. Workflows manage the jobs that you have defined in your configuration and the order in which they run.

## Table Usage Guide

The `circleci_workflow` table provides insights into Workflows within CircleCI. As an engineer or developer, explore workflow-specific details through this table, including status, project details, pipeline information, and more. Utilize it to uncover information about workflows, such as those with failed jobs, the run order of jobs within a workflow, and the overall status of workflows.

## Examples

### Workflow status of a pipeline
Explore the status history of a specific pipeline to understand its progression over time. This can be useful in identifying patterns or issues in the pipeline's workflow.

```sql
select
  id,
  created_at,
  status
from
  circleci_workflow
where
  pipeline_id = 'f43cc52a-c7eb-4a72-a05f-399c8577bb3e'
order by
  created_at desc;
```

### Workflow duration of a pipeline
Analyze the duration of a specific pipeline's workflow in CircleCI. This can be useful in assessing the efficiency of the pipeline, identifying potential areas for optimization.

```sql
select
  id,
  project_slug,
  extract(seconds
from
  (
    stopped_at - created_at
  )
) as duration_in_seconds
from
  circleci_workflow
where
  pipeline_id = 'f43cc52a-c7eb-4a72-a05f-399c8577bb3e';
```