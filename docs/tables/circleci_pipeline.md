---
title: "Steampipe Table: circleci_pipeline - Query CircleCI Pipelines using SQL"
description: "Allows users to query CircleCI Pipelines, specifically the details of each pipeline run, providing insights into build processes, status, and potential issues."
---

# Table: circleci_pipeline - Query CircleCI Pipelines using SQL

CircleCI is a continuous integration and delivery platform that automates the build, test, and deploy processes for software applications. The Pipeline is a key feature of CircleCI, representing an individual instance of a run workflow, including its status, triggers, and associated metadata. It provides developers with a detailed view of their application's build process, enabling them to quickly identify and resolve issues.

## Table Usage Guide

The `circleci_pipeline` table provides insights into each pipeline run within CircleCI. As a developer or DevOps engineer, explore pipeline-specific details through this table, including status, triggers, and associated metadata. Utilize it to monitor your application's build process, identify bottlenecks or failures, and optimize your continuous integration and delivery workflows.

**Important Notes**
 - You must specify `project_slug` in the `where` clause to query this table.

## Examples

### Error details of the pipelines of a project
Discover the segments that contain errors within a specific project's pipelines. This can be useful in quickly identifying and addressing issues to improve project performance and efficiency.

```sql+postgres
select
  id,
  number,
  errors
from
  circleci_pipeline
where
  project_slug = 'gh/fluent-cattle/sp-plugin-test'
  and errors is not null
  and jsonb_array_length(errors) > 0;
```

```sql+sqlite
select
  id,
  number,
  errors
from
  circleci_pipeline
where
  project_slug = 'gh/fluent-cattle/sp-plugin-test'
  and errors is not null
  and json_array_length(errors) > 0;
```

### Number of pipelines per project
Analyze the settings to understand the distribution of pipelines across various projects. This can help in identifying projects with heavy pipeline utilization, aiding in resource allocation and optimization.

```sql+postgres
select
  pr.slug,
  count(pl.*)
from
  circleci_project pr
  join
    circleci_pipeline pl
    on pr.slug = pl.project_slug
group by
  pr.slug;
```

```sql+sqlite
select
  pr.slug,
  count(pl.project_slug)
from
  circleci_project pr
  join
    circleci_pipeline pl
    on pr.slug = pl.project_slug
group by
  pr.slug;
```

### Pipelines of a project by state
Explore the distribution of pipeline states within a specific project. This can help in identifying patterns or issues related to the project's pipeline states.

```sql+postgres
select
  state,
  count(*)
from
  circleci_pipeline
where
  project_slug = 'gh/fluent-cattle/sp-plugin-test'
group by
  state;
```

```sql+sqlite
select
  state,
  count(*)
from
  circleci_pipeline
where
  project_slug = 'gh/fluent-cattle/sp-plugin-test'
group by
  state;
```