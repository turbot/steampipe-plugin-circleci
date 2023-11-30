---
title: "Steampipe Table: circleci_project - Query CircleCI Projects using SQL"
description: "Allows users to query CircleCI Projects, specifically project-level details and configurations, providing insights into project management and workflow orchestration."
---

# Table: circleci_project - Query CircleCI Projects using SQL

CircleCI is a continuous integration and delivery platform that automates the build, test, and deploy process for software applications. It allows developers to rapidly release code by automating the build, test, and delivery process. It integrates seamlessly with GitHub, Bitbucket, and other version control systems, making it a popular choice for software development teams.

## Table Usage Guide

The `circleci_project` table provides insights into projects within CircleCI. As a DevOps engineer, explore project-specific details through this table, including vcs type, username, project name, and default branch. Utilize it to uncover information about projects, such as those with specific configurations, the default branches for each project, and the version control systems in use.

## Examples

### List checkout keys of a project
Explore the different checkout keys associated with a specific project to gain insights into their attributes such as fingerprint, preference, public key, time, type, and login. This can be particularly useful for project management and security purposes, such as tracking key usage and identifying preferred keys.

```sql
select
  k ->> 'fingerprint' as fingerprint,
  k ->> 'preferred' as preferred,
  k ->> 'public_key' as public_key,
  k ->> 'time' as time,
  k ->> 'type' as type,
  k ->> 'login' as login
from
  circleci_project p,
  jsonb_array_elements(p.checkout_keys) k
where
  p.slug = 'gh/fluent-cattle/sp-plugin-test'
order by
  k ->> 'time';
```

### Projects with builds running on the main branch
Explore which projects are currently executing builds on the main branch. This can be useful in identifying active development efforts and monitoring build statuses in real-time.

```sql
select
  concat(username, '/', reponame) as repository,
  'main' as branch,
  jsonb_array_length(branches -> 'main' -> 'running_builds') as running_builds
from
  circleci_project
where
  (
    branches -> 'main'
  )
  is not null;
```

### Project's last successful build (main branch)
Analyze the settings to understand the last successful build for a project's main branch in CircleCI. This is useful for tracking the progress and stability of your main branch over time.

```sql
select
  concat(username, '/', reponame) as repository,
  branches -> 'main' -> 'last_success' -> 'build_num' as build_num
from
  circleci_project
where
  (
    branches -> 'main'
  )
  is not null;
```