---
title: "Steampipe Table: circleci_context - Query CircleCI Contexts using SQL"
description: "Allows users to query CircleCI Contexts, specifically to gather detailed information about each context including its name, ID, and created at timestamp."
---

# Table: circleci_context - Query CircleCI Contexts using SQL

CircleCI Contexts is a feature within CircleCI that allows you to create and manage contexts. Contexts are used to share environment variables across multiple projects. These environment variables are encrypted and stored securely, allowing you to manage access to sensitive information within your pipelines.

## Table Usage Guide

The `circleci_context` table provides insights into contexts within CircleCI. As a DevOps engineer, explore context-specific details through this table, including names, IDs, and creation timestamps. Utilize it to uncover information about contexts, such as their associated environment variables, the projects they are shared with, and the management of access to sensitive data.

## Examples

### Contexts across my organizations
Explore the various contexts within your organizations to understand when they were created and their correlation with specific organization slugs. This can help in managing and organizing your resources more efficiently.

```sql
select
  id,
  name,
  organization_slug,
  created_at
from
  circleci_context;
```

### Context environment variables of an organization
Discover the environment variables within a specific organization's context to understand when they were created or last updated. This can help in maintaining the organization's configuration and track any changes made over time.

```sql
select
  c.name as context,
  v.context_id,
  v.variable,
  v.created_at,
  v.updated_at
from
  circleci_context c
  join
    circleci_context_environment_variable v
    on v.context_id = c.id
where
  organization_slug = 'gh/fluent-cattle';
```