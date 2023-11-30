---
title: "Steampipe Table: circleci_insights_workflow_run - Query CircleCI Workflow Runs using SQL"
description: "Allows users to query CircleCI Workflow Runs, providing detailed insights into each workflow run's performance and status."
---

# Table: circleci_insights_workflow_run - Query CircleCI Workflow Runs using SQL

CircleCI is a continuous integration and delivery platform that automates the build, test, and deploy processes for software. The Workflow Runs in CircleCI are individual executions of a pipeline, which include one or more jobs configured in the `.circleci/config.yml` file. Workflow Runs provide detailed information about the execution of jobs, including status, duration, and outcome.

## Table Usage Guide

The `circleci_insights_workflow_run` table provides insights into Workflow Runs within CircleCI. If you are a DevOps engineer or a software developer, you can use this table to monitor and analyze the performance of your software's build, test, and deploy processes. The table can also aid in identifying any issues or bottlenecks in these processes, thereby helping you optimize your software's continuous integration and delivery pipeline.

## Examples

### Get average duration and deployment count of a project for each month
Analyze the performance of a specific project over time by determining the average duration and number of successful deployments each month. This can help assess the efficiency of the project's workflow and identify potential areas for improvement.

```sql
select
  project_slug,
  workflow_name,
  id,
  to_char(created_at, 'YYYY-MM') as year_month,
  avg(duration) as average_duration,
  count(id) as deployment_count
from
  circleci_insights_workflow_run
where
  workflow_name = 'default'
  and project_slug = 'gh/companyname/projectname'
  and branch = 'main'
  and status = 'success'
group by
  project_slug,
  year_month;
```

### List workflows created in the last 30 days
Explore the recent workflows created in your project over the past month. This is useful for tracking project progress and assessing the duration of each workflow.

```sql
select
  workflow_name,
  branch,
  id,
  project_slug,
  duration
from
  circleci_insights_workflow_run
where
  workflow_name = 'default'
  and project_slug = 'gh/companyname/projectname'
  and created_at >= current_date - interval '30' day;
```

### List workflows which are in failed state
Uncover the details of failed workflows within a specific project. This query is useful to identify bottlenecks, analyze project performance, and implement corrective measures.

```sql
select
  workflow_name,
  branch,
  id,
  project_slug,
  duration
from
  circleci_insights_workflow_run
where
  workflow_name = 'default'
  and project_slug = 'gh/companyname/projectname'
  and status = 'failed';
```