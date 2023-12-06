---
title: "Steampipe Table: circleci_organization - Query CircleCI Organizations using SQL"
description: "Allows users to query CircleCI Organizations, specifically the details about each organization, providing insights into the structure and configurations of the organizations."
---

# Table: circleci_organization - Query CircleCI Organizations using SQL

A CircleCI Organization is a workspace for teams to collaborate on projects. It can include multiple projects and users. The organization in CircleCI provides a centralized way to manage projects and users, and also allows the configuration of settings at the organization level.

## Table Usage Guide

The `circleci_organization` table provides insights into organizations within CircleCI. As a DevOps engineer, explore organization-specific details through this table, including vcs settings, trial usage, and associated metadata. Utilize it to uncover information about organizations, such as their settings, the usage of trial periods, and the verification of vcs settings.

## Examples

### Organizations I have access to
Explore the different organizations you have access to and their associated version control systems, allowing you to better manage your resources and align your operations with your access privileges.

```sql+postgres
select
  name,
  vcs_type
from
  circleci_organization
order by
  name;
```

```sql+sqlite
select
  name,
  vcs_type
from
  circleci_organization
order by
  name;
```

### Context environment variables not updated for more then 90 days across my organizations
Determine areas in your organization where context environment variables have not been updated for more than 90 days. This is useful for maintaining up-to-date configurations and ensuring optimal performance.

```sql+postgres
select
  c.organization_slug,
  v.context_id,
  c.name as context,
  v.variable,
  v.created_at,
  v.updated_at
from
  circleci_context_environment_variable v
  join
    circleci_context c
    on c.id = v.context_id
where
  updated_at < current_date - interval '90' day;
```

```sql+sqlite
select
  c.organization_slug,
  v.context_id,
  c.name as context,
  v.variable,
  v.created_at,
  v.updated_at
from
  circleci_context_environment_variable v
  join
    circleci_context c
    on c.id = v.context_id
where
  updated_at < date('now','-90 day');
```