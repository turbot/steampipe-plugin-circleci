---
title: "Steampipe Table: circleci_build - Query CircleCI Builds using SQL"
description: "Allows users to query CircleCI Builds, specifically providing insights into build details such as status, workflows, and execution time."
---

# Table: circleci_build - Query CircleCI Builds using SQL

CircleCI is a Continuous Integration and Delivery platform that helps software teams rapidly release code with confidence. It automates the build, test, and delivery of applications, allowing developers to concentrate on creating new features and fixing bugs. CircleCI provides a rich set of features including parallel execution, exclusive job execution, and insights into your pipelines.

## Table Usage Guide

The `circleci_build` table provides insights into builds within CircleCI. As a DevOps engineer, explore build-specific details through this table, including status, workflows, and execution time. Utilize it to uncover information about builds, such as those with failed tests, the workflows associated with each build, and the execution time of each build.

## Examples

### Last 10 successful builds
Analyze the settings to understand the most recent successful builds in your circleci project. This allows you to keep track of the progress and success rate of your builds, which can aid in improving future build processes.

```sql+postgres
select
  username as "organization",
  reponame,
  branch,
  build_time_millis,
  status,
  author_name,
  build_url
from
  circleci_build
where
  status = 'success'
order by
  stop_time desc limit 10;
```

```sql+sqlite
select
  username as "organization",
  reponame,
  branch,
  build_time_millis,
  status,
  author_name,
  build_url
from
  circleci_build
where
  status = 'success'
order by
  stop_time desc limit 10;
```

### Number of failed builds in a repository
Determine the areas in which the number of failed builds in a repository is high. This can help in identifying problematic repositories that may require extra attention or resources.

```sql+postgres
select
  concat(username, '/', reponame) as repository,
  count(1) as failed_builds
from
  circleci_build b
where
  status = 'failed'
group by
  concat(username, '/', reponame)
order by
  failed_builds desc;
```

```sql+sqlite
select
  username || '/' || reponame as repository,
  count(1) as failed_builds
from
  circleci_build b
where
  status = 'failed'
group by
  username || '/' || reponame
order by
  failed_builds desc;
```

### Average execution time duration of successful builds of a repository (in seconds)
Analyze the performance of a specific repository by determining the average time taken for successful builds. This information can be useful in pinpointing efficiency issues or assessing the effectiveness of recent changes.

```sql+postgres
select
  ROUND(avg(build_time_millis / 1000)) as average_duration
from
  circleci_build
where
  status = 'success'
  and username = 'fluent-cattle'
  and reponame = 'sp-plugin-test'
group by
  status;
```

```sql+sqlite
select
  ROUND(avg(build_time_millis / 1000)) as average_duration
from
  circleci_build
where
  status = 'success'
  and username = 'fluent-cattle'
  and reponame = 'sp-plugin-test'
group by
  status;
```